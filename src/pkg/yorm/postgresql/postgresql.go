package postgresql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"io"
)

type PostgreSql struct {
	Host     string
	Port     int
	UserName string
	Password string
	TimeZone string
	DbName   string
}

func (p PostgreSql) Device() string {
	return "postgres"
}

func (p *PostgreSql) Initialize() (db *sql.DB, err error) {
	return sql.Open(p.Device(), p.Dsn())
}

func (p *PostgreSql) BindVarTo(w io.Writer, i int) {
	w.Write([]byte{'$'})
	w.Write([]byte(fmt.Sprintf("%d", i)))
}


func NewPostgreSql(host string, port int, userName string, pwd string, db string) *PostgreSql {
	return &PostgreSql{
		Host:     host,
		Port:     port,
		UserName: userName,
		Password: pwd,
		DbName:   db,
	}
}

func (p *PostgreSql) Dsn() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s", p.UserName, p.Password, p.DbName, p.Port, p.TimeZone)
}
