package grammarPractice

import (
	f "fmt"
)

func ForTest() {
	for i:=0; i<10; i++ {
		f.Println(i, "first")
	}

	for i:=0; i<10; {
		i = i + 1;
		f.Println(i, "second")
	}
}