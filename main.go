package main

import "fmt"

type node struct {
	val  int
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

func printNodes(root *node) {
	for root != nil {
		fmt.Println(root.val)
		root = root.next
	}
}

func recursionHandler(root *node) *node {
	if root == nil || root.next == nil {
		return root
	}
	newRoot := recursionHandler(root.next)
	root.next.next = root
	root.next = nil
	return newRoot
}

func loopHandler(root *node) *node {
	var prev *node
	cur := root

	for cur != nil {
		tmp := cur.next
		cur.next = prev
		prev = cur
		cur = tmp
	}

	return prev
}

func searchLessOrEqual(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			if mid-1 >=0 && arr[mid-1] <= target {
				return mid - 1
			}
			right = mid - 1
		}
	}

	if right == len(arr)-1 {
		return right
	}

	return -1
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	for _, i := range []int{-1, 1, 3, 5, 7, 9, 11, 13, 15} {
		idx := searchLessOrEqual(arr, i)
		if idx == -1 {
			fmt.Println(-1)
		} else {
			fmt.Println(arr[idx])
		}
	}
}
