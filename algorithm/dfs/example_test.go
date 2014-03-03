package dfs

import (
	"fmt"

	"github.com/gyuho/goraph/graph/gsd"
)

func Example_DFS() {
	testgraph4 := `
S|A,15|B,14|C,9
A|S,15|B,5|D,20|T,44
B|S,14|A,5|D,30|E,18
C|S,9|E,24
D|A,20|B,30|E,2|F,11|T,16
E|B,18|C,24|D,2|F,6|T,19
F|D,11|E,6|T,6
T|A,44|D,16|F,6|E,19
`
	g4 := gsd.ParseToGraph(testgraph4)
	fmt.Println(DFS(g4))
	// S → A → B → D → E → F → T → C

	testgraph5 := `
A|B,7|C,9|F,20
B|A,7|C,10|D,15
C|A,9|B,10|D,11|E,30|F,2
D|B,15|C,11|E,2
E|C,30|D,2|F,9
F|A,20|C,2|E,9
`
	g5 := gsd.ParseToGraph(testgraph5)
	fmt.Println(DFS(g5))
	// A → B → C → D → E → F

	// Output:
	// S → A → B → D → E → F → T → C
	// A → B → C → D → E → F
}
