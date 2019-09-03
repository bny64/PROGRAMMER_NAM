package grammarPractice4

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

type counter5 struct {
	i int64
}

func (c *counter5) increment() {
	atomic.AddInt64(&c.i, 1)
}

func (c *counter5) display() {
	fmt.Println(c.i)
}

func SyncAtomicTest() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := counter5{i: 0}
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.increment()
		}()

	}

	wg.Wait()
	c.display()
}
