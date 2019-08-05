package grammarPractice3

import "fmt"

type Coster interface {
	Cost() float64
}

type Item struct {
	name     string
	price    float64
	quantity int
}

type DiscountItem struct {
	Item
	discountRate float64
}

func displayCost(c Coster) {
	fmt.Println("cost :", c.Cost())
}

func (t Item) Cost() float64 {
	return t.price * float64(t.quantity)
}

func (t DiscountItem) Cost() float64 {
	return t.Item.Cost() * (1.0 - t.discountRate/100)
}

func PolymorPhismTest() {
	shoes := Item{"Sports Shoes", 30000, 2}
	eventShoes := DiscountItem{
		Item{"Women's Walking Shoes", 50000, 3},
		10.00,
	}
	/* Item타입과 DiscountItem타입은 Cost() float64 메서드를 가지므로
	Coster 인터페이스로 사용할 수 있다. */
	displayCost(shoes)
	displayCost(eventShoes)
}
