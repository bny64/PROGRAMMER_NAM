package controllers

import (
	"github.com/revel/revel"	
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"strings"
	_ "log"
)

type GormController struct {
	*revel.Controller
	Txn *gorm.DB
}

var Gdb *gorm.DB

func InitDB(){	
	_ = GetDBConn()
}

func GetDBConn() *gorm.DB {	
	var err error

	connectingString := getConnectingString()	
	
	Gdb, err = gorm.Open("mysql", connectingString)
		
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	Gdb.DB()
	Gdb.LogMode(true)

	return Gdb
}

func getParamString(param string, defaultValue string) string {	
	p := revel.Config.StringDefault(param, "")	
	if p=="" {
		if defaultValue == "" {
			fmt.Println("could not : ", param)
		} else {
			return defaultValue
		}
	}
	return p
}

func getConnectingString() string {
	
	host := getParamString("db.host", "")
	port := getParamString("db.port", "3306")
	user := getParamString("db.user", "")
	pass := getParamString("db.password", "")
	dbname := getParamString("db.name", "bny_mysql")
	protocol := getParamString("db.protocol", "tcp")
	dbargs := getParamString("dbargs", " ")

	if strings.Trim(dbargs, " ") != "" {
		dbargs = "?" + dbargs
	}else {
		dbargs = ""
	}
	
	return fmt.Sprintf("%s:%s@%s(%s:%s)/%s%s?charset=utf8&parseTime=True", user, pass, protocol, host, port, dbname, dbargs)
	//parseTime 넣어줘야 go의 time.Time을 mysql에서 dateTime 자동변환이 가능해짐.
}

func (c *GormController) SetDB() revel.Result {	
	c.Txn = Gdb
	return nil
}