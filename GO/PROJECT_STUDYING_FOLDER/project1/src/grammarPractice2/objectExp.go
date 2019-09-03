package grammarPractice2

import "fmt"

type Item struct {
	name     string
	price    float64
	quantity int
}

//javascript의 prototype이랑 비슷?
func (t Item) Cost() float64 {
	return t.price * float64(t.quantity)
}

func ObjectExpTest() {
	shirt := Item{
		name:     "Men's Slim-Fit Short",
		price:    25000,
		quantity: 3,
	}
	fmt.Println(shirt.Cost())
}
