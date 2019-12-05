package daily_exercise

func mergeLinkedList(r1 *node, r2 *node) *node {
	if r1 == nil {
		return r2
	}
	if r2 == nil {
		return r1
	}

	if r1.val <= r2.val {
		r1.next = mergeLinkedList(r1.next, r2)
		return r1
	} else {
		r2.next = mergeLinkedList(r1, r2.next)
		return r2
	}
}
