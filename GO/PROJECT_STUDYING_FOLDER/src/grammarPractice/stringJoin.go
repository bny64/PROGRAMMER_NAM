package grammarPractice

import (
	"fmt"
	"strings"
	"unicode"
)

const (
	_ = 1<<15 - 1
)

func JoinTest() {
	fmt.Println(MaxInt8)
	strArr := []string{}
	for i := 0; i < MaxInt8; i++ {
		if s, ok := nextString1(i); ok {
			strArr = append(strArr, s)
		}
	}
	fmt.Println(strings.Join(strArr, ""))
}

func nextString1(i int) (s string, ok bool) {
	if unicode.IsLetter(rune(i)) {
		return string(i), true
	}
	return "", false
}
