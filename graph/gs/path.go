package gs

import (
	"log"

	slice "github.com/gyuho/goraph/gosequence"
)

// Path returns true if there is a path between two Vertices.
func (g *Graph) Path(src, end *Vertex) bool {
	if src == nil || end == nil {
		log.Fatal("Wrong Vertex Passed!")
	}

	src.Color = "gray"

	queue := slice.NewSequence() // Q = ∅
	queue.PushBack(src)          // ENQUEUE(Q, s)

	for queue.Len() != 0 {
		u := queue.PopFront().(*Vertex)
		ovs := u.GetOutVertices()
		for _, vtx := range *ovs {
			if vtx == nil {
				continue
			}
			vt := vtx.(*Vertex)
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

// PathRecur returns true if there is a path between two Vertices.
func (g *Graph) PathRecur(src, end *Vertex) bool {
	if src == nil || end == nil {
		log.Fatal("Wrong Vertex Passed!")
	}

	rb := false

	if src.Color == "white" {
		g.pathRecur(src, end, &rb)
	}

	return rb
}

// pathRecur returns true if there is a path between two Vertices.
func (g *Graph) pathRecur(src, end *Vertex, rb *bool) {
	if src == end {
		*rb = true
		return
	}

	src.Color = "gray"

	ovs := src.GetOutVertices()
	for _, vtx := range *ovs {
		if vtx == nil {
			continue
		}
		vt := vtx.(*Vertex)
		if vt.Color == "white" {
			g.pathRecur(vt, end, rb)
		}
	}
	src.Color = "black"
}
