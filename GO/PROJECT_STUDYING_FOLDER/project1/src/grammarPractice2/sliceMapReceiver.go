package grammarPractice2

import "fmt"

type numberMap map[string]int

func (m numberMap) add(key string, value int) {
	m[key] = value
}

func (m numberMap) remove(key string) {
	delete(m, key)
}

func SliceMapReceiverTest() {
	m := make(numberMap)
	m["one"] = 1
	m["two"] = 2
	m.add("three", 3)
	fmt.Println(m)
	m.remove("two")
	fmt.Println(m)

	/* 슬라이스와 맵은 참조타입이므로 포인터 변수를 정해주지 않아도 리시버의 값을 수정할 수 있다 */
}
