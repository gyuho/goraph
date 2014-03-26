package kruskal

import (
	"log"
	"os"
	"os/exec"

	"github.com/gyuho/goraph/graph/gsd"
)

// MSTString returns the MST result in DOT format.
func MSTString(g *gsd.Graph) string {
	mstedges := MST(g)
	result := "graph mst {" + "\n"
	for _, edge := range mstedges {
		result += "\t" + edge.Src.ID + " -- " + edge.Dst.ID + "\n"
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
