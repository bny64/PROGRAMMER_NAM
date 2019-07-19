package grammarPractice

import (
	"fmt"
)

func IfTest() {
	a := 10
	b := 30

	if a > 20 || b < 40 {
		fmt.Println("a > 20 || b < 40")
		fmt.Println(a, b)
	}

	if c := 50; c < 55 {
		fmt.Println("c:=50; c < 55")
		fmt.Println(c)
	}

	c := 100
	switch c {
	case 100:
		fmt.Println("c는 80이상")
		fallthrough; 
		//switch문에서 조건을 만나면 바로 빠져나오기 때문에
		//다음 case로 넘길 때는 fallthrough를 사용.
		//다음 case를 비교하는 것이 아니라 비교하지 않고 직접 case 안으로 들어가서 실행됨.
		//따라서 c는 120이하라는 메세지가 출력됨.
	case 120:
		fmt.Println("c는 120이하")		
	}
	//변수를 쓰지 않는 switch인 경우
	switch {
	case c > 99:
		fmt.Println("c는 99이상")
	case c>101:
		fmt.Println("c는101이상")
	}
	//if문보다 가독성이 좋음.
	//초기화 구문도 가능
	switch d:=300; {
	case d>200:
		fmt.Println("d는 200이상")
	}
	
}
