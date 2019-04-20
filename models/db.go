package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const DB_HOST = "127.0.0.1"
const DB_PORT = 3306
const DB_NAME = "test"
const DB_USER = "root"
const DB_PWD = "123456"

var (
	db orm.Ormer
)

func GetConnection(model interface{}, dbAlias string) orm.Ormer {
	orm.Debug = true // 是否开启调试模式 调试模式下会打印出sql语句
	var dbDsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", DB_USER, DB_PWD, DB_HOST, DB_PORT, DB_NAME)
	orm.RegisterDataBase(dbAlias, "mysql", dbDsn, 30)
	orm.RegisterModel(model)
	db = orm.NewOrm()
	return db
}
