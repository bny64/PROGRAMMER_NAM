package grammarPractice2

import "fmt"

type rmItem struct {
	name     string
	price    float64
	quantity int
}

func (t rmItem) Cost() float64 {
	return t.price * float64(t.quantity)
}

type DiscountItem struct {
	rmItem
	discountRate float64
}

func ReuseMethodTest() {
	shoes := Item{"Women's Walking Shoes", 30000, 2}
	eventShoes := DiscountItem{
		rmItem{"Sports Shoes", 50000, 3},
		10.00,
	}
	fmt.Println(shoes.Cost())      //60000
	fmt.Println(eventShoes.Cost()) //150000
	/* DiscountItem타입에는 Cost() 메서드가 정의되어 있지 않지만
	DiscountItem의 임베디드 필드인 rmItem에는 Cost() 메서드가 정의되어 있으므로
	DiscountItem에서도 Cost()메서드를 사용할 수 있다. */
}
