package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func inOrder(root *treeNode, res *[]int) {
	if root == nil {
		return
	}
	inOrder(root.left, res)
	*res = append(*res, root.val)
	inOrder(root.right, res)
}

func preOrder(root *treeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.val)
	preOrder(root.left, res)
	preOrder(root.right, res)
}

func getVal(maxNodes int, exist map[int]bool) int {
	var val int
	for {
		val = rand.Intn(10*maxNodes) + 1
		if !exist[val] {
			exist[val] = true
			break
		}
	}
	return val
}

func newTree(maxNodes int, exist map[int]bool) *treeNode {
	var root *treeNode
	var nodeNum int
	if maxNodes == 0 {
		return root
	}
	var queue []*treeNode
	if rand.Float64() > 0.5 {
		root = &treeNode{val: getVal(maxNodes, exist)}
		nodeNum++
	} else {
		return root
	}
	if nodeNum >= maxNodes {
		return root
	}
	queue = append(queue, root)

	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[i]
			if rand.Float64() > 0.5 {
				node.left = &treeNode{val: getVal(maxNodes, exist)}
				nodeNum++
				if nodeNum >= maxNodes {
					return root
				}
				queue = append(queue, node.left)
			}

			if rand.Float64() > 0.5 {
				node.right = &treeNode{val: getVal(maxNodes, exist)}
				nodeNum++
				if nodeNum >= maxNodes {
					return root
				}
				queue = append(queue, node.right)
			}
		}
		queue = queue[n:]
	}
	return root
}

func TestRebuildBinaryTree(t *testing.T) {
	rand.Seed(time.Now().Unix())

	for i := 0; i < 10; i++ {
		maxNodes := rand.Intn(64)
		exist := make(map[int]bool)
		root := newTree(maxNodes, exist)
		var pre, io []int
		preOrder(root, &pre)
		inOrder(root, &io)

		newRoot := reBuildBinaryTree(pre, io)

		var newPre, newIo []int
		preOrder(newRoot, &newPre)
		inOrder(newRoot, &newIo)

		for i, val := range pre {
			if val != newPre[i] {
				t.Error("error")
			}
		}
		if len(pre) != len(newPre) {
			t.Error("error")
		}
		fmt.Println("pre: ", pre)

		for i, val := range io {
			if val != newIo[i] {
				t.Error("error")
			}
		}
		if len(io) != len(newIo) {
			t.Error("error")
		}
		fmt.Println("io: ", io)

		time.Sleep(time.Second)
	}
}
