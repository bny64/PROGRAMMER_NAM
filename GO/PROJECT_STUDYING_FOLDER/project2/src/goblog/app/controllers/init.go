package controllers

import (
	"github.com/revel/revel"	
)

func init(){
	
	revel.OnAppStart(InitDB)	//앱 시작하기 전 DB 초기화	
	revel.InterceptMethod((*GormController).SetDB, revel.BEFORE)
}