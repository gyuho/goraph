// Package bfs implements Breadth First Search algorithm (or BFS).
package bfs

/*
Pseudo Code CLRS page 595

BFS(G, s)
	for each vertex u ∈ g.V - {s}
		u.color = WHITE
		u.d = ∞
		u.π = NIL
	// this is already done
	// when instantiating the graph
	// and instead of InVertices
	// we can just create another slice
	// inside Graph (Prev)
	// in order not to modify the original graph

	s.color = GRAY
	s.d = 0
	s.π = NIL
	Q = ∅

	ENQUEUE(Q, s)

	while Q ≠ ∅
		u = DEQUEUE(Q)
		for each v ∈ g.Adj[u]
			if v.color == WHITE
				v.color = GRAY
				v.d = u.d + 1
				v.π = u
				ENQUEUE(Q, v)
		u.color = BLACK
*/

import (
	"fmt"

	"github.com/gyuho/goraph/graph/gsd"
	slice "github.com/gyuho/gosequence"
)

// type Graph gsd.Graph
// type Vertex gsd.Vertex

// BFS does Breadth First Search and return the result in visited order.
func BFS(g *gsd.Graph, src *gsd.Vertex) string {
	src.Color = "gray"
	src.StampD = 0

	queue := slice.NewSequence() // Q = ∅
	queue.PushBack(src)          // ENQUEUE(Q, s)

	s := ""
	for queue.Len() != 0 {
		u := queue.PopFront().(*gsd.Vertex)
		ovs := u.GetOutVertices()

		for _, vtx := range *ovs {
			if vtx == nil {
				continue
			}
			vt := vtx.(*gsd.Vertex)
			if vt.Color == "white" {
				vt.Color = "gray"
				vt.StampD = u.StampD + 1

				// v.π = u
				vt.Prev.PushBack(u)
				queue.PushBack(vt)
			}
		}
		u.Color = "black"
		s += fmt.Sprintf("%v → ", u.ID)
		// s += fmt.Sprintf("%v(timestamp: %v) → ", u.ID, u.StampD)
	}

	return s[:len(s)-5]
}

// Path returns true if there is a path between two Vertices.
func Path(g *gsd.Graph, src, end *gsd.Vertex) bool {
	if src == nil || end == nil {
		panic("Wrong Vertex Passed!")
	}
	src.Color = "gray"

	queue := slice.NewSequence() // Q = ∅
	queue.PushBack(src)          // ENQUEUE(Q, s)

	for queue.Len() != 0 {
		u := queue.PopFront().(*gsd.Vertex)
		ovs := u.GetOutVertices()
		for _, vtx := range *ovs {
			if vtx == nil {
				continue
			}
			vt := vtx.(*gsd.Vertex)
			if vt == end {
				return true
			}
			if vt.Color == "white" {
				vt.Color = "gray"
				queue.PushBack(vt)
			}
		}
		u.Color = "black"
	}
	return false
}
