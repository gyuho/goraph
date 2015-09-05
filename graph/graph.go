package graph

import (
	"bytes"
	"fmt"
	"sync"
)

// Graph contains graph data, represented in adjacency list and slice.
// Make sure to use Pointer when we need to update the struct with receiver.
// (https://golang.org/doc/faq#methods_on_values_or_pointers)
type Graph struct {
	sync.Mutex

	// NodeMap is a hash-map for all Nodes in the graph.
	NodeMap map[*Node]bool

	// maintain nodeID in order not to have duplicate Node IDs in the graph.
	nodeID map[string]bool
}

// New returns a new Graph.
func New() *Graph {
	return &Graph{
		NodeMap: make(map[*Node]bool),
		nodeID:  make(map[string]bool),
		// without this
		// panic: assignment to entry in nil map
	}
}

// Init initializes the graph Data.
func (g *Graph) Init() {
	// (X) g = New()
	// this only updates the pointer
	//
	*g = *New()
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
}

// NewNode returns a new Node.
func NewNode(id string) *Node {
	return &Node{
		ID:         id,
		Color:      "white",
		WeightTo:   make(map[*Node]float32),
		WeightFrom: make(map[*Node]float32),
	}
}

// AddNode adds a Node to a graph Data.
// It returns true if the Node is added the graph Data.
func (g *Graph) AddNode(nd *Node) bool {

	if nd == nil {
		return false
	}

	g.Lock()
	defer g.Unlock()
	if _, ok := g.nodeID[nd.ID]; ok {
		return false
	}
	if _, ok := g.NodeMap[nd]; ok {
		return false
	}
	g.nodeID[nd.ID] = true
	g.NodeMap[nd] = true
	return true
}

// GetNodeSize returns the size of Node of the graph Data.
func (g Graph) GetNodeSize() int {
	return len(g.NodeMap)
}

//GetNodeByID finds a Node by ID.
func (g Graph) GetNodeByID(id string) *Node {
	for nd := range g.NodeMap {
		if nd.ID == id {
			return nd
		}
	}
	return nil
}

// Connect adds an edge from src(source) to dst(destination) Node, to a graph Data.
// This doese not connect from dst to src.
func (g *Graph) Connect(src, dst *Node, weight float32) {

	if src == nil || dst == nil {
		return
	}

	// if we do not want to allow a cycle
	// if src.ID == dst.ID {
	// 	return
	// }

	// add to Data
	if !g.AddNode(src) {
		src = g.GetNodeByID(src.ID)
		// this only updates the pointer
		//
		// this updates the value
		// *src = *(g.GetNodeByID(src.ID))
	}
	if !g.AddNode(dst) {
		dst = g.GetNodeByID(dst.ID)
	}

	g.Lock()
	defer g.Unlock()

	// if v, ok := src.WeightTo[dst]; !ok {
	// (X) src.WeightTo[dst] = v + weight

	// update src Node
	src.Lock()
	src.WeightTo[dst] = weight
	src.Unlock()

	// update dst Node
	dst.Lock()
	dst.WeightFrom[src] = weight
	dst.Unlock()
}

// GetEdgeWeight returns the weight value of an edge from src to dst Node.
func (g *Graph) GetEdgeWeight(src, dst *Node) float32 {
	if src == nil || dst == nil {
		return 0.0
	}

	src.Lock()
	defer src.Unlock()

	if _, ok := src.WeightTo[dst]; !ok {
		return 0.0
	}

	return src.WeightTo[dst]
}

// DeleteNode deletes a Node from the graph Data.
// This deletes all the related edges too.
func (g *Graph) DeleteNode(nd *Node) {

	if nd == nil {
		return
	}

	// delete edges from each Node
	for elem := range g.NodeMap {
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
	g.Lock()
	if _, ok := g.NodeMap[nd]; ok {
		delete(g.NodeMap, nd)
	}
	if _, ok := g.nodeID[nd.ID]; ok {
		delete(g.nodeID, nd.ID)
	}
	g.Unlock()

	nd = nil
}

// DeleteEdge deletes an Edge from src to dst from the graph Data.
// This does not delete Nodes.
func (g *Graph) DeleteEdge(src, dst *Node) {

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
func (g Graph) String() string {
	buf := new(bytes.Buffer)
	if g.GetNodeSize() == 0 {
		return "Graph is empty."
	}
	buf.WriteString(fmt.Sprintf("Graph has %d Nodes\n", g.GetNodeSize()))
	for nd := range g.NodeMap {
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

// String describes Node.
func (nd Node) String() string {
	return fmt.Sprintf("[%s / %d Outgoing / %d Incoming Edges]",
		nd.ID,
		len(nd.WeightTo),
		len(nd.WeightFrom),
	)
}

// UpdateEdgeWeight overwrites the edge weight from src to dst Node.
func (g *Graph) UpdateEdgeWeight(src, dst *Node, weight float32) {
	if src == nil || dst == nil {
		return
	}
	src.Lock()
	src.WeightTo[dst] = weight
	src.Unlock()

	dst.Lock()
	dst.WeightFrom[src] = weight
	dst.Unlock()
}
