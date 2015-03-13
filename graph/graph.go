package graph

import (
	"bytes"
	"fmt"
	"log"
	"sync"
)

// Use Pointer when we need to update the struct with receiver
// https://golang.org/doc/faq#methods_on_values_or_pointers

// Data contains graph data, represented in adjacency list and slice.
type Data struct {
	sync.Mutex
	VertexMap map[*Vertex]bool

	// maintain vertexID to prevent having duplicate Vertex ID
	vertexID map[string]bool
}

// NewData returns a new Data.
func NewData() *Data {
	return &Data{
		VertexMap: make(map[*Vertex]bool),
		vertexID:  make(map[string]bool),
		// without this
		// panic: assignment to entry in nil map
	}
}

// Init initializes the graph Data.
func (d *Data) Init() {
	// (X) d = NewData()
	// this only updates the pointer
	//
	// Do this.
	*d = *NewData()
}

// Vertex is a vertex(node) in Graph.
type Vertex struct {
	// ID of Vertex is assumed to be unique among vertices.
	ID string

	// Color is used for graph traversal.
	Color string

	// Stamp stores stamp records for graph algorithms.
	Stamp map[string]float64

	sync.Mutex
	WeightTo   map[*Vertex]float64
	WeightFrom map[*Vertex]float64
}

// NewVertex returns a new Vertex.
func NewVertex(id string) *Vertex {
	return &Vertex{
		WeightTo:   make(map[*Vertex]float64),
		WeightFrom: make(map[*Vertex]float64),
		ID:         id,
		Color:      "white",
		Stamp:      make(map[string]float64),
	}
}

// AddVertex adds a vertex to a graph Data.
func (d *Data) AddVertex(vtx *Vertex) bool {
	d.Lock()
	defer d.Unlock()
	if _, ok := d.vertexID[vtx.ID]; ok {
		return false
	}
	if _, ok := d.VertexMap[vtx]; ok {
		return false
	}
	d.vertexID[vtx.ID] = true
	d.VertexMap[vtx] = true
	return true
}

// GetVertexSize returns the size of Vertex of the graph Data.
func (d Data) GetVertexSize() int64 {
	return int64(len(d.VertexMap))
}

// FindVertexByID finds a Vertex by ID.
func (d Data) FindVertexByID(id string) *Vertex {
	for vtx := range d.VertexMap {
		if vtx.ID == id {
			return vtx
		}
	}
	return nil
}

// Connect adds an edge from src(source) to dst(destination) Vertex, to a graph Data.
// This doese not connect from dst to src.
func (d *Data) Connect(src, dst *Vertex, weight float64) {

	// do not allow a circle
	// if src.ID == dst.ID {
	// 	return
	// }

	// add to Data
	if !d.AddVertex(src) {
		// log.Printf("`%s` was previously added to Data\n", src.ID)
		src = d.FindVertexByID(src.ID)
	}
	if !d.AddVertex(dst) {
		// log.Printf("`%s` was previously added to Data\n", dst.ID)
		dst = d.FindVertexByID(dst.ID)
	}

	d.Lock()
	defer d.Unlock()

	// update src Vertex
	if v, ok := src.WeightTo[dst]; !ok {
		src.WeightTo[dst] = weight
	} else {
		log.Printf("Duplicate(Parallel) Edge. Overwriting the Weight value: %s --> %.3f --> %s (new weight: %.3f)\n", src.ID, v, dst.ID, v+weight)
		src.WeightTo[dst] = v + weight
	}

	// update dst Vertex
	if v, ok := dst.WeightFrom[src]; !ok {
		dst.WeightFrom[src] = weight
	} else {
		log.Printf("Duplicate(Parallel) Edge. Overwriting the Weight value: %s --> %.3f --> %s (new weight: %.3f)\n", src.ID, v, dst.ID, v+weight)
		dst.WeightFrom[src] = v + weight
	}

}

// GetEdgeWeight returns the weight value of an edge from src to dst Vertex.
func (d Data) GetEdgeWeight(src, dst *Vertex) float64 {
	if _, ok := src.WeightTo[dst]; !ok {
		return 0.0
	}
	return src.WeightTo[dst]
}

// UpdateEdgeWeight overwrites the edge weight from src to dst Vertex.
func (d Data) UpdateEdgeWeight(src, dst *Vertex, weight float64) {
	src.WeightTo[dst] = weight
}

// DeleteVertex deletes a Vertex from the graph Data.
// This deletes all the related edges too.
func (d *Data) DeleteVertex(vtx *Vertex) {

	// delete edges from each Vertex
	for elem := range d.VertexMap {
		if elem == vtx {
			continue
		}
		elem.Lock()
		delete(elem.WeightFrom, vtx)
		delete(elem.WeightTo, vtx)
		elem.Unlock()
	}

	// delete from Data(graph)
	d.Lock()
	delete(d.VertexMap, vtx)
	delete(d.vertexID, vtx.ID)
	d.Unlock()

	vtx = nil
}

// DeleteEdge deletes an Edge from src to dst from the graph Data.
// This does not delete Vertices.
func (d *Data) DeleteEdge(src, dst *Vertex) {
	src.Lock()
	delete(src.WeightTo, dst)
	src.Unlock()

	dst.Lock()
	delete(dst.WeightFrom, src)
	dst.Unlock()
}

// String describes the graph Data.
func (d Data) String() string {
	buf := new(bytes.Buffer)
	if d.GetVertexSize() == 0 {
		return "Graph is empty."
	}
	buf.WriteString(fmt.Sprintf("Graph has %d vertices\n", d.GetVertexSize()))
	for vtx := range d.VertexMap {
		vtxLabel := fmt.Sprintf("Vertex: %s | ", vtx.ID)
		if len(vtx.WeightTo) != 0 {
			for dst, weight := range vtx.WeightTo {
				buf.WriteString(fmt.Sprintf(vtxLabel+"Outgoing Edge: [%s] -- %.3f --> [%s]\n", vtx.ID, weight, dst.ID))
			}
		}
		if len(vtx.WeightFrom) != 0 {
			for src, weight := range vtx.WeightFrom {
				buf.WriteString(fmt.Sprintf(vtxLabel+"Incoming Edge: [%s] -- %.3f --> [%s]\n", src.ID, weight, vtx.ID))
			}
		}
	}
	return buf.String()
}
