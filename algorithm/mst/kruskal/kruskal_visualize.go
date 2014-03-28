package kruskal

import (
	"log"
	"os"
	"os/exec"
	"strconv"

	"github.com/gyuho/goraph/graph/gsd"
)

// MSTString returns the MST result in DOT format.
func MSTString(g *gsd.Graph) string {
	mstedges, _ := MST(g)
	result := "graph KruskalMST {" + "\n"

	for _, edge := range mstedges {
		wt := g.GetEdgeWeight(edge.Src, edge.Dst)[0]
		wts := strconv.FormatFloat(wt, 'f', -1, 64)
		result += "\t" + edge.Src.ID + " -- " + edge.Dst.ID + " [label=" + wts + ", color=blue]" + "\n"
	}

	for _, edge := range *g.GetEdges() {
		result += "\t" + edge.(*gsd.Edge).Src.ID + " -- " + edge.(*gsd.Edge).Dst.ID + "\n"
	}

	result += "}"
	return result
}

// ShowMST shows the Minimum Spanning Tree.
func ShowMST(g *gsd.Graph, outputfile string) {
	str := MSTString(g)
	file, err := os.Create(outputfile)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString(str)
	cmd := exec.Command("open", outputfile)
	er := cmd.Run()
	if er != nil {
		log.Fatal(er)
	}
}
