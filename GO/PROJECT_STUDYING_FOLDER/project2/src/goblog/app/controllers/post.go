package controllers

import (
	"github.com/revel/revel"
	"goblog/app/models"
	"log"
	"time"	
)

type Post struct {
	*revel.Controller //*revel.Controller를 임베디드 필드로 지정
	GormController
}

func (c Post) Index() revel.Result {		
	log.Println("PostController - Index()")
	var posts []models.Post
	c.Txn.Find(&posts) //전체 조회
	return c.Render(posts)
}

func (c Post) New() revel.Result {
	log.Println("PostController - New()")
	post := models.Post{}
	return c.Render(post)

}

func (c Post) Create(title, body string) revel.Result {
	log.Println("PostController - Create()")
	post := models.Post{Title: title, Body : body, CreatedAt:time.Now(), UpdatedAt:time.Now()}
	
	c.Txn.Create(&post)//db insert

	c.Flash.Success("포스트 작성 완료")
	return c.Redirect("/posts")
}