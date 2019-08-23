package grammarPractice5

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("음수는 사용할 수 없습니다.")
		//return 0, fmt.Errorf("음수(%g)는 사용할 수 없습니다.", f)
	}
	//math.Sqrt() 사각형의 대각선 길이 구하는 함수.
	return math.Sqrt(f), nil
}

func Error2Test() {
	if f, err := Sqrt(-1); err != nil {
		fmt.Printf("Error : %s\n", err)
	} else {
		fmt.Println(f)
	}
}
