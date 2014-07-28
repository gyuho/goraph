package kosaraju

import "github.com/gyuho/goraph/graph/gs"

// SCC returns the Strongly Connected Components
// using Kosaraju's algorithm.
func SCC(g, gr *gs.Graph) [][]string {
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
			if !contains(val.(*gs.Vertex).ID, Stack) {
				Stack = DFSandSCC(g, val.(*gs.Vertex))
				break
			}
		}

		// Perform a depth-first search starting at v
		// Each time that depth-first search finishes
		// expanding a vertex u, push u onto S
		// Stack = DFSandSCC(g, vtx)
	}

	// Reverse the directions of all arcs
	// to obtain the transpose graph.
	// gr

	result := [][]string{}
	// While S is nonempty
	for len(Stack) != 0 {
		top := Stack[len(Stack)-1]
		Stack = Stack[:len(Stack)-1 : len(Stack)-1]
		rs := DFSandSCC(gr, gr.FindVertexByID(top))
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

// contains returns true if vtx exists in the slice sl.
func contains(vtx string, sl []string) bool {
	for _, val := range sl {
		if val == vtx {
			return true
		}
	}
	return false
}
