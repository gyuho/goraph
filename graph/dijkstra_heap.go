package graph

import "container/heap"

type nodeDistance struct {
	node     *Node
	distance float32
}

// container.Heap's Interface needs sort.Interface, Push, Pop to be implemented

// nodeDistanceHeap is a min-heap of nodeDistances.
type nodeDistanceHeap []nodeDistance

func (h nodeDistanceHeap) Len() int           { return len(h) }
func (h nodeDistanceHeap) Less(i, j int) bool { return h[i].distance < h[j].distance } // Min-Heap
func (h nodeDistanceHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *nodeDistanceHeap) Push(x interface{}) {
	*h = append(*h, x.(nodeDistance))
}

func (h *nodeDistanceHeap) Pop() interface{} {
	heapSize := len(*h)
	lastNode := (*h)[heapSize-1]
	*h = (*h)[0 : heapSize-1]
	return lastNode
}

func (h *nodeDistanceHeap) updateDistance(node *Node, val float32) {
	// for _, elem := range *h {
	for i := 0; i < len(*h); i++ {
		// elem := (*h)[i] (X)
		if (*h)[i].node == node {
			(*h)[i].distance = val
			break
		}
	}
}

// Dijkstra returns the shortest path using Dijkstra algorithm with a min-priority queue.
// This algorithm does not work with negative weight edges.
// (http://en.wikipedia.org/wiki/Dijkstra%27s_algorithm)
//
//	1  function Dijkstra(Graph, source):
//	2      dist[source] ← 0                      // Initialization
//	3      for each vertex v in Graph:
//	4          if v ≠ source
//	5              dist[v] ← infinity            // Unknown distance from source to v
//	6              prev[v] ← undefined           // Predecessor of v
//	7          end if
//	8          Q.add_with_priority(v, dist[v])
//	9      end for
//	10
//	11     while Q is not empty:               // The main loop
//	12         u ← Q.extract_min()            // Remove and return best vertex
//	13         for each neighbor v of u:
//	14             alt = dist[u] + length(u, v)
//	15             if alt < dist[v]
//	16                 dist[v] ← alt
//	17                 prev[v] ← u
//	18                 Q.decrease_priority(v, alt)
//	19             end if
//	20         end for
//	21     end while
//	22     return prev[]
//
func (g *Graph) Dijkstra(src, dst *Node) ([]*Node, map[*Node]float32) {

	mapToDistance := make(map[*Node]float32)
	mapToDistance[src] = 0.0

	minHeap := &nodeDistanceHeap{}

	// initialize mapToDistance
	for nd := range g.NodeMap {
		if nd != src {
			mapToDistance[nd] = 2147483646.0
		}
		ndd := nodeDistance{}
		ndd.node = nd
		ndd.distance = mapToDistance[nd]
		heap.Push(minHeap, ndd)
	}

	mapToPrevID := make(map[string]string)
	heap.Init(minHeap)

	for minHeap.Len() != 0 {
		elem := heap.Pop(minHeap)
		for ov, weight := range elem.(nodeDistance).node.WeightTo {
			if mapToDistance[ov] > mapToDistance[elem.(nodeDistance).node]+weight {
				mapToDistance[ov] = mapToDistance[elem.(nodeDistance).node] + weight
				minHeap.updateDistance(ov, mapToDistance[elem.(nodeDistance).node]+weight)
				heap.Init(minHeap)

				mapToPrevID[ov.ID] = elem.(nodeDistance).node.ID
			}
		}
		for iv, weight := range elem.(nodeDistance).node.WeightTo {
			if mapToDistance[iv] > mapToDistance[elem.(nodeDistance).node]+weight {
				mapToDistance[iv] = mapToDistance[elem.(nodeDistance).node] + weight
				minHeap.updateDistance(iv, mapToDistance[elem.(nodeDistance).node]+weight)
				heap.Init(minHeap)

				mapToPrevID[iv.ID] = elem.(nodeDistance).node.ID
			}
		}
	}

	pathSlice := []*Node{dst}
	id := dst.ID
	for mapToPrevID[id] != src.ID {
		prevID := mapToPrevID[id]
		id = prevID
		copied := make([]*Node, len(pathSlice)+1) // push front
		copied[0] = g.GetNodeByID(prevID)
		copy(copied[1:], pathSlice)
		pathSlice = copied
	}
	copied := make([]*Node, len(pathSlice)+1) // push front
	copied[0] = g.GetNodeByID(src.ID)
	copy(copied[1:], pathSlice)
	pathSlice = copied

	return pathSlice, mapToDistance
}
