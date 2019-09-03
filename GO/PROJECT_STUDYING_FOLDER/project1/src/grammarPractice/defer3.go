package grammarPractice

import "fmt"

func defer3Enter(s string) string {
	fmt.Println("entering : ", s)
	return s
}

func defer3Leave(s string){
	fmt.Println("leaving:",s)
}

func defer3A(){
	defer defer3Leave(defer3Enter("a"))
	fmt.Println("in a")
}

func defer3B(){
	defer defer3Leave(defer3Enter("b"))
	fmt.Println("in b")
	defer3A()
}

func Defer3Test(){
	defer3B()
	/* 
		entering:b
		in b
		entering:a
		in a
		leaving: a
		leaving: b
	*/
}