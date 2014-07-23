package gs

import (
	slice "github.com/gyuho/goraph/gosequence"
	"github.com/gyuho/goraph/parsex/dotx"
	"github.com/gyuho/goraph/parsex/jsonx"
)

// Graph is a graph represented in adjacency list, but implemented in slice.
type Graph struct {
	Vertices *slice.Sequence
	Edges    *slice.Sequence
}

// NewGraph returns a pointer to a new graph.
func NewGraph() *Graph {
	return &Graph{
		slice.NewSequence(),
		slice.NewSequence(),
	}
}

// Vertex is a vertex(node) in Graph.
type Vertex struct {
	ID    string
	Color string

	// InVertices is a slice of vertices that goes into this vertex
	// (vertices that precede this vertex in graph)
	InVertices *slice.Sequence

	// OutVertices is a slice of vertices that go out of this vertex
	OutVertices *slice.Sequence

	// StampD is a time stamp to record the distance
	// from source vertex
	StampD int64

	// StampF is another timestamp to be used in other algorithms
	StampF int64
}

// NewVertex returns a pointer to a new Vertex.
func NewVertex(id string) *Vertex {
	return &Vertex{
		ID:          id,
		Color:       "white",
		InVertices:  slice.NewSequence(),
		OutVertices: slice.NewSequence(),
		StampD:      9999999999,
		StampF:      9999999999,
	}
}

// Edge is an edge(arc) in a graph that has direction from one to another vertex.
type Edge struct {
	// Src is the source vertex that the edge starts from
	Src *Vertex

	// Dst is the destination vertex that the edge goes to
	Dst *Vertex

	// Weight contains the weight value in float64 format
	Weight float64
}

// NewEdge returns a new edge from src to dst.
func NewEdge(src, dst *Vertex, weight float64) *Edge {
	return &Edge{
		src,
		dst,
		weight,
	}
}

// FromJSON parses a JSON file to a graph data structure.
func FromJSON(fpath, gname string) *Graph {
	nodes := jsonx.GetNodes(fpath, gname)
	gmap := jsonx.GetGraphMap(fpath, gname)
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

			// gs always has one element in the weight value slice.
			// g.Connect(src, dst, weights[0])
			for _, weight := range gmap[srcID][dstID] {
				g.Connect(src, dst, weight)
			}
		}
	}
	return g
}

// FromDOT parses a DOT file to a graph data structure.
func FromDOT(fpath string) *Graph {
	nodes := dotx.GetNodes(fpath)
	_, gmap := dotx.GetGraphMap(fpath)
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
				g.Connect(src, dst, weight)
			}
		}
	}
	return g
}
