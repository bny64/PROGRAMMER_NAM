package grammarPractice5

import (
	"errors"
	"fmt"
)

func Error1Test() {
	errNotFound := errors.New("Not found error")

	fmt.Println("error : ", errNotFound)
	fmt.Println("error : ", errNotFound.Error())
}
