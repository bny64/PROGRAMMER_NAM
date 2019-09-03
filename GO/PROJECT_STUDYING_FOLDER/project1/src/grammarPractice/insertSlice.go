package grammarPractice

import "fmt"

func InsertSliceTest() {
	s := []int{1, 2, 3, 4, 5}
	s = insert1(s, []int{-3, -2}, 0)
	fmt.Println(s)

	s = insert1(s, []int{0}, 2)
	fmt.Println(s)

	s = insert1(s, []int{6, 7}, len(s))
	fmt.Println(s)

}

func insert1(s, new []int, index int) []int {
	return append(s[:index], append(new, s[index:]...)...)
}

func insert2(s, new []int, index int) []int {
	result := make([]int, len(s)+len(new))
	position := copy(result, s[:index])
	position += copy(result[position:], new)
	copy(result[position:], s[index:])
	return result
}
