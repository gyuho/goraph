package kosaraju

/*
Pseudo Code CLRS page 604

DFS(G)
	for each vertex u ∈ G.V
		u.color = WHITE
		u.π = NIL
	// this is already done
	// when instantiating the graph
	// and instead of InVertices
	// we can just create another slice
	// inside Graph (Prev)
	// in order not to modify the original graph

	time = 0

	for each vertex u ∈ G.V
		if u.color == WHITE
			DFS-Visit(G, u)

---

DFS-Visit(G, u)
	time = time + 1
	u.d = time
	u.color = GRAY

	for each v ∈ G.Adj[u]
		if v.color == WHITE
			v.π = u
			DFS-Visit(G, v)

	u.color = BLACK
	time = time + 1
	u.f = time
*/

import (
	"fmt"

	"github.com/gyuho/goraph/graph/gsd"
	slice "github.com/gyuho/gosequence"
)

// Breadth First Search: Queue
// Depth First Search: Stack, Recursion

// DFS does Depth First Search and return the result in visited order.
func DFS(g *gsd.Graph) string {
	// time = 0
	var stamp int64 = 0

	// for each vertex u ∈ g.V
	result := slice.NewSequence()
	vertices := g.GetVertices()
	for _, vtx := range *vertices {
		if vtx == nil {
			continue
		}
		if vtx.(*gsd.Vertex).Color == "white" {
			DFSVisit(g, vtx.(*gsd.Vertex), stamp, result)
		}
	}

	s := ""
	for _, v := range *result {
		s += fmt.Sprintf("%v → ", v)
	}
	return s[:len(s)-5]
}

// DFSs performs DFS with a starting Vertex.
func DFSs(g *gsd.Graph, start *gsd.Vertex) string {
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

	/*
		vertices := g.GetVertices()
		for _, vtx := range *vertices {
			if vtx == nil {
				continue
			}
			if vtx.(*gsd.Vertex).Color == "white" {
				DFSVisit(g, vtx.(*gsd.Vertex), stamp, result)
			}
		}
	*/

	s := ""
	for _, v := range *result {
		s += fmt.Sprintf("%v → ", v)
	}
	return s[:len(s)-5]
}

// DFS_SCC performs dfsSCC.
func DFS_SCC(g *gsd.Graph, start *gsd.Vertex) *slice.Sequence {
	stack := slice.NewSequence()
	dfsSCC(g, start, stack)
	return stack
}

// dfsSCC performs DFS and returns the result in stack.
// We put vertices on a stack as we finish the recursive step.
func dfsSCC(g *gsd.Graph, start *gsd.Vertex, stack *slice.Sequence) {
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

	for i := len(*result) - 1; i >= 0; i-- {
		stack.PushBack((*result)[i].(string))
	}
}

// DFSVisit recursively visits the vertices in the graph.
func DFSVisit(g *gsd.Graph, src *gsd.Vertex, stamp int64, result *slice.Sequence) {
	stamp += 1
	src.StampD = stamp
	src.Color = "gray"

	ovs := src.GetOutVertices()
	for _, vtx := range *ovs {
		if vtx == nil {
			continue
		}
		if vtx.(*gsd.Vertex).Color == "white" {
			vtx.(*gsd.Vertex).Prev.PushBack(src)
			DFSVisit(g, vtx.(*gsd.Vertex), stamp, result)
		}
	}

	src.Color = "black"
	result.PushFront(src.ID)

	stamp += 1
	src.StampF = stamp
}

// DFS_SCC_Reverse performs dfsSCCReverse.
func DFS_SCC_Reverse(g *gsd.Graph, start *gsd.Vertex) *slice.Sequence {
	stack := slice.NewSequence()
	dfsSCCReverse(g, start, stack)
	return stack
}

// dfsSCCReverse performs DFSReverse.
func dfsSCCReverse(g *gsd.Graph, start *gsd.Vertex, stack *slice.Sequence) {
	if start == nil {
		panic("Wrong Start Vertex Passed!")
	}

	// time = 0
	var stamp int64 = 0

	// for each vertex u ∈ g.V
	result := slice.NewSequence()

	if start.Color == "white" {
		DFSVisitReverse(g, start, stamp, result)
	}

	for i := len(*result) - 1; i >= 0; i-- {
		stack.PushBack((*result)[i].(string))
	}
}

// DFSVisitReverse recursively visits the vertices in a reverse order.
func DFSVisitReverse(g *gsd.Graph, src *gsd.Vertex, stamp int64, result *slice.Sequence) {
	stamp += 1
	src.StampD = stamp
	src.Color = "gray"

	// ovs := src.GetOutVertices()
	ovs := src.GetInVertices()
	for _, vtx := range *ovs {
		if vtx == nil {
			continue
		}
		if vtx.(*gsd.Vertex).Color == "white" {
			vtx.(*gsd.Vertex).Prev.PushBack(src)
			DFSVisit(g, vtx.(*gsd.Vertex), stamp, result)
		}
	}

	src.Color = "black"
	result.PushFront(src.ID)

	stamp += 1
	src.StampF = stamp
}
