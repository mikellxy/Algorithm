package daily_exercise

import (
	"math/rand"
	"testing"
	"time"
)

func getTarget() (uint64, int) {
	var target uint64
	var num int
	for i := 0; i < 64; i++ {
		if rand.Float64() >= 0.5 {
			target |= 1 << uint(i)
			num++
		}
	}
	return target, num
}

func TestGetOnesNum(t *testing.T) {

	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		target, num := getTarget()
		ones := getOnesNum(target)
		if num != ones {
			t.Fatalf("The num of ones in %d is %d, but should be %d", target, ones, num)
		}
	}
}
