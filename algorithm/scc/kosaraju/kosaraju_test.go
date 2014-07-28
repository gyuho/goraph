package kosaraju

import (
	"fmt"
	"testing"

	"github.com/gyuho/goraph/graph/gs"
)

func TestSCC(t *testing.T) {
	g15 := gs.FromJSON("../../../files/testgraph.json", "testgraph.015")
	gr15 := gs.FromJSONT("../../../files/testgraph.json", "testgraph.015")
	rs := SCC(g15, gr15)
	fmt.Println(rs)
	// [[B E A] [D C] [G F] [H]]

	if len(rs) != 4 {
		t.Errorf("expected 4 but %v", rs)
	}
	//
	//
	// TODO
	// g16 := gs.FromJSON("../../../files/testgraph.json", "testgraph.016")
	// gr16 := gs.FromJSONT("../../../files/testgraph.json", "testgraph.016")
	// fmt.Println(SCC(g16, gr16))
	// [[B F G A] [D H C] [I] [E J]]
}
