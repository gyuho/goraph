package bfs

import (
	"fmt"

	slice "github.com/gyuho/goraph/gosequence"
	"github.com/gyuho/goraph/graph/gs"
)

// BFS does Breadth First Search and return the result in visited order.
func BFS(g *gs.Graph, src *gs.Vertex) string {
	src.Color = "gray"
	src.StampD = 0

	queue := slice.NewSequence() // Q = ∅
	queue.PushBack(src)          // ENQUEUE(Q, s)

	s := ""
	for queue.Len() != 0 {
		u := queue.PopFront().(*gs.Vertex)
		ovs := u.GetOutVertices()

		for _, vtx := range *ovs {
			if vtx == nil {
				continue
			}
			vt := vtx.(*gs.Vertex)
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
func Path(g *gs.Graph, src, end *gs.Vertex) bool {
	if src == nil || end == nil {
		panic("Wrong Vertex Passed!")
	}
	src.Color = "gray"

	queue := slice.NewSequence() // Q = ∅
	queue.PushBack(src)          // ENQUEUE(Q, s)

	for queue.Len() != 0 {
		u := queue.PopFront().(*gs.Vertex)
		ovs := u.GetOutVertices()
		for _, vtx := range *ovs {
			if vtx == nil {
				continue
			}
			vt := vtx.(*gs.Vertex)
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
