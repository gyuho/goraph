package gl

import "container/list"

// Vertexer is a set of methods for gs graph data structure
// that receives Vertex as an input.
type Vertexer interface {
	// GetOutVertices returns a slice of adjacent vertices from the receiver Vertex.
	GetOutVertices() *list.List

	// GetOutVerticesSize returns the size of the receiver Vertex's OutVertices.
	GetOutVerticesSize() int

	// GetInVertices returns a slice of adjacent vertices that come into the receiver Vertex.
	GetInVertices() *list.List

	// GetInVerticesSize returns the size of the receiver Vertex's InVertices.
	GetInVerticesSize() int

	// GetPrev returns a slice of Prev.
	GetPrev() *list.List

	// GetPrevSize returns the size of thereceiver Vertex's Prev.
	GetPrevSize() int

	// AddPrevVertex adds the vertex v to a receiver's Prev.
	AddPrevVertex(vtx *Vertex)

	// AddInVertex adds a vertex to a receiver's InVertices.
	AddInVertex(vtx *Vertex)

	// AddOutVertex adds a vertex to a receiver's OutVertices.
	AddOutVertex(vtx *Vertex)

	// DeleteInVertex removes the input vertex from the receiver's InVertices.
	DeleteInVertex(vtx *Vertex)

	// DeleteOutVertex removes the input vertex from the receiver's OutVertices.
	DeleteOutVertex(vtx *Vertex)
}
