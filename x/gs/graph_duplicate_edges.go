package gs

import (
	"github.com/gyuho/goraph/parsex/dotxd"
	"github.com/gyuho/goraph/parsex/jsonxd"
)

// This allows duplicate edges.
// There can be multiple edges from one to the other node.
// "Connect" and "GetEdgeWeight" are defined different.
// And "DeleteEdge" works different than others.

// ConnectDupl connects the vertex v to A, not A to v.
// When there is more than one edge, it adds up the weight values.
func (g *Graph) ConnectDupl(A, B *Vertex, weight float64) {
	// if there is already an edge from A to B
	// if g.ImmediateDominate(A, B) {

	// we just add another edge
	edge := NewEdge(A, B, weight)
	g.AddEdge(edge)
	A.AddOutVertex(B)
	B.AddInVertex(A)
}

// FromJSONDupl parses a JSON file to a graph data structure,
// with duplicate edges being allowed.
func FromJSONDupl(fpath, gname string) *Graph {
	nodes := jsonxd.GetNodes(fpath, gname)
	gmap := jsonxd.GetGraphMapDupl(fpath, gname)
	// map[string]map[string][]float64

	g := NewGraph()
	// source vertex
	for _, srcID := range nodes {
		src := g.CreateAndAddToGraph(srcID)

		// destination vertex
		for dstID := range gmap[srcID] {
			dst := g.CreateAndAddToGraph(dstID)

			// This is not constructing the bi-directional edge automatically.
			// We need to specify the edge direction in graph data.
			// weights, _ := gmap[srcID][dstID]
			//
			// gs always has one element in the weight value slice.
			// g.Connect(src, dst, weights[0])
			//
			for _, weight := range gmap[srcID][dstID] {
				g.ConnectDupl(src, dst, weight)
			}
		}
	}
	return g
}

// FromDOTDupl parses a DOT file to a graph data structure.
func FromDOTDupl(fpath string) *Graph {
	nodes := dotxd.GetNodes(fpath)
	_, gmap := dotxd.GetGraphMapDupl(fpath)
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
				g.ConnectDupl(src, dst, weight)
			}
		}
	}
	return g
}

// GetEdgeWeightDupl returns the slice of weight values
// of the edge from source to destination vertex.
// In case we need to allow duplicate edges,
// we return a slice of weights.
func (g Graph) GetEdgeWeightDupl(src, dst *Vertex) []float64 {
	slice := g.GetEdges()
	result := []float64{}
	for _, edge := range *slice {
		if edge.(*Edge).Src == src && edge.(*Edge).Dst == dst {
			result = append(result, edge.(*Edge).Weight)
		}
	}
	return result
}
