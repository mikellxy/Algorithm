package heap

const (
	MaxHeap = 0 // 大顶堆
	MinHeap = 1 // 小顶堆
)

// 放入堆中的对象需要实现HeapItem接口
type Item interface {
	Score() int
}

type Config struct {
	Items []Item
	Type  int
}

type Heap struct {
	arr []Item
	len int
	hm  map[Item]int // 保存堆内元素当前在arr中的下标
	typ int
}

func NewHeap(config *Config) *Heap {
	n := len(config.Items)
	h := &Heap{
		arr: make([]Item, len(config.Items)+1),
		len: n,
		hm:  make(map[Item]int),
		typ: config.Type,
	}
	if n > 0 {
		// 数组的第一个元素不是堆的元素
		copy(h.arr[1:], config.Items[:])
		for i := 1; i <= h.len; i++ {
			h.hm[h.arr[i]] = i
		}
		// 建堆的时候，已存在的节点可能都不满足堆的性质，所以需要检查每一个节点
		h.build()
	}
	return h
}

func (h *Heap) GetLen() int {
	return h.len
}

func (h *Heap) GetType() int {
	return h.typ
}

func (h *Heap) build() {
	for i := h.len / 2; i >= 1; i -- {
		j := i
		for {
			lastIndex := h.heapify(j)
			if j == lastIndex {
				break
			}
			j = lastIndex
		}
	}
}

// 向堆内放入一个元素
func (h *Heap) Add(item Item) {
	h.arr = append(h.arr, item)
	h.len++
	h.addToMap(item, h.len)

	// 植入元素，只需要不断向上检查父节点进行堆化即可
	i := h.len / 2
	for i > 0 {
		lastIndex := h.heapify(i)
		if lastIndex == i {
			return
		}
		// 发生了交换，继续检查上一层的父节点
		i = i / 2
	}
}

// 堆化，先判断是大顶堆还是小顶堆
func (h *Heap) heapify(i int) int {
	if h.typ == MaxHeap {
		return h.maxheapify(i)
	}
	return h.minheapify(i)
}

func (h *Heap) minheapify(i int) int {
	// 最后一次跟父节点交换值的子节点的下标
	// 如果没有交换过则返回的最初父节点的坐标
	// 调用方可以知道此次堆化是否进行了元素交换
	lastIndex := i
	if 2*i <= h.len && h.arr[2*i].Score() < h.arr[lastIndex].Score() {
		lastIndex = 2 * i
	}
	if 2*i+1 <= h.len && h.arr[2*i+1].Score() < h.arr[lastIndex].Score() {
		lastIndex = 2*i + 1
	}
	if lastIndex != i {
		h.exchange(lastIndex, i)
	}
	return lastIndex
}

func (h *Heap) maxheapify(i int) int {
	lastIndex := i
	if 2*i <= h.len && h.arr[2*i].Score() > h.arr[lastIndex].Score() {
		lastIndex = 2 * i
	}
	if 2*i+1 <= h.len && h.arr[2*i+1].Score() > h.arr[lastIndex].Score() {
		lastIndex = 2*i + 1
	}
	if lastIndex != i {
		h.exchange(lastIndex, i)
	}
	return lastIndex
}

// 判断堆是否为空
func (h *Heap) Empty() bool {
	return h.len == 0
}

// 取出堆顶元素
func (h *Heap) Top() Item {
	if h.Empty() {
		return nil
	}
	ret := h.arr[1]
	h.deleteFromMap(ret)
	h.arr[1] = h.arr[h.len]
	h.addToMap(h.arr[1], 1)
	h.arr = h.arr[:h.len]
	h.len--

	// 自上而下堆化，因为之前已经是堆了，每次交换只会破坏一个子树
	i := 1
	for i <= h.len/2 {
		lastIndex := h.heapify(i)
		if lastIndex == i {
			break
		}
		// 发生了交换，继续堆化被破坏了性质的子树
		i = lastIndex
	}
	return ret
}

// 堆内某个元素的Score返回的值更新了的话
// 需要调用这个接口，进行必要的堆化
func (h *Heap) Update(item Item) {
	index, ok := h.hm[item]
	if !ok {
		return
	}

	// 对于下层，相当于删除了一个元素
	i := index
	for i <= h.len/2 {
		lastIndex := h.heapify(i)
		if lastIndex == i {
			break
		}
		i = lastIndex
	}

	// 对于上层，相当于插入了一个元素
	i = index / 2
	for i > 0 {
		lastIndex := h.heapify(i)
		if lastIndex == i {
			break
		}
		i = i / 2
	}
}

// 记录元素在底层数组中的下标
func (h *Heap) addToMap(item Item, index int) {
	h.hm[item] = index
}

// 删除元素在底层数组中的下标记录
func (h *Heap) deleteFromMap(item Item) {
	delete(h.hm, item)
}

// 交换两个节点的值
func (h *Heap) exchange(first int, second int) {
	h.arr[first], h.arr[second] = h.arr[second], h.arr[first]
	h.hm[h.arr[first]] = first
	h.hm[h.arr[second]] = second
}

func (h *Heap) Clear() {
	h.arr = h.arr[:1]
	h.len = 0
}
