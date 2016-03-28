package goraph_test

import (
	"fmt"
	"log"
	"os"

	"github.com/gyuho/goraph"
)

func ExampleNewGraph() {
	f, err := os.Open("testdata/graph.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	g, err := goraph.NewGraphFromJSON(f, "graph_00")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(g.String())

	// Output for g.String() but it's unordered because it's map:
	// F -- 11.000 -→ D
	// F -- 6.000 -→ E
	// F -- 6.000 -→ T
	// E -- 6.000 -→ F
	// E -- 19.000 -→ T
	// E -- 18.000 -→ B
	// E -- 24.000 -→ C
	// E -- 2.000 -→ D
	// S -- 100.000 -→ A
	// S -- 14.000 -→ B
	// S -- 200.000 -→ C
	// B -- 14.000 -→ S
	// B -- 5.000 -→ A
	// B -- 30.000 -→ D
	// B -- 18.000 -→ E
	// C -- 9.000 -→ S
	// C -- 24.000 -→ E
	// T -- 16.000 -→ D
	// T -- 6.000 -→ F
	// T -- 19.000 -→ E
	// T -- 44.000 -→ A
	// A -- 5.000 -→ B
	// A -- 20.000 -→ D
	// A -- 44.000 -→ T
	// A -- 15.000 -→ S
	// D -- 11.000 -→ F
	// D -- 16.000 -→ T
	// D -- 20.000 -→ A
	// D -- 30.000 -→ B
	// D -- 2.000 -→ E
}
