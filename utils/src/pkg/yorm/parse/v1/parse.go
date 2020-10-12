package parse_v1

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/timchine/pole/pkg/yorm/parse/util"
	"reflect"
	"strconv"
	"strings"
)

type ParseV1 struct {
	modelPathMap util.ModelPathMap
	util.Binder
	index int
}

func NewParseV1(b util.Binder) *ParseV1 {
	return &ParseV1{
		modelPathMap: make(util.ModelPathMap),
		Binder:       b,
	}
}

func (p *ParseV1) Parse(mp map[string]interface{}) (*util.ModelPathMap, error) {
	var (
		err error
	)
	for key, model := range mp {
		if sli, ok := model.([]interface{}); ok {
			err = p.makeComplex(key, sli)
			if err != nil {
				return nil, err
			}
		} else {
			sql, params := p.makeModelAndParams(model)
			if sql != "" {
				err = p.modelPathMap.AddModel(key, sql, params, p.Version())
				if err != nil {
					return nil, err
				}
			}
		}
	}
	return &p.modelPathMap, nil
}

func (p ParseV1) Version() string {
	return "v1.0"
}

func (p *ParseV1) ParseModelComplex(c interface{}, query util.Query) (*strings.Builder, []interface{}) {
	var (
		sqlBuild = &strings.Builder{}
		vs       []interface{}
	)
	if csl, ok := c.([]interface{}); ok {
		for _, model := range csl {
			if sql, ok := model.(string); ok {
				s, ps := p.makeModelAndParams(sql)
				for _, po := range ps {
					v, err := query.GetValue(po)
					if err != nil {
						return nil, nil
					}
					vs = append(vs, v.Interface())
				}
				p.index += len(ps)
				sqlBuild.WriteString(s)
				sqlBuild.WriteString(" ")
			} else if sql, ok := model.(map[interface{}]interface{}); ok {
				var (
					b bool
				)
				if con, ok := sql["if"]; ok {
					b = evalIf(con, query)
					if b {
						if then, ok := sql["then"]; ok {
							s, ps := p.makeModelAndParams(then)
							p.index += len(ps)
							for _, po := range ps {
								v, err := query.GetValue(po)
								if err != nil {
									return nil, nil
								}
								vs = append(vs, v.Interface())
							}
							sqlBuild.WriteString(s)
							sqlBuild.WriteString(" ")
						}

					}
				}
				if con, ok := sql["for"]; ok {
					var (
						sep  string
						then string
					)
					if s, ok := sql["sep"]; ok {
						sep, _ = s.(string)
					}
					if s, ok := sql["then"]; ok {
						then,_ = s.(string)
					}
					builder, pas, err := evalFor(con, p.index, sep, then, query, p.Binder)
					if err != nil {
						return nil, nil
					}
					sqlBuild.WriteString(" ")
					sqlBuild.WriteString(builder.String())
					sqlBuild.WriteString(" ")
					vs = append(vs, pas...)
					p.index += len(pas)

				}
			}
		}
	}
	return sqlBuild, vs
}

func evalFor(con interface{}, index int, sep string, then string, query util.Query, bind util.Binder) (*strings.Builder, []interface{}, error) {
	if c, ok := con.(string); ok {
		var (
			keyName  string
			vName    string
			sqlBuild = &strings.Builder{}
			params   []interface{}
		)
		c = strings.Replace(c, " ", "", -1)
		cs := strings.Split(c, "=>")
		if len(cs) == 2 {
			css := strings.Split(cs[0], ",")
			if len(css) == 1 {
				vName = css[0]
			} else {
				vName = css[1]
				keyName = css[0]
			}
			qo, err := query.GetValue(cs[1])
			if err != nil {
				return sqlBuild, nil, err
			}
			switch qo.Type().Kind() {
			case reflect.Map:
				iter := qo.MapRange()
				i := 0
				for iter.Next() {
					i++
					var (
						t bool
						b int
					)
					for j := 0; j < len(then); j++ {
						switch then[j] {
						case '$':
							t = true
							sqlBuild.WriteString(then[b:j])
							b = j + 1
							break
						case '{':
							b = j + 1
							break
						case '}':
							if t {
								key := then[b:j]
								if key == keyName {
									params = append(params, iter.Key().Interface())
								} else {
									keys := strings.Split(key, ".")
									if len(keys) > 1 {
										if keys[0] == vName {
											v := iter.Value()
											for _, k := range keys[1:] {
												if v.Type().Kind() == reflect.Ptr {
													v = v.Elem()
												}
												v = v.FieldByName(k)
												if !v.IsValid() {
													return nil, nil, fmt.Errorf("unfind value %s", k)
												}
											}
											params = append(params, v.Interface())
										} else {
											return sqlBuild, nil, fmt.Errorf("unknown key %s", keys[0])
										}
									}
								}
								index += 1
								bind.BindVarTo(sqlBuild, index)
							}
							b = j + 1
							break
						}
					}
					if b < len(then) {
						sqlBuild.WriteString(then[b:])
					}
					if i < qo.Len() {
						sqlBuild.WriteString(sep)
					}
				}
			case reflect.Slice:
				for i := 0; i < qo.Len(); i++ {
					var (
						t bool
						b int
					)
					for j := 0; j < len(then); j++ {
						switch then[j] {
						case '$':
							t = true
							sqlBuild.WriteString(then[b:j])
							b = j + 1
							break
						case '{':
							b = j + 1
							break
						case '}':
							if t {
								key := then[b:j]
								if key == keyName {
									params = append(params, reflect.ValueOf(i).Interface())
								} else {
									keys := strings.Split(key, ".")
									if len(keys) > 1 {
										if keys[0] == vName {
											v := qo.Index(i)
											for _, k := range keys[1:] {
												if v.Type().Kind() == reflect.Ptr {
													v = v.Elem()
												}
												v = v.FieldByName(k)
												if !v.IsValid() {
													return nil, nil, fmt.Errorf("unfind value %s", k)
												}
											}
											params = append(params, v.Interface())
										} else {
											return sqlBuild, nil, fmt.Errorf("unknown key %s", keys[0])
										}
									}
								}
								index += 1
								bind.BindVarTo(sqlBuild, index)
							}
							b = j + 1
							break
						}
					}
					if b < len(then) {
						sqlBuild.WriteString(then[b:])
					}
					if i < qo.Len()-1 {
						sqlBuild.WriteString(sep)
					}

				}
			}
		}
		return sqlBuild, params, nil
	}
	return nil, nil, errors.New("syntax error")
}

