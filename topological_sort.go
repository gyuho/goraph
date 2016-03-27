package goraph

// TopologicalSort does topological sort(ordering) with DFS.
// It returns true if the graph is a DAG (no cycle, with a topological sort).
// False if the graph is not a DAG (cycle, with no topological sort).
//
//	 0. TopologicalSort(G)
//	 1.
//	 2. 	L = Empty list that will contain the sorted nodes
//	 3. 	isDAG = true
//	 4.
//	 5. 	for each vertex v in G:
//	 6.
//	 7. 		if v.color == "white":
//	 8.
//	 9. 			topologicalSortVisit(v, L, isDAG)
//	10.
//	11.
//	12.
//	13.
//	14. topologicalSortVisit(v, L, isDAG)
//	15.
//	16. 	if v.color == "gray":
//	17. 		isDAG = false
//	18. 		return
//	19.
//	20. 	if v.color == "white":
//	21.
//	22. 		v.color = "gray":
//	23.
//	24.			for each child vertex w of v:
//	25. 			topologicalSortVisit(w, L, isDAG)
//	26.
//	27. 		v.color = "black"
//	28.			L.push_front(v)
//
func TopologicalSort(g Graph) ([]string, bool) {

	// L = Empty list that will contain the sorted nodes
	L := []string{}
	isDAG := true
	color := make(map[string]string)
	for v := range g.GetVertices() {
		color[v] = "white"
	}

	// for each vertex v in G:
	for v := range g.GetVertices() {
		// if v.color == "white":
		if color[v] == "white" {
			// topologicalSortVisit(v, L, isDAG)
			topologicalSortVisit(g, v, &L, &isDAG, &color)
		}
	}

	return L, isDAG
}

func topologicalSortVisit(
	g Graph,
	vtx string,
	L *[]string,
	isDAG *bool,
	color *map[string]string,
) {

	// if v.color == "gray":
	if (*color)[vtx] == "gray" {
		// isDAG = false
		*isDAG = false
		return
	}

	// if v.color == "white":
	if (*color)[vtx] == "white" {
		// v.color = "gray":
		(*color)[vtx] = "gray"

		// for each child vertex w of v:
		cmap, err := g.GetChildren(vtx)
		if err != nil {
			panic(err)
		}
		for w := range cmap {
			// topologicalSortVisit(w, L, isDAG)
			topologicalSortVisit(g, w, L, isDAG, color)
		}

		// v.color = "black"
		(*color)[vtx] = "black"

		// L.push_front(v)
		temp := make([]string, len(*L)+1)
		temp[0] = vtx
		copy(temp[1:], *L)
		*L = temp
	}
}
