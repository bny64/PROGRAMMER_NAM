package grammarPractice

import "fmt"

const (
	Running = 1 << iota
	Waiting
	Send
	Receive
)

/*
	내부에서 사용할 경우 함수이름 소문자로 시작
	외부에서 사용할 경우 함수이름 대문자로 시작
	go install <패키지명> 명령을 실행할 경우
	프로젝트명/pkg 디렉토리안에 패키지가 생성된다.
*/
func VariableTest(){
	stat := Running | Send	
	display(stat)
}

func display(stat int){
	if stat&Running == Running {
		fmt.Println("Running")
	}

	if stat&Waiting == Waiting {
		fmt.Println("Waiting")
	}

	if stat&Send == Send {
		fmt.Println("Send")
	}

	if stat&Receive == Receive {
		fmt.Println("Receive")
	}
}