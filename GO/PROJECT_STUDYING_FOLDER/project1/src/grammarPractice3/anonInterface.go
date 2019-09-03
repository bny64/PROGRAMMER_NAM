package grammarPractice3

import "fmt"

func display(s interface{ show() }) {
	s.show()
}

/*
type rect struct{ width, height float64 }
*/
func (r rect) show() {
	fmt.Printf("width: %f, height: %f\n", r.width, r.height)
}

type circle struct{ radius float64 }

func (c circle) show() {
	fmt.Printf("radius : %f\n", c.radius)
}

func AnonIterfaceTest() {
	r := rect{3, 4}
	c := circle{2.5}
	display(r)
	display(c)
}
