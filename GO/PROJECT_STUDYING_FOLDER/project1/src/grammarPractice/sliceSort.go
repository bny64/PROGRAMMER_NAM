package grammarPractice

import (
	"fmt"
	"sort"
)

func SliceSortTest() {
	a := []string{"kim", "park", "lee", "bae"}

	sort.Strings(a)

	fmt.Println(a)
}
