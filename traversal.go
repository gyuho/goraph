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
func BFS(g Graph, name string) []string {
	if g.GetNode(name) == nil {
		return nil
	}

	q := []string{name}
	visited := make(map[string]bool)
	visited[name] = true
	rs := []string{name}

	// while Q is not empty:
	for len(q) != 0 {

		u := q[0]
		q = q[1:len(q):len(q)]

		// for each vertex w adjacent to u:
		cmap, _ := g.GetTargets(g.GetNode(u).ID())
		for _, w := range cmap {
			// if w is not visited yet:
			if _, ok := visited[w.String()]; !ok {
				q = append(q, w.String())  // Q.push(w)
				visited[w.String()] = true // label w as visited

				rs = append(rs, w.String())
			}
		}
		pmap, _ := g.GetSources(g.GetNode(u).ID())
		for _, w := range pmap {
			// if w is not visited yet:
			if _, ok := visited[w.String()]; !ok {
				q = append(q, w.String())  // Q.push(w)
				visited[w.String()] = true // label w as visited

				rs = append(rs, w.String())
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
func DFS(g Graph, name string) []string {
	if g.GetNode(name) == nil {
		return nil
	}

	s := []string{name}
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
			cmap, _ := g.GetTargets(g.GetNode(u).ID())
			for _, w := range cmap {
				// if w is not visited yet:
				if _, ok := visited[w.String()]; !ok {
					s = append(s, w.String()) // S.push(w)
				}
			}
			pmap, _ := g.GetSources(g.GetNode(u).ID())
			for _, w := range pmap {
				// if w is not visited yet:
				if _, ok := visited[w.String()]; !ok {
					s = append(s, w.String()) // S.push(w)
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
func DFSRecursion(g Graph, name string) []string {
	if g.GetNode(name) == nil {
		return nil
	}

	visited := make(map[string]bool)
	rs := []string{}

	dfsRecursion(g, name, visited, &rs)

	return rs
}

func dfsRecursion(g Graph, name string, visited map[string]bool, rs *[]string) {
	// base case of recursion
	//
	// if v is visited:
	if _, ok := visited[name]; ok {
		return
	}

	// label v as visited
	visited[name] = true
	*rs = append(*rs, name)

	// for each vertex u adjacent to v:
	cmap, _ := g.GetTargets(g.GetNode(name).ID())
	for _, u := range cmap {
		// if u is not visited yet:
		if _, ok := visited[u.String()]; !ok {
			// recursive DFS(G, u)
			dfsRecursion(g, u.String(), visited, rs)
		}
	}
	pmap, _ := g.GetSources(g.GetNode(name).ID())
	for _, u := range pmap {
		// if u is not visited yet:
		if _, ok := visited[u.String()]; !ok {
			// recursive DFS(G, u)
			dfsRecursion(g, u.String(), visited, rs)
		}
	}
}
