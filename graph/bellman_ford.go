package graph

import (
	"fmt"
	"math"
)

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

	fmt.Println("distance", distance)
	fmt.Println("prev", prev)
	return path, distance, nil
}
