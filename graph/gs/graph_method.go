package gs

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	slice "github.com/gyuho/goraph/gosequence"
)

// ToMAP converts a receiver graph data structure to a map.
func (g Graph) ToMAP() map[string]map[string][]float64 {
	rm := make(map[string]map[string][]float64)

	gVts := g.GetVertices()
	for _, src := range *gVts {

		srcNode := src.(*Vertex)
		srcNodeID := src.(*Vertex).ID

		outVts := srcNode.GetOutVertices()
		tmap := make(map[string][]float64)
		for _, ov := range *outVts {
			ovNode := ov.(*Vertex)
			ovNodeID := ov.(*Vertex).ID
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

// OpenToOverwrite creates or opens a file for overwriting.
// Make sure to close the file.
func OpenToOverwrite(fpath string) *os.File {
	file, err := os.OpenFile(fpath, os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		// log.Fatal(err)
		file, err = os.Create(fpath)
		if err != nil {
			log.Fatal(err)
		}
	}
	return file
}

// WriteToFile writes the input string slice into a text file.
func WriteToFile(fpath, str string) {
	file := OpenToOverwrite(fpath)
	defer file.Close()
	txt := bufio.NewWriter(file)
	_, err := txt.WriteString(str)
	if err != nil {
		log.Fatal(err)
	}
	defer txt.Flush()
}

// ToJSONFile converts a graph to a JSON file.
func (g Graph) ToJSONFile(fpath string) {
	jstr := g.ToJSON()
	WriteToFile(fpath, jstr)
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
	WriteToFile(fpath, dstr)
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

// GetEdge returns the Edge from src to dst Vertex.
// (Assume that there is no duplicate Edge for now.)
func (g Graph) GetEdge(src, dst *Vertex) *Edge {
	slice := g.GetEdges()
	for _, edge := range *slice {
		//
		// This does not work when used with CopyGraph and SameGraph
		// because the copy has different pointer values.
		// (X) if edge.(*Edge).Src == src && edge.(*Edge).Dst == dst {
		//
		if edge.(*Edge).Src.ID == src.ID && edge.(*Edge).Dst.ID == dst.ID {
			return edge.(*Edge)
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
	slice := g.GetVertices()
	for _, v := range *slice {
		if v.(*Vertex).ID == id {
			return v.(*Vertex)
		}
	}
	// nil is used as pointer
	// null pointer
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
	_, exist := A.OutVertices.Find(B)
	if exist {
		return true
	}
	return false
}

// ShowPrev shows the Prev of Vertex.
// This is useful when debugging Dijkstra shortest path algorithm.
func (g Graph) ShowPrev(src *Vertex) string {
	sl := []string{}
	for _, p := range *src.Prev {
		if p == nil {
			continue
		}
		sl = append(sl, p.(*Vertex).ID)
	}
	return "Prev of " + src.ID + ": " + strings.Join(sl, ", ")
}

// GetEdgeWeight returns the weight value of the edge from source to destination vertex.
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
