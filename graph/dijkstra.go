package graph

import (
	"container/heap"
	"math"
)

type vertexDistance struct {
	vertex   string
	distance float64
}

// container.Heap's Interface needs sort.Interface, Push, Pop to be implemented

// vertexDistanceHeap is a min-heap of vertexDistances.
type vertexDistanceHeap []vertexDistance

func (h vertexDistanceHeap) Len() int           { return len(h) }
func (h vertexDistanceHeap) Less(i, j int) bool { return h[i].distance < h[j].distance } // Min-Heap
func (h vertexDistanceHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *vertexDistanceHeap) Push(x interface{}) {
	*h = append(*h, x.(vertexDistance))
}

func (h *vertexDistanceHeap) Pop() interface{} {
	heapSize := len(*h)
	lastVertex := (*h)[heapSize-1]
	*h = (*h)[0 : heapSize-1]
	return lastVertex
}

func (h *vertexDistanceHeap) updateDistance(vtx string, val float64) {
	for i := 0; i < len(*h); i++ {
		if (*h)[i].vertex == vtx {
			(*h)[i].distance = val
			break
		}
	}
}

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
//	27. 	path = []
//	28. 	u = target
//	29. 	while prev[u] is defined:
//	30. 		path.push_front(u)
//	31. 		u = prev[u]
//	32.
//	33. 	return path, prev
//
func Dijkstra(g Graph, source, target string) ([]string, map[string]float64, error) {

	minHeap := &vertexDistanceHeap{}

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
				heap.Init(minHeap)
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
