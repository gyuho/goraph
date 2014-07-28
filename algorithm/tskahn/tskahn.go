package tskahn

import (
	"fmt"

	slice "github.com/gyuho/goraph/gosequence"
	"github.com/gyuho/goraph/graph/gs"
)

// TSKahn returns the topological sort; if it does not exist, it returns false.
func TSKahn(g *gs.Graph) (string, bool) {
	// vertices_with_no_incoming_edge
	niedges := slice.NewSequence()
	vertices := g.GetVertices()
	for _, vtx := range *vertices {
		if vtx == nil {
			continue
		}
		if vtx.(*gs.Vertex).InVertices.Len() == 0 {
			niedges.PushBack(vtx.(*gs.Vertex))
		}
	}

	result := slice.NewSequence()

	// while  S.length  !=  0
	for niedges.Len() != 0 {

		// remove a node n from S
		n := niedges.PopFront()

		// add n to the tail of L
		result.PushBack(n.(*gs.Vertex))
		// for each vertex m that comes out of n
		ovs := n.(*gs.Vertex).GetOutVertices()

		// for _, m := range *ovs {
		// once deleted, the original slice size decreases
		// since we use pointers
		// Make sure to use the copy, not the original!
		todelete := []interface{}{}
		for _, m := range *ovs {
			todelete = append(todelete, m)
		}

		for _, m := range todelete {
			g.DeleteEdge(n.(*gs.Vertex), m.(*gs.Vertex))
		}
		for _, m := range todelete {
			if m == nil {
				continue
			}
			// if m has no other incoming edges
			if m.(*gs.Vertex).InVertices.Len() == 0 {
				// add m to S
				niedges.PushBack(m.(*gs.Vertex))
			}
		}
	}

	if g.GetEdgesSize() != 0 {
		return "No Topological Sort (Not a DAG, there is a cycle)", false
	}

	var rs string
	for _, v := range *result {
		rs += fmt.Sprintf("%v → ", v.(*gs.Vertex).ID)
	}
	return rs[:len(rs)-5], true
}
