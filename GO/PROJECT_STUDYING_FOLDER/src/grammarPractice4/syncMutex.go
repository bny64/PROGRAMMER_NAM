package grammarPractice4

import (
	"fmt"
	"runtime"
	"sync"
)

type counter2 struct {
	i  int64
	mu sync.Mutex //공유 데이터 i를 보호하기 위한 뮤텍스
}

func (c *counter2) increment2() {
	c.mu.Lock()   //i 값을 변경하는 부분(임계 영역)을 뮤텍스로 잠금
	c.i += 1      //공유 데이터 변경
	c.mu.Unlock() //i 값을 변경 완료 후 뮤텍스 잠금 해제
}

func (c *counter2) display2() {
	fmt.Println(c.i)
}

func Mutex1Test() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //최대 CPU

	c := counter2{i: 0}         //카운터 생성
	done := make(chan struct{}) //완료 신호 수신용 구조체 채널

	for i := 0; i < 1000; i++ {
		go func() {
			c.increment2()
			done <- struct{}{} //done 채널에 완료 신호 전송
		}()
	}

	for i := 0; i < 1000; i++ {
		<-done //채널에서 값을 꺼냄
	}

	c.display2()
}
