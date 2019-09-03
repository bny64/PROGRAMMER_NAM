package main

import (
	"net/http"

	sessions "github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

const (
	sessionKey    = "simple_chat_session"
	sessionSecret = "simple_chat_session_secret"
)

var renderer *render.Render

func init() {
	renderer = render.New()
}

func main() {
	n := negroni.Classic()
	store := cookiestore.New([]byte(sessionKey))
	n.Use(sessions.Sessions(sessionKey, store))
	n.Use(LoginRequired("/login", "/auth"))

	router := httprouter.New()

	router.GET("/login", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		renderer.HTML(w, http.StatusOK, "login", nil)
	})

	router.GET("/logout", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		sessions.GetSession(r).Delete(currentUserKey)
		http.Redirect(w, r, "/login", http.StatusFound)
	})

	router.GET("/auth/:action/:provider", loginHandler)
	n.UseHandler(router)
	n.Run(":8080")
}
