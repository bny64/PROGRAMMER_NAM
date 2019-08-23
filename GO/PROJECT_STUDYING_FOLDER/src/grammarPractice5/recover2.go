package grammarPractice5

import "fmt"

func badCall() {
	panic("bad end")
}

func test() {
	defer func() {
		//패닉이 발생했을 때 recover()함수로 회복시키고 e에 할당.
		if e := recover(); e != nil {
			fmt.Printf("Panicking %s\r\n", e)
		}
	}()

	badCall()
	fmt.Printf("After bad call \r\n")
}

func Recover2Test() {
	fmt.Printf("Calling test\r\n")
	test()
	fmt.Printf("Test complete\r\n")
}
