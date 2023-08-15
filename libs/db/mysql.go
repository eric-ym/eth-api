package db

import (
	"database/sql"
	"eth/libs/config"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
	"time"
)

var sqlDB *sql.DB

var DB *gorm.DB

var once sync.Once

func InitDB() {
	once.Do(func() {
		var err error
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.Config().DB.Username,
			config.Config().DB.Password,
			config.Config().DB.Host,
			config.Config().DB.Port,
			config.Config().DB.DBName)
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Error),
		})
		if err != nil {
			zap.S().Panic(err)
		}
		sqlDB, err = DB.DB()
		if err != nil {
			zap.S().Panic(err)
		}

		// SetMaxIdleConn 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(config.Config().DB.MaxIdleConn)

		// SetMaxOpenConn 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(config.Config().DB.MaxIdleConn)

		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		sqlDB.SetConnMaxLifetime(time.Hour)

		err = sqlDB.Ping()
		if err != nil {
			zap.S().Panic(err)
		}
		zap.S().Infof("%v", sqlDB.Stats())

	})
}

func GetDB() *gorm.DB {
	return DB
}
