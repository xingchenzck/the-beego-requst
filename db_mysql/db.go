package controller

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"frist/models"
	"github.com/beego-develop"
)


var Db *sql.DB
//在初始化
func init() {
	fmt.Println("连接MySQL数据库")
	config := beego.AppConfig
	dbDriver := config.String("db_driverName")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp :=config.String("db_ip")
	dbName :=config.String("db_name")

	connUrl :=dbUser + ":" +dbPassword + "@tcp("+dbIp+ ")/" +dbName+"?charsent=utf-8"
	db,err :=sql.Open(dbDriver,connUrl)
	if err !=nil {
		panic("数据错误")
	}
	Db = db
}
func InserUser(user models.User)(int,error) {
	hashMd5 :=md5.New()
	hashMd5.Write([]byte(user.Possword))
	bytes :=hashMd5.Sum(nil)


	Db.Exec("insert into user (nick,password) values (?,?)"),user.Nick,user.Possword
	if err != nil{
		return -1,err
	}

}
func Queryuser()  {
	Db.QueryRow("")
}
