package tarjan

import (
	"fmt"
	"testing"

	"github.com/gyuho/goraph/graph/gsd"
)

func Test_JSON_SCC(test *testing.T) {
	g15 := gsd.JSONGraph("../../../example_files/testgraph.json", "testgraph.015")
	result15 := SCC(g15)
	rc15 := "[[H] [G F] [D C] [E B A]]"
	if fmt.Sprintf("%v", result15) != rc15 {
		test.Errorf("Should return \n%+v\nbut\n%+v", result15, rc15)
	}

	g16 := gsd.JSONGraph("../../../example_files/testgraph.json", "testgraph.016")
	fmt.Println(SCC(g16))
	// [[E J] [I] [H D C] [F G B A]]
}

func Test_JSON_Contains(test *testing.T) {
	g15 := gsd.JSONGraph("../../../example_files/testgraph.json", "testgraph.015")
	a := g15.FindVertexByID("A")
	b := g15.FindVertexByID("B")
	ovs := a.GetOutVertices()
	if !Contains(b, ovs) {
		test.Errorf("Should contain B in OutVertices but %+v", Contains(b, ovs))
	}
}
