package grammarPractice

import "fmt"

func DeferTest(){
	f1()
}

func f1(){
	fmt.Println("f1 - start")
	defer f2()
	fmt.Println("f1 - end")
}


func f2(){
	fmt.Println("f2 - deferred")
}