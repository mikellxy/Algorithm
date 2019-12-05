package daily_exercise

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const MAXSIZE = 100

func getVal(exists map[int]bool) int {
	var val int
	for {
		val = rand.Intn(1000 * MAXSIZE)
		if !exists[val] {
			exists[val] = true
			return val
		}
	}
}

func getSlice() []int {
	exist := make(map[int]bool)
	var ret []int
	var n int
	for {
		if len(ret) == MAXSIZE {
			return ret
		}

		if rand.Float64() >= 0.5 {
			ret = append(ret, getVal(exist))
			n ++
		} else {
			return ret
		}
	}
}

func evenOddPartitionLite(arr []int) []int {
	var ret []int
	for _, v := range arr {
		if v%2 == 1 {
			ret = append(ret, v)
		}
	}
	for _, v := range arr {
		if v%2 == 0 {
			ret = append(ret, v)
		}
	}
	return ret
}

func isSame(a []int, b []int, t *testing.T) {
	if len(a) != len(b) {
		t.Log(fmt.Sprintf("a: %v, b: %v", a, b))
	}

	for i, v := range a {
		if v != b[i] {
			t.Log(fmt.Sprintf("a: %v, b: %v", a, b))
		}
	}
}

func TestEvenOddPartition(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		arr := getSlice()
		t.Log(arr)

		a := evenOddPartition(arr)
		b := evenOddPartitionLite(arr)
		isSame(a, b, t)

		time.Sleep(time.Second)
	}
}
