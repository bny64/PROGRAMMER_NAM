package grammarPractice

import (
	"fmt"
	"strings"
)

func fbcMakeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix){
			return name + suffix
		}
		return name
	}
}

func FBCTest(){
	addZip := fbcMakeSuffix(".zip")
	addTgz := fbcMakeSuffix(".tar.gz")
	fmt.Println(addTgz("go1.5.1.src"))
	fmt.Println(addZip("go1.5.1.windows-amd64"))
}