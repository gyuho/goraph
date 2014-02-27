// Package tsdfs employs DFS for topological sorting,
// but does not detect if the graph is a DAG (Directed Acyclic Graph)
// or not. When the graph is not a DAG, there is no topological sort.
// Use other packages: tsdag or tskahn for more complete solutions.
package tsdfs

import (
	"github.com/gyuho/goraph/algorithm/dfs"
	"github.com/gyuho/goraph/graph/gsd"
)

// TSDFS returns the topological sort of the graph
// but does not detect if it is a DAG or not.
func TSDFS(g *gsd.Graph) string {
	return dfs.DFS(g)
}
