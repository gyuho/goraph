package graph

// BFS does breadth-first search, and returns the list of vertices.
// (https://en.wikipedia.org/wiki/Breadth-first_search)
//
//	 0. BFS(G, v):
//	 1.
//	 2.	let Q be a queue
//	 3.	Q.push(v)
//	 4.	label v as visited
//	 5.
//	 6.	while Q is not empty:
//	 7.
//	 8.		u = Q.dequeue()
//	 9.
//	10.		for each vertex w adjacent to u:
//	11.
//	12.			if w is not visited yet:
//	13.				Q.push(w)
//	14.				label w as visited
//
func BFS(g Graph, vtx string) []string {
	if !g.FindVertex(vtx) {
		return nil
	}

	rs := []string{vtx}
	q := []string{vtx}
	visited := make(map[string]bool)
	visited[vtx] = true

	for len(q) != 0 {

		u := q[0]
		q = q[1:len(q):len(q)]

		// for each vertex w adjacent to u:
		cmap, _ := g.GetChildren(u)
		for w := range cmap {
			// if w is not visited yet:
			if _, ok := visited[w]; !ok {
				q = append(q, w)  // Q.push(w)
				visited[w] = true // label w as visited

				rs = append(rs, w)
			}
		}
		pmap, _ := g.GetParents(u)
		for w := range pmap {
			// if w is not visited yet:
			if _, ok := visited[w]; !ok {
				q = append(q, w)  // Q.push(w)
				visited[w] = true // label w as visited

				rs = append(rs, w)
			}
		}
	}

	return rs
}
