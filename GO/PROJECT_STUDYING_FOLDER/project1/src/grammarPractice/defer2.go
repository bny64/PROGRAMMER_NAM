package grammarPractice

import "fmt"

func defer2Enter(s string){ fmt.Println("entering:", s)}
func defer2Leave(s string){ fmt.Println("leaving:", s)}

func defer2A(){
	defer2Enter("a")
	defer defer2Leave("a")
	fmt.Println("in a")
}
func defer2B(){
	defer2Enter("b")
	defer defer2Leave("b")
	fmt.Println("in b")
	defer2A()
}

func Defer2Test(){
	defer2B()
}

/*
	defer는 FILO 형식
*/