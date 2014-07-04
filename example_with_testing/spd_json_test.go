package example_with_testing

import (
	"fmt"
	"testing"

	"github.com/gyuho/goraph/algorithm/spd"
	"github.com/gyuho/goraph/graph/gsd"
	// go test -v github.com/gyuho/goraph/example_with_testing
	// go test -v /Users/gyuho/go/src/github.com/gyuho/goraph/example_with_testing/spd_json_test.go
)

func Test_JSON_SPD(test *testing.T) {
	fmt.Println("Dijkstra Shortest Path on testgraph4:")
	g4 := gsd.JSONGraph("../example_files/testgraph.json", "testgraph.004")
	fmt.Println(spd.SPD(g4, "S", "T"))
	fmt.Println(g4.ShowPrev("T"))
	fmt.Println(g4.ShowPrev("F"))
	fmt.Println(g4.ShowPrev("E"))
	fmt.Println(g4.ShowPrev("B"))
	fmt.Println(g4.ShowPrev("S"))
	/*
	   S(=0) → B(=14) → E(=32) → F(=38) → T(=44)
	   Prev of T:  A E D F
	   Prev of F:  E
	   Prev of E:  C B
	   Prev of B:  S
	   Prev of S:

	   BackTrack keeps adding the Prev vertex
	   with the biggest StampD, recursively
	   until we reach the source
	*/

	println()
	fmt.Println("Dijkstra Shortest Path on testgraph5:")
	g5 := gsd.JSONGraph("../example_files/testgraph.json", "testgraph.005")
	fmt.Println(spd.SPD(g5, "A", "E"))
	fmt.Println(g5.ShowPrev("E"))
	fmt.Println(g5.ShowPrev("F"))
	fmt.Println(g5.ShowPrev("C"))
	fmt.Println(g5.ShowPrev("A"))
	/*
	   A(=0) → C(=9) → F(=11) → E(=20)
	   Prev of E:  C F
	   Prev of F:  A C
	   Prev of C:  A
	   Prev of A:

	   BackTrack keeps adding the Prev vertex
	   with the biggest StampD, recursively
	   until we reach the source
	*/

	println()
	fmt.Println("Dijkstra Shortest Path on testgraph10:")
	g10 := gsd.JSONGraph("../example_files/testgraph.json", "testgraph.010")
	fmt.Println(spd.SPD(g10, "A", "E"))
	fmt.Println(g10.ShowPrev("E"))
	fmt.Println(g10.ShowPrev("D"))
	fmt.Println(g10.ShowPrev("B"))
	fmt.Println(g10.ShowPrev("C"))
	fmt.Println(g10.ShowPrev("A"))
	/*
	   A(=0) → C(=9) → B(=19) → D(=34) → E(=36)
	   Prev of E:  C D
	   Prev of D:  B
	   Prev of B:  C
	   Prev of C:  A
	   Prev of A:

	   BackTrack keeps adding the Prev vertex
	   with the biggest StampD, recursively
	   until we reach the source
	*/

	println()
	fmt.Println("Dijkstra Shortest Path on testgraph10o:")
	g10o := gsd.JSONGraph("../example_files/testgraph.json", "testgraph.010")
	fmt.Println(spd.SPD(g10o, "E", "A"))
	fmt.Println(g10o.ShowPrev("A"))
	fmt.Println(g10o.ShowPrev("B"))
	fmt.Println(g10o.ShowPrev("C"))
	fmt.Println(g10o.ShowPrev("F"))
	fmt.Println(g10o.ShowPrev("E"))
	/*
	   E(=0) → F(=9) → C(=11) → B(=21) → A(=22)
	   Prev of A:  F B
	   Prev of B:  C
	   Prev of C:  E F
	   Prev of F:  E
	   Prev of E:

	   BackTrack keeps adding the Prev vertex
	   with the biggest StampD, recursively
	   until we reach the source
	*/

	println()
	fmt.Println("Dijkstra Shortest Path on testgraph11:")
	g11 := gsd.JSONGraph("../example_files/testgraph.json", "testgraph.011")
	fmt.Println(spd.SPD(g11, "S", "T"))
	fmt.Println(g11.ShowPrev("T"))
	fmt.Println(g11.ShowPrev("E"))
	fmt.Println(g11.ShowPrev("D"))
	fmt.Println(g11.ShowPrev("B"))
	fmt.Println(g11.ShowPrev("A"))
	fmt.Println(g11.ShowPrev("S"))
	/*
	   S(=0) → A(=11) → B(=16) → D(=46) → E(=49) → T(=68)
	   Prev of T:  A E
	   Prev of E:  D
	   Prev of D:  A B
	   Prev of B:  S A
	   Prev of A:  S
	   Prev of S:

	   BackTrack keeps adding the Prev vertex
	   with the biggest StampD, recursively
	   until we reach the source
	*/

	println()
	fmt.Println("Dijkstra Shortest Path on testgraph11o:")
	g11o := gsd.JSONGraph("../example_files/testgraph.json", "testgraph.011")
	fmt.Println(spd.SPD(g11o, "T", "S"))
	fmt.Println(g11o.ShowPrev("S"))
	fmt.Println(g11o.ShowPrev("B"))
	fmt.Println(g11o.ShowPrev("E"))
	fmt.Println(g11o.ShowPrev("D"))
	fmt.Println(g11o.ShowPrev("T"))
	/*
	   T(=0) → D(=10) → E(=13) → B(=31) → S(=48)
	   Prev of S:  B
	   Prev of B:  D E
	   Prev of E:  T D
	   Prev of D:  T
	   Prev of T:

	   BackTrack keeps adding the Prev vertex
	   with the biggest StampD, recursively
	   until we reach the source
	*/
}
