package controllers

import (
	"github.com/revel/revel"
	"log"
)

type Home struct {
	*revel.Controller
}

func(c Home) Index() revel.Result {
	log.Println("home controller")
	return c.Render()
}