package yorm

import (
	"database/sql"
	"gitee.com/grandeep/org-svc/utils/src/pkg/yorm/parse/util"
	"gorm.io/gorm"
	"strings"
)

type Dialect interface {
	Device() string
	Initialize() (db *gorm.DB, err error)
	util.Binder
}

type DB struct {
	mp *util.ModelPathMap
	*gorm.DB
	conn *sql.DB
	dialect Dialect
}

type DBRepo struct {
	db *DB
	mp *util.ModelPathMap
}

type Page struct {
	Total int
	TotalPage float64
	PageSize int
	PageNum int
	Data []map[string]interface{}
}

func (db *DB) GetRepo(key string) *DBRepo {
	var (
		mp = db.mp
	)
	ps := strings.Split(key, ".")
	for _, v := range ps {
		mp = (*mp)[v].GetChild()
	}
	return &DBRepo{
		db: db,
		mp: mp,
	}
}

func (db *DBRepo) AddQuery(key string, obj interface{}) *Statement {
	var (
		stm = NewStatement(db.mp, db.db)
	)
	return stm.AddQuery(key, obj)
}

func (db *DB) AddQuery(key string, obj interface{}) *Statement {
	var (
		stm = NewStatement(db.mp, db)
	)
	return stm.AddQuery(key, obj)
}

func (db *DB) Exec(key string) *Statement {
	var (
		stm = NewStatement(db.mp, db)
	)
	return stm.Exec(key)
}



func (db *DB) makeModelPathMap(prefix string, m *util.ModelPathMap) {
	var (
		mp  = db.mp
	)
	ps := strings.Split(prefix, ".")
	if prefix != "" && len(ps) > 0 {
		for k,v := range ps {
			if _, ok := (*mp)[v]; ok {
				mp = (*mp)[v].GetChild()
			} else {
				if k + 1 == len(ps) {
					(*mp)[v] = util.NewModelMap(m)
				} else {
					(*mp)[v] = util.NewModelMap(&util.ModelPathMap{})
					mp = (*mp)[v].GetChild()
				}

			}
		}
	}
}



func Open(dialect Dialect) (*DB, error) {
	var (
		db = &DB{}
		err error
	)
	db.dialect = dialect
	db.mp = &util.ModelPathMap{}
	db.DB, err = dialect.Initialize()
	db.conn, err = db.DB.DB()
	if err != nil {
		return nil, err
	}
	return db, nil
}