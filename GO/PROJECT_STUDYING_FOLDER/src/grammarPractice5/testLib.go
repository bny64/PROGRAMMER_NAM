package grammarPractice5

import (
	"testing"
)

func TestLenForMap(t *testing.T) {
	v := map[string]int{"A": 1, "B": 2}
	actual, expected := Len(v), 2
	if actual != expected {
		t.Errorf("%d != %d", actual, expected)
	}
}

func TestLenForString(t *testing.T) {
	v := "one"
	actual, expected := Len(v), 3
	if actual != expected {
		t.Errorf("%d != %d", actual, expected)
	}
}

func TestLenForSlice(t *testing.T) {
	v := []int{5, 0, 4, 1}
	actual, expected := Len(v), 4
	if actual != expected {
		t.Errorf("%d != %d", actual, expected)
	}
}

/* 실행방법
go test
go test -v(옵션: 각 테스트의 진행 상세 상황을 화면에 출력)
*/
