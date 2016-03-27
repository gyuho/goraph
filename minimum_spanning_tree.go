package goraph

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

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
func Kruskal(g Graph) map[Edge]struct{} {

	// A = ∅
	A := make(map[Edge]struct{})

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
	foundEdge := make(map[string]struct{})
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
				foundEdge[fmt.Sprintf("%+v", edge)] = struct{}{}
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
				foundEdge[fmt.Sprintf("%+v", edge)] = struct{}{}
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
			A[edge] = struct{}{}

			// Union(u, v)
			// overwrite v's represent with u's represent
			Union(forests, FindSet(forests, edge.Source), FindSet(forests, edge.Target))
		}
	}

	return A
}

// Prim finds the minimum spanning tree with min-heap (priority queue).
// (http://en.wikipedia.org/wiki/Prim%27s_algorithm)
//
//	 0. Prim(G, source)
//	 1.
//	 2. 	let Q be a priority queue
//	 3. 	distance[source] = 0
//	 4.
//	 5. 	for each vertex v in G:
//	 6.
//	 7. 		if v ≠ source:
//	 8. 			distance[v] = ∞
//	 9. 			prev[v] = undefined
//	10.
//	11. 		Q.add_with_priority(v, distance[v])
//	12.
//	13.
//	14. 	while Q is not empty:
//	15.
//	16. 		u = Q.extract_min()
//	17.
//	18. 		for each adjacent vertex v of u:
//	19.
//	21. 			if v ∈ Q and distance[v] > weight(u, v):
//	22. 				distance[v] = weight(u, v)
//	23. 				prev[v] = u
//	24. 				Q.decrease_priority(v, weight(u, v))
//	25.
//	26.
//	27. 	return tree from prev
//
func Prim(g Graph, source string) map[Edge]struct{} {

	// let Q be a priority queue
	minHeap := &vertexDistanceHeap{}

	// distance[source] = 0
	distance := make(map[string]float64)
	distance[source] = 0.0

	// for each vertex v in G:
	for vtx := range g.GetVertices() {

		// if v ≠ source:
		if vtx != source {
			// distance[v] = ∞
			distance[vtx] = math.MaxFloat64

			// prev[v] = undefined
			// prev[v] = ""
		}

		// Q.add_with_priority(v, distance[v])
		vd := vertexDistance{}
		vd.vertex = vtx
		vd.distance = distance[vtx]

		heap.Push(minHeap, vd)
	}

	heap.Init(minHeap)
	prev := make(map[string]string)

	// while Q is not empty:
	for minHeap.Len() != 0 {

		// u = Q.extract_min()
		u := heap.Pop(minHeap).(vertexDistance)

		// for each adjacent vertex v of u:
		cmap, err := g.GetChildren(u.vertex)
		if err != nil {
			panic(err)
		}
		for v := range cmap {

			isExist := false
			for _, one := range *minHeap {
				if v == one.vertex {
					isExist = true
					break
				}
			}

			// weight(u, v)
			weight, err := g.GetWeight(u.vertex, v)
			if err != nil {
				panic(err)
			}

			// if v ∈ Q and distance[v] > weight(u, v):
			if isExist && distance[v] > weight {

				// distance[v] = weight(u, v)
				distance[v] = weight

				// prev[v] = u
				prev[v] = u.vertex

				// Q.decrease_priority(v, weight(u, v))
				minHeap.updateDistance(v, weight)
				heap.Init(minHeap)
			}
		}
		pmap, err := g.GetParents(u.vertex)
		if err != nil {
			panic(err)
		}
		for uu := range pmap {
			v := u.vertex

			isExist := false
			for _, one := range *minHeap {
				if v == one.vertex {
					isExist = true
					break
				}
			}

			// weight(u, v)
			weight, err := g.GetWeight(uu, v)
			if err != nil {
				panic(err)
			}

			// if v ∈ Q and distance[v] > weight(u, v):
			if isExist && distance[v] > weight {

				// distance[v] = weight(u, v)
				distance[v] = weight

				// prev[v] = u
				prev[v] = uu

				// Q.decrease_priority(v, weight(u, v))
				minHeap.updateDistance(v, weight)
				heap.Init(minHeap)
			}
		}
	}

	tree := make(map[Edge]struct{})
	for k, v := range prev {
		one := Edge{}
		one.Source = v
		one.Target = k
		weight, err := g.GetWeight(v, k)
		if err != nil {
			panic(err)
		}
		one.Weight = weight
		tree[one] = struct{}{}
	}
	return tree
}
