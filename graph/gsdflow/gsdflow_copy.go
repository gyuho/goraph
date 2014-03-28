package gsdflow

// CopyVertex returns the copy of input Vertex.
func CopyVertex(vtx *Vertex) *Vertex {
	cp := NewVertex(vtx.ID)
	cp.Color = vtx.Color
	cp.InVertices = vtx.InVertices.CopyPt()
	cp.OutVertices = vtx.OutVertices.CopyPt()
	cp.StampD = vtx.StampD
	cp.StampF = vtx.StampF
	cp.Prev = vtx.Prev.CopyPt()
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
	cp.Vertices = graph.Vertices.CopyPt()
	cp.Edges = graph.Edges.CopyPt()
	return cp
}
