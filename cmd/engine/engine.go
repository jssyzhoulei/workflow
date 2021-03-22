package engine

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func connect() (*gorm.DB, error) {
	dsn := "root:123456@tcp(localhost:3306)/workflow?charset=utf8mb4&timeout=4s&parseTime=True&loc=Local"

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

}

func InitDB() *gorm.DB {
	db, _ := connect()
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	//_ = db.AutoMigrate(&models.User{}, &models.Group{}, &models.Menu{},
	//				   &models.Permission{}, &models.Role{}, &models.RoleMenuPermission{},
	//				   &models.UserRole{}, &models.WorkFLow{}, &models.WorkNode{},
	//)
	return db
}
