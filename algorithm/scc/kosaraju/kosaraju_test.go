package kosaraju

import (
	"testing"

	"github.com/gyuho/goraph/graph/gsd"
)

func Test_JSON_Contains(test *testing.T) {
	g15 := gsd.JSONGraph("../../../testgraph/testgraph.json", "testgraph.015")
	a := g15.FindVertexByID("A")
	b := g15.FindVertexByID("B")
	ovs := a.GetOutVertices()
	if !Contains(b, ovs) {
		test.Errorf("Should contain B in OutVertices but %+v", Contains(b, ovs))
	}
}
