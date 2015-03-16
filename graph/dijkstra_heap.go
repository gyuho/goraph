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
	for _, elem := range *h {
		if elem.node == node {
			elem.distance = val
			break
		}
	}
}

// Dijkstra returns the shortest path using Dijkstra algorithm with a min-priority queue.
// This algorithm does not work with negative weight edges.
// (http://en.wikipedia.org/wiki/Dijkstra%27s_algorithm).
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
func (d *Data) Dijkstra(src *Node) (map[*Node]float32, map[string]string) {

	mapToDistance := make(map[*Node]float32)
	mapToDistance[src] = 0.0

	mapToPrevID := make(map[string]string)
	minHeap := &nodeDistanceHeap{}

	// initialize mapToDistance
	for nd := range d.NodeMap {
		if nd != src {
			mapToDistance[nd] = 2147483646.0
			mapToPrevID[nd.ID] = ""
		}
		ndd := nodeDistance{}
		ndd.node = nd
		ndd.distance = mapToDistance[nd]
		heap.Push(minHeap, ndd)
	}

	heap.Init(minHeap)

	for minHeap.Len() != 0 {
		minPqElem := heap.Pop(minHeap)
		for ov, weight := range minPqElem.(nodeDistance).node.WeightTo {
			if mapToDistance[ov] > mapToDistance[minPqElem.(nodeDistance).node]+weight {
				mapToDistance[ov] = mapToDistance[minPqElem.(nodeDistance).node] + weight
				mapToPrevID[ov.ID] = minPqElem.(nodeDistance).node.ID
				minHeap.updateDistance(ov, mapToDistance[minPqElem.(nodeDistance).node]+weight)
			}
		}
		for iv, weight := range minPqElem.(nodeDistance).node.WeightTo {
			if mapToDistance[iv] > mapToDistance[minPqElem.(nodeDistance).node]+weight {
				mapToDistance[iv] = mapToDistance[minPqElem.(nodeDistance).node] + weight
				mapToPrevID[iv.ID] = minPqElem.(nodeDistance).node.ID
				minHeap.updateDistance(iv, mapToDistance[minPqElem.(nodeDistance).node]+weight)
			}
		}
	}

	return mapToDistance, mapToPrevID
}
