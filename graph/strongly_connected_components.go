package graph

import "sync"

// Tarjan finds the strongly connected components.
// In the mathematics, a directed graph is "strongly connected"
// if every vertex is reachable from every other node.
// Therefore, a graph is strongly connected if there is a path
// in each direction between each pair of node of a graph.
// Then a pair of vertices u and v is strongly connected to each other
// because there is a path in each direction.
// "Strongly connected components" of an arbitrary graph
// partition into sub-graphs that are themselves strongly connected.
// That is, "strongly connected component" of a directed graph
// is a sub-graph that is strongly connected.
// Formally, "Strongly connected components" of a graph is a maximal
// set of vertices C in G.V such that for all u, v ∈ C, there is a path
// both from u to v, and from v to u.
// (https://en.wikipedia.org/wiki/Tarjan%27s_strongly_connected_components_algorithm)
//
//	 0. Tarjan(G):
//	 1.
//	 2. 	globalIndex = 0 // smallest unused index
//	 3. 	let S be a stack
//	 4. 	result = []
//	 5.
//	 6. 	for each vertex v in G:
//	 7. 		if v.index is undefined:
//	 8. 			tarjan(G, v, globalIndex, S, result)
//	 9.
//	10. 	return result
//	11.
//	12.
//	13. tarjan(G, v, globalIndex, S, result):
//	14.
//	15. 	v.index = globalIndex
//	16. 	v.lowLink = globalIndex
//	17. 	globalIndex++
//	18. 	S.push(v)
//	19.
//	20. 	for each child vertex w of v:
//	21.
//	22. 		if w.index is undefined:
//	23. 			tarjan(G, w, globalIndex, S, result)
//	24. 			v.lowLink = min(v.lowLink, w.lowLink)
//	25.
//	26. 		else if w is in S:
//	27. 			v.lowLink = min(v.lowLink, w.index)
//	28.
//	29. 	// if v is the root
//	30. 	if v.lowLink == v.index:
//	31.
//	32. 		// start a new strongly connected component
//	33. 		component = []
//	34.
//	35. 		while True:
//	36.
//	37. 			u = S.pop()
//	38. 			component.push(u)
//	39.
//	40. 			if u == v:
//	41. 				result.push(component)
//	42. 				break
//
func Tarjan(g Graph) [][]string {

	data := newTarjanData()

	// for each vertex v in G:
	for v := range g.GetVertices() {
		// if v.index is undefined:
		if _, ok := data.index[v]; !ok {
			// tarjan(G, v, globalIndex, S, result)
			tarjan(g, v, data)
		}
	}

	return data.result
}

type tarjanData struct {
	sync.Mutex

	// globalIndex is the smallest unused index
	globalIndex int

	// index is an index of a node to record
	// the order of being discovered.
	index map[string]int

	// lowLink is the smallest index of any index
	// reachable from v, including v itself.
	lowLink map[string]int

	// S is the stack.
	S []string

	// extra map to check if a vertex is in S.
	smap map[string]bool

	result [][]string
}

func newTarjanData() *tarjanData {
	d := tarjanData{}
	d.globalIndex = 0
	d.index = make(map[string]int)
	d.lowLink = make(map[string]int)
	d.S = []string{}
	d.smap = make(map[string]bool)
	d.result = [][]string{}
	return &d
}

func tarjan(
	g Graph,
	vtx string,
	data *tarjanData,
) {

	// TODO: be more completely thread-safe.
	// This is not inherently parallelizable problem,
	// but just to make sure.
	data.Lock()

	// v.index = globalIndex
	data.index[vtx] = data.globalIndex

	// v.lowLink = globalIndex
	data.lowLink[vtx] = data.globalIndex

	// globalIndex++
	data.globalIndex++

	// S.push(v)
	data.S = append(data.S, vtx)
	data.smap[vtx] = true

	data.Unlock()

	// for each child vertex w of v:
	cmap, err := g.GetChildren(vtx)
	if err != nil {
		panic(err)
	}
	for w := range cmap {

		// if w.index is undefined:
		if _, ok := data.index[w]; !ok {

			// tarjan(G, w, globalIndex, S, result)
			tarjan(g, w, data)

			// v.lowLink = min(v.lowLink, w.lowLink)
			data.lowLink[vtx] = min(data.lowLink[vtx], data.lowLink[w])

		} else if _, ok := data.smap[w]; ok {
			// else if w is in S:

			// v.lowLink = min(v.lowLink, w.index)
			data.lowLink[vtx] = min(data.lowLink[vtx], data.index[w])
		}

	}

	data.Lock()

	// if v is the root
	// if v.lowLink == v.index:
	if data.lowLink[vtx] == data.index[vtx] {
		// start a new strongly connected component
		component := []string{}

		// while True:
		for {

			// u = S.pop()
			u := data.S[len(data.S)-1]
			data.S = data.S[:len(data.S)-1 : len(data.S)-1]
			delete(data.smap, u)

			// component.push(u)
			component = append(component, u)

			// if u == v:
			if u == vtx {
				data.result = append(data.result, component)
				break
			}
		}
	}

	data.Unlock()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
