package gl

import (
	"container/list"
	"fmt"

	"strings"

	"github.com/gyuho/goraph/graph/helper"
)

// ToMAP converts a receiver graph data structure to a map.
func (g Graph) ToMAP() map[string]map[string][]float64 {
	rm := make(map[string]map[string][]float64)
	for vtx1 := g.Vertices.Front(); vtx1 != nil; vtx1 = vtx1.Next() {
		srcNode := vtx1.Value.(*Vertex)
		srcNodeID := vtx1.Value.(*Vertex).ID

		tmap := make(map[string][]float64)
		for vtx2 := srcNode.GetOutVertices().Front(); vtx2 != nil; vtx2 = vtx2.Next() {
			ovNode := vtx2.Value.(*Vertex)
			ovNodeID := vtx2.Value.(*Vertex).ID
			wgt := g.GetEdgeWeight(srcNode, ovNode)

			// Not here
			// tmap := make(map[string][]float64)
			tmap[ovNodeID] = []float64{wgt}

			// Not here
			// rm[srcNodeID] = tmap
		}
		rm[srcNodeID] = tmap
	}

	return rm
}

// ToJSON converts a receiver graph data structure to JSON format.
func (g Graph) ToJSON() string {
	/*
	   {
	       "testgraph.017": {
	           "S": {
	               "A": [10],
	               "B": [5],
	               "C": [15]
	           },
	           "A": {
	               "B": [4],
	               "D": [9],
	               "E": [15]
	   }
	*/
	// map[string]map[string][]float64
	rm := g.ToMAP()
	tmplSrcvtx := "\t\t\"%s\": {\n"
	tmplOutvtx := "\t\t\t\"%s\": [%v]"
	rstr := ""
	cn := 0
	for srcNodeID, outMap := range rm {
		vtxSlice := []string{}
		if cn == 0 {
			rstr = rstr + fmt.Sprintf(tmplSrcvtx, srcNodeID)
		} else {
			rstr = rstr + ",\n" + fmt.Sprintf(tmplSrcvtx, srcNodeID)
		}
		cn++
		outSlice := []string{}
		for outNodeID, fs := range outMap {
			fsl := []string{}
			for _, val := range fs {
				fsl = append(fsl, fmt.Sprintf("%v", val))
			}
			fslStr := strings.Join(fsl, ", ")
			tstr := fmt.Sprintf(tmplOutvtx, outNodeID, fslStr)
			outSlice = append(outSlice, tstr)
		}
		vtxSlice = append(vtxSlice, strings.Join(outSlice, ",\n")+"\n\t\t}")
		rstr = rstr + strings.Join(vtxSlice, ",")
	}

	line0 := "{\n"
	line1 := "\t\"goraph\": {\n"
	line2 := "\n\t}\n}"
	// rstr = rstr + lineE

	return line0 + line1 + rstr + line2
}

// ToJSONFile converts a graph to a JSON file.
func (g Graph) ToJSONFile(fpath string) {
	jstr := g.ToJSON()
	helper.WriteToFile(fpath, jstr)
}

// ToDOT converts a receiver graph data structure to DOT format.
func (g Graph) ToDOT() string {
	/*
	   digraph testgraph003 {
	   	S -> A [label=100]
	   	S -> B [label=6]
	   	S -> B [label=14]
	   	...
	*/
	// map[string]map[string][]float64
	rm := g.ToMAP()
	tmpl := "\t%s -> %s [label=%v]"
	lines := []string{}
	for srcNodeID, outMap := range rm {
		for outNodeID, fs := range outMap {
			lines = append(lines, fmt.Sprintf(tmpl, srcNodeID, outNodeID, fs[0]))
		}
	}

	line0 := "digraph goraph {\n"
	lineE := "\n}"
	return line0 + strings.Join(lines, "\n") + lineE
}

// ToDOTFile converts a graph to a DOT file.
func (g Graph) ToDOTFile(fpath string) {
	dstr := g.ToDOT()
	helper.WriteToFile(fpath, dstr)
}

// GetVertices returns the vertex slice.
func (g Graph) GetVertices() *list.List {
	return g.Vertices
}

// GetVerticesSize returns the size of vertex slice in a graph.
func (g Graph) GetVerticesSize() int {
	// dereference
	return g.Vertices.Len()
}

// GetEdges returns the edge slice.
func (g Graph) GetEdges() *list.List {
	return g.Edges
}

