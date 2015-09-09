package graph

import (
	"container/heap"
	"fmt"
	"math"
)

// Dijkstra returns the shortest path using Dijkstra
// algorithm with a min-priority queue. This algorithm
// does not work with negative weight edges.
// (https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm)
//
//	 0. Dijkstra(G, source, target)
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
//	13. 	while Q is not empty:
//	14.
//	15. 		u = Q.extract_min()
//	16. 		if u == target:
//	17. 			break
//	18.
//	19. 		for each child vertex v of u:
//	20.
//	21. 			alt = distance[u] + weight(u, v)
//	22. 			if distance[v] > alt:
//	23. 				distance[v] = alt
//	24. 				prev[v] = u
//	25. 				Q.decrease_priority(v, alt)
//	26.
//	27. 		reheapify(Q)
//	28.
//	29.
//	30. 	path = []
//	31. 	u = target
//	32. 	while prev[u] is defined:
//	33. 		path.push_front(u)
//	34. 		u = prev[u]
//	35.
//	36. 	return path, prev
//
func Dijkstra(g Graph, source, target string) ([]string, map[string]float64, error) {

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

		// if u == target:
		if u.vertex == target {
			break
		}

		// for each child vertex v of u:
		cmap, err := g.GetChildren(u.vertex)
		if err != nil {
			return nil, nil, err
		}
		for v := range cmap {

			// alt = distance[u] + weight(u, v)
			weight, err := g.GetWeight(u.vertex, v)
			if err != nil {
				return nil, nil, err
			}
			alt := distance[u.vertex] + weight

			// if distance[v] > alt:
			if distance[v] > alt {

				// distance[v] = alt
				distance[v] = alt

				// prev[v] = u
				prev[v] = u.vertex

				// Q.decrease_priority(v, alt)
				minHeap.updateDistance(v, alt)
			}
		}
		heap.Init(minHeap)
	}

	// path = []
	path := []string{}

	// u = target
	u := target

	// while prev[u] is defined:
	for {
		if _, ok := prev[u]; !ok {
			break
		}
		// path.push_front(u)
		temp := make([]string, len(path)+1)
		temp[0] = u
		copy(temp[1:], path)
		path = temp

		// u = prev[u]
		u = prev[u]
	}

	// add the source
	temp := make([]string, len(path)+1)
	temp[0] = source
	copy(temp[1:], path)
	path = temp

	return path, distance, nil
}

// BellmanFord returns the shortest path using Bellman-Ford algorithm
// This algorithm works with negative weight edges.
// Time complexity is O(|V||E|).
// (http://courses.csail.mit.edu/6.006/spring11/lectures/lec15.pdf)
// It returns error when there is a negative-weight cycle.
// A negatively-weighted cycle adds up to infinite negative-weight.
//
//	 0. BellmanFord(G, source, target)
//	 1.
//	 2. 	distance[source] = 0
//	 3.
//	 4. 	for each vertex v in G:
//	 5.
//	 6. 		if v ≠ source:
//	 7. 			distance[v] = ∞
//	 8. 			prev[v] = undefined
//	 9.
//	10.
//	11. 	for 1 to |V|-1:
//	12.
//	13. 		for every edge (u, v):
//	14.
//	15. 			alt = distance[u] + weight(u, v)
//	16. 			if distance[v] > alt:
//	17. 				distance[v] = alt
//	18. 				prev[v] = u
//	19.
//	20.
//	21. 	for every edge (u, v):
//	22.
//	23. 		alt = distance[u] + weight(u, v)
//	24. 		if distance[v] > alt:
//	25. 			there is a negative-weight cycle
//	26.
//	27.
//	28. 	path = []
//	29. 	u = target
//	30. 	while prev[u] is defined:
//	31. 		path.push_front(u)
//	32. 		u = prev[u]
//	33.
//	34. 	return path, prev
//
func BellmanFord(g Graph, source, target string) ([]string, map[string]float64, error) {

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
	}

	prev := make(map[string]string)

	// for 1 to |V|-1:
	for i := 1; i <= len(g.GetVertices())-1; i++ {

		// for every edge (u, v):
		for vtx := range g.GetVertices() {

			cmap, err := g.GetChildren(vtx)
			if err != nil {
				return nil, nil, err
			}
			for v := range cmap {
				u := vtx
				// edge (u, v)
				weight, err := g.GetWeight(u, v)
				if err != nil {
					return nil, nil, err
				}

				// alt = distance[u] + weight(u, v)
				alt := distance[u] + weight

				// if distance[v] > alt:
				if distance[v] > alt {
					// distance[v] = alt
					distance[v] = alt

					// prev[v] = u
					prev[v] = u
				}
			}

			pmap, err := g.GetParents(vtx)
			if err != nil {
				return nil, nil, err
			}
			for u := range pmap {
				v := vtx
				// edge (u, v)
				weight, err := g.GetWeight(u, v)
				if err != nil {
					return nil, nil, err
				}

				// alt = distance[u] + weight(u, v)
				alt := distance[u] + weight

				// if distance[v] > alt:
				if distance[v] > alt {
					// distance[v] = alt
					distance[v] = alt

					// prev[v] = u
					prev[v] = u
				}
			}
		}
	}

	// for every edge (u, v):
	for vtx := range g.GetVertices() {

		cmap, err := g.GetChildren(vtx)
		if err != nil {
			return nil, nil, err
		}
		for v := range cmap {
			u := vtx
			// edge (u, v)
			weight, err := g.GetWeight(u, v)
			if err != nil {
				return nil, nil, err
			}

			// alt = distance[u] + weight(u, v)
			alt := distance[u] + weight

			// if distance[v] > alt:
			if distance[v] > alt {
				return nil, nil, fmt.Errorf("there is a negative-weight cycle: %v", g)
			}
		}

		pmap, err := g.GetParents(vtx)
		if err != nil {
			return nil, nil, err
		}
		for u := range pmap {
			v := vtx
			// edge (u, v)
			weight, err := g.GetWeight(u, v)
			if err != nil {
				return nil, nil, err
			}

			// alt = distance[u] + weight(u, v)
			alt := distance[u] + weight

			// if distance[v] > alt:
			if distance[v] > alt {
				return nil, nil, fmt.Errorf("there is a negative-weight cycle: %v", g)
			}
		}
	}

	// path = []
	path := []string{}

	// u = target
	u := target

	// while prev[u] is defined:
	for {
		if _, ok := prev[u]; !ok {
			break
		}
		// path.push_front(u)
		temp := make([]string, len(path)+1)
		temp[0] = u
		copy(temp[1:], path)
		path = temp

		// u = prev[u]
		u = prev[u]
	}

	// add the source
	temp := make([]string, len(path)+1)
	temp[0] = source
	copy(temp[1:], path)
	path = temp

	return path, distance, nil
}
