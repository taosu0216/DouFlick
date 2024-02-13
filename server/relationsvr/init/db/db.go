package db

import (
	"fmt"
	"relationsvr/config"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	mysqldb *gorm.DB
	dbOnce  sync.Once
)

func DBinit() {
	dbOnce.Do(initDB)
}

func initDB() {
	dbConfig := config.GetGlobalConfig().MysqlConfig
	connArgs := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbConfig.UserName,
		dbConfig.PassWord, dbConfig.Host, dbConfig.Port, dbConfig.Database)
	//log.Info("mdb addr:" + connArgs)
	var err error
	mysqldb, err = gorm.Open(mysql.Open(connArgs), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}
	sqlDB, err := mysqldb.DB()
	if err != nil {
		fmt.Println("failed to make sqlDB, err:" + err.Error())
		panic(fmt.Errorf("failed to make sqlDB, err:%s", err.Error()))
	}
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConn) // 设置最大空闲连接
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConn) // 设置最大打开的连接
}
func GetMySqlDB() *gorm.DB {
	return mysqldb
}
func CloseDB() {
	if mysqldb != nil {
		sqlDB, err := mysqldb.DB()
		if err != nil {
			panic("fetch db connection err:" + err.Error())
		}
		_ = sqlDB.Close()
	}
}
