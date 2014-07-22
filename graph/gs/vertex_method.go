package gs

import slice "github.com/gyuho/goraph/gosequence"

// GetOutVertices returns a slice of adjacent vertices from the receiver Vertex.
func (v Vertex) GetOutVertices() *slice.Sequence {
	return v.OutVertices
}

// GetOutVerticesSize returns the size of the receiver Vertex's OutVertices.
func (v Vertex) GetOutVerticesSize() int {
	return v.OutVertices.Len()
}

// GetInVertices returns a slice of adjacent vertices that come into the receiver Vertex.
func (v Vertex) GetInVertices() *slice.Sequence {
	return v.InVertices
}

// GetInVerticesSize returns the size of the receiver Vertex's InVertices.
func (v Vertex) GetInVerticesSize() int {
	// dereference
	return v.InVertices.Len()
}

// AddInVertex adds a vertex to a receiver's InVertices.
func (v *Vertex) AddInVertex(vtx *Vertex) {
	v.InVertices.PushBack(vtx)
}

// AddOutVertex adds a vertex to a receiver's OutVertices.
func (v *Vertex) AddOutVertex(vtx *Vertex) {
	v.OutVertices.PushBack(vtx)
}

// DeleteInVertex removes the input vertex from the receiver's InVertices.
func (v *Vertex) DeleteInVertex(vtx *Vertex) {
	v.InVertices.FindAndDelete(vtx)
}

// DeleteOutVertex removes the input vertex from the receiver's OutVertices.
func (v *Vertex) DeleteOutVertex(vtx *Vertex) {
	v.OutVertices.FindAndDelete(vtx)
}
