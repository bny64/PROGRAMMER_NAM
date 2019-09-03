package grammarPractice5

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

func TitleCase(s string) string {
	return strings.Title(s)
}

func ReflectFuncTest() {
	caption := "go is an open source programming language"

	//TitleCase를 바로 호출
	title := TitleCase(caption)
	fmt.Println(title)

	//TitleCase를 동적 호출
	titleFuncValue := reflect.ValueOf(TitleCase) //
	fmt.Println(titleFuncValue)
	fmt.Println(reflect.TypeOf(titleFuncValue))
	values := titleFuncValue.Call([]reflect.Value{reflect.ValueOf(caption)})
	fmt.Println("------\n")
	log.Printf("reflect.ValueOf(caption) : %s", reflect.ValueOf(caption))
	log.Printf("[]reflect.Value{reflect.ValueOf(caption)} : %s", []reflect.Value{reflect.ValueOf(caption)})
	log.Println(reflect.TypeOf([]reflect.Value{reflect.ValueOf(caption)}))
	title = values[0].String()
	fmt.Println(title)
}
