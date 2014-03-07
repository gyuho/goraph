// Package gs implements graph using adjacency list and slice data structure.
package gs

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	slice "github.com/gyuho/gosequence"
	"github.com/gyuho/gson/jgd"
)

// Graph is a graph represented in adjacency list, but implemented in slice.
type Graph struct {
	Vertices *slice.Sequence
	Edges    *slice.Sequence
}

// NewGraph returns a new graph.
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

	// slice of vertices that goes into this vertex
	// (vertices that precede this vertex in graph)
	InVertices *slice.Sequence

	// slice of vertices that go out of this vertex
	OutVertices *slice.Sequence

	// time stamp to record the distance
	// from source vertex
	StampD int64

	// another timestamp to be used in other algorithms
	StampF int64
}

// NewVertex returns a new Vertex.
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

// Edge is an edge(arc) in a graph
// that has direction from one to another vertex.
type Edge struct {
	Src *Vertex
	// source vertex that this edge starts from

	Dst *Vertex
	// destination vertex that this edge goes to

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

// GetOutVertices returns a slice of adjacent vertices from vertex v.
func (v Vertex) GetOutVertices() *slice.Sequence {
	return v.OutVertices
}

// GetOutVerticesSize returns the size of the vertex v's OutVertices
func (v Vertex) GetOutVerticesSize() int {
	// dereference
	return v.OutVertices.Len()
}

// GetInVertices returns a slice of adjacent vertices that goes to vertex v.
func (v Vertex) GetInVertices() *slice.Sequence {
	return v.InVertices
}

// GetInVerticesSize returns the size of the vertex v's InVertices.
func (v Vertex) GetInVerticesSize() int {
	// dereference
	return v.InVertices.Len()
}

// FindVertexByID returns the vertex with input ID
// , or return nil if it doesn't exist.
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

// CreateAndAddToGrammar finds the vertex with the ID, or create it.
func (g *Graph) CreateAndAddToGraph(id string) *Vertex {
	vtx := g.FindVertexByID(id)
	if vtx == nil {
		vtx = NewVertex(id)
		// then add this vertex to the graph
		g.AddVertex(vtx)
	}
	return vtx
}

// AddInVertex adds the vertex v to the vertex i's InVertices.
func (i *Vertex) AddInVertex(v *Vertex) {
	i.InVertices.PushBack(v)
}

// AddOutVertex adds the vertex v to
// the vertex o's OutVertices.
func (o *Vertex) AddOutVertex(v *Vertex) {
	o.OutVertices.PushBack(v)
}

// AddVertex adds the vertex v to
// the graph g's Vertices.
func (g *Graph) AddVertex(v *Vertex) {
	g.Vertices.PushBack(v)
}

// AddEdge adds the edge e to
// the graph g's Edges.
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

// GetEdgeWeight returns the slice of weight values
// of the edge from source to destination vertex.
// In case we need to allow duplicate edges,
// we can use slice.
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

// ParseToGraph parses string data to return a new graph.
func ParseToGraph(str string) *Graph {
	validID := regexp.MustCompile(`\t{1,}`)
	str = validID.ReplaceAllString(str, " ")
	str = strings.TrimSpace(str)
	lines := strings.Split(str, "\n")

	g := NewGraph()

	for _, line := range lines {
		fields := strings.Split(line, "|")

		// srcID in string format
		srcID := fields[0]

		// source vertex
		src := g.CreateAndAddToGraph(srcID)

		edgepairs := fields[1:]

		for _, pair := range edgepairs {
			if len(strings.Split(pair, ",")) == 1 {
				// to skip the lines below
				// and go back to the for-loop
				continue
			}
			dstID := strings.Split(pair, ",")[0]
			dst := g.CreateAndAddToGraph(dstID)
			// This is not constructing the bi-directional edge automatically.
			// We need to input bi-directional graph data.
			weight := strToFloat(strings.Split(pair, ",")[1])
			g.Connect(src, dst, weight)
		}
	}
	return g
}

// DeleteInVertex removes the input vertex v
// from the vertex i's InVertices.
func (i *Vertex) DeleteInVertex(v *Vertex) {
	i.InVertices.FindAndDelete(v)
}

// DeleteOutVertex removes the input vertex v
// from the vertex o's OutVertices.
func (o *Vertex) DeleteOutVertex(v *Vertex) {
	o.OutVertices.FindAndDelete(v)
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

// DeleteVertex removes the input vertex A
// from the graph g's Vertices.
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

// strToFloat converts string to float64.
func strToFloat(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println(err.Error())
		panic("Fail")
	}
	return f
}

// JSONGraph parses JSON file to a graph.
func JSONGraph(filename, graph string) *Graph {
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
				g.Connect(src, dst, weight)
			}
		}
	}
	return g
}
