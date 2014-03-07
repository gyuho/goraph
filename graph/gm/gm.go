// Package gm implements graph using adjacency list and map data structure.
package gm

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gyuho/gson/jgd"
)

// string is just to map the ID to the Vertex.
// Do not map "Vertex" to string
// since VMAP will be used as slice
// and the map does not allow duplicate keys.
type VMAP map[string]*Vertex
type EMAP map[int]*Edge

// Graph is a graph represented in adjacency list, but implemented in slice.
type Graph struct {
	Vertices VMAP
	Edges    EMAP
}

// NewGraph returns a new graph.
func NewGraph() *Graph {
	return &Graph{
		make(VMAP),
		make(EMAP),
	}
}

// Vertex is a vertex(node) in Graph.
type Vertex struct {
	ID    string
	Color string

	// slice of vertices that goes into this vertex
	// (vertices that precede this vertex in graph)
	InVertices VMAP

	// slice of vertices that go out of this vertex
	OutVertices VMAP

	// time stamp to record the distance
	// from source vertex
	stampD int64

	// another timestamp to be used in other algorithms
	stampF int64
}

// NewVertex returns a new Vertex.
func NewVertex(id string) *Vertex {
	return &Vertex{
		ID:          id,
		Color:       "white",
		InVertices:  make(VMAP),
		OutVertices: make(VMAP),
		stampD:      9999999999,
		stampF:      9999999999,
	}
}

// Edge is an edge(arc) in a graph
// that has direction from one to another vertex.
type Edge struct {
	SrcDst map[string]string
	// Source vertex is mapped to Destination vertex

	Weight float64
}

// NewEdge returns a new edge from src to dst.
func NewEdge(src, dst *Vertex, weight float64) *Edge {
	tm := make(map[string]string)
	tm[src.ID] = dst.ID
	return &Edge{
		tm,
		weight,
	}
}

// GetVertices returns the vertex slice.
func (g Graph) GetVertices() VMAP {
	return g.Vertices
}

// GetVerticesSize returns the size of vertex slice in a graph.
func (g Graph) GetVerticesSize() int {
	return len(g.GetVertices())
}

// GetEdges returns the edge slice.
func (g Graph) GetEdges() EMAP {
	return g.Edges
}

// GetEdgesSize returns the size of edge slice in a graph.
func (g Graph) GetEdgesSize() int {
	return len(g.GetEdges())
}

// CreateAndAddToGrammar finds the vertex with the ID, or create it.
func (g *Graph) CreateAndAddToGraph(id string) *Vertex {
	m := g.Vertices
	vtx, ok := m[id]
	if !ok {
		vtx = NewVertex(id)
		// then add this vertex to the graph
		g.AddVertex(vtx)
	}
	return vtx
}

// AddVertex adds the vertex v to
// the graph g's Vertices.
// Only to be used when there is no duplicate vertex.
func (g *Graph) AddVertex(v *Vertex) {
	g.GetVertices()[v.ID] = v
}

// AddEdge adds the edge e to the graph g's Edges.
// Only to be used when there is no duplicate edge.
func (g *Graph) AddEdge(e *Edge) {
	g.GetEdges()[len(g.GetEdges())] = e
}

// AddInVertex adds the vertex B to the vertex A's InVertices.
// B goes into A.
func (A *Vertex) AddInVertex(B *Vertex) {
	A.GetInVertices()[B.ID] = B
}

// AddOutVertex adds the vertex B to the vertex A's OutVertices.
// A goes out to B.
func (A *Vertex) AddOutVertex(B *Vertex) {
	A.GetOutVertices()[B.ID] = B
}

// GetOutVertices returns a slice of adjacent vertices from vertex v.
func (v Vertex) GetOutVertices() VMAP {
	return v.OutVertices
}

// GetOutVerticesSize returns the size of the vertex v's OutVertices
func (v Vertex) GetOutVerticesSize() int {
	// dereference
	return len(v.GetOutVertices())
}

// GetInVertices returns a slice of adjacent vertices that goes to vertex v.
func (v Vertex) GetInVertices() VMAP {
	return v.InVertices
}

// GetInVerticesSize returns the size of the vertex v's InVertices.
func (v Vertex) GetInVerticesSize() int {
	// dereference
	return len(v.GetInVertices())
}

// ImmediateDominate returns true if A immediately dominates B.
// That is, true if A can go to B with only one edge.
func (g Graph) ImmediateDominate(A, B *Vertex) bool {
	if _, exist := A.GetOutVertices()[B.ID]; exist {
		return true
	}
	return false
}

// GetEdgeWeight returns the weight value.
// Our graph does not allow the duplicate edges.
// If there is a duplicate edges,
// the value is added to the existent weight.
func (g Graph) GetEdgeWeight(src, dst *Vertex) float64 {
	edges := g.GetEdges()
	for _, edge := range edges {
		if edge.SrcDst[src.ID] == dst.ID {
			return edge.Weight
		}
	}
	return 0.0
}

// UpdateWeight updates the weight value between vertices.
func (g *Graph) UpdateWeight(src, dst *Vertex, value float64) {
	edges := g.GetEdges()
	for _, edge := range edges {
		if edge.SrcDst[src.ID] == dst.ID {
			edge.Weight = value
		}
	}
}

// Connect connects the vertex A to B, not B to A.
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
	delete(i.GetInVertices(), v.ID)
}

// DeleteOutVertex removes the input vertex v
// from the vertex o's OutVertices.
func (o *Vertex) DeleteOutVertex(v *Vertex) {
	delete(o.GetOutVertices(), v.ID)
}

// DeleteEdge deletes the edge from the vertex A to B.
// Note that this only delete one direction from A to B.
func (g *Graph) DeleteEdge(A, B *Vertex) {
	if _, ok := g.Vertices[A.ID]; !ok {
		// To Debug
		// panic(A.ID + " Vertex does not exist! Can't delete the Edge!")
		return
	}

	if _, ok := g.Vertices[B.ID]; !ok {
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
	delete(A.GetOutVertices(), B.ID)

	// delete A from B's InVertices
	delete(B.GetInVertices(), A.ID)

	// Always delete from graph at the end
	// remove the edge from the graph's edge list
	edges := g.GetEdges()
	for k, edge := range edges {
		if edge.SrcDst[A.ID] == B.ID {
			delete(edges, k)
		}
	}
}

// DeleteVertex removes the input vertex A
// from the graph g's Vertices.
func (g *Graph) DeleteVertex(A *Vertex) {
	if _, ok := g.Vertices[A.ID]; !ok {
		// To Debug
		// panic(A.ID + " Vertex does not exist! Can't delete the Vertex!")
		return
	}

	// remove all edges connected to this vertex
	edges := g.GetEdges()
	for k, edge := range edges {
		if _, ok := edge.SrcDst[A.ID]; ok {
			delete(edges, k)
		}
		for key := range edge.SrcDst {
			if edge.SrcDst[key] == A.ID {
				delete(edges, k)
			}
		}
	}

	// remove A from its outgoing vertices' InVertexList
	for _, v := range A.GetOutVertices() {
		delete(v.GetInVertices(), A.ID)
	}

	// remove A from its incoming vertices' OutVertexList
	for _, v := range A.GetInVertices() {
		delete(v.GetOutVertices(), A.ID)
	}

	// do this at the end!
	// remove from the graph
	delete(g.GetVertices(), A.ID)
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
