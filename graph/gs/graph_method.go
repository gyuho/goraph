package gs

import (
	"fmt"

	slice "github.com/gyuho/goraph/gosequence"
)

// ToJSON converts a receiver graph data structure to JSON format.
func (g Graph) ToJSON() string {
	return ""
}

// ToDOT converts a receiver graph data structure to DOT format.
func (g Graph) ToDOT() string {
	return ""
}

// GetVertices returns the vertex slice.
func (g Graph) GetVertices() *slice.Sequence {
	return g.Vertices
}

// GetVerticesSize returns the size of vertex slice in a graph.
func (g Graph) GetVerticesSize() int {
	// dereference
	return g.Vertices.Len()
}

// GetEdges returns the edge slice.
func (g Graph) GetEdges() *slice.Sequence {
	return g.Edges
}

// GetEdgesSize returns the size of edge slice in a graph.
func (g Graph) GetEdgesSize() int {
	// dereference
	return g.Edges.Len()
}

// FindVertexByID returns the vertex with input ID, or return nil if it doesn't exist.
func (g Graph) FindVertexByID(id interface{}) *Vertex {
	slice := g.GetVertices()
	for _, v := range *slice {
		if fmt.Sprintf("%v", v.(*Vertex).ID) == fmt.Sprintf("%v", id) {
			return v.(*Vertex)
		}
	}
	// nil is used as pointer
	// null pointer
	return nil
}

// CreateAndAddToGraph finds the vertex with the ID, or create it.
func (g *Graph) CreateAndAddToGraph(id string) *Vertex {
	vtx := g.FindVertexByID(id)
	if vtx == nil {
		vtx = NewVertex(id)
		// then add this vertex to the graph
		g.AddVertex(vtx)
	}
	return vtx
}

// AddVertex adds the Vertex v to a graph's Vertices.
func (g *Graph) AddVertex(v *Vertex) {
	g.Vertices.PushBack(v)
}

// AddEdge adds the Edge e to a graph's Edges.
func (g *Graph) AddEdge(e *Edge) {
	g.Edges.PushBack(e)
}

// ImmediateDominate returns true if A immediately dominates B.
// That is, true if A can go to B with only one edge.
func (g Graph) ImmediateDominate(A, B *Vertex) bool {
	_, exist := A.OutVertices.Find(B)
	if exist {
		return true
	}
	return false
}

// GetEdgeWeight returns the weight value of the edge from source to destination vertex.
func (g Graph) GetEdgeWeight(src, dst *Vertex) float64 {
	slice := g.GetEdges()
	for _, edge := range *slice {
		if edge.(*Edge).Src == src && edge.(*Edge).Dst == dst {
			return edge.(*Edge).Weight
		}
	}
	return 0.0
}

// UpdateWeight updates the weight value between vertices.
func (g *Graph) UpdateWeight(src, dst *Vertex, value float64) {
	edges := g.GetEdges()
	for _, edge := range *edges {
		if edge.(*Edge).Src == src && edge.(*Edge).Dst == dst {
			edge.(*Edge).Weight = value
		}
	}
}

// Connect connects the vertex v to A, not A to v.
func (g *Graph) Connect(A, B *Vertex, weight float64) {
	// if there is already an edge from A to B
	if g.ImmediateDominate(A, B) {
		c := g.GetEdgeWeight(A, B)
		n := c + weight
		g.UpdateWeight(A, B, n)
	} else {
		edge := NewEdge(A, B, weight)
		g.AddEdge(edge)
		A.AddOutVertex(B)
		B.AddInVertex(A)
	}
}

// DeleteVertex removes the input vertex A from the graph g's Vertices.
func (g *Graph) DeleteVertex(A *Vertex) {
	if g.FindVertexByID(A.ID) == nil {
		// To Debug
		// panic(A.ID + " Vertex does not exist! Can't delete the Vertex!")
		return
	}

	// remove all edges connected to this vertex
	todelete := slice.NewSequence()
	for _, edge := range *g.Edges {
		// if the edge is from A to B
		if edge.(*Edge).Src == A || edge.(*Edge).Dst == A {
			todelete.PushBack(edge.(*Edge))
		}
	}
	for _, e := range *todelete {
		g.Edges.FindAndDelete(e)
	}

	// remove A from its outgoing vertices' InVertices
	for _, v := range *A.OutVertices {
		// can't convert / type-assert on nil
		// so we must check if it is nil
		if v == nil {
			break
		}
		d := v.(*Vertex).GetInVertices()
		d.FindAndDelete(A)
	}

	// remove A from its incoming vertices' OutVertices
	for _, v := range *A.InVertices {
		if v == nil {
			break
		}
		d := v.(*Vertex).GetOutVertices()
		d.FindAndDelete(A)
	}

	// do this at the end!
	// remove from the graph
	g.Vertices.FindAndDelete(A)
}

// DeleteEdge deletes the edge from the vertex A to B.
// Note that this only delete one direction from A to B.
func (g *Graph) DeleteEdge(A, B *Vertex) {
	if g.FindVertexByID(A.ID) == nil {
		// To Debug
		// panic(A.ID + " Vertex does not exist! Can't delete the Edge!")
		return
	}

	if g.FindVertexByID(B.ID) == nil {
		// To Debug
		// panic(B.ID + " Vertex does not exist! Can't delete the Edge!")
		return
	}

	if g.ImmediateDominate(A, B) == false {
		// To Debug
		// panic("No edge from " + A.ID + " to " + B.ID)
		return
	}

	// delete B from A's OutVertices
	A.OutVertices.FindAndDelete(B)

	// delete A from B's InVertices
	B.InVertices.FindAndDelete(A)

	// Always delete from graph at the end
	// remove the edge from the graph's edge list
	todelete := slice.NewSequence()
	for _, edge := range *g.Edges {
		// if the edge is from A to B
		if edge.(*Edge).Src == A && edge.(*Edge).Dst == B {
			todelete.PushBack(edge.(*Edge))
		}
	}
	for _, e := range *todelete {
		g.Edges.FindAndDelete(e)
	}
}
