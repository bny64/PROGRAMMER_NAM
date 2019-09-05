package controllers

import (
	"github.com/revel/revel"
	"goblog/app/models"
	"github.com/jinzhu/gorm"
	"log"
	"time"
	"strconv"
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

func (c Post) Show(id int) revel.Result {
	post := getPost(c.Txn, id)

	return c.Render(post)
}

func getPost(txn *gorm.DB, id int) (models.Post){
	post := models.Post{}
	txn.Where("id = ?", id).First(&post);
	
	return post
}

func (c Post) Edit(id int) revel.Result {
	post := getPost(c.Txn, id)

	return c.Render(post)
}

func (c Post) Update(id int, title, body string) revel.Result {
	post := models.Post{}
	c.Txn.Model(&post).Where("id = ?", id).Updates(map[string]interface{}{"title":title, "body":body})

	c.Flash.Success("포스트 수정 완료")
	return c.Redirect("/posts/"+strconv.Itoa(id))
}

func (c Post) Destroy(id int) revel.Result {
	post := models.Post{}
	c.Txn.Where("id = ?", id).Delete(&post)
	c.Flash.Success("포스트 삭제 완료")
	return c.Redirect("/posts")
}