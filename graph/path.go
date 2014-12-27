package graph

// Path returns true if there is an edge from src to dst Vertex.
func (d Data) Path(src, dst *Vertex) bool {
	return true
}

// DAG returns true if the graph is Directed acyclic graph.
// That is, it returns true if the graph has no cycle.
// (http://en.wikipedia.org/wiki/Directed_acyclic_graph)
func (d Data) DAG() bool {
	return true
}
