package grammarPractice

import "fmt"

func MapTest() {
	numberMap := map[string]int{}
	numberMap["one"] = 1
	numberMap["two"] = 2
	numberMap["three"] = 3
	fmt.Println(numberMap)

	if _, ok := numberMap["three"]; ok {
		fmt.Println("has")
	} else {
		fmt.Println("none")
	}

	numberMap2 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(numberMap2)

	numberMap3 := make(map[string]int, 3)
	numberMap3["one"] = 1
	numberMap3["two"] = 2
	numberMap3["three"] = 3
	fmt.Println(numberMap3)
	delete(numberMap3, "two")
	fmt.Println(numberMap3)

	numberMap4 := make(map[int]int, 3)
	numberMap4[1] = 1
	fmt.Println(numberMap4)
	//가능

	groupMap := make(map[string]string)
	group1 := []int32{1, 4, 6}
	group2 := []int32{2, 4, 5}
	group3 := []int32{4, 6, 7}

	groupMap[string(group1)] = "first"
	groupMap[string(group2)] = "second"
	groupMap[string(group3)] = "third"

	fmt.Println(string(group1))

	for i, v2 := range groupMap {
		fmt.Println(i, v2)
	}

}
