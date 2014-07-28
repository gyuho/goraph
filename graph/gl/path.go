package gl

import (
	"container/list"
	"log"
)

// Path returns true if there is a path between two Vertices.
func (g *Graph) Path(src, end *Vertex) bool {
	if src == nil || end == nil {
		log.Fatal("Wrong Vertex Passed!")
	}
	src.Color = "gray"

	queue := list.New() // Q = ∅
	queue.PushBack(src) // ENQUEUE(Q, s)

	for queue.Len() != 0 {
		u := queue.Front()
		ovs := u.Value.(*Vertex).GetOutVertices()
		for vtx := ovs.Front(); vtx != nil; vtx = vtx.Next() {
			// for _, vtx := range *ovs {
			if vtx == nil {
				continue
			}
			vt := vtx.Value.(*Vertex)
			if vt == end {
				return true
			}
			if vt.Color == "white" {
				vt.Color = "gray"
				queue.PushBack(vt)
			}
		}
		u.Value.(*Vertex).Color = "black"
		queue.Remove(u)
	}
	return false
}
