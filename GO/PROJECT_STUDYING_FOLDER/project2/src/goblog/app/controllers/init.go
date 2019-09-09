package controllers

import (
	"github.com/revel/revel"	
)

func init(){
	
	revel.OnAppStart(InitDB)	//앱 시작하기 전 DB 초기화	
	revel.InterceptMethod((*GormController).SetDB, revel.BEFORE)
	revel.InterceptMethod((*GormController).Begin, revel.BEFORE)
	revel.InterceptMethod((*GormController).Commit, revel.AFTER)
	revel.InterceptMethod((*GormController).Rollback, revel.FINALLY)
	revel.InterceptMethod((*App).SetCurrentUser, revel.BEFORE)
	revel.InterceptMethod(Post.CheckUser, revel.BEFORE)
	revel.InterceptMethod(Comment.CheckUser, revel.BEFORE)

}