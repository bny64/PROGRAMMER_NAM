package grammarPractice

import "fmt"

func ChangeSliceTest() {
	ns1 := []int{1, 2, 3}
	ns2 := append(ns1, ns1...)
	ns3 := append(ns1, 4, 5)
	fmt.Println(ns1)
	fmt.Println(ns2)
	fmt.Println(ns3)
}
