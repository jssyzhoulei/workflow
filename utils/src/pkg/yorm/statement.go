package yorm

import (
	"database/sql"
	"fmt"
	"gitee.com/grandeep/org-svc/utils/src/pkg/yorm/parse"
	"gitee.com/grandeep/org-svc/utils/src/pkg/yorm/parse/util"
	"gitee.com/grandeep/org-svc/utils/src/pkg/yorm/schema"
	"math"
	"reflect"
	"strings"
	"sync"
)

type Session int8

const (
	SessionNull Session = iota
	SelectSession
	UpdateSession
	DeleteSession
	InsertSession
)

type Statement struct {
	mp         *util.ModelPathMap
	db         *DB
	query      util.Query
	err        error
	result     sql.Result
	Rows       *sql.Rows
	Session    Session
	Schema     *schema.Schema
	CacheStore *sync.Map
	page       *Page
}

func NewStatement(mp *util.ModelPathMap, db *DB) *Statement {
	return &Statement{
		mp:    mp,
		query: make(util.Query),
		db:    db,
		CacheStore: &sync.Map{},
	}
}

func (s *Statement) AddQuery(key string, obj interface{}) *Statement {
	s.query[key] = reflect.ValueOf(obj)
	return s
}

func (s *Statement) Exec(k string) *Statement {
	var (
		params   []string
		sqlBuild strings.Builder
		complex  interface{}
	)
	md, err := s.mp.GetModelMap(k)
	if err != nil {
		s.err = err
		return s
	}
	params, sqlBuild, complex, err = md.GetInfo()
	version := md.GetVersion()
	if err != nil {
		s.err = err
		return s
	}
	if complex != nil {
		p := parse.GetParser(version, s.db.dialect)
		sqlBuild, qs := p.ParseModelComplex(complex, s.query)
		sein := getSession(sqlBuild.String())
		s.Session = sein
		if sein == SelectSession {
			var (
				sqlB string
			)
			if s.page != nil {
				var (
					limit int
					offset int
				)
				sqlCount := fmt.Sprintf(`SELECT count(*) as count FROM (%s) as page`, sqlBuild.String())
				rs, _ := s.db.conn.Query(sqlCount, qs...)
				if rs.Next() {
					err = rs.Scan(&s.page.Total)
				}
				defer rs.Close()
				limit = s.page.PageSize
				if limit == 0 {
					limit = 10
				}
				if s.page.PageSize == 0 {
					s.page.PageSize = 1
				}
				s.page.TotalPage = math.Ceil(float64(s.page.Total) / float64(s.page.PageSize))
				offset = (s.page.PageSize - 1) * limit
				sqlB = fmt.Sprintf("%s  LIMIT %d OFFSET %d", sqlBuild.String(), limit, offset)
			} else {
				sqlB = sqlBuild.String()
			}
			s.Rows, s.err = s.db.conn.Query(sqlB, qs...)
		} else {
			stm, err := s.db.conn.Prepare(sqlBuild.String())
			if err != nil {
				s.err = err
				return s
			}
			s.result, s.err = stm.Exec(qs...)
		}

		return s
	} else {
		var (
			qs []interface{}
		)
		for _, v := range params {
			vl, err := s.query.GetValue(v)
			if err != nil {
				s.err = err
				return s
			}
			qs = append(qs, vl.Interface())
		}
		sein := getSession(sqlBuild.String())
		s.Session = sein
		if sein == SelectSession {
			var (
				sqlB string
			)
			if s.page != nil {
				var (
					limit int
					offset int
				)
				sqlCount := fmt.Sprintf(`SELECT count(*) as count FROM ( %s ) as page`, sqlBuild.String())
				rs, _ := s.db.conn.Query(sqlCount, qs...)
				if rs.Next() {
					err = rs.Scan(&s.page.Total)
				}
				defer rs.Close()
				limit = s.page.PageSize
				if limit == 0 {
					limit = 10
				}
				if s.page.PageSize == 0 {
					s.page.PageSize = 1
				}
				s.page.TotalPage = math.Ceil(float64(s.page.Total) / float64(s.page.PageSize))
				offset = (s.page.PageSize - 1) * limit
				sqlB = fmt.Sprintf("%s  LIMIT %d OFFSET %d", sqlBuild.String(), limit, offset)
			} else {
				sqlB = sqlBuild.String()
			}
			s.Rows, s.err = s.db.conn.Query(sqlB, qs...)
		} else {
			stm, err := s.db.conn.Prepare(sqlBuild.String())
			if err != nil {
				s.err = err
				return s
			}
			//qs

			s.result, s.err = stm.Exec(qs...)
		}
		return s
	}

}

func getSession(s string) Session {
	s = strings.TrimSpace(s)
	switch strings.ToUpper(s[:6]) {
	case "SELECT":
		return SelectSession
	case "DELETE":
		return DeleteSession
	case "UPDATE":
		return UpdateSession
	case "INSERT":
		return InsertSession
	}
	return 0
}

func (s *Statement) Scan(obj interface{}) error {
	objV := reflect.ValueOf(obj)
	if objV.Type().Kind() == reflect.Ptr {
		objV = objV.Elem()
	}
	switch s.Session {
	case SelectSession:
		Scan(s.Rows, s, obj)
		defer s.Rows.Close()
	case InsertSession:
		id, err := s.result.LastInsertId()
		if err != nil {
			s.err = err
			return err
		}
		objV.SetInt(id)

	case UpdateSession, DeleteSession:
		af, err := s.result.RowsAffected()
		if err != nil {
			s.err = err
			return err
		}
		objV.SetInt(af)
	}
	defer s.Rows.Close()
	return s.err
}

func (s *Statement) Page(pageSize, pageNum int, key string) (Page,error) {
	s.page = &Page{
		PageNum: pageNum,
		PageSize: pageSize,
	}
	var (
		m []map[string]interface{}
	)
	err := s.Exec(key).Scan(&m)
	s.page.Data = m
	return *s.page, err
}


func (s *Statement) LastInsertId() int64 {
	if s.Session == InsertSession && s.result != nil {
		id, err := s.result.LastInsertId()
		if err != nil {
			s.err = err
		}
		return id
	}
	return 0
}

func (s *Statement) RowsAffected() int64 {
	if (s.Session == UpdateSession || s.Session == DeleteSession) && s.result != nil {
		id, err := s.result.RowsAffected()
		if err != nil {
			s.err = err
		}
		return id
	}
	return 0
}

func (s *Statement) AddError(err error) {
	s.err = err
}
