package controllers

import (
	"github.com/revel/revel"
	"log"
)

func init(){
	
	revel.OnAppStart(InitDB)	//앱 시작하기 전 DB 초기화
	log.Println("init function in init.go")
	revel.InterceptMethod((*GormController).SetDB, revel.BEFORE)
}