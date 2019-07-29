package grammarPractice

import (
	"bytes"
	"fmt"
	"unicode"
)

const (
	MaxInt8 = 1<<15 - 1
)

func BufferTest() {
	var buffer bytes.Buffer
	for i := 0; i < MaxInt8; i++ {
		if s, ok := nextString2(i); ok {
			buffer.WriteString(s)
		}
	}
	fmt.Println(buffer.String(), "\n")
}

func nextString2(i int) (s string, ok bool) {
	if unicode.IsLetter(rune(i)) {
		return string(i), true
	}
	return "", false
}
