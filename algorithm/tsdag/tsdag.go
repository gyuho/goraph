package tsdag

import (
	"fmt"

	slice "github.com/gyuho/goraph/gosequence"
	"github.com/gyuho/goraph/graph/gs"
)

// TSDAG returns topological sort using DFS.
// It returns false if there is a cycle
// which means that there is no topological sort.
func TSDAG(g *gs.Graph) (string, bool) {
	// time = 0
	var stamp int64

	exist := true
	// for each vertex u ∈ g.V
	result := slice.NewSequence()
	vertices := g.GetVertices()
	for _, vtx := range *vertices {
		if vtx == nil {
			continue
		}
		if vtx.(*gs.Vertex).Color == "white" {
			dfsVisit(g, vtx.(*gs.Vertex), stamp, result, &exist)
		}
	}

	if exist == true && result.Len() > 0 {
		var rs string
		for _, v := range *result {
			rs += fmt.Sprintf("%v → ", v)
		}
		return rs[:len(rs)-5], true
	}
	return "No Topological Sort (Not a DAG, there is a cycle)", false
}

// dfsVisit recursively visits the vertices in the graph.
// And it ends when it finds a cycle in a graph.
func dfsVisit(g *gs.Graph, src *gs.Vertex, stamp int64, result *slice.Sequence, exist *bool) {
	// if the vertex has a temporary mark
	// then stop, which means it is not a DAG
	if src.Color == "gray" {
		*exist = false
		result.Init()
		return
	}

	if src.Color == "white" {
		stamp++
		src.StampD = stamp
		src.Color = "gray"

		/*
		   DFS

		   ovs := src.GetOutVertices()
		   for _, vtx := range *ovs {
		   	if vtx.(*gs.Vertex).Color == "white" {
		   		vtx.(*gs.Vertex).Prev.PushBack(src)
		   		dfsVisit(g, vtx.(*gs.Vertex), stamp, result)
		   	}
		   }
		*/

		// Unlike DFS, we need not check if it's white
		ovs := src.GetOutVertices()
		for _, vtx := range *ovs {
			if vtx == nil {
				continue
			}
			//if vtx.(*gs.Vertex).Color == "white" {
			vtx.(*gs.Vertex).Prev.PushBack(src)
			dfsVisit(g, vtx.(*gs.Vertex), stamp, result, exist)
			//}
		}

		src.Color = "black"
		result.PushFront(src.ID)

		stamp++
		src.StampF = stamp
	}
}
