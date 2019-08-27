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

func main() {

	r := &router{make(map[string]map[string]http.HandlerFunc)}

	r.HandleFunc("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "welcome!")
	})

	r.HandleFunc("GET", "/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "about")
	})

	r.HandleFunc("GET", "/users/:id", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "retrieve user")
	})

	r.HandleFunc("GET", "/users/:user_id/addresses/:address_id", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "retrieve user's address")
	})

	r.HandleFunc("POST", "/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "create user")
	})

	r.HandleFunc("POST", "/users/:user_id/addresses", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "create user's address")
	})

	http.ListenAndServe(":8080", r)

	/* http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "welcome!")
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "about")
	})
	http.ListenAndServe(":8080", nil) */
}

func (r *router) HandleFunc(method, pattern string, h http.HandlerFunc) {
	// http 메서드로 등록된 맵이 있는지 확인.
	m, ok := r.handlers[method]
	if !ok {
		//등록된 맵이 없으면 새 맵을 생성
		m = make(map[string]http.HandlerFunc)
		r.handlers[method] = m
	}
	//http 메서드로 등록된 맵에 URL 패턴과 핸들러 함수 등록
	m[pattern] = h
}

func match(pattern, path string) (bool, map[string]string) {
	// 패턴과 패스가 정확히 일치하면 바로 true를 변환
	if pattern == path {
		return true, nil
	}

	//패턴과 패스를 "/"단위로 구분
	patterns := strings.Split(pattern, "/")
	paths := strings.Split(path, "/")

	//패턴과 패스를 "/"로 구분한 후 부분 문자열 집합의 개수가 다르면 false를 반환
	if len(patterns) != len(paths) {
		return false, nil
	}

	//패턴에 일치하는 URL매개변수를 담기 위한 params 맵생성
	params := make(map[string]string)

	// "/"로 구분된 패턴/패스의 각 문자열을 하나씩 비교
	for i := 0; i < len(patterns); i++ {
		switch {
		case patterns[i] == paths[i]:
			//패턴과 패스의 부분 문자열이 일치하면 바로 다음 루프 수행.
		case len(patterns[i]) > 0 && patterns[i][0] == ':':
			//패턴이 ':' 문자로 시작하면 parmas에 URL params를 담은 후 다음 루프 수행.
			params[patterns[i][1:]] = paths[i]
		default:
			//일치하는 경우가 없으면 false를 반환
			return false, nil
		}
	}

	//true와 params를 반환
	return true, params

}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	for pattern, handler := range r.handlers[req.Method] {
		if ok, _ := match(pattern, req.URL.Path); ok {
			//요청 URL에 해당하는 handler 수행
			handler(w, req)
			return
		}
	}

	//요청 URL에 해당하는 handler를 찾기 못하면 NotFound 에러 처리
	http.NotFound(w, req)
	return

	/* if m, ok := r.handlers[req.Method]; ok {
		if h, ok := m[req.URL.Path]; ok {
			h(w, req)
			return
		}
	}
	http.NotFound(w, req) */
}
