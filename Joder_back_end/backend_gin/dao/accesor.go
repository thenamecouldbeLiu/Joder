package dao

import (
	"fmt"
	"time"

	"joder/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var MysqlDb *gorm.DB

func init() {
	var err error

	account := config.Val.MYSQL_ACCNT
	password := config.Val.MYSQL_PASSWORD
	url := config.Val.MYSQL_URL_IP
	schema := config.Val.MY_SQL_SCHEMA
	charset := config.Val.MY_SQL_CHARSET
	maxLifeTime := config.Val.MY_SQL_MAXLIFETIME
	maxConn := config.Val.MY_SQL_MAXCONN
	maxIdel := config.Val.MYSQL_MAXIDEL
	properties := fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=true&loc=Local", account, password, url, schema, charset)
	MysqlDb, err = gorm.Open(mysql.Open(properties), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		panic(err)
	}

	//設定ConnMaxLifetime/MaxIdleConns/MaxOpenConns
	db, err1 := MysqlDb.DB()

	if err1 != nil {
		fmt.Println("get db failed:", err)
		panic(err)
	}

	db.SetConnMaxLifetime(time.Duration(maxLifeTime) * time.Second)
	db.SetMaxIdleConns(maxIdel)
	db.SetMaxOpenConns(maxConn)
	fmt.Println("DB CONNECTED")
}
