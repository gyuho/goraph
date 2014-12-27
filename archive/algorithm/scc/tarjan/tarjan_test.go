package tarjan

import (
	"fmt"
	"testing"

	"github.com/gyuho/goraph/graph/gs"
)

func TestSCC(t *testing.T) {
	g15 := gs.FromJSON("../../../files/testgraph.json", "testgraph.015")
	result15 := SCC(g15)

	fmt.Println(result15)
	// [[H] [F G] [D C] [A E B]]

	if len(result15) != 4 {
		t.Errorf("expected 4 but %v", result15)
	}

	g16 := gs.FromJSON("../../../files/testgraph.json", "testgraph.016")
	result16 := SCC(g16)

	fmt.Println(result16)
	// [[E J] [I] [H D C] [F G B A]]

	if len(result16) != 4 {
		t.Errorf("expected 4 but %v", result16)
	}
}
