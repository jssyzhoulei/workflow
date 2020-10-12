package engine

import (
	"gitee.com/grandeep/org-svc/utils/src/pkg/config"
	"gitee.com/grandeep/org-svc/utils/src/pkg/log"
	"gitee.com/grandeep/org-svc/utils/src/pkg/yorm"
	"gitee.com/grandeep/org-svc/utils/src/pkg/yorm/mysql"
	"go.uber.org/zap"
)

type Engine struct {
	Config *config.Config
	Logger *zap.Logger
	DB *yorm.DB
}

func NewEngine(path string) *Engine {
	var (
		e Engine
		err error
		mysqlConfig mysql.Mysql
		sqlPath string
	)
	e.Config,err = config.NewConfig(path)
	e.Logger = log.Logger()
	if err != nil {
		panic("engine err")
	}
	err = e.Config.GetBind("mysql", &mysqlConfig)
	if err != nil {
		panic(err)
	}
	e.DB, err = yorm.Open(&mysqlConfig)
	if err != nil {
		panic(err)
	}
	sqlPath, err = e.Config.GetString("mysql.sqlPath")
	err = e.DB.LoadSqlYaml(sqlPath)
	if err != nil {
		panic(err)
	}
	e.Logger.Info("engine success")
	return &e
}
