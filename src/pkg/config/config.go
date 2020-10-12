package config

import (
	"errors"
	"fmt"
	"github.com/timchine/pole/pkg/log"
	"gopkg.in/yaml.v2"
	"os"
	"reflect"
	"strings"
)

var (
	configNilError = errors.New("配置文件数据为空")
	configPrtError = errors.New("传入值为指针类型")
	configNotExistError = errors.New("未找到配置")
	configTypeUnMatch = errors.New("类型不匹配")
)

type Config map[interface{}]interface{}

func NewConfig(path string) (*Config, error) {
	var (
		err     error
		file    *os.File
		decoder *yaml.Decoder
		conf    Config
	)
	conf = make(Config)
	file, err = os.Open(path)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	decoder = yaml.NewDecoder(file)
	err = decoder.Decode(&conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}

func (c *Config) GetBind(key string, conf interface{}) error {
	var (
		t reflect.Type
		cv reflect.Value
		ks []string
		con Config
	)
	if c == nil {
		return configNilError
	}
	t = reflect.TypeOf(conf)
	cv = reflect.ValueOf(conf)
	if t.Kind() != reflect.Ptr {
		return configPrtError
	}
	t = t.Elem()
	cv = cv.Elem()
	con = *c
	ks = strings.Split(key, ".")
	for i, k := range ks {
		k = firstToLow(k)
		if v, ok := con[k]; ok {
			switch reflect.TypeOf(v).Kind() {
			case reflect.Map:
				vv := reflect.ValueOf(v)
				mr := vv.MapRange()
				if i == len(ks)-1 {
					if t.Kind() == reflect.Map {
						for mr.Next() {
							cv.SetMapIndex(mr.Key(), mr.Value())
						}
					} else if t.Kind() == reflect.Struct {
						bindToStruct(cv, mr, t)
					} else {
						return configTypeUnMatch
					}
				} else {
					con = v.(Config)
				}
				break
			default:
				if i == len(ks)-1 {
					bindValue(cv, v)
				}
			}
		} else {
			return configNotExistError
		}
	}
	return nil
}

func bindValue(cv reflect.Value, in interface{}) {
	switch cv.Type().Kind() {
	case reflect.String:
		cv.SetString(getString(in))
		break
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		cv.SetInt(getInt(in))
		break
	case reflect.Float32, reflect.Float64:
		cv.SetFloat(getFloat(in))
		break
	case reflect.Bool:
		if s := fmt.Sprintf("%v", in);s == "true" || s == "1" {
			cv.SetBool(true)
		} else {
			cv.SetBool(false)
		}
		break
	case reflect.Slice:
		if v, ok := in.([]interface{}); ok {
			cv.Set(reflect.ValueOf(v))
		}
		break
	}
}

//如果出现字段不匹配只会报警不会反回错误
func bindToStruct(cv reflect.Value, mr *reflect.MapIter, t reflect.Type) {
	for mr.Next() {
		fu := firstToUp(fmt.Sprintf("%v", mr.Key()))
		f, b := t.FieldByName(fu)
		if b {
			switch f.Type.Kind() {
			case reflect.Struct:
				if mrv, ok := mr.Value().Interface().(Config); ok {
					mv := reflect.ValueOf(mrv)
					bindToStruct(cv.FieldByName(fu), mv.MapRange(), cv.FieldByName(fu).Type())
				}
			default:
				bindValue(cv.FieldByName(fu), mr.Value().Interface())
			}
		}
	}
}

//
func getString(in interface{}) string {
	value := reflect.ValueOf(in)
	switch value.Type().Kind() {
	case reflect.String, reflect.Int,reflect.Float64, reflect.Interface:
		return fmt.Sprintf("%v", value)
	case reflect.Bool:
		if value.Bool() {
			return "true"
		} else {
			return "false"
		}
	default:
		log.Logger().Warn("未知类型")
	}
	return ""
}


func getInt(in interface{}) int64 {
	value := reflect.ValueOf(in)
	switch value.Type().Kind() {
	case reflect.Int:
		return value.Int()
		break
	default:
		log.Logger().Warn("类型映射错误")
	}
	return 0
}

func getFloat(in interface{}) float64 {
	value := reflect.ValueOf(in)
	switch value.Type().Kind() {
	case reflect.Float64:
		return value.Float()
		break
	default:
		log.Logger().Warn("未知类型")
	}
	return 0
}

//将字符串首字母无论大小写转小写
func firstToLow(s string) string {
	if len(s) > 0 {
		if 'A' <= s[0] && s[0] <= 'Z' {
			b := make([]byte, len(s))
			b[0] = s[0] + 32
			copy(b[1:], s[1:])
			return string(b)
		} else {
			return s
		}
	}
	return ""
}

func firstToUp(s string) string {
	if len(s) > 0 {
		if 'a' <= s[0] && s[0] <= 'z' {
			b := make([]byte, len(s))
			b[0] = s[0] - 32
			copy(b[1:], s[1:])
			return string(b)
		} else {
			return s
		}
	}
	return ""
}

func (c *Config) GetString(key string) (string, error) {
	var (
		s string
	)
	err := c.GetBind(key, &s)
	return s, err
}

func (c *Config) GetInt(key string)(int, error) {
	var (
		i int
	)
	err := c.GetBind(key, &i)
	return i, err
}

func (c *Config) GetFloat(key string) (float64, error) {
	var (
		f float64
	)
	err := c.GetBind(key, &f)
	return f, err
}

func (c *Config) GetSlice(key string) ([]interface{}, error) {
	var (
		sl []interface{}
	)
	err := c.GetBind(key, &sl)
	return sl, err
}