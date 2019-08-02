package grammarPractice2

import "fmt"

type rect struct {
	width  float64
	height float64
}

type shaper interface {
	area() float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func describe2(s shaper) {
	fmt.Println("area : ", s.area())
}

func StructTest() {
	r := rect{3, 4}
	describe2(r)
}
