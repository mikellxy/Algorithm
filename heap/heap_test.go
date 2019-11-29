package heap

import (
	"fmt"
	"testing"
)

type Value struct {
	val int
}

func (v *Value) Score() int {
	return v.val
}

func (v *Value) String() string {
	return fmt.Sprintf("%d", v.val)
}

func newValue(val int) *Value {
	return &Value{val: val}
}

func TestHeap(t *testing.T) {
	v := newValue(2)
	v2 := newValue(4)
	config := &Config{
		Type: MaxHeap,
		Items: []Item{
			newValue(3),
			newValue(1),
			v,
			newValue(5),
			v2,
		},
	}
	h := NewHeap(config)

	v.val = 20
	h.Update(v)
	v2.val = 40
	h.Update(v2)

	result := []int{40, 20, 5, 3, 1}
	i := 0
	for !h.Empty() {
		r := h.Top()
		if r.Score() != result[i] {
			t.Fatalf("第%d个top value是%d, 应该是%d", i+1, r.Score(), result[i])
		}
		i++
	}
}
