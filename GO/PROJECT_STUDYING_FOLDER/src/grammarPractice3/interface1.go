package grammarPractice3

import "fmt"

type shaper interface {
	area() float64
}

func describe(s shaper) {
	fmt.Println("area : ", s.area())
}

type rect struct{ width, height float64 }

func (r rect) area() float64 {
	return r.width * r.height
}

func Interface1Test() {
	/* rect타입이 shaper 인터페이스에 정의된 메서드 area()를 제공하면
	rect 타입 값을 shaper인터페이스로 사용할 수 있다 */
	r := rect{3, 4}
	describe(r)
}
