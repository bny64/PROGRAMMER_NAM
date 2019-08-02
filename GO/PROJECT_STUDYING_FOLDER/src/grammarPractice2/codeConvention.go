package grammarPractice2

import "fmt"

type Clothes struct {
	name     string
	price    float64
	quantity int
}

func New(name string, price float64, quantity int) *Clothes {
	if price <= 0 || quantity <= 0 || len(name) == 0 {
		return nil
	}
	return &Clothes{name, price, quantity}
}

func (c *Clothes) Name() string {
	return c.name
}

func (c *Clothes) SetName(n string) {
	if len(n) != 0 {
		c.name = n
	}
}

func (c *Clothes) Price() float64 {
	return c.price
}

func (c *Clothes) SetPrice(p float64) {
	if p > 0 {
		c.price = p
	}
}

func (c *Clothes) Quantity() int {
	return c.quantity
}

func (c *Clothes) SetQuantity(q int) {
	if q > 0 {
		c.quantity = q
	}
}

func CodeConventionTest() {
	shirts := New("Men's Slim-Fit Shirt", 25000, 3)
	shirts.SetQuantity(2)
	fmt.Println("Name : ", shirts.Name())
	fmt.Println("Price : ", shirts.Price())
	fmt.Println("Quantity : ", shirts.Quantity())
}
