// Package tarjan implements Tarjan's Strongly Connected Components algorithm.
package tarjan

/*
http://en.wikipedia.org/wiki/Strongly_connected_component

In the mathematical theory of directed graphs,
a graph is said to be "strongly connected"
if every vertex is reachable from every other vertex.

A directed graph is called "strongly connected"
if there is a "path" in each direction
between each pair of vertices of the graph.

a pair of vertices u and v are said to be
strongly connected to each other
if there is a path in each direction between them.

binary relation of being strongly connected is an equivalence relation


"strongly connected components" of an arbitrary directed graph
form a partition into subgraphs that are themselves strongly connected.

"strongly connected component" of a directed graph G
is a subgraph that is strongly connected
, and is maximal with this property

SCC of G is a maximal set of vertices C in V
such that for all u, v ∈ C,
there is a path both from u to v and from v to u.

http://en.wikipedia.org/wiki/Tarjan's_strongly_connected_components_algorithm

*/

// SCC returns the Strongly Connected Components using Tarjan's algorithm.
// (http://en.wikipedia.org/wiki/Tarjan's_strongly_connected_components_algorithm)
/*
func SCC(g *gsd.Graph) string {
	Stack := slice.NewSequence()
	//
	// v.index
	//	numbers the nodes consecutively in the order
	//	in which they are discovered
	//
	// v.lowlink
	//	represents (roughly speaking) the smallest index
	//	of any node known to be reachable from v, including v itself
	//
	// if v.lowlink < v.index
	// 	v must be left on the stack
	//
	// if v.lowlink == v.index
	//	whereas v must be removed as the root
	//	of a strongly connected component
	//
	index := 0
	Vertices := g.GetVertices()
	for _, vtx := range *Vertices {

	}

	for Stack.Len() != 0 {

	}

	return ""
}

func scc(g *gsd.Graph) string {
	Stack := slice.NewSequence()
	Vertices := g.GetVertices()

	for !slice.IsEqual(*Stack, *Vertices) {

	}

	for Stack.Len() != 0 {

	}

	return ""
}
*/
