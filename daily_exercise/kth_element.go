package daily_exercise

func getKthElement(root *node, k int) *node {
	p1, p2 := root, root

	var i int
	for i = 0; i < k-1 ; i++ {
		if p2 == nil || p2.next == nil {
			break
		}
		p2 = p2.next
	}
	if i < k-1 {
		return nil
	}

	for {
		if p2.next == nil {
			return p1
		}
		p1 = p1.next
		p2 = p2.next
	}
}
