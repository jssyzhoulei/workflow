package util

import (
	"errors"
	"fmt"
	"strings"
)

func (m ModelPathMap) AddModel(key string, model string, params []string, version string) error {
	if _, ok := m[key]; ok {
		return fmt.Errorf("sql func %s is exist", key)
	}
	var (
		builder = strings.Builder{}
	)
	builder.WriteString(model)
	m[key] = modelMap{
		model: &sqlModel{
			params: params,
			modelBuilder: builder,
			version: version,
		},
	}
	return nil
}


func (m ModelPathMap) AddModelComplex(key string, complex interface{}, version string) error {
	m[key] = modelMap{
		model: &sqlModel{
			modelComplex: complex,
			version: version,
		},
	}
	return nil
}
//
//func (m ModelPathMap) AddHandle(key string, t HandleType, handle string, then string) {
//	if mp, ok := m[key]; ok {
//
//		mp.model.handles = append(mp.model.handles, func() string {
//			return ""
//		})
//		m[key] = mp
//	}
//}

func (m ModelPathMap) WriteModel(key string, s string) {
	if _, ok := m[key]; ok {
		m[key].model.modelBuilder.WriteString(s)
	} else {
		m[key] = modelMap{
			model: &sqlModel{
				modelBuilder: strings.Builder{},
			},
		}
		m[key].model.modelBuilder.WriteString(s)
	}
}

func (m *ModelPathMap) GetModel(key string) (string, error) {
	var (
		mp = m
	)
	ps := strings.Split(key, ".")
	if len(ps) > 1 {
		for _, v := range ps[:len(ps)-1] {
			if _,ok := (*mp)[v]; ok {
				mp = (*mp)[v].child
			} else {
				return "", errors.New("not exit")
			}
		}
	}

	if model, ok := (*mp)[ps[len(ps)-1]]; ok {
		return model.model.modelBuilder.String(), nil
	}
	return "", errors.New("not exit")
}

func (m *ModelPathMap) GetModelMap(key string) (mdp modelMap, err error) {
	var (
		mp = m
	)
	ps := strings.Split(key, ".")
	if len(ps) > 1 {
		for _, v := range ps[:len(ps)-1] {
			if _,ok := (*mp)[v]; ok {
				mp = (*mp)[v].child
			} else {
				err = fmt.Errorf("sql func %s is not exist", key)
				return
			}
		}
	}
	return (*mp)[ps[len(ps)-1]],nil
}

func (m modelMap) GetChild() *ModelPathMap {
	return m.child
}

func (m modelMap) GetInfo() ([]string, strings.Builder, interface{}, error) {
	var (
		err error
	)
	if m.model != nil {
		return m.model.params,m.model.modelBuilder, m.model.modelComplex, err
	}
	return nil, strings.Builder{} , nil, fmt.Errorf("sql mode not exist")
}

func (m ModelPathMap) ResetModel(key string, model string, params []string, version string) {
	var (
		builder = strings.Builder{}
	)
	builder.WriteString(model)
	m[key] = modelMap{
		model: &sqlModel{
			params: params,
			modelBuilder: builder,
			version: version,
		},
	}
}



func NewModelMap(m *ModelPathMap) modelMap {
	return modelMap{
		child: m,
	}
}
