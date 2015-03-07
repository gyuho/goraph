package graph

import (
	"fmt"
	"log"
	"strings"
	"sync"
)

// Use Pointer when we need to update the struct with receiver
// https://golang.org/doc/faq#methods_on_values_or_pointers

// Data contains graph data, represented in adjacency list and slice.
type Data struct {
	Vertices []*Vertex

	sync.Mutex

	// OutEdges maps each Vertex to its outgoing edges
	OutEdges map[*Vertex][]*Edge

	// InEdges maps each Vertex to its incoming edges
	InEdges map[*Vertex][]*Edge

	// to prevent duplicating vertex IDs
	vertexIDs map[string]bool
}

// Vertex is a vertex(node) in Graph.
type Vertex struct {
	// ID of Vertex is assumed to be unique between vertices.
	ID string

	// Color is used for graph traversal.
	Color string

	sync.Mutex

	// Stamp stores stamp records for several graph algorithms.
	Stamp map[string]float64
}

// Edge is an edge(arc) in a graph that has direction from one to another vertex.
type Edge struct {
	// Vtx can be either source or destination
	Vtx *Vertex

	// Weight contains the weight value in float64.
	// Note that `Weight` is a single floating value.
	// Define with []float64 if we want duplicate edge values.
	Weight float64
}

// NewData returns a new Data.
func NewData() *Data {
	return &Data{
		Vertices:  []*Vertex{},
		OutEdges:  make(map[*Vertex][]*Edge),
		InEdges:   make(map[*Vertex][]*Edge),
		vertexIDs: make(map[string]bool),
	}
}

// NewVertex returns a new Vertex.
func NewVertex(id string) *Vertex {
	return &Vertex{
		ID:    id,
		Color: "",
		Stamp: make(map[string]float64),
	}
}

// AddVertex adds a vertex to a graph Data.
func (d *Data) AddVertex(vtx *Vertex) bool {
	d.Lock()
	if _, ok := d.vertexIDs[vtx.ID]; ok {
		d.Unlock()
		return false
	}
	d.vertexIDs[vtx.ID] = true
	d.Unlock()
	d.Vertices = append(d.Vertices, vtx)
	return true
}

// Connect adds an edge from src to dst Vertex, to a graph Data.
func (d *Data) Connect(src, dst *Vertex, weight float64) {
	isAddedSrc := d.AddVertex(src)
	if !isAddedSrc {
		log.Printf("`%s` was previously added to Data\n", src.ID)
		src = d.FindVertexByID(src.ID)
	} else {
		log.Printf("`%s` is added to Data\n", src.ID)
	}
	isAddedDst := d.AddVertex(dst)
	if !isAddedDst {
		log.Printf("`%s` was previously added to Data\n", dst.ID)
		dst = d.FindVertexByID(dst.ID)
	} else {
		log.Printf("`%s` is added to Data\n", dst.ID)
	}
	edgeSrc := Edge{
		Vtx:    src,
		Weight: weight,
	}
	edgeDst := Edge{
		Vtx:    dst,
		Weight: weight,
	}

	d.Lock()

	// update Outgoing Edges
	if _, ok := d.OutEdges[src]; !ok {
		d.OutEdges[src] = []*Edge{&edgeDst}
	} else {
		// if OutEdges already exists
		isDuplicate := false
		for _, edge := range d.OutEdges[src] {
			// if there is a duplicate(parallel) edge
			if edge.Vtx == dst {
				log.Println("Duplicate(Parallel) Edge Found. Overwriting the Weight value.")
				log.Printf("%v --> %v + %v\n", edge.Weight, edge.Weight, weight)
				edge.Weight += weight
				// d.OutEdges[src][idx] = edge
				isDuplicate = true
				break
			}
		}
		// if this is just another edge from `src` Vertex
		if !isDuplicate {
			d.OutEdges[src] = append(d.OutEdges[src], &edgeDst)
		}
	}

	// update Incoming Edges
	if _, ok := d.InEdges[dst]; !ok {
		d.InEdges[dst] = []*Edge{&edgeSrc}
	} else {
		// if InEdges already exists
		isDuplicate := false
		for _, edge := range d.InEdges[dst] {
			// if there is a duplicate(parallel) edge
			if edge.Vtx == src {
				log.Println("Duplicate(Parallel) Edge Found. Overwriting the Weight value.")
				log.Printf("%v --> %v + %v\n", edge.Weight, edge.Weight, weight)
				edge.Weight += weight

				//
				// if
				// OutEdges map[*Vertex][]Edge
				//
				// `range` iterates over values
				// Make sure to overwrite with assignment
				// d.InEdges[dst][idx] = edge
				isDuplicate = true
				break
			}
		}
		// if this is just another edge to `dst` Vertex
		if !isDuplicate {
			d.InEdges[dst] = append(d.InEdges[dst], &edgeSrc)
		}
	}

	d.Unlock()
}

