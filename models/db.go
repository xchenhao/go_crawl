package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"fmt"
)

const DB_HOST = "127.0.0.1"
const DB_PORT = 3306
const DB_NAME = "test"
const DB_USER = "root"
const DB_PWD = "123456"

var (
	db orm.Ormer
)

func GetConnection(model interface{}) orm.Ormer {
	orm.Debug = true // 是否开启调试模式 调试模式下会打印出sql语句
	var dbDsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", DB_USER, DB_PWD, DB_HOST, DB_PORT, DB_NAME)
	orm.RegisterDataBase("default", "mysql", dbDsn, 30)
	orm.RegisterModel(model)
	db = orm.NewOrm()
	return db
}