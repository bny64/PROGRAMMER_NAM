package grammarPractice5

import (
	"fmt"
	"log"
)

type fType func(int, int) int

func errorHandler(fn fType) fType {
	fmt.Println("errorHandler func in")
	return func(a int, b int) int {
		defer func() {
			fmt.Println("defer func in")
			if err, ok := recover().(error); ok {
				fmt.Println(err)
				fmt.Println(ok)
				log.Printf("run time panic: %v", err)
			}
		}()
		fmt.Println("before return fn(a, b)")
		return fn(a, b)
	}
}

func ClosureErrorTest() {
	fmt.Println(errorHandler(divide)(4, 2))
	fmt.Println("--border--")
	fmt.Println(errorHandler(divide)(3, 0))
}

func divide(a int, b int) int {
	return a / b
}

/* 결과값
errorHandler func in
before return fn(a, b)
defer func in
2
--border--
errorHandler func in
before return fn(a, b)
defer func in
runtime error: integer divide by zero
true
2009/11/10 23:00:00 run time panic: runtime error: integer divide by zero
0
*/
