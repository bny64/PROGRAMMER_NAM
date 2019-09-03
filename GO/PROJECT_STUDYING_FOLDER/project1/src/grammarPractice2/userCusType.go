package grammarPractice2

import "fmt"

type quantity int
type dozen []quantity

func UserCusTypeTest() {
	/* a := int(3); 이렇게 선언도 가능.*/
	var d dozen
	for i := quantity(1); i <= 12; i++ {
		d = append(d, i)
	}
	fmt.Println(d)
}