// Init initializes the graph Data.
func (d *Data) Init() {
	// (X) d = NewData()
	// this only updates the pointer
	//
	// Do this.
	*d = *NewData()
}

// GetVertexSize returns the size of Vertex of the graph Data.
func (d Data) GetVertexSize() int64 {
	return int64(len(d.Vertices))
}

// String describes the graph Data.
func (d Data) String() string {
	if d.GetVertexSize() == 0 {
		return "Graph is empty."
	}
	slice := []string{}
	for _, vtx := range d.Vertices {
		vtxLabel := fmt.Sprintf("Vertex: %s | ", vtx.ID)
		if _, ok := d.OutEdges[vtx]; !ok {
			slice = append(slice, fmt.Sprintf(vtxLabel+"Outgoing Edges: [%s] -- none", vtx.ID))
		} else {
			for _, edge := range d.OutEdges[vtx] {
				slice = append(slice, fmt.Sprintf(vtxLabel+"Outgoing Edges: [%s] -- %.3f --> [%s]", vtx.ID, edge.Weight, edge.Vtx.ID))
			}
		}
		if _, ok := d.InEdges[vtx]; !ok {
			slice = append(slice, fmt.Sprintf(vtxLabel+"Incoming Edges: none --> [%s]", vtx.ID))
		} else {
			for _, edge := range d.InEdges[vtx] {
				slice = append(slice, fmt.Sprintf(vtxLabel+"Incoming Edges: [%s] -- %.3f --> [%s]", edge.Vtx.ID, edge.Weight, vtx.ID))
			}
		}
	}
	return strings.Join(slice, "\n")
}

// FindVertexByID finds a Vertex by ID.
func (d Data) FindVertexByID(id string) *Vertex {
	for _, vtx := range d.Vertices {
		if vtx.ID == id {
			return vtx
		}
	}
	return nil
}

