package sp

import (
	"fmt"
	"testing"

	"github.com/gyuho/goraph/graph/gsd"
)

func Test_SP(test *testing.T) {
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
	fmt.Println("SP on testgraph4:")
	g4 := gsd.ParseToGraph(testgraph4)
	g4s, okg4s := SP(g4, "S", "T")
	g4c := "S(=0) → B(=14) → E(=32) → F(=38) → T(=44)"
	okg4c := true
	if g4s != g4c || okg4s != okg4c {
		test.Errorf("Should be same but\n%v\n%v\n%v", g4s, g4c, okg4s)
	}

	testgraph5 := `
A|B,7|C,9|F,20
B|A,7|C,10|D,15
C|A,9|B,10|D,11|E,30|F,2
D|B,15|C,11|E,2
E|C,30|D,2|F,9
F|A,20|C,2|E,9
`
	fmt.Println("SP on testgraph5:")
	g5 := gsd.ParseToGraph(testgraph5)
	g5s, okg5s := SP(g5, "A", "E")
	g5c := "A(=0) → C(=9) → F(=11) → E(=20)"
	okg5c := true
	if g5s != g5c || okg5s != okg5c {
		test.Errorf("testgraph5 Should be same but\n%v\n%v\n%v", g5s, g5c, okg5s)
	}

	testgraph10 := `
A|C,9|F,20
B|A,1|D,15
C|B,10|E,30
D|C,11|E,2
E|C,30|F,9
F|A,20|C,2
`
	fmt.Println("SP on testgraph10:")
	g10 := gsd.ParseToGraph(testgraph10)
	g10s, okg10s := SP(g10, "A", "E")
	g10c := "A(=0) → C(=9) → B(=19) → D(=34) → E(=36)"
	okg10c := true
	if g10s != g10c || okg10s != okg10c {
		test.Errorf("Should be same but\n%v\n%v\n%v", g10s, g10c, okg10s)
	}

	fmt.Println("SP on testgraph10:")
	g10o := gsd.ParseToGraph(testgraph10)
	g10so, okg10so := SP(g10o, "E", "A")
	g10co := "E(=0) → F(=9) → C(=11) → B(=21) → A(=22)"
	okg10co := true
	if g10so != g10co || okg10so != okg10co {
		test.Errorf("Should be same but\n%v\n%v\n%v", g10so, g10co, okg10so)
	}

	testgraph11 := `
S|A,11|B,17|C,9
A|S,11|B,5|D,50|T,500
B|S,17|D,30
C|S,9
D|A,50|B,30|E,3|F,11
E|B,18|C,27|D,3|T,19
F|D,11|E,6|T,77
T|A,500|D,10|F,77|E,19
`
	fmt.Println("SP on testgraph10:")
	g11 := gsd.ParseToGraph(testgraph11)
	g11s, okg11s := SP(g11, "S", "T")
	g11c := "S(=0) → A(=11) → B(=16) → D(=46) → E(=49) → T(=68)"
	okg11c := true
	if g11s != g11c || okg11s != okg11c {
		test.Errorf("Should be same but\n%v\n%v\n%v", g11s, g11c, okg11s)
	}

	fmt.Println("SP on testgraph11:")
	g11o := gsd.ParseToGraph(testgraph11)
	g11so, okg11so := SP(g11o, "T", "S")
	g11co := "T(=0) → D(=10) → E(=13) → B(=31) → S(=48)"
	okg11co := true
	if g11so != g11co || okg11so != okg11co {
		test.Errorf("Should be same but\n%v\n%v\n%v", g11so, g11co, okg11so)
	}

	testgraph12 := `
S|A,7|B,6
A|C,-3|T,9
B|A,8|T,-4|C,5
C|B,-2
T|S,2|C,7
`
	fmt.Println("SP on testgraph12:")
	g12 := gsd.ParseToGraph(testgraph12)
	s12, ok12 := SP(g12, "S", "T")
	s12c := "S(=0) → A(=7) → C(=4) → B(=2) → T(=-2)"
	ok12c := true
	if s12 != s12c || ok12 != ok12c {
		test.Errorf("Should be same but %v", s12)
	}

	testgraph13 := `
S|A,7|B,6
A|C,-3|T,9
B|A,-8|T,-4|C,5
C|B,-2
T|S,2|C,7
`
	fmt.Println("SP on testgraph13:")
	g13 := gsd.ParseToGraph(testgraph13)
	s13, ok13 := SP(g13, "S", "T")
	s13c := "There is negative weighted cycle (No Shortest Path)"
	ok13c := false
	if s13 != s13c || ok13 != ok13c {
		test.Errorf("Should be same but %v", s13)
	}
}
