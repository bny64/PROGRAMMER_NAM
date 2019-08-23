package grammarPractice4

import "fmt"

type SharedMap struct {
	m map[string]interface{} //실제 값이 저장될 맵
	c chan command           //SharedMap에 명령을 전달하기 위한 채널
}

type command struct {
	key    string             //키
	value  interface{}        //값
	action int                //수행할 액션
	result chan<- interface{} //액션 처리 결과
}

const (
	//SharedMap에서 처리할 수 있는 액션
	set = iota
	get
	remove
	count
)

func (sm SharedMap) Set(k string, v interface{}) {
	//파라미터로 string타입의 k와 interface타입인 v를 받음.
	sm.c <- command{action: set, key: k, value: v}
	//sharedMap의 채널 c로 액션과 키, 값을 보냄.
}

func (sm SharedMap) Get(k string) (interface{}, bool) {
	//파라미터로 string타입의 k를받고 interface와 bool을 리턴시킴.
	callback := make(chan interface{})
	//callback인 interface채널을 선언하고 할당.
	sm.c <- command{action: get, key: k, result: callback}
	//sharedMap의 채널c로 get액션과 키, 결과를 보냄.
	result := (<-callback).([2]interface{})
	return result[0], result[1].(bool)
}

func (sm SharedMap) Remove(k string) {
	sm.c <- command{action: remove, key: k}
}

func (sm SharedMap) Count() int {
	callback := make(chan interface{})
	sm.c <- command{action: count, result: callback}
	return (<-callback).(int)
}

func (sm SharedMap) run() {
	for cmd := range sm.c { //SharedMap의 채널 c(command)에서 채널이 닫히는지 확인. -> cmd
		//채널이 닫히면 for문도 종료 되는 듯??
		switch cmd.action { //채널의 action이 다음과 같을 때
		case set: //set일 때
			sm.m[cmd.key] = cmd.value //채널의 key, value를 SharedMap의 key,value로 입력.
		case get:
			v, ok := sm.m[cmd.key]
			cmd.result <- [2]interface{}{v, ok}
		case remove:
			delete(sm.m, cmd.key)
		case count:
			cmd.result <- len(sm.m)
		}
	}
}

func NewMap() SharedMap {
	//리턴타입이 SharedMap인 NewMap함수
	//SharedMap 할당과 선언.
	sm := SharedMap{
		m: make(map[string]interface{}), //키:string 값:interface인 맵
		c: make(chan command),           //채널
	}
	go sm.run()
	return sm
}

func SharedMapTest() {
	m := NewMap()

	m.Set("foo", "bar")

	t, ok := m.Get("foo")

	if ok {
		bar := t.(string)
		fmt.Println("bar: ", bar)
	}

	m.Remove("foo")

	fmt.Println("Count: ", m.Count())
}
