package yorm

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"github.com/timchine/pole/pkg/yorm/schema"
	"reflect"
	"time"
)

func Scan(rows *sql.Rows, stmp *Statement, obj interface{}) {
	objV := reflect.ValueOf(obj)
	if objV.Kind() == reflect.Ptr {
		objV = objV.Elem()
	}
	columns, _ := rows.Columns()
	values := make([]interface{}, len(columns))

	switch dest := obj.(type) {
	case map[string]interface{}, *map[string]interface{}:
		if rows.Next() {
			columnTypes, _ := rows.ColumnTypes()
			prepareValues(values, columnTypes, columns)

			stmp.AddError(rows.Scan(values...))

			mapValue, ok := dest.(map[string]interface{})
			if !ok {
				if v, ok := dest.(*map[string]interface{}); ok {
					mapValue = *v
				}
			}
			scanIntoMap(mapValue, values, columns)
		}
	case *[]map[string]interface{}:
		columnTypes, _ := rows.ColumnTypes()
		for rows.Next() {
			prepareValues(values, columnTypes, columns)
			stmp.AddError(rows.Scan(values...))

			mapValue := map[string]interface{}{}
			scanIntoMap(mapValue, values, columns)
			*dest = append(*dest, mapValue)
		}
	case *int, *int64, *uint, *uint64, *float32, *float64, *string, *time.Time:
		for rows.Next() {
			stmp.AddError(rows.Scan(dest))
		}
	default:
		Schema := stmp.Schema
		switch objV.Kind() {
		case reflect.Slice, reflect.Array:
			var (
				reflectValueType = objV.Type().Elem()
				isPtr            = reflectValueType.Kind() == reflect.Ptr
				fields           = make([]*schema.Field, len(columns))
			)

			if isPtr {
				reflectValueType = reflectValueType.Elem()
			}

			objV.Set(reflect.MakeSlice(objV.Type(), 0, 0))

			Schema, _ = schema.Parse(obj, stmp.CacheStore)

			for idx, column := range columns {
				if field := Schema.LookUpField(changeString(column)); field != nil && field.Readable {
					fields[idx] = field
				} else if field := Schema.LookUpField(column); field != nil && field.Readable {
					fields[idx] = field
				} else {
					values[idx] = &sql.RawBytes{}
				}
			}

			// pluck values into slice of data
			isPluck := false
			if len(fields) == 1 {
				if _, ok := reflect.New(reflectValueType).Interface().(sql.Scanner); ok {
					isPluck = true
				} else if reflectValueType.Kind() != reflect.Struct || reflectValueType.ConvertibleTo(schema.TimeReflectType) {
					isPluck = true
				}
			}

			for rows.Next() {

				elem := reflect.New(reflectValueType).Elem()
				if isPluck {
					stmp.AddError(rows.Scan(elem.Addr().Interface()))
				} else {
					for idx, field := range fields {
						if field != nil {
							values[idx] = reflect.New(reflect.PtrTo(field.IndirectFieldType)).Interface()
						}
					}

					stmp.AddError(rows.Scan(values...))

					for idx, field := range fields {
						if field != nil {
							field.Set(elem, values[idx])
						}
					}
				}

				if isPtr {
					objV.Set(reflect.Append(objV, elem.Addr()))
				} else {
					objV.Set(reflect.Append(objV, elem))
				}
			}
		case reflect.Struct:
			Schema, _ = schema.Parse(obj, stmp.CacheStore)
			if rows.Next() {
				for idx, column := range columns {
					if field := Schema.LookUpField(changeString(column)); field != nil && field.Readable {
						values[idx] = reflect.New(reflect.PtrTo(field.IndirectFieldType)).Interface()
					} else if field := Schema.LookUpField(column); field != nil && field.Readable {
						values[idx] = reflect.New(reflect.PtrTo(field.IndirectFieldType)).Interface()
					} else {
						values[idx] = &sql.RawBytes{}
					}
				}
				stmp.AddError(rows.Scan(values...))

				for idx, column := range columns {
					if field := Schema.LookUpField(changeString(column)); field != nil && field.Readable {
						field.Set(objV, values[idx])
					} else if field := Schema.LookUpField(column); field != nil && field.Readable {
						field.Set(objV, values[idx])
					}
				}
			}
		}
	}
}

func prepareValues(values []interface{}, columnTypes []*sql.ColumnType, columns []string) {
	if len(columnTypes) > 0 {
		for idx, columnType := range columnTypes {
			if columnType.ScanType() != nil {
				values[idx] = reflect.New(reflect.PtrTo(columnType.ScanType())).Interface()
			} else {
				values[idx] = new(interface{})
			}
		}
	} else {
		for idx := range columns {
			values[idx] = new(interface{})
		}
	}
}

func scanIntoMap(mapValue map[string]interface{}, values []interface{}, columns []string) {
	for idx, column := range columns {
		if reflectValue := reflect.Indirect(reflect.Indirect(reflect.ValueOf(values[idx]))); reflectValue.IsValid() {
			mapValue[column] = reflectValue.Interface()
			if valuer, ok := mapValue[column].(driver.Valuer); ok {
				mapValue[column], _ = valuer.Value()
			} else if b, ok := mapValue[column].(sql.RawBytes); ok {
				mapValue[column] = string(b)
			}
		} else {
			mapValue[column] = nil
		}
	}
}

func changeString(s string) string {
	bs := bytes.Buffer{}
	var b bool
	for i := 0; i < len(s); i++ {
		if s[i] > 'Z' && i == 0 {
			bs.WriteByte(s[i] - 32)
		} else {
			if s[i] < 'A' || s[i] > 'z' || ('Z' < s[i] && s[i] < 'a') {
				b = true
				i++
			}
			if b {
				bs.WriteByte(s[i] - 32)
				b = false
			} else {
				bs.WriteByte(s[i])
			}
		}

	}
	return bs.String()
}
