package gt

import (
	"reflect"
	"testing"
)

var testgraph1 string = `
S|A,100|C,200|B,14
A|S,15|B,5|D,20|T,44
B|S,14|A,5|D,30|E,18
C|S,9|E,24
D|A,20|B,30|E,2|F,11|T,16
E|B,18|C,24|D,2|F,6|T,19
F|D,11|E,6|T,6
T|A,44|D,16|F,6|E,19
`

var testgraph2 string = `
S|A,100|C,200|B,14
A|S,15|B,5|D,20|T,44
B|S,14|A,5|D,30|E,18
C|S,9|E,24
E|B,18|D,2|F,6|T,19
F|D,11|E,6|T,6
T|A,44|D,16|F,6|E,19
`

func Test_NewGraph(test *testing.T) {
	a := NewGraph(5)
	r := reflect.TypeOf(a)
	c := &Graph{}
	if r != reflect.TypeOf(c) {
		test.Error("Type Error")
	}
}

func Test_ParseToGraph(test *testing.T) {
	g1 := ParseToGraph(testgraph1)
	if g1.GetVerticesSize() != 8 {
		test.Error("The graph should be of vertex size 8 but:", g1.GetVerticesSize())
	}

	g2 := ParseToGraph(testgraph2)
	if g2.GetVerticesSize() != 8 {
		test.Error("The graph should be of vertex size 8 but:", g2.GetVerticesSize())
	}
}

func Test_GetVertices(test *testing.T) {
	g1 := ParseToGraph(testgraph1)
	vs1 := g1.GetVertices()
	if len(vs1) != 8 {
		test.Errorf("The graph should be of 8 vertices but: %v", vs1)
	}

	g2 := ParseToGraph(testgraph2)
	vs2 := g2.GetVertices()
	if len(vs2) != 8 {
		test.Errorf("The graph should be of 8 vertices but: %v", vs2)
	}
}

func Test_GetEdgesSize(test *testing.T) {
	g1 := ParseToGraph(testgraph1)
	if g1.GetEdgesSize() != 30 {
		test.Error("The graph should be of edge size 30 but:", g1.GetEdgesSize())
	}

	g2 := ParseToGraph(testgraph2)
	if g2.GetEdgesSize() != 24 {
		test.Error("The graph should be of edge size 24 but:", g1.GetEdgesSize())
	}
}

func Test_GetEdgeWeight(test *testing.T) {
	g1 := ParseToGraph(testgraph1)
	if g1.GetEdgeWeight("S", "C") != 200 {
		test.Error("Should return 200 but:", g1.GetEdgeWeight("S", "C"))
	}

	g2 := ParseToGraph(testgraph2)
	if g2.GetEdgeWeight("D", "B") != 0.0 {
		test.Error("Should return 0.0 but:", g2.GetEdgeWeight("D", "B"))
	}
	if g2.GetEdgeWeight("E", "D") != 2.0 {
		test.Error("Should return 2.0 but:", g2.GetEdgeWeight("E", "D"))
	}
}

func Test_RemoveEdge(test *testing.T) {
	g1 := ParseToGraph(testgraph1)
	g1.RemoveEdge("S", "C")
	if g1.GetEdgesSize() != 29 {
		test.Error("Should return 29 but:", g1.GetEdgesSize())
	}

	g2 := ParseToGraph(testgraph2)
	g2.RemoveEdge("D", "B")
	if g2.GetEdgesSize() != 24 {
		test.Error("Should return 24 but:", g1.GetEdgesSize())
	}

	g2.RemoveEdge("E", "D")
	if g2.GetEdgesSize() != 23 {
		test.Error("Should return 23 but:", g1.GetEdgesSize())
	}
}

func Test_HasEdge(test *testing.T) {
	g1 := ParseToGraph(testgraph1)
	if g1.HasEdge("S", "C") != true {
		test.Error("Should return true but:", g1.HasEdge("S", "C"))
	}

	g2 := ParseToGraph(testgraph2)
	if g2.HasEdge("D", "B") != false {
		test.Error("Should return false but:", g2.HasEdge("D", "B"))
	}

	if g2.HasEdge("E", "D") != true {
		test.Error("Should return true but:", g2.HasEdge("E", "D"))
	}
}

func Test_GetInVertices(test *testing.T) {
	g1 := ParseToGraph(testgraph1)
	if len(g1.GetInVertices("S")) != 3 {
		test.Error("Should return 3 but:", len(g1.GetInVertices("S")))
	}

	g2 := ParseToGraph(testgraph2)
	if len(g2.GetInVertices("B")) != 3 {
		test.Error("Should return 3 but:", len(g2.GetInVertices("B")))
	}

	if len(g2.GetInVertices("D")) != 5 {
		test.Error("Should return 5 but:", len(g2.GetInVertices("D")))
	}
}

func Test_GetOutVertices(test *testing.T) {
	g1 := ParseToGraph(testgraph1)
	if len(g1.GetOutVertices("S")) != 3 {
		test.Error("Should return 3 but:", len(g1.GetOutVertices("S")))
	}

	g2 := ParseToGraph(testgraph2)
	if len(g2.GetOutVertices("B")) != 4 {
		test.Error("Should return 4 but:", len(g2.GetOutVertices("B")))
	}

	if len(g2.GetOutVertices("D")) != 0 {
		test.Error("Should return 0 but:", len(g2.GetOutVertices("D")))
	}

	if len(g2.GetOutVertices("C")) != 2 {
		test.Error("Should return 2 but:", len(g2.GetOutVertices("C")))
	}
}

func Test_Initialize(test *testing.T) {
	g1 := ParseToGraph(testgraph1)
	g1.Initialize(99999.999)

	cnum := true
	mat := g1.Matrix
	for r := range mat {
		for c := range mat[r] {
			if mat[r][c] != 99999.999 {
				cnum = false
			}
		}
	}

	if !cnum {
		test.Error("Should return true but %v", mat)
	}
}
