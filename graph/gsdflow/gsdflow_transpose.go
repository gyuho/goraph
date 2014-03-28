package gsdflow

import (
	"github.com/gyuho/gson/jgd"
)

// JSONGraphT parses JSON file to a graph in a reverse order.
// It returns the tranposed Graph of the original data.
func JSONGraphT(filename, graph string) *Graph {
	nodes := jgd.GetNodes(filename, graph)
	gmap := jgd.MapGraph(filename, graph)
	// map[string]map[string][]float64

	g := NewGraph()
	for _, srcID := range nodes {
		// source vertex
		src := g.CreateAndAddToGraph(srcID)
		for dstID := range gmap[srcID] {
			dst := g.CreateAndAddToGraph(dstID)
			// This is not constructing the bi-directional edge automatically.
			// We need to input bi-directional graph data.
			for _, weight := range gmap[srcID][dstID] {
				// g.Connect(src, dst, weight)
				// Connect reversely
				g.Connect(dst, src, weight)
			}
		}
	}
	return g
}

// Transpose transposes the original Graph.
func (g *Graph) Transpose() *Graph {
	// To change the direction of edges
	edges := g.GetEdges()
	/*
	   type Edge struct {
	   	Src *Vertex
	   	Dst *Vertex
	   	Weight float64
	   }
	*/
	for _, edge := range *edges {
		src := edge.(*Edge).Src
		dst := edge.(*Edge).Dst
		temp := *src
		*src = *dst
		*dst = temp
	}

	// To update the Vertex
	// we can just swap InVertices and OutVertices
	vertices := g.GetVertices()
	/*
	   InVertices *slice.Sequence
	   OutVertices *slice.Sequence
	*/
	for _, vtx := range *vertices {
		in := vtx.(*Vertex).GetInVertices()
		out := vtx.(*Vertex).GetOutVertices()
		temp := *in
		*in = *out
		*out = temp
	}

	return g
}
