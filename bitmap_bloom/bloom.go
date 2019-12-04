package main

import "fmt"

type hashFunc func(val, l int) (int, int)

type bloom struct {
	arr       []int64
	hashFuncs []hashFunc
}

func (b *bloom) set(val int) {
	l := len(b.arr)

	for _, hf := range b.hashFuncs {
		n, offset := hf(val, l)
		b.arr[n] |= 1 << uint(offset)
	}
}

func (b *bloom) has(val int) bool {
	l := len(b.arr)

	for _, hf := range b.hashFuncs {
		n, offset := hf(val, l)
		if b.arr[n] & (1 << uint(offset)) == 0 {
			return false
		}
	}
	return true
}

func hf1(val, l int) (int, int) {
	n := val / 64
	n = n % l

	return n, val / 64
}

func hf2(val, l int) (int, int) {
	n := 2 * (val + 1) / 64
	n = n % l

	return n, 2 * val / 64
}

func hf3(val, l int) (int, int) {
	n := 3 * (val + 4) / 64
	n = n % l

	return n, 3 * val / 64
}

func main() {
	bl := &bloom{
		arr:       make([]int64, 10),
		hashFuncs: []hashFunc{hf1, hf2, hf3},
	}

	bl.set(191)
	bl.set(728)

	fmt.Println(bl.has(191))
	fmt.Println(bl.has(782))
	fmt.Println(bl.has(999))
}
