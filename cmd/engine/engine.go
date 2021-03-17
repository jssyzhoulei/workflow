package engine

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

//
//import (
//	"github.com/jssyzhoulei/workflow/logger"
//	"github.com/jssyzhoulei/workflow/utils/src/pkg/config"
//	"github.com/jssyzhoulei/workflow/utils/src/pkg/yorm"
//	"github.com/jssyzhoulei/workflow/utils/src/pkg/yorm/mysql"
//	"go.uber.org/zap"
//)
//
//type Engine struct {
//	Config *config.Config
//	Logger *zap.Logger
//	DB     *yorm.DB
//}
//
//func NewEngine(path string) *Engine {
//	var (
//		e           Engine
//		err         error
//		mysqlConfig mysql.Mysql
//		//sqlPath string
//	)
//	e.Config, err = config.NewConfig(path)
//	e.Logger = log.Logger
//	if err != nil {
//		panic("engine err")
//	}
//	err = e.Config.GetBind("mysql", &mysqlConfig)
//	if err != nil {
//		panic(err)
//	}
//	e.DB, err = yorm.Open(&mysqlConfig)
//	if err != nil {
//		panic(err)
//	}
//	//e.DB.AutoMigrate(&models.User{}, &models.Group{}, &models.Menu{}, &models.Permission{}, &models.Quota{}, &models.Role{}, &models.RoleMenuPermission{}, &models.UserRole{})
//	//sqlPath, err = e.Config.GetString("mysql.sqlPath")
//	//err = e.DB.LoadSqlYaml(sqlPath)
//	//if err != nil {
//	//	panic(err)
//	//}
//	e.Logger.Info("engine success")
//	return &e
//}

func connect() (*gorm.DB, error) {
	dsn := "root:123456@tcp(localhost:3306)/workflow?charset=utf8&timeout=4s&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	//db, err := gorm.Open("mysql", dataSourceName)
	//if err != nil {
	//	return nil, err
	//}
	//
	//db.SetLogger(logrus.StandardLogger())
	//
	//err = db.DB().Ping()
	//if err != nil {
	//	db.Close()
	//	return nil, err
	//}
	//db.LogMode(viper.GetBool("mysql.grus.dbLogMode"))
	//logrus.Debugf("database connected [%s](%s:%d)", cfg.database, cfg.host, cfg.port)
	//return db, nil
}

func InitDB() {
	db, _ := connect()
	//_ = db.AutoMigrate(&models.User{}, &models.Group{}, &models.Menu{}, &models.Permission{}, &models.Quota{}, &models.Role{}, &models.RoleMenuPermission{}, &models.UserRole{})
	DB = db
}
