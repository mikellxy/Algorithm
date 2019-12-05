package daily_exercise

func handler(root1 *treeNode, root2 *treeNode) bool {
	if root2 == nil {
		return true
	}
	if root1 == nil {
		return false
	}
	if root1.val != root2.val {
		return false
	}
	return handler(root1.left, root2.left) && handler(root1.right, root2.right)
}

func isSubStruct(root1 *treeNode, root2 *treeNode) bool {
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.val == root2.val {
		if handler(root1, root2) {
			return true
		}
	}
	return isSubStruct(root1.left, root2) || isSubStruct(root1.right, root2)
}

