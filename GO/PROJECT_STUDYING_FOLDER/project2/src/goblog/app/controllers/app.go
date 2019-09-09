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

/* 
로그아웃 하게 되면 다른 컨트롤러에서는 c.CurrentUser는 nil이 찍히는데
setCurrentUser 함수를 통과할 때 c.CurrentUser가 이전에 담아놓은 User의 정보가 찍히는지
알 수가 없다. 나중에 더 공부하고 포인터를 더 상세하게 이해하게 되면 더 보도록 하자.
*/

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
		log.Println("app.go : CreateSession Func : c.Session : ", c.Session)
		log.Println("app.go : CreateSession Func : c.CurrentUser : ", c.CurrentUser)
		return c.Redirect(Post.Index)
	}

	//세션 정보를 모두 제거하고 홈으로 이동
	for k := range c.Session {
		delete(c.Session, k)
	}

	c.Flash.Out["username"] = username
	c.Flash.Error("Login failed")
	log.Println("app.go : CreateSession Func : c.Session : ", c.Session)
	log.Println("app.go : CreateSession Func : c.CurrentUser : ", c.CurrentUser)
	return c.Redirect(Home.Index)
}

func (c App) DestroySession() revel.Result{
	 
	for k := range c.Session {
		delete(c.Session, k)		
	}
	
	log.Println("app.go : DestroySession Func : c.Session : ", c.Session)
	log.Println("app.go : DestroySession Func : c.CurrentUser : ", c.CurrentUser)
	

	return c.Redirect(Home.Index)
}

func (c *App) SetCurrentUser() revel.Result {
	//뷰에서 currentUser를 사용할 수 있게 RenderArgs에 currentUser 추가
	defer func(){
		if c.CurrentUser != nil {
			c.ViewArgs["currentUser"] = c.CurrentUser
		} else {
			delete(c.ViewArgs, "currentUser")
		}
	}()
	log.Println("app.go : setCurrentUser Func : c.Session : ", c.Session)
	log.Println("app.go : setCurrentUser Func : c.CurrentUser : ", c.CurrentUser)	
	
	username, ok := c.Session["username"]
	
	if !ok || username == ""{
		c.CurrentUser = nil
		return nil
	}
	
	authKey, ok := c.Session["authKey"]
	if !ok || authKey == ""{
		c.CurrentUser = nil
		return nil
	}
	
	log.Println("app.go : setCurrentUser Func2 : c.Session : ", c.Session)
	log.Println("app.go : setCurrentUser Func : c.CurrentUser : ", c.CurrentUser)	

	if match := revel.Verify(username.(string), authKey.(string)); match {
		var user models.User
		c.Txn.Where(&models.User{Username:username.(string)}).First(&user)
		if &user != nil {
			c.CurrentUser = &user
		}		
	}

	log.Println("app.go : setCurrentUser Func3 : c.Session : ", c.Session)
	log.Println("app.go : setCurrentUser Func3 : c.CurrentUser : ", c.CurrentUser)
	return nil
}
