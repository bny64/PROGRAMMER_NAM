package controllers

import (
	"goblog/app/routes"
	"github.com/revel/revel"
	"goblog/app/models"
	"github.com/jinzhu/gorm"
	"log"
	"time"
	_ "strconv"
)

type Post struct {
	App
}

func (c Post) CheckUser() revel.Result{
	
	log.Println("post.go : CheckUser Func : c.Session : ", c.Session)
	log.Println("post.go : CheckUser Func : c.CurrentUser : ", c.CurrentUser)
	
	switch c.MethodName {
	case "Index", "Show":		
		return nil
	}
	
	if c.CurrentUser == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(App.Login)
	}

	if c.CurrentUser.Role != "admin" {
		c.Response.Status = 401
		c.Flash.Error("You're not admin")
		return c.Redirect(App.Login)
	}
	return nil
}

func (c Post) Index() revel.Result {		
	
	log.Println("post.go : Index Func : c.Session : ", c.Session)
	log.Println("post.go : Index Func : c.CurrentUser : ", c.CurrentUser)

	var posts []models.Post
	if err := c.Txn.Find(&posts).Error; err != nil {
		panic(err)
	} //전체 조회

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
	
	if err := c.Txn.Create(&post).Error; err != nil {//db insert
		panic(err)
	}

	c.Flash.Success("포스트 작성 완료")
	return c.Redirect(routes.Post.Index())
}

func (c Post) Show(id int) revel.Result {
	post := getPost(c.Txn, id)
	
	log.Println("post.go : Show Func : c.Session : ", c.Session)
	log.Println("post.go : Show Func : c.CurrentUser : ", c.CurrentUser)

	return c.Render(post)
}

func getPost(txn *gorm.DB, id int) (models.Post){
	post := models.Post{}
	if err := txn.Where("id = ?", id).First(&post).Error; err != nil {
		panic(err)
	}
	
	post.Comments = getComments(txn, id)

	return post
}

func getComments(txn *gorm.DB, postId int) (comments []models.Comment){
	var comment []models.Comment
	
	if err:= txn.Where("post_id = ?", postId).Find(&comment).Order("created_at desc").Error; err != nil {
		panic(err)
	}

	for _, val := range comment {
		comments = append(comments, val)
	}
	
	return
}

func (c Post) Edit(id int) revel.Result {
	post := getPost(c.Txn, id)

	return c.Render(post)
}

func (c Post) Update(id int, title, body string) revel.Result {
	log.Println("PostController Update()")
	post := models.Post{}
	if err := c.Txn.Model(&post).Where("id = ?", id).Updates(map[string]interface{}{"title":title, "body":body}).Error; err != nil {
		panic(err)
	}

	c.Flash.Success("포스트 수정 완료")
	return c.Redirect(routes.Post.Show(id))
}

func (c Post) Destroy(id int) revel.Result {
	post := models.Post{}
	if err := c.Txn.Where("id = ?", id).Delete(&post).Error; err != nil {
		panic(err)
	}
	c.Flash.Success("포스트 삭제 완료")
	return c.Redirect(routes.Post.Index())
}