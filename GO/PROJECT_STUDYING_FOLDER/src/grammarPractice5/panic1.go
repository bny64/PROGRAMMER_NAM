package grammarPractice5

import "fmt"

func Panic1Test() {
	fmt.Println("Starting the program")
	panic("A server error occurred: stopping the program!")
	//패닉이 발생되면 프로그램은 종료된다.
	fmt.Println("Ending the program")
}
