package main

import (
	"fmt"
	"net/http"
	"strings"
)

type router struct {
	//키 : http 메서드
	//값 : URL패턴별로 실행할 HandlerFunc
	handlers map[string]map[string]http.HandlerFunc
}

type Handler interface {
	//Handler 인터페이스는 ServeHTTP 사용 가능
	//웹 요청을 받아서 처리하는 http.Handler 인터페이스 사용 가능
	ServeHTTP(http.ResponseWriter, *http.Request)
}

func (r *router) HandleFunc(method, pattern string, h http.HandlerFunc) {
	//http 메서드로 등록된 맵이 있는지 확인
	m, ok := r.handlers[method] //m => map타입 key:string,value:func
	if !ok {
		//등록된 맵이 없으면 새 맵을 생성
		m = make(map[string]http.HandlerFunc)
		r.handlers[method] = m
		//http 메서드로 등록된 맵에 URL 패턴과 핸들러 함수 등록
	}
	m[pattern] = h
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for pattern, handler := range r.handlers[req.Method] {
		if ok, _ := match(pattern, req.URL.Path); ok {
			handler(w, req)
			return
		}
	}
	http.NotFound(w, req)
	return
}

func match(pattern, path string) (bool, map[string]string) {
	if pattern == path {
		return true, nil
	}

	patterns := strings.Split(pattern, "/")
	paths := strings.Split(path, "/")

	if len(patterns) != len(paths) {
		return false, nil
	}

	params := make(map[string]string)
	fmt.Println(patterns, paths)

	for i := 0; i < len(patterns); i++ {
		fmt.Println("for loop : ", patterns[i])
	}

	for i := 0; i < len(patterns); i++ {
		switch {
		case patterns[i] == paths[i]:
			fmt.Println("patterns[i]==paths[i] : ", patterns[i], paths[i])
		case len(patterns[i]) > 0 && patterns[i][0] == ':':
			params[patterns[i][1:]] = paths[i]
			fmt.Println("params[patterns[i][1:]] = paths[i] : ", params[patterns[i][1:]], paths[i])
		default:
			return false, nil
		}
	}
	fmt.Println("params : ", params)
	return true, params
}
