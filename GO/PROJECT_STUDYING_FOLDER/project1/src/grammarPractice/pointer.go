package grammarPractice

import "fmt"

type rect struct{ w, h float64 }

func PointerTest() {
	var p *int
	var pp **int

	i := 1
	p = &i
	pp = &p
	fmt.Println(i, *p, **pp)

	var j int = 1
	var q *int = &j
	var s *rect = &rect{1, 2}

	fmt.Println(q)
	fmt.Println(s)
	fmt.Println(*s)
}