// DeleteVertex deletes a Vertex from the graph Data.
func (d *Data) DeleteVertex(vtx *Vertex) {

	// delete from d.Vertices
	for idx, elem := range d.Vertices {
		if elem == vtx {
			copy(d.Vertices[idx:], d.Vertices[idx+1:])
			d.Vertices[len(d.Vertices)-1] = nil // zero value of type or nil
			d.Vertices = d.Vertices[:len(d.Vertices)-1 : len(d.Vertices)-1]
			break
		}
	}

	// delete edges from neighbor vertex's outgoing, incoming edges
	for _, edge1 := range d.OutEdges[vtx] {
		for idx, edge2 := range d.OutEdges[edge1.Vtx] {
			if edge2.Vtx == vtx {
				copy(d.OutEdges[edge1.Vtx][idx:], d.OutEdges[edge1.Vtx][idx+1:])
				d.OutEdges[edge1.Vtx][len(d.OutEdges[edge1.Vtx])-1] = nil
				d.OutEdges[edge1.Vtx] = d.OutEdges[edge1.Vtx][:len(d.OutEdges[edge1.Vtx])-1 : len(d.OutEdges[edge1.Vtx])-1]
				break
			}
		}
		for idx, edge2 := range d.InEdges[edge1.Vtx] {
			if edge2.Vtx == vtx {
				copy(d.InEdges[edge1.Vtx][idx:], d.InEdges[edge1.Vtx][idx+1:])
				d.InEdges[edge1.Vtx][len(d.InEdges[edge1.Vtx])-1] = nil
				d.InEdges[edge1.Vtx] = d.InEdges[edge1.Vtx][:len(d.InEdges[edge1.Vtx])-1 : len(d.InEdges[edge1.Vtx])-1]
				break
			}
		}
	}
	for _, edge1 := range d.InEdges[vtx] {
		for idx, edge2 := range d.OutEdges[edge1.Vtx] {
			if edge2.Vtx == vtx {
				copy(d.OutEdges[edge1.Vtx][idx:], d.OutEdges[edge1.Vtx][idx+1:])
				d.OutEdges[edge1.Vtx][len(d.OutEdges[edge1.Vtx])-1] = nil
				d.OutEdges[edge1.Vtx] = d.OutEdges[edge1.Vtx][:len(d.OutEdges[edge1.Vtx])-1 : len(d.OutEdges[edge1.Vtx])-1]
				break
			}
		}
		for idx, edge2 := range d.InEdges[edge1.Vtx] {
			if edge2.Vtx == vtx {
				copy(d.InEdges[edge1.Vtx][idx:], d.InEdges[edge1.Vtx][idx+1:])
				d.InEdges[edge1.Vtx][len(d.InEdges[edge1.Vtx])-1] = nil
				d.InEdges[edge1.Vtx] = d.InEdges[edge1.Vtx][:len(d.InEdges[edge1.Vtx])-1 : len(d.InEdges[edge1.Vtx])-1]
				break
			}
		}
	}

	// delete from maps
	d.Lock()
	delete(d.OutEdges, vtx)
	delete(d.InEdges, vtx)
	delete(d.vertexIDs, vtx.ID)
	d.Unlock()
}

// DeleteEdge deletes an Edge from src to dst from the graph Data.
// This does not delete Vertices.
func (d *Data) DeleteEdge(src, dst *Vertex) {
	// delete an edge from OutEdges
	for idx, edge := range d.OutEdges[src] {
		if edge.Vtx == dst {
			copy(d.OutEdges[src][idx:], d.OutEdges[src][idx+1:])
			d.OutEdges[src][len(d.OutEdges[src])-1] = nil // zero value of type or nil
			d.OutEdges[src] = d.OutEdges[src][:len(d.OutEdges[src])-1 : len(d.OutEdges[src])-1]
			break
		}
	}
	// delete an edge from InEdges
	for idx, edge := range d.InEdges[dst] {
		if edge.Vtx == src {
			copy(d.InEdges[dst][idx:], d.InEdges[dst][idx+1:])
			d.InEdges[dst][len(d.InEdges[dst])-1] = nil
			d.InEdges[dst] = d.InEdges[dst][:len(d.InEdges[dst])-1 : len(d.InEdges[dst])-1]
			break
		}
	}
}

// GetEdgeWeight returns the weight value of an edge from src to dst Vertex.
func (d Data) GetEdgeWeight(src, dst *Vertex) float64 {
	for _, edge := range d.OutEdges[src] {
		if edge.Vtx == dst {
			return edge.Weight
		}
	}
	return 0.0
}

// UpdateEdgeWeight overwrites an Edge's weight value.
func (d *Data) UpdateEdgeWeight(src, dst *Vertex, weight float64) {
	// `range` iterates over values
	// Make sure to overwrite with assignment
	for idx, edge := range d.OutEdges[src] {
		if edge.Vtx == dst {
			edge.Weight = weight
			d.OutEdges[src][idx] = edge
			break
		}
	}
}
