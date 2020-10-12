package mysql

import (
	"database/sql"
	"fmt"
	"io"
	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	Host     string
	Port     int
	UserName string
	Password string
	TimeZone string
	DbName   string
}

func (p Mysql) Device() string {
	return "mysql"
}

func (p *Mysql) Initialize() (db *sql.DB, err error) {
	return sql.Open(p.Device(), p.Dsn())
}

func (p *Mysql) BindVarTo(w io.Writer, i int) {
	w.Write([]byte{'?'})
}

func NewMysql(host string, port int, userName string, pwd string, db string) *Mysql {
	return &Mysql{
		Host:     host,
		Port:     port,
		UserName: userName,
		Password: pwd,
		DbName:   db,
	}
}

func (p *Mysql) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=%s", p.UserName, p.Password, p.Host, p.Port, p.DbName, p.TimeZone)
}
