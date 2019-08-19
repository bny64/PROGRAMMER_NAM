package grammarPractice4

import (
	"fmt"
	"time"
)

func GoDefaultTest() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default: //모든 채널이 가능상태가 아닐 때 default케이스 수행함.
			fmt.Println(".")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
