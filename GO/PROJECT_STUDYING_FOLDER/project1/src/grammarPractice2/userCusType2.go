package grammarPractice2

import "fmt"

type quantity2 int
type costCalculator func(quantity2, float64) float64

func describe(q quantity2, price float64, c costCalculator) {
	fmt.Printf("quantity2: %d, price: %0.0f, cost : %0.0f\n",
		q, price, c(q, price))
}

func UserCusType2Test() {
	var offBy10Percent, offBy1000Won costCalculator

	offBy10Percent = func(q quantity2, price float64) float64 {
		return float64(q) * price * 0.9
	}

	offBy1000Won = func(q quantity2, price float64) float64 {
		return float64(q)*price - 1000
	}

	describe(3, 10000, offBy10Percent)
	describe(3, 10000, offBy1000Won)
}
