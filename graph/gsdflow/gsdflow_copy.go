package gsdflow

// CopyVertex returns the copy of input Vertex.
func CopyVertex(vtx *Vertex) *Vertex {
	cp := NewVertex(vtx.ID)
	cp.Color = vtx.Color
	cp.InVertices = vtx.InVertices.CopySeqPt()
	cp.OutVertices = vtx.OutVertices.CopySeqPt()
	cp.StampD = vtx.StampD
	cp.StampF = vtx.StampF
	cp.Prev = vtx.Prev.CopySeqPt()
	return cp
}

// CopyEdge returns the copy of input Edge.
func CopyEdge(edge *Edge) *Edge {
	cs := CopyVertex(edge.Src)
	ds := CopyVertex(edge.Dst)
	nw := edge.Weight
	return NewEdge(cs, ds, nw)
}

// CopyGraph returns the copy of input Graph.
func CopyGraph(graph *Graph) *Graph {
	cp := NewGraph()
	cp.Vertices = graph.Vertices.CopySeqPt()
	cp.Edges = graph.Edges.CopySeqPt()
	return cp
}
