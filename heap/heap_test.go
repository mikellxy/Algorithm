package heap

import (
	"testing"
)

type Value struct {
	val int
}

func (v *Value) Score() int {
	return v.val
}

func newValue(val int) *Value {
	return &Value{val: val}
}

func TestHeap(t *testing.T) {
	v := newValue(2)
	config := &Config{
		Type: MaxHeap,
		Items: []Item{
			newValue(3),
			newValue(1),
			v,
			newValue(5),
			newValue(4),
		},
	}
	h := NewHeap(config)
	v.val = 20
	h.Update(v)

	result := []int{20, 5, 4, 3, 1}
	i := 0
	for !h.Empty() {
		r := h.Top()
		if r.Score() != result[i] {
			t.Fatalf("第%d个top value是%d, 应该是%d", i+1, r.Score(), result[i])
		}
		i++
	}
}
