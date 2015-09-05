package graph

// Clone clones(deep copy) the graph Data. (changing the cloned Data would not affect the original Data.)
// It traverses every single node with depth-first-search.
func (g *Graph) Clone() *Graph {
	clonedData := New()
	src := &Node{}
	for nd := range g.NodeMap {
		src = nd
		break
	}
	g.cloneDfs(src, clonedData)
	return clonedData
}

func (g *Graph) cloneDfs(src *Node, clonedData *Graph) {
	if src.Color == "black" {
		return
	}

	src.Color = "black"
	srcClone := NewNode(src.ID)

	for ov, weight := range src.WeightTo {
		ovClone := NewNode(ov.ID)
		clonedData.Connect(srcClone, ovClone, weight)
		if ov.Color == "white" {
			g.cloneDfs(ov, clonedData)
		}
	}

	for iv, weight := range src.WeightFrom {
		ivClone := NewNode(iv.ID)
		clonedData.Connect(ivClone, srcClone, weight)
		if iv.Color == "white" {
			g.cloneDfs(iv, clonedData)
		}
	}
}
