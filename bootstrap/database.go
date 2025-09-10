package bootstrap

import (
	"errors"
	"fmt"
	"gohub/pkg/config"
	"gohub/pkg/database"
	"gorm.io/gorm/logger"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDB() {

	var dbConfig gorm.Dialector
	switch config.MustGet[string]("database.connection") {
	case "mysql":
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
			config.MustGet[string]("database.mysql.username"),
			config.MustGet[string]("database.mysql.password"),
			config.MustGet[string]("database.mysql.host"),
			config.MustGet[string]("database.mysql.port"),
			config.MustGet[string]("database.mysql.database"),
			config.MustGet[string]("database.mysql.charset"),
		)

		dbConfig = mysql.New(mysql.Config{
			DSN: dsn,
		})
	default:
		panic(errors.New("database connection not supported"))
	}

	// 连接数据库, 并设置 GORM 的日志模式
	database.Connect(dbConfig, logger.Default.LogMode(logger.Info))

	// 设置最大连接数
	if maxOpenConnections, ok := config.Get[int]("database.mysql.max_open_connections"); ok {
		database.SQLDB.SetMaxOpenConns(maxOpenConnections)
	} else {
		panic(errors.New("database max_open_connections not set"))
	}

	// 设置最大空闲连接数
	if maxIdleConnections, ok := config.Get[int]("database.mysql.max_idle_connections"); ok {
		database.SQLDB.SetMaxIdleConns(maxIdleConnections)
	} else {
		panic(errors.New("database max_idle_connections not set"))
	}

	// 设置每个连接的过期时间
	if maxLifeSeconds, ok := config.Get[int]("database.mysql.max_life_seconds"); ok {
		database.SQLDB.SetConnMaxLifetime(time.Duration(maxLifeSeconds) * time.Second)
	} else {
		panic(errors.New("database max_life_seconds not set"))
	}
}
