package grammarPractice

import (
	f "fmt"
)

func ForLabelTest(){
	x := 5
	table := [][]int{{1,5,9}, {2,6,5,13}, {5,3,7,4}}
	//found := false
	LOOP:
	for row:=0; row < len(table); row++{
		f.Println(row)
		for col :=0; col < len(table[row]); col++{
			if table[row][col] ==x {
				//found = true
				f.Printf("found %d(row:%d, col:%d)\n", x, row, col)
				continue LOOP;
			}
		}
		/* if found {
			break
		} */
	}
}