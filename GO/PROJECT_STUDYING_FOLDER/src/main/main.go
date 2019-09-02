package main

import (
	"net/http"
	_ "net/http"

	"github.com/codegangsta/negroni"
	sessions "github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"
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

	router := httprouter.New()

	router.GET("/login", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		renderer.HTML(w, http.StatusOK, "login", nil)
	})

	router.GET("/logout", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		sessions.GetSession(r).Delete(keyCurrentUser)
		http.Redirect(w, r, "/login", http.StatusFound)
	})

	router.GET("/auth/:action/:provider", loginHandler)
}
