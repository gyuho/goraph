// Package prim implements Prim's Minimum Spanning Tree algorithm.
package prim

import (
	"container/heap"
	"strconv"

	"github.com/gyuho/goraph/graph/gsd"
)

/*
Prim Algorithm uses Min-Priority Queue
like Dijkstra algorithm shortest path algorithm.
(https://github.com/gyuho/goraph/tree/master/algorithm/spd)

Tree starts from an arbitrary root vertex  r
The tree grows until it spans all the vertices in V

All vertices that are not in the tree reside
in a min-priority queue Q based on the key attribute

v.key  is the minimum weight of any edge
from v to a vertex in the tree
v.key  is ∞  if there is no such edge

CLRS p.634

for  each vertex  u ∈ G.V
	u.key = ∞
	u.π = NIL

r.key = 0
Q = G.V

while  Q ≠ ø
	u = Extract-Min(Q)
	for each v ∈ G.Adj[u]
		if v ∈ Q  and  v.key > w(u, v)
			v.π = u
			v.key = w(u, v)
*/

// MST implements Prim's Minimum Spanning Tree algorithm.
// It returns the Minimum Spanning Tree in DOT file format
// and the total weights of Minimum Spanning Tree.
func MST(g *gsd.Graph) (string, float64) {
	// for  each vertex  u ∈ G.V
	// 		u.key = ∞
	// 		u.π = NIL
	vertices := g.GetVertices()
	for _, vtx := range *vertices {
		vtx.(*gsd.Vertex).StampF = 9999999999
		vtx.(*gsd.Vertex).Prev.Init()
	}
	// When instantiated, gsd already sets with:
	// 		StampF:      9999999999,
	// 		Prev:        slice.NewSequence(),

	// r.key = 0
	root := (*vertices)[0].(*gsd.Vertex)
	root.StampF = 0 // use StampF as key

	// Q = G.V
	// var minHeap VertexSlice
	minHeap := make(VertexSlice, 0)
	for _, vtx := range *vertices {
		if vtx == nil {
			continue
		}
		heap.Push(&minHeap, vtx.(*gsd.Vertex))
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
		u := minHeap.Pop().(*gsd.Vertex)

		// for each v ∈ G.Adj[u]
		// traverse the OutVertices of Vertex u
		ovs := u.GetOutVertices()
		for _, vtx := range *ovs {
			// if v ∈ Q  and  v.key > w(u, v)
			if minHeap.Contains(vtx.(*gsd.Vertex)) && vtx.(*gsd.Vertex).StampF > float64ToInt64(g.GetEdgeWeight(u, vtx.(*gsd.Vertex))[0]) {
				// for each OutVertex v, v.Prev = u
				// v.π = u
				// update using Stack
				vtx.(*gsd.Vertex).Prev.Init()
				if vtx.(*gsd.Vertex).Prev.Len() != 0 {
					panic("Not Emptied")
				}
				vtx.(*gsd.Vertex).Prev.PushBack(u)

				// Do NOT use this! This generates duplicate Prev!
				// We need to overwrite the Prev using Init method
				/*
					if vtx.(*gsd.Vertex).Prev.Len() == 0 {
						// fmt.Println(vtx.(*gsd.Vertex).ID, u.ID)
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
				*/

				// for each OutVertex v, v.StampF = Weight(u, v)
				// v.key = w(u, v)
				vtx.(*gsd.Vertex).StampF = float64ToInt64(g.GetEdgeWeight(u, vtx.(*gsd.Vertex))[0])
			}
		}
	}

	result := "graph PrimMST {" + "\n"
	var total float64
	for _, mvt := range *g.GetVertices() {
		if mvt.(*gsd.Vertex).Prev.Len() != 0 {
			for _, vt := range *mvt.(*gsd.Vertex).Prev {
				wt := g.GetEdgeWeight(mvt.(*gsd.Vertex), vt.(*gsd.Vertex))[0]
				wts := strconv.FormatFloat(wt, 'f', -1, 64)
				result += "\t" + mvt.(*gsd.Vertex).ID + " -- " + vt.(*gsd.Vertex).ID + " [label=" + wts + ", color=blue]" + "\n"
				total += wt
			}
		}
	}
	for _, edge := range *g.GetEdges() {
		result += "\t" + edge.(*gsd.Edge).Src.ID + " -- " + edge.(*gsd.Edge).Dst.ID + "\n"
	}
	result += "}"
	return result, total
}

// Min-Heap's first element is the minimum
type VertexSlice []*gsd.Vertex

func (vs VertexSlice) Len() int {
	return len(vs)
}

func (vs VertexSlice) Less(i, j int) bool {
	return vs[i].StampF < vs[j].StampF
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

// Contains returns true if vtx exists in the VertexSlice.
func (vs VertexSlice) Contains(vtx *gsd.Vertex) bool {
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
