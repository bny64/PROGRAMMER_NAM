package grammarPractice4

import (
	"fmt"
	"runtime"
)

type counter struct {
	i int64
}

func (c *counter) increment() {
	c.i += 1
}

func (c *counter) display() {
	fmt.Println(c.i)
}

func main() {
	//모든 CPU사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := counter{i: 0}          //카운터 구조체 생성
	done := make(chan struct{}) //완료 신호 수신용 채널

	for i := 0; i < 1000; i++ {
		go func() { //고루틴
			c.increment()
			done <- struct{}{}
		}()
	}

	for i := 0; i < 1000; i++ {
		<-done
	}

	c.display()
	//책에서 나온 값 994
	//실제 값 1000
	//고 코드가 수정됨.
}
