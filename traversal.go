package goraph

// BFS does breadth-first search, and returns the list of vertices.
// (https://en.wikipedia.org/wiki/Breadth-first_search)
//
//	 0. BFS(G, v):
//	 1.
//	 2. 	let Q be a queue
//	 3. 	Q.push(v)
//	 4. 	label v as visited
//	 5.
//	 6. 	while Q is not empty:
//	 7.
//	 8. 		u = Q.dequeue()
//	 9.
//	10. 		for each vertex w adjacent to u:
//	11.
//	12. 			if w is not visited yet:
//	13. 				Q.push(w)
//	14. 				label w as visited
//
func BFS(g Graph, vtx string) []string {

	if !g.FindVertex(vtx) {
		return nil
	}

	q := []string{vtx}
	visited := make(map[string]bool)
	visited[vtx] = true
	rs := []string{vtx}

	// while Q is not empty:
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

	s := []string{vtx}
	visited := make(map[string]bool)
	rs := []string{}

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

// DFSRecursion does depth-first search recursively.
//
//	 0. DFS(G, v):
//	 1.
//	 2. 	if v is visited:
//	 3. 		return
//	 4.
//	 5. 	label v as visited
//	 6.
//	 7. 	for each vertex u adjacent to v:
//	 8.
//	 9. 		if u is not visited yet:
//	10. 			recursive DFS(G, u)
//
func DFSRecursion(g Graph, vtx string) []string {

	if !g.FindVertex(vtx) {
		return nil
	}

	visited := make(map[string]bool)
	rs := []string{}

	dfsRecursion(g, vtx, visited, &rs)

	return rs
}

func dfsRecursion(g Graph, vtx string, visited map[string]bool, rs *[]string) {

	// base case of recursion
	//
	// if v is visited:
	if _, ok := visited[vtx]; ok {
		return
	}

	// label v as visited
	visited[vtx] = true
	*rs = append(*rs, vtx)

	// for each vertex u adjacent to v:
	cmap, _ := g.GetChildren(vtx)
	for u := range cmap {
		// if u is not visited yet:
		if _, ok := visited[u]; !ok {
			// recursive DFS(G, u)
			dfsRecursion(g, u, visited, rs)
		}
	}
	pmap, _ := g.GetParents(vtx)
	for u := range pmap {
		// if u is not visited yet:
		if _, ok := visited[u]; !ok {
			// recursive DFS(G, u)
			dfsRecursion(g, u, visited, rs)
		}
	}
}
