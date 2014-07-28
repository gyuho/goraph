package dfs

import (
	"fmt"

	slice "github.com/gyuho/goraph/gosequence"
	"github.com/gyuho/goraph/graph/gs"
)

// Breadth First Search: Queue
// Depth First Search: Stack, Recursion

// DFS does Depth First Search and return the result in visited order.
func DFS(g *gs.Graph) string {
	// time = 0
	// var stamp int64 = 0
	var stamp int64

	// for each vertex u ∈ g.V
	result := slice.NewSequence()
	vertices := g.GetVertices()
	for _, vtx := range *vertices {
		if vtx == nil {
			continue
		}
		if vtx.(*gs.Vertex).Color == "white" {
			dfsVisit(g, vtx.(*gs.Vertex), stamp, result)
		}
	}

	var rs string
	for _, v := range *result {
		rs += fmt.Sprintf("%v → ", v)
	}
	return rs[:len(rs)-5]
}

// dfsVisit recursively visits the vertices in the graph.
func dfsVisit(g *gs.Graph, src *gs.Vertex, stamp int64, result *slice.Sequence) {
	stamp++
	src.StampD = stamp
	src.Color = "gray"

	ovs := src.GetOutVertices()
	for _, vtx := range *ovs {
		if vtx == nil {
			continue
		}
		if vtx.(*gs.Vertex).Color == "white" {
			vtx.(*gs.Vertex).Prev.PushBack(src)
			dfsVisit(g, vtx.(*gs.Vertex), stamp, result)
		}
	}

	src.Color = "black"
	result.PushFront(src.ID)

	stamp++
	src.StampF = stamp
}
