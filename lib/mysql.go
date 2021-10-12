package lib

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/liu-jiangyuan/api/config"
	"time"
)

var (
	Engine *xorm.Engine
	err error
)

func init() {
	Engine, err = xorm.NewEngine("mysql", config.Config.DbSource)
	if err != nil {
		panic(err)
	}
	if err = Engine.Ping();err != nil {
		panic(err)
	}
	Engine.ShowSQL(true)
	Engine.ShowExecTime(true)
	Engine.SetMaxIdleConns(config.Config.Mysql.MaxIdle)
	Engine.SetMaxOpenConns(config.Config.Mysql.MaxOpen)
	Engine.SetConnMaxLifetime(config.Config.Mysql.MaxLifeTime * time.Hour)
}