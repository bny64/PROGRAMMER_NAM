package controllers

import (
	"goblog/app/models"
	"golang.org/x/crypto/bcrypt"
	"github.com/revel/revel"
	"log"
)

type App struct {
	GormController
	CurrentUser *models.User	
}

func (c App) Login() revel.Result {
	return c.Render()
}

func (c App) CreateSession(username, password string)revel.Result {
	var user models.User

	c.Txn.Where(&models.User{Username:username}).First(&user)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		authKey := revel.Sign(user.Username)
		c.Session["authKey"] = authKey
		c.Session["username"] = user.Username
		c.Flash.Success("Welcome, " + user.Name)
		return c.Redirect(Post.Index)
	}

	//세션 정보를 모두 제거하고 홈으로 이동
	for k := range c.Session {
		delete(c.Session, k)
	}

	c.Flash.Out["username"] = username
	c.Flash.Error("Login failed")
	return c.Redirect(Home.Index)
}

func (c App) DestroySession() revel.Result{
	log.Println("AppController DestroySession()")
	for k := range c.Session {
		delete(c.Session, k)		
	}
	
	c.CurrentUser = nil

	return c.Redirect(Home.Index)
}

func (c *App) setCurrentUser() revel.Result {
	//뷰에서 currentUser를 사용할 수 있게 RenderArgs에 currentUser 추가
	defer func(){
		if c.CurrentUser != nil {
			c.ViewArgs["currentUser"] = c.CurrentUser
		} else {
			delete(c.ViewArgs, "currentUser")
		}
	}()

	username, ok := c.Session["username"]
	if !ok || username == ""{
		return nil
	}

	authKey, ok := c.Session["authKey"]
	if !ok || authKey == ""{
		return nil
	}

	if match := revel.Verify(username.(string), authKey.(string)); match {
		var user models.User
		c.Txn.Where(&models.User{Username:username.(string)}).First(&user)
		if &user != nil {
			c.CurrentUser = &user
		}		
	}
	log.Println("user : ", c.CurrentUser)
	log.Println("session : ", c.Session)
	return nil
}
