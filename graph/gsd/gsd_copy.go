package gsd

import slice "github.com/gyuho/gosequence"

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
	vs := slice.NewSequence()
	for _, vtx := range *(graph.GetVertices()) {
		vs.PushBack(CopyVertex(vtx.(*Vertex)))
	}
	es := slice.NewSequence()
	for _, edge := range *(graph.GetEdges()) {
		es.PushBack(CopyEdge(edge.(*Edge)))
	}
	cp.Vertices = vs
	cp.Edges = es
	return cp
}
