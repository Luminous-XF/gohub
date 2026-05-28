package database

import (
	"database/sql"
	"fmt"

	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// DB 对象
var DB *gorm.DB
var SQLDB *sql.DB

// Connect 连接数据库
func Connect(dbConfig gorm.Dialector, _logger gormlogger.Interface) {
	// 使用 gorm.Open 连接数据库
	var err error
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
	})

	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	// 获取底层的 sqlDB
	SQLDB, err = DB.DB()
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
}
