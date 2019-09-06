package controllers

import (
	"goblog/app/routes"
	"goblog/app/models"
	"time"
	"github.com/revel/revel"
	"log"	
)

type Comment struct {
	App
}

func (c Comment) CheckUser() revel.Result {
	if c.MethodName != "Destroy" {
		return nil
	}

	if c.CurrentUser == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(App.Login)
	}

	if c.CurrentUser.Role != "admin" {
		c.Response.Status = 401
		c.Flash.Error("You are not admin")
		return c.Redirect(App.Login)
	}
	return nil
}

func (c Comment) Create(postId int, body, commenter string) revel.Result {
	log.Println("CommentController Create()")
	comment := models.Comment{PostId:postId, Body:body, Commenter:commenter, CreatedAt:time.Now(), UpdatedAt:time.Now()}

	if err := c.Txn.Create(&comment).Error; err != nil {
		panic(err)
	}

	c.Flash.Success("댓글 작성 완료")
	return c.Redirect(routes.Post.Show(postId))
}

func (c Comment) Destroy(postId, id int) revel.Result {
	log.Println("CommentController Destroy()")
	comment := models.Comment{}

	if err := c.Txn.Where("id = ?", id).Delete(&comment).Error; err != nil {
		panic(err)
	}

	c.Flash.Success("댓글 삭제 완료")
	return c.Redirect(routes.Post.Show(postId))
}