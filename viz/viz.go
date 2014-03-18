// Package viz visualizes graphs with Graphviz.
// (http://www.graphviz.org/)
package viz

import (
	"log"
	"os/exec"

	"github.com/gyuho/goraph/viz/dot"
)

func Show(filename, graphID, outputfile string) {
	dot.Convert(filename, graphID, outputfile)
	// func Command(name string, arg ...string) *Cmd
	// Command returns the Cmd struct to execute the named program with the given arguments.
	cmd := exec.Command("open", outputfile)
	er := cmd.Run()
	if er != nil {
		log.Fatal(er)
	}
}
