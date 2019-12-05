package daily_exercise

import (
	"fmt"
	"testing"
)

func TestTwoDArraySeach(t *testing.T) {
	arr := [][]int{
		{1, 3, 5, 7, 9},
		{10, 12, 14, 16, 18},
		{19, 21, 23, 25, 27},
		{28, 30, 32, 34, 36},
	}

	table := map[int]bool{
		0:  false,
		38: false,
	}

	for _, row := range arr {
		for _, val := range row {
			table[val] = true
			table[val+1] = false
		}
	}

	fmt.Println("*****", table)

	for val, res := range table {
		r := twoDArraySeach(arr, val)
		if r != res {
			t.Fatalf("result of searching %d in the array is %v, but should be %v", val, r, res)
		}
	}

}
