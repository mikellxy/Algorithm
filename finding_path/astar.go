package finding_path

import (
	"Algorithm/heap"
	"fmt"
	"math"
)

// 求解过程中，保存图上每个点的处理信息
type vertex struct {
	id   int // 顶点编号
	dist int // 从出发顶点到此顶点经过的距离
	x, y int // 顶点在地图上的坐标
	f    int // 估价函数的值, dist + 此顶点到终点的曼哈顿距离(abs(x - xt) + abs(y - yt))
}

// 实现优先级队列Item接口，按估值函数的值排序
func (v *vertex) Score() int {
	return v.f
}

// 图上的点
type point struct {
	x, y  int     // 图上点坐标
	sides []*side // 这个点到其他可达点的边的信息
}

// 给point添加边
func (p *point) addSide(s, t, w int) {
	p.sides = append(p.sides, &side{
		sid: s,
		tid: t,
		w:   w,
	})
}

type graph struct {
	v      int      // 顶点数量
	points []*point // 图上每个顶点到其所有可达顶点的边
}

func (g *graph) addSide(s, t, w int) {
	g.points[s].addSide(s, t, w)
}

func newGraph(v int) *graph {
	return &graph{
		v:      v,
		points: make([]*point, v),
	}
}

// 图上顶点之间的边
type side struct {
	sid, tid int // 边的两个顶点
	w        int // 边的权重，例如距离、红绿灯时间、拥堵状况等
}

type astar struct {
	graph  *graph
	vertex []*vertex
	v      int // 处理的顶点列表，下表为顶点编号
}

func newAstar(g *graph) *astar {
	a := &astar{
		graph:  g,
		v:      g.v,
		vertex: make([]*vertex, g.v),
	}
	for i, p := range g.points {
		a.vertex[i] = &vertex{
			id:   i,
			x:    p.x,
			y:    p.y,
			dist: -1,
		}
	}
	return a
}

func (a *astar) getPath(s, t int) {
	inHeap := make(map[int]bool)
	pathes := make([]int, a.v)
	h := heap.NewHeap(&heap.Config{
		Type: heap.MinHeap,
	})
	a.vertex[s].dist = 0
	h.Add(a.vertex[s])

	for h.GetLen() > 0 {
		vtx := h.Top().(*vertex)
		pt := a.graph.points[vtx.id]
		for _, s := range a.graph.points[vtx.id].sides {
			nextVtx := a.vertex[s.tid]
			nextPt := a.graph.points[nextVtx.id]
			dist := vtx.dist + s.w
			if nextVtx.dist == -1 || dist < nextVtx.dist {
				nextVtx.dist = dist
				nextVtx.f = ManHattan(pt.x, pt.y, nextPt.x, nextPt.y)
				if inHeap[nextVtx.id] {
					h.Update(nextVtx)
				} else {
					h.Add(nextVtx)
					inHeap[nextVtx.id] = true
				}
				pathes[nextVtx.id] = vtx.id
			}

			if nextVtx.id == t {
				h.Clear()
			}
		}
	}
}

func (a *astar) printPath(s, t int, pathes []int) {
	if s == t {
		fmt.Println(s)
		return
	}
	a.printPath(s, pathes[t], pathes)
	fmt.Println(t)
}

func ManHattan(xs, ys, xt, yt int) int {
	mht := math.Abs(float64(xs-xt)) + math.Abs(float64(ys-yt))
	return int(mht)
}
