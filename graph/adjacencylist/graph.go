package adjacencylist

import "container/list"

// Graph is a graph represented in adjacency list, but implemented in slice.
type Graph struct {
	Nodes *list.List
	Edges *list.List
}

// New returns a pointer to a new graph.
func New() *Graph {
	return &Graph{
		list.New(),
		list.New(),
	}
}

// Node is a Vertex(node) in a graph.
type Node struct {
	ID    string
	Color string

	// OutgoingNodes is a slice of Nodes that go out of this Node
	OutgoingNodes *list.List

	// IncomingNodes is a slice of Nodes that goes into this Node
	// (Nodes that precede this Node in graph)
	IncomingNodes *list.List
}

// NewNode returns a pointer to a new Node.
func NewNode(id string) *Node {
	return &Node{
		ID:            id,
		Color:         "white",
		OutgoingNodes: list.New(),
		IncomingNodes: list.New(),
	}
}

// Edge is an edge(arc) in a graph that has direction from one to another Node.
type Edge struct {
	// Src is the source Node that the edge starts from.
	Src *Node

	// Dst is the destination Node that the edge goes to.
	Dst *Node

	// Weight contains the weight value in float32 format.
	Weight float32
}

// NewEdge returns a new edge from src to dst.
func NewEdge(src, dst *Node, weight float32) *Edge {
	return &Edge{
		src,
		dst,
		weight,
	}
}

// GetNodeByID returns the Node with input ID, or return nil if it doesn't exist.
func (g Graph) GetNodeByID(id string) *Node {
	for nd := g.Nodes.Front(); nd != nil; nd = nd.Next() {
		if nd.Value.(*Node).ID == id {
			return nd.Value.(*Node)
		}
	}
	return nil
}

// ImmediateDominate returns true if A immediately dominates B.
// That is, true if A can go to B with only one edge.
func (a *Node) ImmediateDominate(b *Node) bool {
	for nd := a.OutgoingNodes.Front(); nd != nil; nd = nd.Next() {
		if nd.Value.(*Node) == b {
			return true
		}
	}
	return false
}

// GetEdgeWeight returns the weight value of the edge from source to destination vertex.
func (g Graph) GetEdgeWeight(src, dst *Node) float32 {
	for edge := g.Edges.Front(); edge != nil; edge = edge.Next() {
		if edge.Value.(*Edge).Src == src && edge.Value.(*Edge).Dst == dst {
			return edge.Value.(*Edge).Weight
		}
	}
	return 0.0
}

// UpdateWeight updates the weight value between Nodes.
func (g *Graph) UpdateWeight(src, dst *Node, value float32) {
	for edge := g.Edges.Front(); edge != nil; edge = edge.Next() {
		if edge.Value.(*Edge).Src == src && edge.Value.(*Edge).Dst == dst {
			edge.Value.(*Edge).Weight = value
		}
	}
}

// AddNode adds the Node v to a graph's Nodes.
func (g *Graph) AddNode(nd *Node) {
	g.Nodes.PushBack(nd)
}

// AddEdge adds the Edge e to a graph's Edges.
func (g *Graph) AddEdge(e *Edge) {
	g.Edges.PushBack(e)
}

// Connect connects the vertex v to A, not A to v.
// When there is more than one edge, it adds up the weight values.
func (g *Graph) Connect(a, b *Node, weight float32) {
	if g.GetNodeByID(a.ID) == nil {
		g.AddNode(a)
	}
	if g.GetNodeByID(b.ID) == nil {
		g.AddNode(b)
	}
	// if there is already an edge from A to B
	if a.ImmediateDominate(b) {
		g.UpdateWeight(a, b, weight)
	} else {
		edge := NewEdge(a, b, weight)
		g.AddEdge(edge)
		a.OutgoingNodes.PushBack(b)
		b.IncomingNodes.PushBack(a)
	}
}

