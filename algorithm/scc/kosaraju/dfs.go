package kosaraju

import (
	"fmt"
	"log"

	slice "github.com/gyuho/goraph/gosequence"
	"github.com/gyuho/goraph/graph/gs"
)

// Breadth First Search: Queue
// Depth First Search: Stack, Recursion

// DFS does Depth First Search and return the result in visited order.
func DFS(g *gs.Graph) string {
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

// DFSStart performs DFS with a starting Vertex.
func DFSStart(g *gs.Graph, src *gs.Vertex) string {
	if src == nil {
		log.Fatal("Wrong Start Vertex Passed!")
	}

	var stamp int64

	// for each vertex u ∈ g.V
	result := slice.NewSequence()

	if src.Color == "white" {
		dfsVisit(g, src, stamp, result)
	}

	/*
		vertices := g.GetVertices()
		for _, vtx := range *vertices {
			if vtx == nil {
				continue
			}
			if vtx.(*gs.Vertex).Color == "white" {
				dfsVisit(g, vtx.(*gs.Vertex), stamp, result)
			}
		}
	*/

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

// DFSandSCC performs dfsSCC.
func DFSandSCC(g *gs.Graph, start *gs.Vertex) []string {
	stack := []string{}
	dfsSCC(g, start, &stack)
	return stack
}

// dfsSCC performs DFS and returns the result in stack.
// We put vertices on a stack as we finish the recursive step.
func dfsSCC(g *gs.Graph, start *gs.Vertex, stack *[]string) {
	if start == nil {
		log.Fatal("Wrong Start Vertex Passed!")
	}

	// time = 0
	var stamp int64

	// for each vertex u ∈ g.V
	result := slice.NewSequence()

	if start.Color == "white" {
		dfsVisit(g, start, stamp, result)
	}

	for i := result.Len() - 1; i >= 0; i-- {
		*stack = append(*stack, ((*result)[i].(string)))
	}
}
