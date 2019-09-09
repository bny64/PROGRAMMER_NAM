package controllers

import (
	"github.com/revel/revel"
	"log"
)

type Home struct {
	App
	*revel.Controller
}

func(c Home) Index() revel.Result {
	
	log.Println("HOME INDEX : ", c.Session)
	log.Println("HOME INDEX : ", c.CurrentUser)

	return c.Render()
}