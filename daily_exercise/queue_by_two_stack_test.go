package daily_exercise

import "testing"

func TestQueueByTwoStack(t *testing.T) {
	q := newQueueByTwoStack()

	items := []int{1, 2, 3, 4, 5, 6, 7}
	for _, item := range items {
		q.push(item)
	}

	i := 0
	for {
		if item, err := q.pop(); err != nil {
			if item != items[i] {
				t.Fatalf("the %dth item pop is %d, but should be %d", i+1, item, items[i])
			}
			i++
		} else {
			return
		}
	}
}
