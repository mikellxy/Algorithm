package daily_exercise

import (
	"fmt"
	"testing"
)

func TestGetKthElement(t *testing.T) {
	root := newLinkedList([]int{1,2,3})
	ele := getKthElement(root, 1)
	fmt.Println("kth node: ", ele)

	ele = getKthElement(root, 2)
	fmt.Println("kth node: ", ele)

	ele = getKthElement(root, 3)
	fmt.Println("kth node: ", ele)

	ele = getKthElement(root, 4)
	fmt.Println("kth node: ", ele)

	ele = getKthElement(root, 5)
	fmt.Println("kth node: ", ele)
}