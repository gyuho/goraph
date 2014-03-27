// Package kosaraju implements Kosaraju's Strongly Connected Components algorithm.
package kosaraju

import (
	"github.com/gyuho/goraph/graph/gsd"
	slice "github.com/gyuho/gosequence"
)

/*
http://www.geeksforgeeks.org/strongly-connected-components/

http://en.wikipedia.org/wiki/Kosaraju%27s_algorithm
It makes use of the fact that the transpose graph
(the same graph with the direction of every edge reversed)
has exactly the same strongly connected components as the original graph.


Let G be a directed graph
Let S be an empty stack

While S does not contain all vertices:
	Choose an arbitrary vertex v not in S.
		Perform a DFS starting at v.
			Each time that DFS finishes
			expanding a vertex u, push u onto S.

Reverse the directions of all arcs to obtain the transpose graph.

While S is nonempty:
	Pop the top vertex v from S.
		Perform a DFS starting at v in the transpose graph.

		The set of visited vertices will give
		the strongly connected component containing v

		record this and
		remove all these vertices from the graph G and the stack S.
*/

// SCC returns the Strongly Connected Components using Kosaraju's algorithm.
func SCC(g *gsd.Graph) string {
	// Let G be a directed graph and S be an empty stack.
	Stack := slice.NewSequence()
	Vertices := g.GetVertices()

	// While S does not contain all vertices:
	for !slice.IsEqual(*Stack, *Vertices) {
		// Choose an arbitrary vertex v not in S
		var vtx *gsd.Vertex
		for _, val := range *Vertices {
			if !Contains(val.(*gsd.Vertex), Stack) {
				vtx = val.(*gsd.Vertex)
				break
			}
		}
		_ = vtx

		// Perform a depth-first search starting at v

		// Each time that depth-first search finishes
		// expanding a vertex u, push u onto S
	}

	for Stack.Len() != 0 {

	}

	return ""
}

// Contains returns true if vtx exists in the slice sl.
func Contains(vtx *gsd.Vertex, sl *slice.Sequence) bool {
	for _, val := range *sl {
		if val.(*gsd.Vertex).ID == vtx.ID {
			return true
		}
	}
	return false
}
