package grammarPractice4

import (
	"fmt"
	"time"
)

func GoroutineTimeoutTest() {
	quit := make(chan struct{})
	done := process(quit) //string 채널
	timeout := time.After(1 * time.Second)

	select {
	//done채널과 timeout채널을 동시에 대기하다가 먼저 수신되는 케이스를 수행.
	case d := <-done: //done채널로 process()함수의 실행 결과를 받아옴.
		fmt.Println(d)
	case <-timeout: //timeout채널로 10밀리초 후의 시간을 받아온다.
		fmt.Println("Timeout!")
		quit <- struct{}{}
	}
}

//파라미터 : 수신 전용 채널
func process(quit <-chan struct{}) chan string {
	done := make(chan string)
	go func() {
		go func() {
			time.Sleep(10 * time.Second)

			done <- "Complete!"
		}()

		<-quit //채널에 종료 신호 전송
		return
	}()
	return done
}
