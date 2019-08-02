package grammarPractice2

import "fmt"

type csItem struct {
	name     string
	price    float64
	quantity int
}

func (t csItem) cost() float64 {
	return t.price * float64(t.quantity)
}

func ConnectStructTest() {
	var t csItem
	t.name = "Men's Slim-Fit Shirt"
	t.price = 25000
	t.quantity = 3

	fmt.Println(t.name)
	fmt.Println(t.price)
	fmt.Println(t.quantity)
	fmt.Println(t.cost())
}
