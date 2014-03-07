// Package tskahn finds topological sort
// based on algorithm by Arthur Kahn(1962).
// (Reference: http://dl.acm.org/citation.cfm?doid=368996.369025)
package tskahn

/*
DAG_Kahn(G)
S as a set of nodes with no incoming edges
while  S.length  !=  0
	remove a node n from S
	add n to the tail of L
	for each vertex m that comes out of n
		remove the edge from n to m
		if m has no other incoming edges
			add m to S

if graph still has edges
	panic to return error
	(Graph has at least one cycle)
else
	return L (topological sort)
*/

import (
	"fmt"

	"github.com/gyuho/goraph/graph/gsd"
	slice "github.com/gyuho/gosequence"
)

func TSKahn(g *gsd.Graph) (string, bool) {
	// vertices_with_no_incoming_edge
	niedges := slice.NewSequence()
	vertices := g.GetVertices()
	for _, vtx := range *vertices {
		if vtx == nil {
			continue
		}
		if vtx.(*gsd.Vertex).InVertices.Len() == 0 {
			niedges.PushBack(vtx.(*gsd.Vertex))
		}
	}

	result := slice.NewSequence()

	// while  S.length  !=  0
	for niedges.Len() != 0 {

		// remove a node n from S
		n := niedges.PopFront()

		// add n to the tail of L
		result.PushBack(n.(*gsd.Vertex))
		// for each vertex m that comes out of n
		ovs := n.(*gsd.Vertex).GetOutVertices()

		// for _, m := range *ovs {
		// once deleted, the original slice size decreases
		// since we use pointers
		// Make sure to use the copy, not the original!
		todelete := []interface{}{}
		for _, m := range *ovs {
			todelete = append(todelete, m)
		}

		for _, m := range todelete {
			g.DeleteEdge(n.(*gsd.Vertex), m.(*gsd.Vertex))
		}
		for _, m := range todelete {
			if m == nil {
				continue
			}
			// if m has no other incoming edges
			if m.(*gsd.Vertex).InVertices.Len() == 0 {
				// add m to S
				niedges.PushBack(m.(*gsd.Vertex))
			}
		}
	}

	if g.GetEdgesSize() != 0 {
		return "No Topological Sort (Not a DAG, there is a cycle)", false
	} else {
		s := ""
		for _, v := range *result {
			s += fmt.Sprintf("%v → ", v.(*gsd.Vertex).ID)
		}
		return s[:len(s)-5], true
	}
}
