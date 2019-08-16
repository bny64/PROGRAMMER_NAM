package grammarPractice4

import (
	"fmt"
	"time"
)

func GoChannelTest() {
	fmt.Println("main 함수 시작", time.Now())

	done := make(chan bool)
	go long2(done)
	go short2(done)

	<-done //값 수신
	<-done

	fmt.Println("main 함수 종료", time.Now())
	/*
		chan<- string 송신 전용 채널
		<-chan string 수신 전용 채널
		ex)
		일반 채널 :
		var name1 chan bool
		name1 := make(chan bool)
		받기 전용, 보내기 전용 채널
		var name2 chan<- bool
		var name3 <-chan bool
		name2 := make(chan<- bool)
		name3 := make(<-chan bool)
	*/
}

func long2(done chan bool) {
	fmt.Println("long 함수 시작", time.Now())
	time.Sleep(3 * time.Second) //3초 대기
	fmt.Println("long 함수 종료", time.Now())
	done <- true //채널에 값 전송
}

func short2(done chan bool) {
	fmt.Println("short 함수 시작", time.Now())
	time.Sleep(1 * time.Second) //1초 대기
	fmt.Println("short 함수 종료", time.Now())
	done <- true
}
