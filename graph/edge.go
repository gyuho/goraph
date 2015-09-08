package graph

// Edge is an Edge from Source to Target.
type Edge struct {
	Source string
	Target string
	Weight float64
}

type EdgeSlice []Edge

func (e EdgeSlice) Len() int           { return len(e) }
func (e EdgeSlice) Less(i, j int) bool { return e[i].Weight < e[j].Weight }
func (e EdgeSlice) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
