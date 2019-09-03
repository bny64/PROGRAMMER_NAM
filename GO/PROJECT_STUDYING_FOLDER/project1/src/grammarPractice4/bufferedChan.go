package grammarPractice4

import (
	"fmt"
)

func bufferedChanTest() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	go func() {
		fmt.Println("go func")
		c <- 3
		/*
			고루틴은 채널 c에 값을 전송할 수 있을 때까지
			대기하다가, 채널에 들어온 첫 번째 값을 수신해가는
			즉시 바로 채널에 값을 전송한다.
		*/
	}()
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("before <- c")
	fmt.Println(<-c)
}
