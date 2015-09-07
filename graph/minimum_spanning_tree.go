package graph

import (
	"fmt"
	"sort"
)

// Edge is an Edge from Source to Target.
type Edge struct {
	Source string
	Target string
	Weight float64
}

type EdgeSlice []Edge

func (e EdgeSlice) Len() int           { return len(e) }
func (e EdgeSlice) Less(i, j int) bool { return e[i].Weight < e[j].Weight }
func (e EdgeSlice) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

// Kruskal finds the minimum spanning tree with disjoint-set data structure.
// (http://en.wikipedia.org/wiki/Kruskal%27s_algorithm)
//
//	 0. Kruskal(G)
//	 1.
//	 2. 	A = ∅
//	 3.
//	 4. 	for each vertex v in G:
//	 5. 		MakeDisjointSet(v)
//	 6.
//	 7. 	edges = get all edges
//	 8. 	sort edges in ascending order of weight
//	 9.
//	10. 	for each edge (u, v) in edges:
//	11. 		if FindSet(u) ≠ FindSet(v):
//	12. 			A = A ∪ {(u, v)}
//	13. 			Union(u, v)
//	14.
//	15. 	return A
//
func Kruskal(g Graph) map[Edge]bool {

	// A = ∅
	A := make(map[Edge]bool)

	// disjointSet maps a member Vertex to a represent.
	// (https://en.wikipedia.org/wiki/Disjoint-set_data_structure)
	forests := NewForests()

	// for each vertex v in G:
	for v := range g.GetVertices() {
		// MakeDisjointSet(v)
		MakeDisjointSet(forests, v)
	}

	// edges = get all edges
	edges := []Edge{}
	foundEdge := make(map[string]bool)
	for vtx := range g.GetVertices() {
		cmap, err := g.GetChildren(vtx)
		if err != nil {
			panic(err)
		}
		for c := range cmap {
			// edge (vtx, c)
			weight, err := g.GetWeight(vtx, c)
			if err != nil {
				panic(err)
			}
			edge := Edge{}
			edge.Source = vtx
			edge.Target = c
			edge.Weight = weight
			if _, ok := foundEdge[fmt.Sprintf("%+v", edge)]; !ok {
				edges = append(edges, edge)
				foundEdge[fmt.Sprintf("%+v", edge)] = true
			}
		}

		pmap, err := g.GetParents(vtx)
		if err != nil {
			panic(err)
		}
		for p := range pmap {
			// edge (p, vtx)
			weight, err := g.GetWeight(p, vtx)
			if err != nil {
				panic(err)
			}
			edge := Edge{}
			edge.Source = p
			edge.Target = vtx
			edge.Weight = weight
			if _, ok := foundEdge[fmt.Sprintf("%+v", edge)]; !ok {
				edges = append(edges, edge)
				foundEdge[fmt.Sprintf("%+v", edge)] = true
			}
		}
	}

	// sort edges in ascending order of weight
	sort.Sort(EdgeSlice(edges))

	// for each edge (u, v) in edges:
	for _, edge := range edges {
		// if FindSet(u) ≠ FindSet(v):
		if FindSet(forests, edge.Source).represent != FindSet(forests, edge.Target).represent {

			// A = A ∪ {(u, v)}
			A[edge] = true

			// Union(u, v)
			// overwrite v's represent with u's represent
			Union(forests, FindSet(forests, edge.Source), FindSet(forests, edge.Target))
		}
	}

	return A
}
