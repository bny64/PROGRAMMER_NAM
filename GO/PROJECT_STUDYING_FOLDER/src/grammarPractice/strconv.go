package grammarPractice

import (
	f "fmt"
	s "strconv"
)

func StrconvTest(ss string){
	if v, err := s.Atoi(ss); err != nil {
		f.Printf("%s는 정수가 아닙니다.\n", ss)
	}else{
		f.Printf("정수 값은 %d입니다.\n", v)
	}
}