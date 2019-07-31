package grammarPractice

import "fmt"

func ArrayTransFuncTest() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//multiply(numbers, 5)
	fmt.Println(numbers)
}

func multiply(numbers []int, factor int) {
	for i := range numbers {
		numbers[i] *= factor
	}
}
