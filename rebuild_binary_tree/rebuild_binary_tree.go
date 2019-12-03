package main

import (
	"fmt"
)

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

// preOrd 前序遍历数组
// inOrder 中序遍历数组
func reBuildBinaryTree(preOrd []int, inOrd []int) *treeNode {

	if len(preOrd) != len(inOrd) {
		panic("pre-order traversal and in-order traversal have different size")
	}

	if len(preOrd) == 0 || len(inOrd) == 0 {
		return nil
	}
	root := &treeNode{val: preOrd[0]}
	i := 0
	for ; i < len(inOrd); i++ {
		if inOrd[i] == preOrd[0] {
			root.left = reBuildBinaryTree(preOrd[1:i+1], inOrd[:i])
			root.right = reBuildBinaryTree(preOrd[i+1:], inOrd[i+1:])
			break
		}
	}
	if i == len(inOrd) {
		fmt.Println("===", preOrd, inOrd)
		panic(fmt.Sprintf("Val: %d, in pre-order traversal, but not in in-order traversal", preOrd[0]))
	}

	return root
}


