package grammarPractice2

import "fmt"

type Option struct {
	name  string
	value string
}

type eItem struct {
	name     string
	price    float64
	quantity int
	Option
}

func EmbeddedTest() {
	shoes := eItem{"Sports Shoes", 30000, 2, Option{"color", "red"}}

	fmt.Println(shoes)

	fmt.Println(shoes.name)
	fmt.Println(shoes.value)
	fmt.Println(shoes.Option.name)
}
