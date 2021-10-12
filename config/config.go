package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"time"
)

type config struct {
	Mysql    *mysql
	DbSource string
}

type mysql struct {
	Host        string  `ini:"host"`
	Port        int     `ini:"port"`
	User        string  `ini:"user"`
	Password    string  `ini:"password"`
	Database    string  `ini:"database"`
	MaxOpen     int     `ini:"maxOpen"`
	MaxIdle     int     `ini:"maxIdle"`
	MaxLifeTime time.Duration     `ini:"maxLifeTime"`
}

var Config = &config{}

func init() {
	c, err := ini.LoadSources(ini.LoadOptions{
		SkipUnrecognizableLines: true,
		AllowShadows:true,
		AllowBooleanKeys: true,
	}, "./config/config.ini")
	if err != nil {
		panic(err)
	}

	db := new(mysql)

	if err := c.Section("mysql").MapTo(db) ; err == nil {
		Config.Mysql = db
		Config.DbSource = fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8",db.User,db.Password,db.Host,db.Port,db.Database)
	}
}