package schema

import (
	"errors"
	"fmt"
	"go/ast"
	"reflect"
	"sync"
)

// ErrUnsupportedDataType unsupported data type
var ErrUnsupportedDataType = errors.New("unsupported data type")

type Schema struct {
	Name                      string
	ModelType                 reflect.Type
	PrioritizedPrimaryField   *Field
	DBNames                   []string
	PrimaryFields             []*Field
	PrimaryFieldDBNames       []string
	Fields                    []*Field
	FieldsByName              map[string]*Field
	FieldsWithDefaultDBValue  []*Field // fields with default value assigned by database
	BeforeCreate, AfterCreate bool
	BeforeUpdate, AfterUpdate bool
	BeforeDelete, AfterDelete bool
	BeforeSave, AfterSave     bool
	AfterFind                 bool
	err                       error
	cacheStore                *sync.Map
	TagMap map[string]string
}

func (schema Schema) String() string {
	return fmt.Sprintf("%v.%v", schema.ModelType.PkgPath(), schema.ModelType.Name())
}

func (schema Schema) MakeSlice() reflect.Value {
	slice := reflect.MakeSlice(reflect.SliceOf(reflect.PtrTo(schema.ModelType)), 0, 0)
	results := reflect.New(slice.Type())
	results.Elem().Set(slice)
	return results
}

func (schema Schema) LookUpField(name string) *Field {
	if field, ok := schema.FieldsByName[name]; ok {
		return field
	}
	return nil
}

type Tabler interface {
	TableName() string
}

// get data type from dialector
func Parse(dest interface{}, cacheStore *sync.Map) (*Schema, error) {
	if dest == nil {
		return nil, fmt.Errorf("%w: %+v", ErrUnsupportedDataType, dest)
	}

	modelType := reflect.ValueOf(dest).Type()
	for modelType.Kind() == reflect.Slice || modelType.Kind() == reflect.Array || modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}

	if modelType.Kind() != reflect.Struct {
		if modelType.PkgPath() == "" {
			return nil, fmt.Errorf("%w: %+v", ErrUnsupportedDataType, dest)
		}
		return nil, fmt.Errorf("%w: %v.%v", ErrUnsupportedDataType, modelType.PkgPath(), modelType.Name())
	}

	if v, ok := cacheStore.Load(modelType); ok {
		return v.(*Schema), nil
	}

	schema := &Schema{
		Name:           modelType.Name(),
		ModelType:      modelType,
		FieldsByName:   map[string]*Field{},
		cacheStore:     cacheStore,
	}

	defer func() {
		if schema.err != nil {
			cacheStore.Delete(modelType)
		}
	}()

	for i := 0; i < modelType.NumField(); i++ {
		if fieldStruct := modelType.Field(i); ast.IsExported(fieldStruct.Name) {
			if field := schema.ParseField(fieldStruct); field.EmbeddedSchema != nil {
				schema.Fields = append(schema.Fields, field.EmbeddedSchema.Fields...)
			} else {
				schema.Fields = append(schema.Fields, field)
			}
		}
	}
	for _, field := range schema.Fields {
		if _, ok := schema.FieldsByName[field.Name]; !ok {
			schema.FieldsByName[field.Tag.Get("yorm")] = field
			schema.FieldsByName[field.Name] = field
		}

		field.setupValuerAndSetter()
	}
	return schema, schema.err
}
