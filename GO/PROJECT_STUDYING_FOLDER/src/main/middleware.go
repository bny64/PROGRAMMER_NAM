package main

import (
	"log"
	"time"
)

//HandleFunc을 매개변수로 받아 새로운 HandlerFunc을 반환
type Middleware func(next HandlerFunc) HandlerFunc

func logHandler(next HandlerFunc) HandlerFunc {
	return func(c *Context) {
		t := time.Now()
		next(c)
		log.Printf("[%s] %q %v\n",
			c.Request.Method,
			c.Request.URL.String(),
			time.Now().Sub(t))
	}
}
