package spfw

import (
	"testing"

	"github.com/gyuho/goraph/graph/gt"
)

func Test_SPFW(test *testing.T) {
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
	g4 := gt.ParseToGraph(testgraph4)
	g4s, g4m := SPFW(g4, "S", "T")
	g4c := 44.0
	if g4s != g4c {
		test.Errorf("Should be same but\n%v\n%v\n%v", g4s, g4c, g4m)
	}

	testgraph5 := `
A|B,7|C,9|F,20
B|A,7|C,10|D,15
C|A,9|B,10|D,11|E,30|F,2
D|B,15|C,11|E,2
E|C,30|D,2|F,9
F|A,20|C,2|E,9
`
	g5 := gt.ParseToGraph(testgraph5)
	g5s, g5m := SPFW(g5, "A", "E")
	g5c := 20.0
	if g5s != g5c {
		test.Errorf("testgraph5 Should be same but\n%v\n%v\n%v", g5s, g5c, g5m)
	}

	testgraph10 := `
A|C,9|F,20
B|A,1|D,15
C|B,10|E,30
D|C,11|E,2
E|C,30|F,9
F|A,20|C,2
`
	g10 := gt.ParseToGraph(testgraph10)
	g10s, g10m := SPFW(g10, "A", "E")
	g10c := 36.0
	if g10s != g10c {
		test.Errorf("Should be same but\n%v\n%v\n%v", g10s, g10c, g10m)
	}

	g10o := gt.ParseToGraph(testgraph10)
	g10so, g10m := SPFW(g10o, "E", "A")
	g10co := 22.0
	if g10so != g10co {
		test.Errorf("Should be same but\n%v\n%v\n%v", g10so, g10co, g10m)
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
	g11 := gt.ParseToGraph(testgraph11)
	g11s, g11m := SPFW(g11, "S", "T")
	g11c := 68.0
	if g11s != g11c {
		test.Errorf("Should be same but\n%v\n%v\n%v", g11s, g11c, g11m)
	}

	g11o := gt.ParseToGraph(testgraph11)
	g11so, g11om := SPFW(g11o, "T", "S")
	g11co := 48.0
	if g11so != g11co {
		test.Errorf("Should be same but\n%v\n%v\n%v", g11so, g11co, g11om)
	}
}
