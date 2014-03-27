package prim

import (
	"log"
	"os"
	"os/exec"

	"github.com/gyuho/goraph/graph/gsd"
)

// ShowMST shows the Minimum Spanning Tree.
func ShowMST(g *gsd.Graph, outputfile string) {
	str, _ := MST(g)
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
