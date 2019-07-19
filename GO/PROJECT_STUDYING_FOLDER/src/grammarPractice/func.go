package grammarPractice

func FuncTest(param1 int, param2 string) (int, string) {
	param1 += 30
	param2 += "hello"
	return param1, param2
}

func FuncTest2(param1 int, param2 string) (i int, s string){
	i = param1 + 30
	s = param2 + "hello"
	return
}