package graph

// Edge is an Edge from Source to Target.
type Edge struct {
	Source string
	Target string
	Weight float64
}

type edgeSlice []Edge

func (e edgeSlice) Len() int           { return len(e) }
func (e edgeSlice) Less(i, j int) bool { return e[i].Weight < e[j].Weight }
func (e edgeSlice) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

// Kruskal finds the minimum spanning tree with disjoint-set data structure.
// (http://en.wikipedia.org/wiki/Kruskal%27s_algorithm)
//
//	 0. Kruskal(G)
//	 1.
//	 2. 	A = ∅
//	 3.
//	 4. 	for each vertex v in G:
//	 5. 		MakeDisjointSet(v)
//
// func Kruskal(g Graph) map[Edge]bool {

// }
