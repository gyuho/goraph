// Package gld implements graph using adjacency
// list and linked list data structure. This is same as
// the package gl except that this allows duplicate edges.
// There can be multiple edges from one to the other node.
// "Connect" and "GetEdgeWeight" are defined different.
// And "DeleteEdge" works different than others.
package gld

import (
	"container/list"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gyuho/gson/jgd"
)

// Graph is a graph represented in adjacency list. Vertices and edges
// are stored in linked list. Look at how linked list is implemented
// in Go's source code. Edges only needs to be handled by a graph
// , not by each vertex. Vertex contains a list of incoming and outgoing
// vertices, as in linked list. Let's suffix with T to differentiate with
// other types of graphs.
type Graph struct {
	Vertices *list.List
	Edges    *list.List
}

// NewGraph returns a new graph.
func NewGraph() *Graph {
	return &Graph{
		list.New(),
		list.New(),
	}
}

// Vertex is a vertex(node) in Graph.
type Vertex struct {
	ID    string
	Color string

	// list of vertices that goes into this vertex
	// (vertices that precede this vertex in graph)
	InVertices *list.List

	// list of vertices that go out of this vertex
	OutVertices *list.List

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
		InVertices:  list.New(),
		OutVertices: list.New(),
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

// GetVertices returns the vertex list.
func (g Graph) GetVertices() *list.List {
	return g.Vertices
}

// GetEdges returns the edge list.
func (g Graph) GetEdges() *list.List {
	return g.Edges
}

// GetVerticesSize returns the size of vertex list in a graph.
func (g Graph) GetVerticesSize() int {
	return g.Vertices.Len()
}

// GetEdgesSize returns the size of edge list in a graph.
func (g Graph) GetEdgesSize() int {
	return g.GetEdges().Len()
}

// GetOutVertices returns a list of adjacent vertices from vertex v.
func (v Vertex) GetOutVertices() *list.List {
	return v.OutVertices
}

// GetInVertices returns a list of adjacent vertices that goes to vertex v.
func (v Vertex) GetInVertices() *list.List {
	return v.InVertices
}

// FindVertexByID returns the vertex with input ID
// , or return nil if it doesn't exist.
func (g Graph) FindVertexByID(id interface{}) *Vertex {
	for vtx := g.Vertices.Front(); vtx != nil; vtx = vtx.Next() {
		// NOT  vtx.Value.(Vertex).ID
		if fmt.Sprintf("%v", vtx.Value.(*Vertex).ID) == fmt.Sprintf("%v", id) {
			return vtx.Value.(*Vertex)
		}
	}
	return nil
}

// CreateAndAddToGrammar finds the vertex with the ID, or create it.
func (g *Graph) CreateAndAddToGraph(id string) *Vertex {
	vtx := g.FindVertexByID(id)
	if vtx == nil {
		vtx = NewVertex(id)
		// then add this vertex to the graph
		g.GetVertices().PushBack(vtx)
	}
	return vtx
}

// ImmediateDominate returns true if A immediately dominates B.
// That is, true if A can go to B with only one edge.
func (g Graph) ImmediateDominate(A, B *Vertex) bool {
	for vtx := A.GetOutVertices().Front(); vtx != nil; vtx = vtx.Next() {
		if vtx.Value.(*Vertex) == B {
			return true
		}
	}
	return false
}

// GetEdgeWeight returns the slice of weight values
// of the edge from source to destination vertex.
// In case we need to allow duplicate edges,
// we return a slice of weights.
func (g Graph) GetEdgeWeight(src, dst *Vertex) []float64 {
	result := []float64{}
	for edge := g.GetEdges().Front(); edge != nil; edge = edge.Next() {
		if edge.Value.(*Edge).Src == src && edge.Value.(*Edge).Dst == dst {
			result = append(result, edge.Value.(*Edge).Weight)
		}
	}
	return result
}

// UpdateWeight updates the weight value.
func (g *Graph) UpdateWeight(src, dst *Vertex, value float64) {
	for edge := g.GetEdges().Front(); edge != nil; edge = edge.Next() {
		if edge.Value.(*Edge).Src == src && edge.Value.(*Edge).Dst == dst {
			edge.Value.(*Edge).Weight = value
		}
	}
}

// Connect connects the vertex v to A, not A to v.
func (g *Graph) Connect(A, B *Vertex, weight float64) {
	// if there is already an edge from A to B
	// if g.ImmediateDominate(A, B) {

	// we just add another edge
	edge := NewEdge(A, B, weight)
	g.GetEdges().PushBack(edge)
	A.GetOutVertices().PushBack(B)
	B.GetInVertices().PushBack(A)
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

	// delete A from B's InVertices
	var next1 *list.Element
	for vtx := B.GetInVertices().Front(); vtx != nil; vtx = next1 {
		next1 = vtx.Next()
		if vtx.Value.(*Vertex) == A {
			B.GetInVertices().Remove(vtx)
		}
	}

	// delete B from A's OutVertices
	var next2 *list.Element
	for vtx := A.GetOutVertices().Front(); vtx != nil; vtx = next2 {
		next2 = vtx.Next()
		if vtx.Value.(*Vertex) == B {
			A.GetOutVertices().Remove(vtx)
		}
	}

	// Always delete from graph at the end
	// remove the edge from the graph's edge list
	var next3 *list.Element
	for edge := g.GetEdges().Front(); edge != nil; edge = next3 {
		next3 = edge.Next()
		// if the edge is from A to B
		if edge.Value.(*Edge).Src == A && edge.Value.(*Edge).Dst == B {
			// don't do this
			// edge.Value.(*Edge).Src = nil
			// edge.Value.(*Edge).Dst = nil
			g.GetEdges().Remove(edge)
		}
	}
}

// DeleteVertex deletes a Vertex from the graph.
func (g *Graph) DeleteVertex(A *Vertex) {
	if g.FindVertexByID(A.ID) == nil {
		// To Debug
		// panic(A.ID + " Vertex does not exist! Can't delete the Vertex!")
		return
	}

	// remove all edges connected to this vertex
	var next1 *list.Element
	for edge := g.GetEdges().Front(); edge != nil; edge = next1 {
		next1 = edge.Next()
		if edge.Value.(*Edge).Src == A || edge.Value.(*Edge).Dst == A {
			// remove from the graph
			g.GetEdges().Remove(edge)
		}
	}

	// remove from graph
	var next2 *list.Element
	for vtx := g.GetVertices().Front(); vtx != nil; vtx = next2 {
		next2 = vtx.Next()
		if vtx.Value.(*Vertex) == A {
			// remove from the graph
			g.GetVertices().Remove(vtx)
		}
	}

	// remove A from its outgoing vertices' InVertices
	// need to traverse all outgoing vertices
	var next3 *list.Element
	for vtx1 := A.GetOutVertices().Front(); vtx1 != nil; vtx1 = next3 {
		next3 = vtx1.Next()
		// traverse each vertex's InVertices
		var next4 *list.Element
		for vtx2 := vtx1.Value.(*Vertex).GetInVertices().Front(); vtx2 != nil; vtx2 = next4 {
			next4 = vtx2.Next()
			// check if the vertex is the one we delete
			if vtx2.Value.(*Vertex) == A {
				vtx1.Value.(*Vertex).GetInVertices().Remove(vtx2)
			}
		}
	}

	// remove A from its incoming vertices' OutVertices
	// need to traverse all incoming vertices
	var next5 *list.Element
	for vtx1 := A.GetInVertices().Front(); vtx1 != nil; vtx1 = next5 {
		next5 = vtx1.Next()
		// traverse each vertex's OutVertices
		var next6 *list.Element
		for vtx2 := vtx1.Value.(*Vertex).GetOutVertices().Front(); vtx2 != nil; vtx2 = next6 {
			next6 = vtx2.Next()
			// check if the vertex is the one we delete
			if vtx2.Value.(*Vertex) == A {
				vtx1.Value.(*Vertex).GetOutVertices().Remove(vtx2)
			}
		}
	}
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
