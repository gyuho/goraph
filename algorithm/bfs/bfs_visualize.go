package bfs

/*
Pseudo Code CLRS page 595

BFS(G, s)
	for each vertex u ∈ g.V - {s}
		u.color = WHITE
		u.d = ∞
		u.π = NIL
	// this is already done
	// when instantiating the graph
	// and instead of InVertices
	// we can just create another slice
	// inside Graph (Prev)
	// in order not to modify the original graph

	s.color = GRAY
	s.d = 0
	s.π = NIL
	Q = ∅

	ENQUEUE(Q, s)

	while Q ≠ ∅
		u = DEQUEUE(Q)
		for each v ∈ g.Adj[u]
			if v.color == WHITE
				v.color = GRAY
				v.d = u.d + 1
				v.π = u
				ENQUEUE(Q, v)
		u.color = BLACK
*/

import (
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/gyuho/goraph/graph/gsd"
)

// DeleteNonAlnum removes all alphanumeric characters.
func DeleteNonAlnum(str string) string {
	// alphanumeric (== [0-9A-Za-z])
	// \s is a white space character
	validID := regexp.MustCompile(`[^[:alnum:]\s]`)
	return validID.ReplaceAllString(str, "")
}

// ShowBFS shows the traversed BFS.
func ShowBFS(g *gsd.Graph, src *gsd.Vertex, filename string) string {
	result := "graph " + DeleteNonAlnum(filename) + " {" + "\n"
	for _, edge := range *g.GetEdges() {
		wt := strconv.FormatFloat(edge.(*gsd.Edge).Weight, 'f', -1, 64)
		result += "\t" + edge.(*gsd.Edge).Src.ID + " -- " + edge.(*gsd.Edge).Dst.ID + " [label=" + wt + "]" + "\n"
	}
	tb := BFS(g, src)
	tb = strings.Replace(tb, "→", "--", -1)
	result += "\t" + tb + " [label=BFS, color=blue]" + "\n"
	result += "}"

	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString(result)
	cmd := exec.Command("open", filename)
	er := cmd.Run()
	if er != nil {
		log.Fatal(er)
	}

	return result
}
