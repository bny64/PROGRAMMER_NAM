package grammarPractice5

import (
	"fmt"
	"log"
	"math"
	"time"
)

type sqrt2Error struct {
	time    time.Time //에러가 발생한 시간
	value   float64   //에러를 발생시킨 값
	message string    //에러 메세지
}

func (e sqrt2Error) Error() string {
	//error 인터페이스에 정의된 Error() 메서드 구현
	return fmt.Sprintf("[%v]Error - %s(value: %g)", e.time, e.message, e.value)
}

func sqrt2(f float64) (float64, error) {
	//매개변수로 전달된 값이 유효한 값이 아닐 때 sqrt2Error를 반환

	if f < 0 {
		return 0, sqrt2Error{time: time.Now(), value: f, message: "음수는 사용할 수 없습니다."}
	}
	if math.IsInf(f, 1) {
		return 0, sqrt2Error{time: time.Now(), value: f, message: "무한대 값은 사용할 수 없습니다"}
	}
	if math.IsNaN(f) {
		return 0, sqrt2Error{time: time.Now(), value: f, message: "잘못된 수 입니다."}
	}

	return math.Sqrt(f), nil
}

func CustomErrorTypeTest() {
	v, err := sqrt2(9) //9의 제곱근. 정상 동작
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(v)

	v, err = sqrt2(-9) //-1의 제곱근. 에러 발생
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(v)
}

//사용자 정의 에러 타입 확인
/*
if e, ok := err.(SqrtError); ok {
	fmt.Println("Sqrt Error", e)
}

switch e:=err.(type) {
case SqrtError:
	fmt.Println("Sqrt Error", e)
default:
	fmt.Println("Default Error", e)
} */