//只能判断简单逻辑
func evalIf(con interface{}, query util.Query) bool {
	if c, ok := con.(string); ok {
		c = strings.Replace(c, " ", "", -1)
		l := 0

		var (
			isPa bool
			isRe bool
			re   []byte
			pa   string
			co   []byte
			b    bool
			bb   bool
		)

		for i := 0; i < len(c); i++ {
			switch c[i] {
			case '>', '<', '!', '=':
				if !isPa {
					pa = string(re)
					re = nil
					isPa = true
					co = append(co, c[i])
				} else {
					co = append(co, c[i])
				}
				if !isRe {
					l = i
				}
				break
			case '|':
				if isRe {
					continue
				}
				bb = b
				isPa = false
				if re != nil {
					b = compare(pa, re, co, query)
					re = nil
				}
				b = b || bb
				if b {
					return b
				}
				l = i + 1
				i = i + 1
				break
			case '&':
				if isRe {
					continue
				}
				bb = b
				isPa = false
				if re != nil {
					b = compare(pa, re, co, query)
					re = nil
				}
				b = b && bb
				if !b {
					return b
				}
				l = i + 1
				i = i + 1
				break
			case '(':
				l = i + 1
				isPa = false
				isRe = true
				break
			case ')':
				if !evalIf(c[l:i], query) {
					b = false
				} else {
					b = true
				}
				l = i
				isPa = false
				re = nil
				isRe = false
				break
			default:
				re = append(re, c[i])
			}
		}
		if re != nil {
			b = compare(pa, re, co, query)
		}
		return b
	}
	return false
}

func compare(pa string, re []byte, co []byte, query util.Query) bool {
	var (
		b bool
	)
	if co == nil {
		q, err := query.GetValue(string(re))
		if err != nil {
			return false
		}
		if q.Type().Kind() == reflect.Bool {
			return q.Bool()
		} else {
			return false
		}
	}
	q, err := query.GetValue(pa)
	if err != nil {
		return false
	}
	for _, v := range co {
		switch v {
		case '!':
			if fmt.Sprintf("%v", q) != string(re) {
				return true
			} else {
				return false
			}
		case '=':
			if fmt.Sprintf("%v", q) == string(re) {
				return true
			}
		case '>':
			qf, err := strconv.ParseFloat(fmt.Sprintf("%v", q), 64)
			if err != nil {
				return false
			}
			rf, err := strconv.ParseFloat(string(re), 64)
			if err != nil {
				return false
			}
			if qf < rf {
				b = false
			}
		case '<':
			qf, err := strconv.ParseFloat(fmt.Sprintf("%v", q), 64)
			if err != nil {
				return false
			}
			rf, err := strconv.ParseFloat(string(re), 64)
			if err != nil {
				return false
			}
			if qf > rf {
				b = false
			}
		}
	}

	return b
}

func (p *ParseV1) makeModelAndParams(h interface{}) (string, []string) {
	s, ok := h.(string)
	if !ok {
		return "", nil
	}
	if p.Binder == nil {
		return "", nil
	}
	var (
		l     = len(s)
		t     bool
		b     int
		model = &bytes.Buffer{}
		ps    []string
	)
	for i := 0; i < l; i++ {
		switch s[i] {
		case '$':
			t = true
			model.WriteString(s[b:i])
			b = i + 1
			break
		case '{':
			b = i + 1
			break
		case '}':
			if t {
				ps = append(ps, s[b:i])
				p.Binder.BindVarTo(model, len(ps)+p.index)
			}
			b = i + 1
			break
		}
	}
	if b < l {
		model.WriteString(s[b:l])
	}
	return model.String(), ps
}

func (p *ParseV1) makeComplex(key string, arr []interface{}) error {
	var (
		handleCase bool
	)
	for _, v := range arr {
		if _, ok := v.(map[interface{}]interface{}); ok {
			return p.modelPathMap.AddModelComplex(key, arr, p.Version())
		} else {
			vs, ok := v.(string)
			if ok {
				p.modelPathMap.WriteModel(key, vs)
				p.modelPathMap.WriteModel(key, " ")
			}
		}
	}
	if !handleCase {
		sql, err := p.modelPathMap.GetModel(key)

		if err == nil {
			sql, params := p.makeModelAndParams(sql)
			p.modelPathMap.ResetModel(key, sql, params, p.Version())
		}
	}
	return nil
}
