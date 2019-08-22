package grammarPractice4

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
	sm.c <- command{action: set, key: k, value: v}
}

func (sm SharedMap) Get(k string) (interface{}, bool) {
	callback := make(chan interface{})
	sm.c <- command{action: get, key: k, result: callback}
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
