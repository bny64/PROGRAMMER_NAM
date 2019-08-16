package grammarPractice4

import (
	"fmt"
)

func fibonacci(c, quit chan int) {
	fmt.Println("fibonacci 안")
	x, y := 0, 1
	for {
		select {
		case c <- x:
			fmt.Println("c <- x")
			x, y = y, x+y
			fmt.Println(x, y)
		case <-quit:
			fmt.Println("<-quit")
			fmt.Println("quit")
			return
		}
	}
}

func goSelectTest() {
	c := make(chan int)
	quit := make(chan int)
	fmt.Println("before fo func")
	go func() {
		fmt.Println("next go func()")
		for i := 0; i < 10; i++ {
			fmt.Println("for 안 ", i)
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
	/*
		select문은 하나의 고루틴이 여러 채널과 통신할 때 사용한다.
		case로 여러 채널을 대기시키고 있다가 실행 가능 상태가 된
		채널이 있으면 해당 케이스를 수행한다.
	*/
	/*
			결과
			before fo func
		fibonacci 안
		next go func()
		for 안  0
		0
		for 안  1
		c <- x
		1 1
		c <- x
		1 2
		1
		for 안  2
		1
		for 안  3
		c <- x
		2 3
		c <- x
		3 5
		2
		for 안  4
		3
		for 안  5
		c <- x
		5 8
		c <- x
		8 13
		5
		for 안  6
		8
		for 안  7
		c <- x
		13 21
		c <- x
		21 34
		13
		for 안  8
		21
		for 안  9
		c <- x
		34 55
		c <- x
		55 89
		34
		<-quit
		quit
	*/
}
