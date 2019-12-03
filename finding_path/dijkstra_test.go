package finding_path

import "testing"

func TestShortestRoute(t *testing.T) {
	graph := newGraph(3)
	graph.addEdge(0, 1, 1)
	graph.addEdge(1,2,2)
	graph.addEdge(0, 2, 3)

	h := newHandler(graph)
	h.FindBestRoute(0, 2)
}
