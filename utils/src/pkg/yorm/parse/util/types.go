package util

import (
	"fmt"
	"io"
	"reflect"
	"strings"
)

type sqlModel struct {
	version string
	params   []string
	modelBuilder strings.Builder
	modelComplex interface{}
}

type modelMap struct {
	model *sqlModel
	child *ModelPathMap
}

func (m modelMap) GetVersion() string {
	if m.model != nil {
		return m.model.version
	}
	return ""
}

type ModelPathMap map[string]modelMap


type Binder interface {
	BindVarTo(w io.Writer, i int)
}

type Query map[string]reflect.Value

func (q Query) GetValue(key string) (reflect.Value, error) {
	var (
		v reflect.Value
		ok bool
	)
	ss := strings.Split(key, ".")
	if len(ss) > 0 && key != ""{
		if v, ok = q[ss[0]]; ok {
			for _, k := range ss[1:] {
				if v.Type().Kind() == reflect.Ptr {
					v = v.Elem()
				}
				v = v.FieldByName(k)
				if !v.IsValid() {
					return reflect.Value{}, fmt.Errorf("unfind value %s", k)
				}
			}
			return v, nil
		} else {
			return reflect.Value{}, fmt.Errorf("unfind query")
		}
	}
	return reflect.Value{}, fmt.Errorf("unfind query")
}