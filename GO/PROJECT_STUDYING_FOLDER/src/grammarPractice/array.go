package grammarPractice

import "fmt"

func ArrayTest() {
	var a [5]int
	b := [3]int{1, 2, 3}
	c := [3]int{1, 2}
	d := [...]int{1, 2, 3, 4, 5}
	e := [3][3]int{
		{1, 2, 3},
		{3, 4, 5},
	}
	fmt.Printf("%-10T %d %v\n", a, len(a), a)
	fmt.Printf("%-10T %d %v\n", b, len(b), b)
	fmt.Printf("%-10T %d %v\n", c, len(c), c)
	fmt.Printf("%-10T %d %v\n", d, len(d), d)
	fmt.Printf("%-10T %d %v\n", e, len(e), e)
}
