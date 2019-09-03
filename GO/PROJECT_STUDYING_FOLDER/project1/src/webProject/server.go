package main

import "net/http"

type Server struct {
	*router
	middlewares  []Middleware
	startHandler HandlerFunc
}

func NewServer() *Server {
	r := &router{make(map[string]map[string]HandlerFunc)}

	s := &Server{router: r}
	s.middlewares = []Middleware{
		logHandler,
		recoverHandler,
		staticHandler,
		parseFormHandler,
		parseJsonBodyHandler,
	}
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := &Context{
		Params:         make(map[string]interface{}),
		ResponseWriter: w,
		Request:        r,
	}
	for k, v := range r.URL.Query() {
		c.Params[k] = v[0]
	}
	s.startHandler(c)
}

func (s *Server) Use(middlewares ...Middleware){
	s.middlewares = append(s.middlewares, middlewares...)
}