package grammarPractice

import "fmt"

func callback(y int, f func(int, int)) {
	f(y, 2)
}

func add(a, b int) {
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
}

func FBPTest() {
	callback(1, add)
}
