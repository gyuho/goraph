package graph

// Edge connects from Src to Dst with weight.
type Edge struct {
	Src    *Node
	Dst    *Node
	Weight float32
}

// GetEdges returns all edges of a graph.
func (g *Graph) GetEdges() []Edge {
	rs := []Edge{}
	for nd1 := range g.NodeMap {
		for nd2, v := range nd1.WeightTo {
			one := Edge{}
			one.Src = nd1
			one.Dst = nd2
			one.Weight = v
			rs = append(rs, one)
		}
	}
	return rs
}
