package grammarPractice

import (
	"fmt"
	"unicode/utf8"
)

func StringTest() {
	s1 := "hello"

	fmt.Println([]rune(s1))
	fmt.Println([]byte(s1))

	s2 := "안녕하세요"

	fmt.Println([]rune(s2))
	fmt.Println([]byte(s2))

	fmt.Println(string([]rune{50504, 45397, 54616, 49464, 50836}))

	for _, c := range s2 {
		fmt.Println(c)
		fmt.Println(string(c))
	}

	fmt.Println(len(s2))
	fmt.Println(utf8.RuneCountInString(s2))
}
