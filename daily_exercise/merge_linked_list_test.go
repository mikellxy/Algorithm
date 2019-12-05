package daily_exercise

import (
	"fmt"
	"testing"
)

func printNodes(root *node) {
	for root != nil {
		fmt.Println(root.val)
		root = root.next
	}
}

func TestMergeLinkedList(t *testing.T) {
	r1 := newLinkedList([]int{1, 3, 5, 7, 9})
	r2 := newLinkedList([]int{2, 4, 6, 8, 10})

	r := mergeLinkedList(r1, r2)
	printNodes(r)
}
