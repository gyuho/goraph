// Package tsdag finds the topological sort.
package tsdag

/*
http://en.wikipedia.org/wiki/Topological_sorting

L ← Empty list that will contain the sorted nodes

while there are unmarked nodes do
  select an unmarked node n
	visit(n)

function visit(node n)
	if n has a temporary mark then stop (not a DAG)
	if n is not marked (i.e. has not been visited yet) then
	  mark n temporarily
	  for each node m with an edge from n to m do
	    visit(m)
		mark n permanently
		add n to head of L
*/

import (
	"fmt"

	"github.com/gyuho/goraph/graph/gsd"
	slice "github.com/gyuho/gosequence"
)

// TSDAG returns topological sort using DFS.
// It returns false if there is a cycle
// which means that there is no topological sort.
func TSDAG(g *gsd.Graph) (string, bool) {
	// time = 0
	var stamp int64 = 0

	exist := true
	// for each vertex u ∈ g.V
	result := slice.NewSequence()
	vertices := g.GetVertices()
	for _, vtx := range *vertices {
		if vtx == nil {
			continue
		}
		if vtx.(*gsd.Vertex).Color == "white" {
			DFSVisit(g, vtx.(*gsd.Vertex), stamp, result, &exist)
		}
	}

	if exist == true && result.Len() > 0 {
		s := ""
		for _, v := range *result {
			s += fmt.Sprintf("%v → ", v)
		}
		return s[:len(s)-5], true
	} else {
		return "No Topological Sort (Not a DAG, there is a cycle)", false
	}
}

// DFSVisit recursively visits the vertices in the graph.
// And it ends when it finds a cycle in a graph.
func DFSVisit(g *gsd.Graph, src *gsd.Vertex, stamp int64, result *slice.Sequence, exist *bool) {
	// if the vertex has a temporary mark
	// then stop, which means it is not a DAG
	if src.Color == "gray" {
		*exist = false
		result.Init()
		return
	}

	if src.Color == "white" {
		stamp += 1
		src.StampD = stamp
		src.Color = "gray"

		/*
		   DFS

		   ovs := src.GetOutVertices()
		   for _, vtx := range *ovs {
		   	if vtx.(*gsd.Vertex).Color == "white" {
		   		vtx.(*gsd.Vertex).Prev.PushBack(src)
		   		DFSVisit(g, vtx.(*gsd.Vertex), stamp, result)
		   	}
		   }
		*/

		// Unlike DFS, we need not check if it's white
		ovs := src.GetOutVertices()
		for _, vtx := range *ovs {
			if vtx == nil {
				continue
			}
			//if vtx.(*gsd.Vertex).Color == "white" {
			vtx.(*gsd.Vertex).Prev.PushBack(src)
			DFSVisit(g, vtx.(*gsd.Vertex), stamp, result, exist)
			//}
		}

		src.Color = "black"
		result.PushFront(src.ID)

		stamp += 1
		src.StampF = stamp
	}
}
