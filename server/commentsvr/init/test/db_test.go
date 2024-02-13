package test

import (
	"commentsvr/config"
	"commentsvr/init/db"
	"testing"
)

func TestMysqlInit(t *testing.T) {
	err := config.Init()
	if err != nil {
		t.Error("config init error")
	}
	db.DBinit()
	mysql := db.GetMySqlDB()
	if mysql == nil {
		t.Error("mysql init error")
	}
}