// DeleteNode removes the input Node A from the graph g's Nodes.
func (g *Graph) DeleteNode(a *Node) {
	if g.GetNodeByID(a.ID) == nil {
		// To Debug
		// panic(a.ID + " Node does not exist! Can't delete the Node!")
		return
	}

	// remove all edges connected to this Node
	var next1 *list.Element
	for edge := g.Edges.Front(); edge != nil; edge = next1 {
		next1 = edge.Next()
		if edge.Value.(*Edge).Src == a || edge.Value.(*Edge).Dst == a {
			// remove from the graph
			g.Edges.Remove(edge)
		}
	}

	// remove from graph
	var next2 *list.Element
	for nd := g.Nodes.Front(); nd != nil; nd = next2 {
		next2 = nd.Next()
		if nd.Value.(*Node) == a {
			// remove from the graph
			g.Nodes.Remove(nd)
		}
	}

	// remove a from its outgoing Nodes' IncomingNodes
	// need to traverse all outgoing Nodes
	var next3 *list.Element
	for nd1 := a.OutgoingNodes.Front(); nd1 != nil; nd1 = next3 {
		next3 = nd1.Next()
		// traverse each Node's IncomingNodes
		var next4 *list.Element
		for nd2 := nd1.Value.(*Node).IncomingNodes.Front(); nd2 != nil; nd2 = next4 {
			next4 = nd2.Next()
			// check if the Node is the one we delete
			if nd2.Value.(*Node) == a {
				nd1.Value.(*Node).IncomingNodes.Remove(nd2)
			}
		}
	}

	// remove a from its incoming Nodes' OutNodes
	// need to traverse all incoming Nodes
	var next5 *list.Element
	for nd1 := a.IncomingNodes.Front(); nd1 != nil; nd1 = next5 {
		next5 = nd1.Next()
		// traverse each Node's OutNodes
		var next6 *list.Element
		for nd2 := nd1.Value.(*Node).OutgoingNodes.Front(); nd2 != nil; nd2 = next6 {
			next6 = nd2.Next()
			// check if the Node is the one we delete
			if nd2.Value.(*Node) == a {
				nd1.Value.(*Node).OutgoingNodes.Remove(nd2)
			}
		}
	}
}

// DeleteEdge deletes the edge from the Node A to B.
// Note that this only delete one direction from A to B.
func (g *Graph) DeleteEdge(a, b *Node) {
	if g.GetNodeByID(a.ID) == nil {
		// To Debug
		// panic(a.ID + " Node does not exist! Can't delete the Edge!")
		return
	}

	if g.GetNodeByID(b.ID) == nil {
		// To Debug
		// panic(b.ID + " Node does not exist! Can't delete the Edge!")
		return
	}

	if a.ImmediateDominate(b) == false {
		// To Debug
		// panic("No edge from " + a.ID + " to " + b.ID)
		return
	}

	// delete b from a's OutNodes
	var nextElem1 *list.Element
	for nd := a.OutgoingNodes.Front(); nd != nil; nd = nextElem1 {
		nextElem1 = nd.Next()
		if nd.Value.(*Node) == b {
			a.OutgoingNodes.Remove(nd)
		}
	}

	// delete a from b's IncomingNodes
	var nextElem2 *list.Element
	for nd := b.IncomingNodes.Front(); nd != nil; nd = nextElem2 {
		nextElem2 = nd.Next()
		if nd.Value.(*Node) == a {
			b.IncomingNodes.Remove(nd)
		}
	}

	// Always delete from graph at the end
	// remove the edge from the graph's edge list
	var nextElem3 *list.Element
	for edge := g.Edges.Front(); edge != nil; edge = nextElem3 {
		nextElem3 = edge.Next()
		// if the edge is from a to b
		if edge.Value.(*Edge).Src == a && edge.Value.(*Edge).Dst == b {
			// don't do this
			// edge.Value.(*Edge).Src = nil
			// edge.Value.(*Edge).Dst = nil
			g.Edges.Remove(edge)
		}
	}
}
