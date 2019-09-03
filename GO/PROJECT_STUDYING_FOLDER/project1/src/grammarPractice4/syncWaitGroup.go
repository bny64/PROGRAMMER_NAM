package grammarPractice4

import (
	"fmt"
	"runtime"
	"sync"
)

type counter4 struct {
	i  int64
	mu sync.Mutex
}

func (c *counter4) increment4() {
	c.mu.Lock()
	c.i += 1
	c.mu.Unlock()
}

func (c *counter4) display4() {
	fmt.Println(c.i)
}

func SyncWaitGroupTest() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := counter4{i: 0}
	wg := sync.WaitGroup{} //WaitGroup 생성

	for i := 0; i < 1000; i++ {
		wg.Add(1) //WaitGroup의 고루틴 갯수 1 증가
		go func() {
			defer wg.Done() //고루틴 종료 시 Done() 처리
			c.increment4()
		}()
		//Add 메서드로 추가한 고루틴의 개수와 Done 메서드를 호출한 횟수는 같아야 함.
	}
	wg.Wait()
	c.display4()
}
