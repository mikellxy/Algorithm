package shortest_route

import (
	"Algorithm/heap"
	"fmt"
)

// 有向有权图
type Graph struct {
	v   int       // 顶点个数
	adj [][]*Edge // 临界表，为每个顶点保存一个他的边的列表
}

func newGraph(v int) *Graph {
	return &Graph{
		v:   v,
		adj: make([][]*Edge, v),
	}
}

func (g *Graph) addEdge(s, t, w int) {
	g.adj[s] = append(g.adj[s], &Edge{
		sid: s,
		tid: t,
		w: w,
	})
}

// 从顶点s到顶点t的路径
type Edge struct {
	sid int // 边的起点编号
	tid int // 边的终点编号
	w   int // 路径的权重
}

func (e *Edge) String() string {
	return fmt.Sprintf("from <%d> to <%d> = %d", e.sid, e.tid, e.w)
}

type Vertex struct {
	id   int
	dist int
}

// 实现heap.Item接口
func (v *Vertex) Score() int {
	return v.dist
}

type handler struct {
	graph    *Graph
	vertexes []*Vertex
	heap     *heap.Heap
	v        int
}

func newHandler(graph *Graph) *handler {
	return &handler{
		graph:    graph,
		v:        graph.v,
		vertexes: make([]*Vertex, graph.v),
	}

}

func (h *handler) Init() {
	for i := 0; i < h.v; i++ {
		h.vertexes[i] = &Vertex{id: i}
	}
	config := &heap.Config{
		Type: heap.MinHeap,
	}
	h.heap = heap.NewHeap(config)
}

func (h *handler) FindBestRoute(s, t int) {
	h.Init()

	inQueue := make(map[*Vertex]bool)
	route := make([]int, h.graph.v)
	h.vertexes[s].dist = 0
	for i, v := range h.vertexes {
		if i != s {
			v.dist = -1
		}
	}

	h.heap.Add(h.vertexes[s])

	for !h.heap.Empty() {
		// 取出更新过最短路径的vertex
		vertex := h.heap.Top().(*Vertex)
		if vertex.id == t {
			break
		}
		for _, e := range h.graph.adj[vertex.id] {
			// 遍历vertex可到的所有nextVertex
			nextVertex := h.vertexes[e.tid]
			dist := vertex.dist + e.w
			if nextVertex.dist == -1 || dist < nextVertex.dist {
				nextVertex.dist = dist
				route[nextVertex.id] = vertex.id
				if _, ok := inQueue[nextVertex]; ok {
					h.heap.Update(nextVertex)
				} else {
					h.heap.Add(nextVertex)
					inQueue[nextVertex] = true
				}
			}
		}
	}

	h.PrintRoute(s, t, route)
}

func (h *handler) PrintRoute(s, t int, route []int) {
	if s == t {
		fmt.Println(s)
		return
	}
	h.PrintRoute(s, route[t], route)
	fmt.Println(t)
}