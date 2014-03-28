// Package spd finds the shortest path using Dijkstra algorithm.
// It does not work with negative edges.
package spd

/*
Dijkstra(G, source, target)
	for each vertex v ∈ G.V
		v.d = ∞
		v.π = nil
	// this is already done
	// when instantiating the graph
	// and instead of InVertices
	// we can just create another slice
	// inside Graph (Prev)
	// in order not to modify the original graph

	source.d = 0

	// Min-Priority queue Q, keyed by their d values
	Q = G.V

	while Q ≠ ∅
		Min-Heapify(Q)
		u = Extract-Min(Q)
		if u.d = ∞
			break
		for each vertex v ∈ G.Adj[u]
			if v.d > u.d + w(u,v)
				v.d = u.d + w(u,v)
				v.π = u
*/

import (
	"container/heap"
	"fmt"

	"github.com/gyuho/goraph/graph/gsd"
	slice "github.com/gyuho/gosequence"
)

// SPD finds the shortest path from src to dst vertex.
func SPD(g *gsd.Graph, src, dst string) string {
	start := g.FindVertexByID(src)
	terminal := g.FindVertexByID(dst)

	start.StampD = 0

	// Min-Priority queue Q, keyed by their d values
	// Q = G.V
	// var minHeap VertexSlice
	minHeap := make(VertexSlice, 0)

	vertices := g.GetVertices()
	for _, vtx := range *vertices {
		if vtx == nil {
			continue
		}
		heap.Push(&minHeap, vtx.(*gsd.Vertex))
	}
	// Build-Min-Heap
	// first element with smallest timestamp
	// heap.Init(&minHeap)

	// while Q ≠ ∅
	for minHeap.Len() != 0 {
		/*
		   Min-Priority queue Q, keyed by their d values
		   u = Extract-Min(Q)
		   first one is the start vertex since we initialized it to 0

		   Reorder the vertex in the Queue
		   Min-Heapify(Q)
		   without this, the algorithm won't work

		   We need to Heapify here, for every loop
		*/
		heap.Init(&minHeap)
		u := minHeap.Pop().(*gsd.Vertex)

		if u.StampD == 9999999999 {
			break
		}

		ovs := u.GetOutVertices()
		for _, vtx := range *ovs {
			weights := g.GetEdgeWeight(u, vtx.(*gsd.Vertex))
			for _, wt := range weights {
				if vtx.(*gsd.Vertex).StampD > u.StampD+int64(wt) {
					vtx.(*gsd.Vertex).StampD = u.StampD + int64(wt)

					// v.π = u
					// update using Stack
					if vtx.(*gsd.Vertex).Prev.Len() == 0 {
						vtx.(*gsd.Vertex).Prev.PushBack(u)
					} else {
						ex := false
						ivs := vtx.(*gsd.Vertex).Prev
						for _, vs := range *ivs {
							// if fmt.Sprintf("%v", vs.(*gsd.Vertex).ID) == fmt.Sprintf("%v", u.ID) {
							if vs.(*gsd.Vertex) == u {
								ex = true
							}
						}

						if ex == false {
							vtx.(*gsd.Vertex).Prev.PushBack(u)
						}
					}

					// not a good place to Heapify
					// because we are inside the if-condition
					// heap.Init(&minHeap)
				}
			}
		}
	}

	result := slice.NewSequence()
	TrackSPD(g, start, terminal, result)

	s := ""
	for _, v := range *result {
		s += fmt.Sprintf("%v(=%v) → ", v.(*gsd.Vertex).ID, v.(*gsd.Vertex).StampD)
	}
	return s[:len(s)-5]
}

