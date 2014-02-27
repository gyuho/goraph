package tsdfs

import (
	"fmt"
	"testing"

	"github.com/gyuho/goraph/graph/gsd"
)

func Test_TSDFS(test *testing.T) {
	testgraph6 := `
D|B,1|C,1
E|C,1|F,1
B|A,1
A|F,1
`
	fmt.Println("Topological Sort on testgraph6:")
	g6 := gsd.ParseToGraph(testgraph6)
	g6s := TSDFS(g6)
	g6c := "E → D → C → B → A → F"
	if g6s != g6c {
		test.Errorf("Should be same but\n%v\n%v", g6s, g6c)
	}

	testgraph7 := `
C|D,1|E,1
B|D,1
A|E,1|H,1
D|F,1|G,1|H,1
E|G,1
`
	fmt.Println("Topological Sort on testgraph7:")
	g7 := gsd.ParseToGraph(testgraph7)
	g7s := TSDFS(g7)
	g7c := "A → B → C → E → D → H → G → F"
	if g7s != g7c {
		test.Errorf("Should be same but\n%v\n%v", g7s, g7c)
	}
}
