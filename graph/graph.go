package graph

import (
	"bytes"
	"fmt"
	"sync"
)

// Data contains graph data, represented in adjacency list and slice.
// Make sure to use Pointer when we need to update the struct with receiver.
// (https://golang.org/doc/faq#methods_on_values_or_pointers)
type Data struct {
	sync.Mutex

	// NodeMap is a hash-map for all Nodes in the graph.
	NodeMap map[*Node]bool

	// maintain nodeID in order not to have duplicate Node IDs in the graph.
	nodeID map[string]bool
}

// NewData returns a new Data.
func NewData() *Data {
	return &Data{
		NodeMap: make(map[*Node]bool),
		nodeID:  make(map[string]bool),
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

// Node is a Node(node) in Graph.
type Node struct {

	// ID of Node is assumed to be unique among Nodes.
	ID string

	// Color is used for graph traversal.
	Color string

	sync.Mutex

	// WeightTo maps its Node to outgoing Nodes with its edge weight (outgoing edges from its Node).
	WeightTo map[*Node]float32

	// WeightFrom maps its Node to incoming Nodes with its edge weight (incoming edges to its Node).
	WeightFrom map[*Node]float32

	// Stamp stores stamp records for graph algorithms.
	// Stamp map[string]float32
}

// NewNode returns a new Node.
func NewNode(id string) *Node {
	return &Node{
		ID:         id,
		Color:      "white",
		WeightTo:   make(map[*Node]float32),
		WeightFrom: make(map[*Node]float32),
		// Stamp: make(map[string]float32),
	}
}

// AddNode adds a Node to a graph Data.
// It returns true if the Node is added the graph Data.
func (d *Data) AddNode(nd *Node) bool {

	if nd == nil {
		return false
	}

	d.Lock()
	defer d.Unlock()
	if _, ok := d.nodeID[nd.ID]; ok {
		return false
	}
	if _, ok := d.NodeMap[nd]; ok {
		return false
	}
	d.nodeID[nd.ID] = true
	d.NodeMap[nd] = true
	return true
}

// GetNodeSize returns the size of Node of the graph Data.
func (d Data) GetNodeSize() int {
	return len(d.NodeMap)
}

//GetNodeByID finds a Node by ID.
func (d Data) GetNodeByID(id string) *Node {
	for nd := range d.NodeMap {
		if nd.ID == id {
			return nd
		}
	}
	return nil
}

// Connect adds an edge from src(source) to dst(destination) Node, to a graph Data.
// This doese not connect from dst to src.
func (d *Data) Connect(src, dst *Node, weight float32) {

	if src == nil || dst == nil {
		return
	}

	// do not allow a circle
	// if src.ID == dst.ID {
	// 	return
	// }

	// add to Data
	if !d.AddNode(src) {
		// log.Printf("`%s` was previously added to Data\n", src.ID)
		src = d.GetNodeByID(src.ID)
	}
	if !d.AddNode(dst) {
		// log.Printf("`%s` was previously added to Data\n", dst.ID)
		dst = d.GetNodeByID(dst.ID)
	}

	d.Lock()
	defer d.Unlock()

	// update src Node
	if v, ok := src.WeightTo[dst]; !ok {
		src.WeightTo[dst] = weight
	} else {
		// log.Printf("Duplicate(Parallel) Edge. Overwriting the Weight value: %s --> %.3f --> %s (new weight: %.3f)\n", src.ID, v, dst.ID, v+weight)
		src.WeightTo[dst] = v + weight
	}

	// update dst Node
	if v, ok := dst.WeightFrom[src]; !ok {
		dst.WeightFrom[src] = weight
	} else {
		// log.Printf("Duplicate(Parallel) Edge. Overwriting the Weight value: %s --> %.3f --> %s (new weight: %.3f)\n", src.ID, v, dst.ID, v+weight)
		dst.WeightFrom[src] = v + weight
	}

}

// GetEdgeWeight returns the weight value of an edge from src to dst Node.
func (d Data) GetEdgeWeight(src, dst *Node) float32 {
	if src == nil || dst == nil {
		return 0.0
	}
	if _, ok := src.WeightTo[dst]; !ok {
		return 0.0
	}
	return src.WeightTo[dst]
}

// DeleteNode deletes a Node from the graph Data.
// This deletes all the related edges too.
func (d *Data) DeleteNode(nd *Node) {

	if nd == nil {
		return
	}

	// delete edges from each Node
	for elem := range d.NodeMap {
		if elem == nd {
			continue
		}
		elem.Lock()
		if _, ok := elem.WeightFrom[nd]; ok {
			delete(elem.WeightFrom, nd)
		}
		if _, ok := elem.WeightTo[nd]; ok {
			delete(elem.WeightTo, nd)
		}
		elem.Unlock()
	}

	// delete from Data(graph)
	d.Lock()
	if _, ok := d.NodeMap[nd]; ok {
		delete(d.NodeMap, nd)
	}
	if _, ok := d.nodeID[nd.ID]; ok {
		delete(d.nodeID, nd.ID)
	}
	d.Unlock()

	nd = nil
}

// DeleteEdge deletes an Edge from src to dst from the graph Data.
// This does not delete Nodes.
func (d *Data) DeleteEdge(src, dst *Node) {

	if src == nil || dst == nil {
		return
	}

	src.Lock()
	if _, ok := src.WeightTo[dst]; ok {
		delete(src.WeightTo, dst)
	}
	src.Unlock()

	dst.Lock()
	if _, ok := dst.WeightFrom[src]; ok {
		delete(dst.WeightFrom, src)
	}
	dst.Unlock()

}

// String describes the graph Data.
func (d Data) String() string {
	buf := new(bytes.Buffer)
	if d.GetNodeSize() == 0 {
		return "Graph is empty."
	}
	buf.WriteString(fmt.Sprintf("Graph has %d Nodes\n", d.GetNodeSize()))
	for nd := range d.NodeMap {
		ndLabel := fmt.Sprintf("Node: %s | ", nd.ID)
		if len(nd.WeightTo) != 0 {
			for dst, weight := range nd.WeightTo {
				buf.WriteString(fmt.Sprintf(ndLabel+"Outgoing Edge: [%s] -- %.3f --> [%s]\n", nd.ID, weight, dst.ID))
			}
		}
		if len(nd.WeightFrom) != 0 {
			for src, weight := range nd.WeightFrom {
				buf.WriteString(fmt.Sprintf(ndLabel+"Incoming Edge: [%s] -- %.3f --> [%s]\n", src.ID, weight, nd.ID))
			}
		}
	}
	return buf.String()
}

// UpdateEdgeWeight overwrites the edge weight from src to dst Node.
func (d Data) UpdateEdgeWeight(src, dst *Node, weight float32) {
	if src == nil || dst == nil {
		return
	}
	src.WeightTo[dst] = weight
}
