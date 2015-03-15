package graph

// Clone clones(deep copy) the graph Data. (changing the cloned Data would not affect the original Data.)
// It traverses every single node with depth-first-search.
func (d *Data) Clone() *Data {
	clonedData := NewData()
	src := &Node{}
	for nd := range d.NodeMap {
		src = nd
		break
	}
	d.cloneDfs(src, clonedData)
	return clonedData
}

func (d *Data) cloneDfs(src *Node, clonedData *Data) {
	if src.Color == "black" {
		return
	}

	src.Color = "black"
	srcClone := NewNode(src.ID)

	for ov, weight := range src.WeightTo {
		ovClone := NewNode(ov.ID)
		clonedData.Connect(srcClone, ovClone, weight)
		if ov.Color == "white" {
			d.cloneDfs(ov, clonedData)
		}
	}

	for iv, weight := range src.WeightFrom {
		ivClone := NewNode(iv.ID)
		clonedData.Connect(ivClone, srcClone, weight)
		if iv.Color == "white" {
			d.cloneDfs(iv, clonedData)
		}
	}
}
