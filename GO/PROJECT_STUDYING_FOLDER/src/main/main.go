package main

import (
	"grammarPractice"
	f "fmt"
)

func main(){
	//grammarPractice.VariableTest();
	//grammarPractice.IfTest();
	//grammarPractice.ForTest();
	//grammarPractice.ForLabelTest();
	a, b := grammarPractice.FuncTest(100, "hi")
	f.Println(a, b)
}
