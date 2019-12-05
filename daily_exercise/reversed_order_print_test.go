package daily_exercise

import "testing"

func TestReversedPrint(t *testing.T) {
	root := &node{val: 1}
	current := root
	for _, val := range []int{2, 3, 4, 5, 6, 7, 8} {
		current.next = &node{val: val}
		current = current.next
	}

	reversedPrint(root)
}
