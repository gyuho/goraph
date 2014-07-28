package bfs

import (
	"fmt"

	slice "github.com/gyuho/goraph/gosequence"
	"github.com/gyuho/goraph/graph/gs"
)

// BFS does Breadth First Search and return the result in visited order.
// BFS traverses graphs in an arbitrary order.
func BFS(g *gs.Graph, src *gs.Vertex) string {
	src.Color = "gray"
	src.StampD = 0

	queue := slice.NewSequence() // Q = ∅
	queue.PushBack(src)          // ENQUEUE(Q, s)

	var rs string
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
		rs += fmt.Sprintf("%v → ", u.ID)
		// rs += fmt.Sprintf("%v(timestamp: %v) → ", u.ID, u.StampD)
	}

	return rs[:len(rs)-5]
}
