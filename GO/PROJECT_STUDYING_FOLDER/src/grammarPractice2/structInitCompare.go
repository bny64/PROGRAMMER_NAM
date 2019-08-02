package grammarPractice2

import "fmt"

type sicRect struct{ w, h float64 }

func StructInitCompareTest() {
	r1 := sicRect{1, 2}
	r2 := new(sicRect)
	r2.w, r2.h = 3, 4
	r3 := &sicRect{}
	r4 := &rect{5, 6}

	fmt.Println(r1) //{1 2}
	fmt.Println(r2) //&{3 4}
	fmt.Println(r3) //&{0 0}
	fmt.Println(r4) //&{5, 6}
	/* new(Type) = &Type{} 성립됨. */
}
