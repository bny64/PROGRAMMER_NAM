package grammarPractice2

import "fmt"

type rect7 struct{ width, height float64 }

func (r rect7) area7() float64 {
	return r.width * r.height
}

func (r *rect7) resize(w, h float64) {
	r.width += w
	r.height += h
}

func MethodFucExpTest() {
	r := rect7{3, 4}
	fmt.Println("area :", r.area7())
	r.resize(10, 10)
	fmt.Println("area :", r.area7())
	/* area7() 메서드의 함수 표현식
	서명 : func(rect) float64*/
	areaFn := rect7.area7
	/* resize() 메서드의 함수 표현식
	서명 : func(*rect, float64, float64) */
	resizeFn := (*rect7).resize

	fmt.Println("area:", areaFn(r))
	resizeFn(&r, -10, -10)
	fmt.Println("area :", areaFn(r))
}