// TrackSPD recursively backtracks the shortest path.
// It recursively adds the Prev vertex with the biggest StampD.
// The recursion ends when we reach the start vertex.
func TrackSPD(g *gsd.Graph, start, target *gsd.Vertex, result *slice.Sequence) {
	// Add target first
	if result.Len() == 0 {
		result.PushFront(target)
	}

	// End recursion when we have NON-connected graph (len = 0)
	// End recursion when we get to Source that has no Prev
	if target.Prev.Len() == 0 {
		return
	}

	// find the Prev vertex with the biggest StampD
	ps := target.Prev
	bg := ps.Front()
	for _, vtx := range *ps {
		if bg.(*gsd.Vertex).StampD < vtx.(*gsd.Vertex).StampD {
			bg = vtx
		}
	}
	// now we know what is the biggest one

	// if it does not exist in the result, add it to result
	exist := false
	for _, vtx := range *result {
		// if fmt.Sprintf("%v", .ID) == fmt.Sprintf("%v", bg.(*gsd.Vertex).ID) {
		if vtx.(*gsd.Vertex) == bg.(*gsd.Vertex) {
			exist = true
		}
	}

	if exist == false {
		result.PushFront(bg)
		TrackSPD(g, start, bg.(*gsd.Vertex), result)
	}
}

// Min-Heap's first element is the minimum
type VertexSlice []*gsd.Vertex

func (vs VertexSlice) Len() int {
	return len(vs)
}

func (vs VertexSlice) Less(i, j int) bool {
	return vs[i].StampD < vs[j].StampD
}

func (vs VertexSlice) Swap(i, j int) {
	vs[i], vs[j] = vs[j], vs[i]
}

func (vs *VertexSlice) Push(x interface{}) {
	*vs = append(*vs, x.(*gsd.Vertex))
}

func (vs *VertexSlice) Pop() interface{} {
	old := *vs
	n := len(old)
	x := old[0]
	*vs = old[1:n]
	return x
}

func spd(g *gsd.Graph, src, dst string) string {
	start := g.FindVertexByID(src)
	terminal := g.FindVertexByID(dst)

	start.StampD = 0

	// Min-Priority queue Q, keyed by their d values
	// Q = G.V
	// var minHeap VertexSlice
	minHeap := make(VertexSlice, 0)

	vertices := g.GetVertices()
	for _, vtx := range *vertices {
		if vtx == nil {
			continue
		}
		heap.Push(&minHeap, vtx.(*gsd.Vertex))
	}
	// Build-Min-Heap
	// first element with smallest timestamp
	// heap.Init(&minHeap)

	// while Q ≠ ∅
	for minHeap.Len() != 0 {
		/*
		   Min-Priority queue Q, keyed by their d values
		   u = Extract-Min(Q)
		   first one is the start vertex since we initialized it to 0

		   Reorder the vertex in the Queue
		   Min-Heapify(Q)
		   without this, the algorithm won't work

		   We need to Heapify here, for every loop
		*/
		heap.Init(&minHeap)
		u := minHeap.Pop().(*gsd.Vertex)

		if u.StampD == 9999999999 {
			break
		}

		ovs := u.GetOutVertices()
		for _, vtx := range *ovs {
			weights := g.GetEdgeWeight(u, vtx.(*gsd.Vertex))
			for _, wt := range weights {
				if vtx.(*gsd.Vertex).StampD > u.StampD+int64(wt) {
					vtx.(*gsd.Vertex).StampD = u.StampD + int64(wt)

					// v.π = u
					// update using Stack
					if vtx.(*gsd.Vertex).Prev.Len() == 0 {
						vtx.(*gsd.Vertex).Prev.PushBack(u)
					} else {
						ex := false
						ivs := vtx.(*gsd.Vertex).Prev
						for _, vs := range *ivs {
							// if fmt.Sprintf("%v", vs.(*gsd.Vertex).ID) == fmt.Sprintf("%v", u.ID) {
							if vs.(*gsd.Vertex) == u {
								ex = true
							}
						}

						if ex == false {
							vtx.(*gsd.Vertex).Prev.PushBack(u)
						}
					}

					// not a good place to Heapify
					// because we are inside the if-condition
					// heap.Init(&minHeap)
				}
			}
		}
	}

	result := slice.NewSequence()
	TrackSPD(g, start, terminal, result)

	s := ""
	for _, v := range *result {
		s += fmt.Sprintf("%v → ", v.(*gsd.Vertex).ID)
	}
	return s[:len(s)-5]
}
