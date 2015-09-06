package graph

// DFS does depth-first search, and returns the list of vertices.
// (https://en.wikipedia.org/wiki/Depth-first_search)
//
//	 0. DFS(G, v):
//	 1.
//	 2. 	let S be a stack
//	 3. 	S.push(v)
//	 4.
//	 5. 	while S is not empty:
//	 6.
//	 7. 		u = S.pop()
//	 8.
//	 9. 		if u is not visited yet:
//	10.
//	11. 			label u as visited
//	12.
//	13. 			for each vertex w adjacent to u:
//	14.
//	15. 				if w is not visited yet:
//	16. 					S.push(w)
//
func DFS(g Graph, vtx string) []string {

	if !g.FindVertex(vtx) {
		return nil
	}

	rs := []string{}
	s := []string{vtx}
	visited := make(map[string]bool)

	// while S is not empty:
	for len(s) != 0 {

		u := s[len(s)-1]
		s = s[:len(s)-1 : len(s)-1]

		// if u is not visited yet:
		if _, ok := visited[u]; !ok {
			// label u as visited
			visited[u] = true

			rs = append(rs, u)

			// for each vertex w adjacent to u:
			cmap, _ := g.GetChildren(u)
			for w := range cmap {
				// if w is not visited yet:
				if _, ok := visited[w]; !ok {
					s = append(s, w) // S.push(w)
				}
			}
			pmap, _ := g.GetParents(u)
			for w := range pmap {
				// if w is not visited yet:
				if _, ok := visited[w]; !ok {
					s = append(s, w) // S.push(w)
				}
			}
		}
	}

	return rs
}
