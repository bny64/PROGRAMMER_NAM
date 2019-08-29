package main

import "net/http"

type router struct {
	//http.HandlerFunc(http.RsponseWriter, *http.Request)
	handlers map[string]map[string]http.HandlerFunc
}

func (r *router) HandleFunc(method, pattern string, h http.HandlerFunc) {
	m, ok := r.handlers[method] //method -> GET,POST and so on

	if !ok { //등록된 맵이 없으면
		m = make(map[string]http.HandlerFunc)
		r.handlers[method] = m
	}

	m[pattern] = h //h -> http.HandlerFunc
}
