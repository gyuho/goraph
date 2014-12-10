package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
)

// Use Pointer when we need to update the struct with receiver
// https://golang.org/doc/faq#methods_on_values_or_pointers

// Work Flow
// 1. Create a graph `Data`.
// 2. Create a `Vertex`.
// 3. Add a `Vertex` to a graph Data.
// 4. Connect with an Edge with `AddEdge`

// Data contains graph data, represented in adjacency list and slice.
type Data struct {
	Vertices []*Vertex
	Edges    []*Edge
}

type Data1 struct {
	Vertices []*Vertex
	sync.Mutex

	// edgeFrom maps each source Vertex to its destination(target)
	// Vertices with weight values
	edgeFrom map[*Vertex]map[*Vertex]float64

	// edgeTo maps each destination Vertex to its source(incoming)
	// Vertices with weight values
	edgeTo map[*Vertex]map[*Vertex]float64
}

// In progress
// Use hash to map each vertex to ~
// http://en.wikipedia.org/wiki/Adjacency_list

// Vertex is a vertex(node) in Graph.
type Vertex struct {
	// ID of Vertex is assumed to be unique between vertices.
	ID string

	// Color is used for graph traversal.
	Color string

	sync.Mutex

	// record stores records for graph algorithms.
	record map[string]float64
}

// Edge is an edge(arc) in a graph that has a direction
// from `Source` vertex to `Destination` vertex.
type Edge struct {
	Source      *Vertex
	Destination *Vertex

	// Weight contains the weight value in float64.
	Weight float64
}

// NewData returns a new Data.
func NewData() *Data {
	return &Data{
		Vertices: []*Vertex{},
		Edges:    []*Edge{},
	}
}

// NewVertex returns a new Vertex.
func NewVertex(id string) *Vertex {
	return &Vertex{
		ID:     id,
		Color:  "",
		record: make(map[string]float64),
	}
}

// AddVertex adds a vertex to a graph Data.
func (d *Data) AddVertex(vtx *Vertex) {
	d.Vertices = append(d.Vertices, vtx)
}

// AddRecord adds a record to a Vertex.
func (v *Vertex) AddRecord(overWrite bool, key string, value float64) error {
	v.Mutex.Lock()
	if !overWrite {
		if val, ok := v.record[key]; ok {
			return fmt.Errorf("%s already exists with %f | %+v", key, val, v.record)
		}
	}
	v.record[key] = value
	v.Mutex.Unlock()
	return nil
}

// DeleteRecord deletes a record from a Vertex, by its key.
func (v *Vertex) DeleteRecord(key string) {
	v.Mutex.Lock()
	delete(v.record, key)
	v.Mutex.Unlock()
}

// Connect adds an edge from src to dst Vertex, to a graph Data.
func (d *Data) Connect(src, dst *Vertex, weight float64) {
	isDuplicate := false
	for _, edge := range d.Edges {
		if edge.Source == src {
			if edge.Destination == dst {
				log.Printf("Overwriting Edge Weight:\n%s -- [Weight %f] --> %s\n",
					edge.Source.ID, edge.Weight, edge.Destination.ID)
				log.Printf("%s -- [Weight %f] --> %s\n",
					edge.Source.ID, weight, edge.Destination.ID)
				edge.Weight = weight
				isDuplicate = true
			}
		}
	}
	if !isDuplicate {
		newEdge := Edge{
			Source:      src,
			Destination: dst,
			Weight:      weight,
		}
		d.Edge = append(d.Edge, &newEdge)
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

// String describes the graph Data.
func (d Data) String() string {
	if len(d.Vertices) == 0 {
		return "Graph is empty."
	}
	slice := []string{}
	for _, vtx := range d.Vertices {
		slice = append(slice, fmt.Sprintf("Vertex: %s", vtx.ID))
		d.Mutex.Lock()
		if _, ok := d.OutEdges[vtx]; !ok {
			slice = append(slice, fmt.Sprintf("No Outgoing Edge from %s", vtx.ID))
		} else {
			for _, edge := range d.OutEdges[vtx] {
				slice = append(slice, fmt.Sprintf("Outgoing Edges: [%s] -- %f --> [%s]\n", edge.Vtx.ID, edge.Weight, vtx.ID))
			}
		}
		if _, ok := d.InEdges[vtx]; !ok {
			slice = append(slice, fmt.Sprintf("No Incoming Edge from %s", vtx.ID))
		} else {
			for _, edge := range d.InEdges[vtx] {
				slice = append(slice, fmt.Sprintf("Incoming Edges: [%s] -- %f --> [%s]\n", edge.Vtx.ID, edge.Weight, vtx.ID))
			}
		}
		slice = append(slice, "\n")
		d.Mutex.Unlock()
	}
	return strings.Join(slice, "\n")
}

// FindVertexByID finds a Vertex by ID.
func (d Data) FindVertexByID(id string) *Vertex {
}

// DeleteVertex deletes a Vertex from the graph Data.
func (d *Data) DeleteVertex(vtx *Vertex) {

}

// DeleteEdge deletes an Edge from src to dst from the graph Data.
func (d *Data) DeleteEdge(src, dst *Vertex) {

}

// Clone clones the graph Data.
// It does `Deep Copy`.
// That is, changing the cloned Data would not affect the original Data.
func (d *Data) Clone() *Data {

}
