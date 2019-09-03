package grammarPractice

import "fmt"

func SliceTest() {
	var a []int
	b := []int{}
	c := []int{1, 2, 3}
	d := [][]int{
		{1, 2},
		{3, 4, 5},
	}
	e := make([]int, 0)
	f := make([]int, 3, 5)

	fmt.Printf("%-10T %d %d %v\n", a, len(a), cap(a), a)
	fmt.Printf("%-10T %d %d %v\n", b, len(b), cap(b), b)
	fmt.Printf("%-10T %d %d %v\n", c, len(c), cap(c), c)
	fmt.Printf("%-10T %d %d %v\n", d, len(d), cap(d), d)
	fmt.Printf("%-10T %d %d %v\n", e, len(e), cap(e), e)
	fmt.Printf("%-10T %d %d %v\n", f, len(f), cap(f), f)
}
