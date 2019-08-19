package grammarPractice4

import (
	"fmt"
	"runtime"
	"sync"
)

const initialValue = -500

type counter3 struct {
	i    int64
	mu   sync.Mutex //공유 데이터 i를 보호하기 위한 뮤텍스
	once sync.Once  //한 번만 수행할 함수를 지정하기 위한 Once 구조체
}

func (c *counter3) increment3() {
	c.once.Do(func() {
		c.i = initialValue //단 한 번만 실행됨.
	})

	c.mu.Lock() //i 값을 변경하는 부분(임계 영역)을 뮤텍스로 잠금.
	c.i += 1
	c.mu.Unlock()
}

func (c *counter3) display3() {
	fmt.Println(c.i)
}

func SyncOnceTest() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := counter3{i: 0}
	done := make(chan struct{})

	for i := 0; i < 1000; i++ {
		go func() {
			c.increment3()
			done <- struct{}{}
		}()
	}

	for i := 0; i < 1000; i++ {
		<-done
	}
	c.display3()
}
