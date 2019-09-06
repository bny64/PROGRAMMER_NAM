package controllers

import (
	"github.com/revel/revel"	
	"github.com/jinzhu/gorm"	
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"strings"
	_ "log"
	"goblog/app/models"	
	_ "strings"
	"golang.org/x/crypto/bcrypt"
)

type GormController struct {
	*revel.Controller
	Txn *gorm.DB
}

var Gdb *gorm.DB

const (
	DefaultName,
	DefaultRole,
	DefaultUsername,
	DefaultPassword = "Admin","admin","admin","admin"
)

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
	initTable()
	return Gdb
}

func initTable(){
	
	Gdb.AutoMigrate(&models.Post{}, &models.Comment{}, &models.User{})
	//패키지에서 함수, 구조체 단순히 문자열로 받을 수 있는지?
	//ex) models["func or struct"]

	bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(DefaultPassword), bcrypt.DefaultCost)

	Gdb.Where(models.User{Name:DefaultName, Role:DefaultRole, Username:DefaultUsername}).Attrs(models.User{Password : string(bcryptPassword)}).FirstOrCreate(&models.User{})
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

func(c *GormController) Begin() revel.Result {
	c.Txn = Gdb.Begin()
	return nil
}

func (c *GormController) Rollback() revel.Result {
	if c.Txn != nil {
		c.Txn.Rollback()
		c.Txn = nil
	}
	return nil
}

func (c *GormController) Commit() revel.Result {
	if c.Txn != nil {
		c.Txn.Commit()
		c.Txn = nil
	}
	return nil
}