package gsdflow

// GetEdge returns the Edge from src to dst Vertex.
// (Assume that there is no duplicate Edge for now.)
func (g Graph) GetEdge(src, dst *Vertex) *Edge {
	slice := g.GetEdges()
	for _, edge := range *slice {
		if edge.(*Edge).Src == src && edge.(*Edge).Dst == dst {
			return edge.(*Edge)
		}
	}
	return nil
}

// GetEdgeWeight returns the slice of weight values
// of the edge from source to destination vertex.
// In case we need to allow duplicate edges,
// we return a slice of weights.
func (g Graph) GetEdgeWeight(src, dst *Vertex) []float64 {
	slice := g.GetEdges()
	result := []float64{}
	for _, edge := range *slice {
		if edge.(*Edge).Src == src && edge.(*Edge).Dst == dst {
			result = append(result, edge.(*Edge).Weight)
		}
	}
	return result
}

// GetEdgeFlow returns the slice of flow values.
func (g Graph) GetEdgeFlow(src, dst *Vertex) []float64 {
	slice := g.GetEdges()
	result := []float64{}
	for _, edge := range *slice {
		if edge.(*Edge).Src == src && edge.(*Edge).Dst == dst {
			result = append(result, edge.(*Edge).Flow)
		}
	}
	return result
}

// UpdateWeight updates the weight value between vertices.
func (g *Graph) UpdateWeight(src, dst *Vertex, value float64) {
	edge := g.GetEdge(src, dst)
	if edge != nil {
		edge.Weight = value
	}
}

// UpdateFlow updates the flow value between vertices.
func (g *Graph) UpdateFlow(src, dst *Vertex, value float64) {
	edge := g.GetEdge(src, dst)
	if edge != nil {
		if edge.Weight < value {
			panic("Over Capicity")
		}
		edge.Flow = value
	}
}

// AddFlow updates the flow value between vertices.
func (g *Graph) AddFlow(src, dst *Vertex, value float64) {
	edge := g.GetEdge(src, dst)
	if edge != nil {
		if edge.Weight < edge.Flow+value {
			panic("Over Capicity")
		}
		edge.Flow = edge.Flow + value
	}
}

// SubFlow updates the flow value between vertices.
func (g *Graph) SubFlow(src, dst *Vertex, value float64) {
	edge := g.GetEdge(src, dst)
	if edge != nil {
		if edge.Flow-value < 0 {
			panic("Below Zero")
		}
		edge.Flow = edge.Flow - value
	}
}

// IsFull returns true if Flow reaches the capacity(Weight).
func (g *Graph) IsFull(src, dst *Vertex) bool {
	edge := g.GetEdge(src, dst)
	return edge.Weight <= edge.Flow
}
