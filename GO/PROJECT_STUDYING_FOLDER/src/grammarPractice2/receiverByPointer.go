package grammarPractice2

import "fmt"

type quantity3 int

func (q quantity3) greaterThan(i int) bool {
	return int(q) > i
}

func (q *quantity3) increment() { *q++ }
func (q *quantity3) decrement() { *q-- }

func ReceiverByPointerTest() {
	//리시버는 값에 의한 호출방식으로 리시버 변수의 값이 복사되어 메서드 내부로 전달되고
	//메서드 내부에서는 리시버변수의 값을 변경할 수 없기 때문에
	//리시버 타입에 포인터를 사용하여 지정해야 한다.
	//포인터를 사용하면 리시버 변수의 주소값이 들어오기 때문에.
	q := quantity3(3)
	q.increment()
	fmt.Printf("Is q(%d) greater than %d? %t \n", q, 3, q.greaterThan(3))
	q.decrement()
	fmt.Printf("Is q(%d) greater than %d? %t \n", q, 3, q.greaterThan(3))
}
