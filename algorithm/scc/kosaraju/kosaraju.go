// Package kosaraju implements Kosaraju's Strongly Connected Components algorithm.
package kosaraju

import (
	"github.com/gyuho/goraph/graph/gsd"
	slice "github.com/gyuho/gosequence"
)

/*
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

// SCC returns the Strongly Connected Components
// using Kosaraju's algorithm.
func SCC(g, gr *gsd.Graph) [][]string {
	// Let G be a directed graph and S be an empty stack.
	Stack := []string{}
	Vertices := g.GetVertices()

	// While S does not contain all vertices:
	for len(Stack) != Vertices.Len() {
		// Choose an arbitrary vertex v not in S
		for _, val := range *Vertices {
			if val == nil || Stack == nil {
				continue
			}
			if !Contains(val.(*gsd.Vertex).ID, Stack) {
				Stack = DFS_SCC(g, val.(*gsd.Vertex))
				break
			}
		}

		// Perform a depth-first search starting at v
		// Each time that depth-first search finishes
		// expanding a vertex u, push u onto S
		// Stack = DFS_SCC(g, vtx)
	}

	// Reverse the directions of all arcs
	// to obtain the transpose graph.
	// gr

	result := [][]string{}
	// While S is nonempty
	for len(Stack) != 0 {
		top := Stack[len(Stack)-1]
		Stack = Stack[:len(Stack)-1 : len(Stack)-1]
		rs := DFS_SCC(gr, gr.FindVertexByID(top))
		sl := []string{}
		for _, val := range rs {
			sl = append(sl, val)
		}
		if len(sl) != 0 {
			result = append(result, sl)
		}
	}

	return result
}

// DFS_SCC performs dfsSCC.
func DFS_SCC(g *gsd.Graph, start *gsd.Vertex) []string {
	stack := []string{}
	dfsSCC(g, start, &stack)
	return stack
}

// dfsSCC performs DFS and returns the result in stack.
// We put vertices on a stack as we finish the recursive step.
func dfsSCC(g *gsd.Graph, start *gsd.Vertex, stack *[]string) {
	if start == nil {
		panic("Wrong Start Vertex Passed!")
	}

	// time = 0
	var stamp int64 = 0

	// for each vertex u ∈ g.V
	result := slice.NewSequence()

	if start.Color == "white" {
		DFSVisit(g, start, stamp, result)
	}

	for i := result.Len() - 1; i >= 0; i-- {
		*stack = append(*stack, ((*result)[i].(string)))
	}
}

// Contains returns true if vtx exists in the slice sl.
func Contains(vtx string, sl []string) bool {
	for _, val := range sl {
		if val == vtx {
			return true
		}
	}
	return false
}
