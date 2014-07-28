package tsdfs

import (
	"github.com/gyuho/goraph/algorithm/dfs"
	"github.com/gyuho/goraph/graph/gs"
)

// TSDFS returns the topological sort of the graph
// but does not detect if it is a DAG or not.
func TSDFS(g *gs.Graph) string {
	return dfs.DFS(g)
}
