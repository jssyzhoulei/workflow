package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	mysql2 "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
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

func (p *Mysql) Initialize() (db *gorm.DB, err error) {
	return gorm.Open(mysql2.Open(p.Dsn()), &gorm.Config{})
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
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=%s&parseTime=true", p.UserName, p.Password, p.Host, p.Port, p.DbName, p.TimeZone)
}
