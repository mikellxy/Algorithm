package daily_exercise

import "errors"

type node struct {
	val int
	next *node
}

func newLinkedList(arr []int) *node {
	if len(arr) == 0 {
		return nil
	}

	root := &node{val: arr[0]}
	cur := root

	for _, val := range arr[1:] {
		cur.next = &node{val: val}
		cur = cur.next
	}

	return root
}

type stack struct {
	storage []int
	len int
}

func (s *stack) push(item int) {
	s.storage = append(s.storage, item)
	s.len++
}

func (s *stack) Top() (int, error) {
	if s.len == 0 {
		return 0, errors.New("the stack is empty")
	}
	item := s.storage[len(s.storage)-1]
	s.storage = s.storage[:len(s.storage)-1]
	s.len--
	return item, nil
}

type treeNode struct {
	val int
	left *treeNode
	right *treeNode
}
