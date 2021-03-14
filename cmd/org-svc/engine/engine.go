package engine
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