// GetEdge returns the Edge from src to dst Vertex.
// (Assume that there is no duplicate Edge for now.)
func (g Graph) GetEdge(src, dst *Vertex) *Edge {
	for edge := g.GetEdges().Front(); edge != nil; edge = edge.Next() {
		if edge.Value.(*Edge).Src.ID == src.ID && edge.Value.(*Edge).Dst.ID == dst.ID {
			return edge.Value.(*Edge)
		}
	}
	return nil
}

// GetEdgesSize returns the size of edge slice in a graph.
func (g Graph) GetEdgesSize() int {
	// dereference
	return g.Edges.Len()
}

// FindVertexByID returns the vertex with input ID, or return nil if it doesn't exist.
func (g Graph) FindVertexByID(id string) *Vertex {
	// slice := g.GetVertices()
	for vtx := g.Vertices.Front(); vtx != nil; vtx = vtx.Next() {
		// NOT  vtx.Value.(Vertex).ID
		if fmt.Sprintf("%v", vtx.Value.(*Vertex).ID) == fmt.Sprintf("%v", id) {
			return vtx.Value.(*Vertex)
		}
	}
	return nil
}

// CreateAndAddToGraph finds the vertex with the ID, or create it.
func (g *Graph) CreateAndAddToGraph(id string) *Vertex {
	vtx := g.FindVertexByID(id)
	if vtx == nil {
		vtx = NewVertex(id)
		// then add this vertex to the graph
		g.AddVertex(vtx)
	}
	return vtx
}

// AddVertex adds the Vertex v to a graph's Vertices.
func (g *Graph) AddVertex(v *Vertex) {
	g.Vertices.PushBack(v)
}

// AddEdge adds the Edge e to a graph's Edges.
func (g *Graph) AddEdge(e *Edge) {
	g.Edges.PushBack(e)
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

// ShowPrev shows the Prev of Vertex.
// This is useful when debugging Dijkstra shortest path algorithm.
func (g Graph) ShowPrev(src *Vertex) string {
	sl := []string{}
	for vtx := src.Prev.Front(); vtx != nil; vtx = vtx.Next() {
		if vtx == nil {
			continue
		}
		sl = append(sl, vtx.Value.(*Vertex).ID)
	}

	return "Prev of " + src.ID + ": " + strings.Join(sl, ", ")
}

// GetEdgeWeight returns the weight value of the edge from source to destination vertex.
func (g Graph) GetEdgeWeight(src, dst *Vertex) float64 {
	for edge := g.GetEdges().Front(); edge != nil; edge = edge.Next() {
		if edge.Value.(*Edge).Src == src && edge.Value.(*Edge).Dst == dst {
			return edge.Value.(*Edge).Weight
		}
	}
	return 0.0
}

// UpdateWeight updates the weight value between vertices.
func (g *Graph) UpdateWeight(src, dst *Vertex, value float64) {
	for edge := g.GetEdges().Front(); edge != nil; edge = edge.Next() {
		if edge.Value.(*Edge).Src == src && edge.Value.(*Edge).Dst == dst {
			edge.Value.(*Edge).Weight = value
		}
	}
}

// Connect connects the vertex v to A, not A to v.
// When there is more than one edge, it adds up the weight values.
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

// DeleteVertex removes the input vertex A from the graph g's Vertices.
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
	var nextElem1 *list.Element
	for vtx := A.GetOutVertices().Front(); vtx != nil; vtx = nextElem1 {
		nextElem1 = vtx.Next()
		if vtx.Value.(*Vertex) == B {
			A.GetOutVertices().Remove(vtx)
		}
	}

	// delete A from B's InVertices
	var nextElem2 *list.Element
	for vtx := B.GetInVertices().Front(); vtx != nil; vtx = nextElem2 {
		nextElem2 = vtx.Next()
		if vtx.Value.(*Vertex) == A {
			B.GetInVertices().Remove(vtx)
		}
	}

	// Always delete from graph at the end
	// remove the edge from the graph's edge list
	var nextElem3 *list.Element
	for edge := g.GetEdges().Front(); edge != nil; edge = nextElem3 {
		nextElem3 = edge.Next()
		// if the edge is from A to B
		if edge.Value.(*Edge).Src == A && edge.Value.(*Edge).Dst == B {
			// don't do this
			// edge.Value.(*Edge).Src = nil
			// edge.Value.(*Edge).Dst = nil
			g.GetEdges().Remove(edge)
		}
	}
}
