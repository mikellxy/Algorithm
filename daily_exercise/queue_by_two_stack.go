package daily_exercise

import "errors"

type queueByTwoStack struct {
	s1  *stack
	s2  *stack
	len int
}

func (q *queueByTwoStack) push(item int) {
	q.s1.push(item)
	q.len++
}

func (q *queueByTwoStack) pop() (int, error) {
	if q.len == 0 {
		return 0, errors.New("the queue is empty")
	}

	if q.s2.len == 0 {
		n := q.s1.len
		for i := 0; i < n; i++ {
			val, _ := q.s1.Top()
			q.s2.push(val)
		}
	}

	item, _ := q.s2.Top()
	q.len--
	return item, nil
}

func newQueueByTwoStack() *queueByTwoStack {
	return &queueByTwoStack{s1: new(stack), s2: new(stack)}
}