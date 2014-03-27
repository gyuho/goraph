// Package viz visualizes graphs with Graphviz.
// (http://www.graphviz.org/)
package viz

import (
	"log"
	"os/exec"

	"github.com/gyuho/goraph/viz/dot"
)

// Show converts the JSON file to DOT file format
// and opens the visualized DOT file.
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

// ShowReverse converts the JSON file to DOT file format
// and opens the visualized DOT file in a reverse order.
func ShowReverse(filename, graphID, outputfile string) {
	dot.ConvertReverse(filename, graphID, outputfile)
	// func Command(name string, arg ...string) *Cmd
	// Command returns the Cmd struct to execute the named program with the given arguments.
	cmd := exec.Command("open", outputfile)
	er := cmd.Run()
	if er != nil {
		log.Fatal(er)
	}
}
