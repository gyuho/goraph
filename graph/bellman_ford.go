package graph

// BellmanFord returns the shortest path using Bellman-Ford algorithm
// This algorithm works with negative weight edges. Time complexity is O (V E).
// (http://courses.csail.mit.edu/6.006/spring11/lectures/lec15.pdf)
// It returns false when there is a negative-weight cycle.
// A negatively-weighted cycle adds up to infinite negative-weight.
//
//	1  for v in V:
//	2  	v.d = ∞
//	3  	v.π = nil
//	4
//	5  source.d = 0
//	6
//	7  for 1 to |V|-1:
//	8  	for every edge (u,v):
//	10  		relax(u,v)
//	11
//	12  relax(u,v):
//	13  	if v.d > u.d + w(u,v)
//	14  		v.d = u.d + w(u,v)
//	15  		v.π = u
//	16
//	17  for every edge (u,v):
//	18  	if v.d > u.d + w(u,v)
//	19  		there is a negative-weight cycle
//
func (d *Data) BellmanFord(src, dst *Node) ([]*Node, map[*Node]float32, bool) {

	mapToDistance := make(map[*Node]float32)

	// initialize mapToDistance
	for nd := range d.NodeMap {
		mapToDistance[nd] = 2147483646.0
	}

	mapToDistance[src] = 0.0
	mapToPrevID := make(map[string]string)

	for i := 1; i <= d.GetNodeSize()-1; i++ {
		// to iterate all edges
		for nd := range d.NodeMap {
			// relax the weight value to the destination
			for ov, weight := range nd.WeightTo {
				if mapToDistance[ov] > mapToDistance[nd]+weight {
					mapToDistance[ov] = mapToDistance[nd] + weight
					mapToPrevID[ov.ID] = nd.ID
				}
			}
			for iv, weight := range nd.WeightFrom {
				if mapToDistance[nd] > mapToDistance[iv]+weight {
					mapToDistance[nd] = mapToDistance[iv] + weight
					mapToPrevID[nd.ID] = iv.ID
				}
			}
		}
	}

	noNegCycle := true
Loop:
	// to iterate all edges
	for nd := range d.NodeMap {
		// relax the weight value to the destination
		for ov, weight := range nd.WeightTo {
			if mapToDistance[ov] > mapToDistance[nd]+weight {
				noNegCycle = false
				break Loop
			}
		}
		for iv, weight := range nd.WeightFrom {
			if mapToDistance[nd] > mapToDistance[iv]+weight {
				noNegCycle = false
				break Loop
			}
		}
	}
	if !noNegCycle {
		return nil, nil, false
	}

	pathSlice := []*Node{dst}
	id := dst.ID
	for mapToPrevID[id] != src.ID {
		prevID := mapToPrevID[id]
		id = prevID
		copied := make([]*Node, len(pathSlice)+1) // push front
		copied[0] = d.GetNodeByID(prevID)
		copy(copied[1:], pathSlice)
		pathSlice = copied
	}
	copied := make([]*Node, len(pathSlice)+1) // push front
	copied[0] = d.GetNodeByID(src.ID)
	copy(copied[1:], pathSlice)
	pathSlice = copied

	return pathSlice, mapToDistance, noNegCycle
}
