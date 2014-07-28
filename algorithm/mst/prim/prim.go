package prim

import (
	"container/heap"
	"log"

	"github.com/gyuho/goraph/graph/gs"
)

// MST implements Prim's Minimum Spanning Tree algorithm.
// It returns the total weights of Minimum Spanning Tree.
func MST(g *gs.Graph) float64 {
	// for  each vertex  u ∈ G.V
	// 		u.key = ∞
	// 		u.π = NIL
	vertices := g.GetVertices()
	for _, vtx := range *vertices {
		vtx.(*gs.Vertex).StampF = 9999999999
		vtx.(*gs.Vertex).Prev.Init()
	}
	// When instantiated, gs already sets with:
	// 		StampF:      9999999999,
	// 		Prev:        slice.NewSequence(),

	// r.key = 0
	root := (*vertices)[0].(*gs.Vertex)
	root.StampF = 0 // use StampF as key

	// Q = G.V
	// var minHeap VertexSlice
	minHeap := make(VertexSlice, 0)
	for _, vtx := range *vertices {
		if vtx == nil {
			continue
		}
		heap.Push(&minHeap, vtx.(*gs.Vertex))
	}
	// Build-Min-Heap
	// first element with smallest timestamp
	// heap.Init(&minHeap)

	// Min-Priority queue Q(minHeap), keyed by their StampF values

	// heap.Init(&minHeap)

	/*
	   while  Q ≠ ø
	   	u = Extract-Min(Q)
	   	for each v ∈ G.Adj[u]
	   		if v ∈ Q  and  v.key > w(u, v)
	   			v.π = u
	   			v.key = w(u, v)
	*/
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

		// u = Extract-Min(Q)
		// Extract the Vertex with Minimum key(StampF)
		u := minHeap.Pop().(*gs.Vertex)

		// for each v ∈ G.Adj[u]
		// traverse the OutVertices of Vertex u
		ovs := u.GetOutVertices()
		for _, vtx := range *ovs {
			// if v ∈ Q  and  v.key > w(u, v)
			if minHeap.Contains(vtx.(*gs.Vertex)) &&
				vtx.(*gs.Vertex).StampF > float64ToInt64(g.GetEdgeWeight(u, vtx.(*gs.Vertex))) {
				// for each OutVertex v, v.Prev = u
				// v.π = u
				// update using Stack
				vtx.(*gs.Vertex).Prev.Init()
				if vtx.(*gs.Vertex).Prev.Len() != 0 {
					log.Fatal("Not Emptied")
				}
				vtx.(*gs.Vertex).Prev.PushBack(u)

				// Do NOT use this! This generates duplicate Prev!
				// We need to overwrite the Prev using Init method
				/*
					if vtx.(*gs.Vertex).Prev.Len() == 0 {
						// fmt.Println(vtx.(*gs.Vertex).ID, u.ID)
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
				*/

				// for each OutVertex v, v.StampF = Weight(u, v)
				// v.key = w(u, v)
				vtx.(*gs.Vertex).StampF = float64ToInt64(g.GetEdgeWeight(u, vtx.(*gs.Vertex)))
			}
		}
	}

	var total float64
	for _, mvt := range *g.GetVertices() {
		if mvt.(*gs.Vertex).Prev.Len() != 0 {
			for _, vt := range *mvt.(*gs.Vertex).Prev {
				wt := g.GetEdgeWeight(mvt.(*gs.Vertex), vt.(*gs.Vertex))
				total += wt
			}
		}
	}

	return total
}

// VertexSlice implements the minimum heap.
// Min-Heap's first element is the minimum
type VertexSlice []*gs.Vertex

func (vs VertexSlice) Len() int {
	return len(vs)
}

func (vs VertexSlice) Less(i, j int) bool {
	return vs[i].StampF < vs[j].StampF
}

func (vs VertexSlice) Swap(i, j int) {
	vs[i], vs[j] = vs[j], vs[i]
}

// Push appends a Vertex to VertexSlice.
func (vs *VertexSlice) Push(x interface{}) {
	*vs = append(*vs, x.(*gs.Vertex))
}

// Pop returns the first element and removes it from VertexSlice.
func (vs *VertexSlice) Pop() interface{} {
	old := *vs
	n := len(old)
	x := old[0]
	*vs = old[1:n]
	return x
}

// Contains returns true if vtx exists in the VertexSlice.
func (vs VertexSlice) Contains(vtx *gs.Vertex) bool {
	for _, val := range vs {
		if val == vtx {
			return true
		}
	}
	return false
}

// float64ToInt64 converts float64 to int64.
func float64ToInt64(num float64) int64 {
	return int64(num)
}
