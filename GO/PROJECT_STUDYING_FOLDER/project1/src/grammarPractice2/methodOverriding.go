package grammarPractice2

import "fmt"

func (t DiscountItem) Cost() float64 {
	return t.rmItem.Cost() * (1.0 - t.discountRate/100)
}

func MethodOverridingTest() {
	shoes := rmItem{"Women's Walking Shoes", 30000, 2}
	eventShoes := DiscountItem{
		rmItem{"Sports Shoes", 50000, 3},
		10.00,
	}

	fmt.Println(shoes.Cost())
	fmt.Println(eventShoes.Cost())
	fmt.Println(eventShoes.rmItem.Cost())
}
