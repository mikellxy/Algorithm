package main

import "fmt"

type bitMap struct {
	arr []int64
}

func (b *bitMap) set(val int) {
	n := (val + 1) / 64
	if n > len(b.arr) {
		arr := make([]int64, n+1)
		copy(arr[:len(b.arr)], b.arr)
		b.arr = arr
	}

	offset := val % 64
	b.arr[n] |= 1 << uint(offset)
}

func (b *bitMap) has(val int) bool {
	n := (val + 1) / 64
	if n + 1 > len(b.arr) {
		return false
	}

	offset := val % 64
	return b.arr[n] & (1 << uint(offset)) != 0
}

func main() {
	bm := new(bitMap)
	bm.set(64)
	bm.set(65)
	fmt.Println(bm.arr)

	fmt.Println(bm.has(64))
	fmt.Println(bm.has(65))
	fmt.Println(bm.has(12000))
}
