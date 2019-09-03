package grammarPractice3

import "fmt"

func display2(s interface{}) {
	fmt.Println(s)
}

func EmptyInterfaceTest() {
	r := rect{3, 4}
	c := circle{2.5}
	display2(r)

	display2(c)
	display2(2.5)
	display2("rect struct")
	/* interface{} 타입은 메서드를 정의하지 않은 인터페이스.
	정의된 메서드가 없어서 어떤 값이라도 interface{}로 사용할 수 있다. */
}
