package initialize

import (
	"fmt"
	"log"
	"time"

	"github.com/paiz4ever/go-web-framework/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDataBase() error {
	switch global.CONFIG.Server.DBType {
	case "mysql":
		if err := initMysql(); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported database type: %s", global.CONFIG.Server.DBType)
	}
	if err := global.DB.AutoMigrate(); err != nil {
		return fmt.Errorf("failed to migrate database: %v", err)
	}
	return nil
}

func initMysql() error {
	m := global.CONFIG.Mysql
	if db, err := gorm.Open(mysql.Open(m.DSN()), &gorm.Config{
		Logger: logger.New(
			log.New(log.Writer(), "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Error,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}); err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		global.DB = db
	}
	return nil
}
