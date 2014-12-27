package spd

import (
	"container/heap"
	"fmt"

	slice "github.com/gyuho/goraph/gosequence"
	"github.com/gyuho/goraph/graph/gs"
)

// SPD finds the shortest path from src to dst vertex.
func SPD(g *gs.Graph, src, dst *gs.Vertex) string {

	src.StampD = 0

	// Min-Priority queue Q, keyed by their d values
	// Q = G.V
	// var minHeap VertexSlice
	minHeap := make(VertexSlice, 0)

	vertices := g.GetVertices()
	for _, vtx := range *vertices {
		if vtx == nil {
			continue
		}
		heap.Push(&minHeap, vtx.(*gs.Vertex))
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
		u := minHeap.Pop().(*gs.Vertex)

		if u.StampD == 9999999999 {
			break
		}

		ovs := u.GetOutVertices()
		for _, vtx := range *ovs {
			//
			// weights := g.GetEdgeWeight(u, vtx.(*gs.Vertex))
			// for _, wt := range weights {
			//
			wt := g.GetEdgeWeight(u, vtx.(*gs.Vertex))
			if vtx.(*gs.Vertex).StampD > u.StampD+int64(wt) {
				vtx.(*gs.Vertex).StampD = u.StampD + int64(wt)

				// v.π = u
				// update using Stack
				if vtx.(*gs.Vertex).Prev.Len() == 0 {
					vtx.(*gs.Vertex).Prev.PushBack(u)
				} else {
					ex := false
					ivs := vtx.(*gs.Vertex).Prev
					for _, vs := range *ivs {
						// if fmt.Sprintf("%v", vs.(*gs.Vertex).ID) == fmt.Sprintf("%v", u.ID) {
						if vs.(*gs.Vertex) == u {
							ex = true
						}
					}

					if ex == false {
						vtx.(*gs.Vertex).Prev.PushBack(u)
					}
				}

				// not a good place to Heapify
				// because we are inside the if-condition
				// heap.Init(&minHeap)
			}
		}
	}

	result := slice.NewSequence()
	TrackSPD(g, src, dst, result)

	s := ""
	for _, v := range *result {
		s += fmt.Sprintf("%v(=%v) → ", v.(*gs.Vertex).ID, v.(*gs.Vertex).StampD)
	}
	return s[:len(s)-5]
}

// TrackSPD recursively backtracks the shortest path.
// It recursively adds the Prev vertex with the biggest StampD.
// The recursion ends when we reach the start vertex.
func TrackSPD(g *gs.Graph, src, dst *gs.Vertex, result *slice.Sequence) {
	// Add target first
	if result.Len() == 0 {
		result.PushFront(dst)
	}

	// End recursion when we have NON-connected graph (len = 0)
	// End recursion when we get to Source that has no Prev
	if dst.Prev.Len() == 0 {
		return
	}

	// find the Prev vertex with the biggest StampD
	ps := dst.Prev
	bg := ps.Front()
	for _, vtx := range *ps {
		if bg.(*gs.Vertex).StampD < vtx.(*gs.Vertex).StampD {
			bg = vtx
		}
	}
	// now we know what is the biggest one

	// if it does not exist in the result, add it to result
	exist := false
	for _, vtx := range *result {
		// if fmt.Sprintf("%v", .ID) == fmt.Sprintf("%v", bg.(*gs.Vertex).ID) {
		if vtx.(*gs.Vertex) == bg.(*gs.Vertex) {
			exist = true
		}
	}

	if exist == false {
		result.PushFront(bg)
		TrackSPD(g, src, bg.(*gs.Vertex), result)
	}
}

// VertexSlice contains vertices.
// Min-Heap's first element is the minimum
type VertexSlice []*gs.Vertex

func (vs VertexSlice) Len() int {
	return len(vs)
}

func (vs VertexSlice) Less(i, j int) bool {
	return vs[i].StampD < vs[j].StampD
}

func (vs VertexSlice) Swap(i, j int) {
	vs[i], vs[j] = vs[j], vs[i]
}

// Push appends an element to VertexSlice.
func (vs *VertexSlice) Push(x interface{}) {
	*vs = append(*vs, x.(*gs.Vertex))
}

// Pop removes the first element from VertexSlice.
func (vs *VertexSlice) Pop() interface{} {
	old := *vs
	n := len(old)
	x := old[0]
	*vs = old[1:n]
	return x
}

func spd(g *gs.Graph, src, dst *gs.Vertex) string {

	src.StampD = 0

	// Min-Priority queue Q, keyed by their d values
	// Q = G.V
	// var minHeap VertexSlice
	minHeap := make(VertexSlice, 0)

	vertices := g.GetVertices()
	for _, vtx := range *vertices {
		if vtx == nil {
			continue
		}
		heap.Push(&minHeap, vtx.(*gs.Vertex))
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
		u := minHeap.Pop().(*gs.Vertex)

		if u.StampD == 9999999999 {
			break
		}

		ovs := u.GetOutVertices()
		for _, vtx := range *ovs {
			wt := g.GetEdgeWeight(u, vtx.(*gs.Vertex))
			if vtx.(*gs.Vertex).StampD > u.StampD+int64(wt) {
				vtx.(*gs.Vertex).StampD = u.StampD + int64(wt)

				// v.π = u
				// update using Stack
				if vtx.(*gs.Vertex).Prev.Len() == 0 {
					vtx.(*gs.Vertex).Prev.PushBack(u)
				} else {
					ex := false
					ivs := vtx.(*gs.Vertex).Prev
					for _, vs := range *ivs {
						// if fmt.Sprintf("%v", vs.(*gs.Vertex).ID) == fmt.Sprintf("%v", u.ID) {
						if vs.(*gs.Vertex) == u {
							ex = true
						}
					}

					if ex == false {
						vtx.(*gs.Vertex).Prev.PushBack(u)
					}
				}

				// not a good place to Heapify
				// because we are inside the if-condition
				// heap.Init(&minHeap)
			}
		}
	}

	result := slice.NewSequence()
	TrackSPD(g, src, dst, result)

	var rs string
	for _, v := range *result {
		rs += fmt.Sprintf("%v → ", v.(*gs.Vertex).ID)
	}
	return rs[:len(rs)-5]
}
