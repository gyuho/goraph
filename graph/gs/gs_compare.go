package gs

// SameVertex returns true if two vertices are the same.
func SameVertex(v1, v2 *Vertex) bool {
	if v1.ID != v2.ID {
		return false
	}
	if v1.Color != v2.Color {
		return false
	}
	if v1.StampD != v2.StampD {
		return false
	}
	if v1.StampF != v2.StampF {
		return false
	}

	if v1.GetOutVerticesSize() != v2.GetOutVerticesSize() {
		return false
	}
	vto1 := v1.GetOutVertices()
	vto2 := v2.GetOutVertices()
	for _, vx1 := range *vto1 {
		exvto := false
		for _, vx2 := range *vto2 {
			if vx1.(*Vertex).ID == vx2.(*Vertex).ID {
				exvto = true
			}
		}
		if exvto == false {
			return false
		}
	}

	if v1.GetInVerticesSize() != v2.GetInVerticesSize() {
		return false
	}
	vti1 := v1.GetInVertices()
	vti2 := v2.GetInVertices()
	for _, vx1 := range *vti1 {
		exvti := false
		for _, vx2 := range *vti2 {
			if vx1.(*Vertex).ID == vx2.(*Vertex).ID {
				exvti = true
			}
		}
		if exvti == false {
			return false
		}
	}

	if v1.GetPrevSize() != v2.GetPrevSize() {
		return false
	}
	pv1 := v1.GetPrev()
	pv2 := v2.GetPrev()
	for _, vx1 := range *pv1 {
		expv := false
		for _, vx2 := range *pv2 {
			if vx1.(*Vertex).ID == vx2.(*Vertex).ID {
				expv = true
			}
		}
		if expv == false {
			return false
		}
	}

	return true
}

// SameEdge returns true if two edges are the same.
func SameEdge(e1, e2 *Edge) bool {
	if e1.Weight != e2.Weight {
		return false
	}
	if !SameVertex(e1.Dst, e2.Dst) {
		return false
	}
	if !SameVertex(e1.Src, e2.Src) {
		return false
	}
	return true
}
