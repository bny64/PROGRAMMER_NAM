package grammarPractice3

import "fmt"

func checkType(v interface{}) {
	switch v.(type) {
	case bool:
		fmt.Printf("%t is a bool\n", v)
	case int, int8, int16, int32, int64:
		fmt.Printf("%d is an int\n", v)
	case uint, uint8, uint16, uint32, uint64:
		fmt.Printf("%d is an unsigned int\n", v)
	case float64:
		fmt.Printf("%f is a float64\n", v)
	case complex64, complex128:
		fmt.Printf("%f is a complex\n", v)
	case string:
		fmt.Printf("%v is unknown type\n", v)
	}
}

func TypeAssertionTest() {
	checkType(3)
	checkType(1.5)
	checkType(complex(1, 5))
	checkType(true)
	checkType("s")
	checkType(struct{}{})
	checkType(nil)
}
