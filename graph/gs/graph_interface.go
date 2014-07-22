package gs

import slice "github.com/gyuho/goraph/gosequence"

// GSInterface is a set of methods for gs graph data structure
// that receives Graph as an input.
type GSInterface interface {
	// ToJSON converts a receiver graph data structure to JSON format.
	ToJSON() string

	// ToDOT converts a receiver graph data structure to DOT format.
	ToDOT() string

	// GetVertices returns the vertex slice.
	GetVertices() *slice.Sequence

	// GetVerticesSize returns the size of vertex slice in a graph.
	GetVerticesSize() int

	// GetEdges returns the edge slice.
	GetEdges() *slice.Sequence

	// GetEdgesSize returns the size of edge slice in a graph.
	GetEdgesSize() int

	// FindVertexByID returns the vertex with input ID, or return nil if it doesn't exist.
	FindVertexByID(id interface{}) *Vertex

	// CreateAndAddToGrammar finds the vertex with the ID, and if it does not exist, create it.
	CreateAndAddToGraph(id string) *Vertex

	// AddVertex adds the Vertex v to a graph's Vertices.
	AddVertex(v *Vertex)

	// AddEdge adds the Edge e to a graph's Edges.
	AddEdge(e *Edge)

	// ImmediateDominate returns true if A immediately dominates B.
	// That is, true if A can go to B with only one edge.
	ImmediateDominate(A, B *Vertex) bool

	// GetEdgeWeight returns the weight value of the edge from source to destination vertex.
	GetEdgeWeight(src, dst *Vertex) float64

	// UpdateWeight updates the weight value between vertices.
	UpdateWeight(src, dst *Vertex, value float64)

	// Connect connects the vertex v to A, not A to v.
	Connect(A, B *Vertex, weight float64)

	// DeleteVertex removes the input vertex A from the graph g's Vertices.
	DeleteVertex(A *Vertex)

	// DeleteEdge deletes the edge from the vertex A to B.
	// Note that this only delete one direction from A to B.
	DeleteEdge(A, B *Vertex)
}
