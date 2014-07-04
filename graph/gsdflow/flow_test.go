package gsdflow

import "testing"

func Test_JSON_GetEdge(test *testing.T) {
	g := JSONGraph("../../example_files/testgraph.json", "testgraph.001")
	a := g.FindVertexByID("A")
	t := g.FindVertexByID("T")
	edge := g.GetEdge(a, t)
	if edge.Weight != 44 {
		test.Errorf("In testgraph1, it should return 44 but %+v", edge)
	}
}

func Test_JSON_GetEdgeWeight(test *testing.T) {
	g := JSONGraph("../../example_files/testgraph.json", "testgraph.001")
	testCases := []struct {
		vertices []string
		weight   float64
	}{
		{[]string{"S", "B"}, 14.0},
		{[]string{"A", "B"}, 5.0},
		{[]string{"A", "D"}, 20.0},
		{[]string{"A", "T"}, 44.0},
		{[]string{"T", "A"}, 44.0},
		{[]string{"D", "E"}, 2.0},
		{[]string{"E", "D"}, 2.0},
		{[]string{"C", "E"}, 24.0},
		{[]string{"B", "E"}, 18.0},
		{[]string{"D", "T"}, 16.0},
		{[]string{"T", "D"}, 16.0},
		{[]string{"F", "E"}, 6.0},
		{[]string{"E", "F"}, 6.0},
		{[]string{"E", "T"}, 19.0},
		{[]string{"S", "C"}, 200.0},
		{[]string{"S", "A"}, 100.0},
	}
	for _, testCase := range testCases {
		weights := g.GetEdgeWeight(g.FindVertexByID(testCase.vertices[0]), g.FindVertexByID(testCase.vertices[1]))
		exist := false
		for _, val := range weights {
			if val == testCase.weight {
				exist = true
			}
		}
		if !exist {
			test.Errorf("In testgraph1, Expected '%#v'. But %#v", testCase.weight, weights)
		}
	}

	var testgraph3 string = `
S|A,100|C,200|B,14|B,6
A|S,15|B,5|D,20|T,44
B|S,14|A,5|D,30|E,18
C|S,9|E,24
D|A,20|B,30|E,2|F,11|T,16
E|B,18|C,24|D,2|F,6|T,19
F|D,11|E,6|T,6
T|A,44|D,16|F,6|E,19
`

	g3 := ParseToGraph(testgraph3)
	testCases3 := []struct {
		vertices []string
		weight   float64
	}{
		{[]string{"S", "B"}, 14.0}, // duplicate edges
		{[]string{"S", "B"}, 6.0},  // duplicate edges
		{[]string{"A", "B"}, 5.0},
		{[]string{"A", "D"}, 20.0},
		{[]string{"A", "T"}, 44.0},
		{[]string{"T", "A"}, 44.0},
		{[]string{"D", "E"}, 2.0},
		{[]string{"E", "D"}, 2.0},
		{[]string{"C", "E"}, 24.0},
		{[]string{"B", "E"}, 18.0},
		{[]string{"D", "T"}, 16.0},
		{[]string{"T", "D"}, 16.0},
		{[]string{"F", "E"}, 6.0},
		{[]string{"E", "F"}, 6.0},
		{[]string{"E", "T"}, 19.0},
		{[]string{"S", "C"}, 200.0},
		{[]string{"S", "A"}, 100.0},
	}
	for _, testCase := range testCases3 {
		weights := g3.GetEdgeWeight(g3.FindVertexByID(testCase.vertices[0]), g3.FindVertexByID(testCase.vertices[1]))
		exist := false
		for _, val := range weights {
			if val == testCase.weight {
				exist = true
			}
		}
		if !exist {
			test.Errorf("In testgraph3, Expected '%#v'. But %#v. %#v, %#v", testCase.weight, weights, testCase.vertices[0], testCase.vertices[1])
		}
	}
}

func Test_JSON_GetEdgeFlow(test *testing.T) {
	g := JSONGraph("../../example_files/testgraph.json", "testgraph.001")
	testCases := []struct {
		vertices []string
		flow     float64
	}{
		{[]string{"S", "B"}, 0},
		{[]string{"A", "B"}, 0},
		{[]string{"A", "D"}, 0},
		{[]string{"A", "T"}, 0},
		{[]string{"T", "A"}, 0},
		{[]string{"D", "E"}, 0},
		{[]string{"E", "D"}, 0},
		{[]string{"C", "E"}, 0},
		{[]string{"B", "E"}, 0},
		{[]string{"D", "T"}, 0},
		{[]string{"T", "D"}, 0},
		{[]string{"F", "E"}, 0},
		{[]string{"E", "F"}, 0},
		{[]string{"E", "T"}, 0},
		{[]string{"S", "C"}, 0},
		{[]string{"S", "A"}, 0},
	}
	for _, testCase := range testCases {
		flows := g.GetEdgeFlow(g.FindVertexByID(testCase.vertices[0]), g.FindVertexByID(testCase.vertices[1]))
		exist := true
		for _, val := range flows {
			if val != testCase.flow {
				exist = false
			}
		}
		if !exist {
			test.Errorf("In testgraph1, Expected '%#v'. But %#v", testCase.flow, flows)
		}
	}

}
func Test_JSON_UpdateWeight(test *testing.T) {
	g := JSONGraph("../../example_files/testgraph.json", "testgraph.001")
	testCases := []struct {
		vertices []string
		weight   float64
	}{
		{[]string{"S", "B"}, 914.0},
		{[]string{"A", "B"}, 95.0},
		{[]string{"A", "D"}, 920.0},
		{[]string{"A", "T"}, 944.0},
		{[]string{"T", "A"}, 944.0},
		{[]string{"D", "E"}, 92.0},
		{[]string{"E", "D"}, 92.0},
		{[]string{"C", "E"}, 924.0},
		{[]string{"B", "E"}, 918.0},
		{[]string{"D", "T"}, 916.0},
		{[]string{"T", "D"}, 916.0},
		{[]string{"F", "E"}, 96.0},
		{[]string{"E", "F"}, 96.0},
		{[]string{"E", "T"}, 919.0},
		{[]string{"S", "C"}, 9200.0},
		{[]string{"S", "A"}, 9100.0},
	}
	for _, testCase := range testCases {
		g.UpdateWeight(g.FindVertexByID(testCase.vertices[0]), g.FindVertexByID(testCase.vertices[1]), testCase.weight)
		weights := g.GetEdgeWeight(g.FindVertexByID(testCase.vertices[0]), g.FindVertexByID(testCase.vertices[1]))
		exist := true
		for _, val := range weights {
			if val != testCase.weight {
				exist = false
			}
		}
		if !exist {
			test.Errorf("In testgraph1, Expected '%#v'. But %#v", testCase.weight, weights)
		}
	}
}

func Test_JSON_UpdateFlow(test *testing.T) {
	g := JSONGraph("../../example_files/testgraph.json", "testgraph.001")
	testCases := []struct {
		vertices []string
		flow     float64
	}{
		{[]string{"S", "B"}, 1},
		{[]string{"A", "B"}, 1},
		{[]string{"A", "D"}, 1},
		{[]string{"A", "T"}, 1},
		{[]string{"T", "A"}, 1},
		{[]string{"D", "E"}, 1},
		{[]string{"E", "D"}, 1},
		{[]string{"C", "E"}, 1},
		{[]string{"B", "E"}, 1},
		{[]string{"D", "T"}, 1},
		{[]string{"T", "D"}, 1},
		{[]string{"F", "E"}, 1},
		{[]string{"E", "F"}, 1},
		{[]string{"E", "T"}, 1},
		{[]string{"S", "C"}, 1},
		{[]string{"S", "A"}, 1},
	}
	for _, testCase := range testCases {
		g.UpdateFlow(g.FindVertexByID(testCase.vertices[0]), g.FindVertexByID(testCase.vertices[1]), testCase.flow)
		flows := g.GetEdgeFlow(g.FindVertexByID(testCase.vertices[0]), g.FindVertexByID(testCase.vertices[1]))
		exist := true
		for _, val := range flows {
			if val != testCase.flow {
				exist = false
			}
		}
		if !exist {
			test.Errorf("In testgraph1, Expected '%#v'. But %#v", testCase.flow, flows)
		}
	}
}

func Test_JSON_AddFlow(test *testing.T) {
	g := JSONGraph("../../example_files/testgraph.json", "testgraph.001")
	testCases := []struct {
		vertices []string
		flow     float64
	}{
		{[]string{"S", "B"}, 1},
		{[]string{"A", "B"}, 1},
		{[]string{"A", "D"}, 1},
		{[]string{"A", "T"}, 1},
		{[]string{"T", "A"}, 1},
		{[]string{"D", "E"}, 1},
		{[]string{"E", "D"}, 1},
		{[]string{"C", "E"}, 1},
		{[]string{"B", "E"}, 1},
		{[]string{"D", "T"}, 1},
		{[]string{"T", "D"}, 1},
		{[]string{"F", "E"}, 1},
		{[]string{"E", "F"}, 1},
		{[]string{"E", "T"}, 1},
		{[]string{"S", "C"}, 1},
		{[]string{"S", "A"}, 1},
	}
	for _, testCase := range testCases {
		g.AddFlow(g.FindVertexByID(testCase.vertices[0]), g.FindVertexByID(testCase.vertices[1]), testCase.flow)
		flows := g.GetEdgeFlow(g.FindVertexByID(testCase.vertices[0]), g.FindVertexByID(testCase.vertices[1]))
		exist := true
		for _, val := range flows {
			if val != testCase.flow {
				exist = false
			}
		}
		if !exist {
			test.Errorf("In testgraph1, Expected '%#v'. But %#v", testCase.flow, flows)
		}
	}
}

func Test_JSON_SubFlow(test *testing.T) {
	g := JSONGraph("../../example_files/testgraph.json", "testgraph.001")
	testCases := []struct {
		vertices []string
		flow     float64
	}{
		{[]string{"S", "B"}, 1},
		{[]string{"A", "B"}, 1},
		{[]string{"A", "D"}, 1},
		{[]string{"A", "T"}, 1},
		{[]string{"T", "A"}, 1},
		{[]string{"D", "E"}, 1},
		{[]string{"E", "D"}, 1},
		{[]string{"C", "E"}, 1},
		{[]string{"B", "E"}, 1},
		{[]string{"D", "T"}, 1},
		{[]string{"T", "D"}, 1},
		{[]string{"F", "E"}, 1},
		{[]string{"E", "F"}, 1},
		{[]string{"E", "T"}, 1},
		{[]string{"S", "C"}, 1},
		{[]string{"S", "A"}, 1},
	}
	for _, testCase := range testCases {
		g.AddFlow(g.FindVertexByID(testCase.vertices[0]), g.FindVertexByID(testCase.vertices[1]), testCase.flow)
		g.SubFlow(g.FindVertexByID(testCase.vertices[0]), g.FindVertexByID(testCase.vertices[1]), testCase.flow)
		flows := g.GetEdgeFlow(g.FindVertexByID(testCase.vertices[0]), g.FindVertexByID(testCase.vertices[1]))
		exist := true
		for _, val := range flows {
			if val != 0 {
				exist = false
			}
		}
		if !exist {
			test.Errorf("In testgraph1, Expected '%#v'. But %#v", testCase.flow, flows)
		}
	}
}

func Test_JSON_IsFull(test *testing.T) {
	g := JSONGraph("../../example_files/testgraph.json", "testgraph.001")
	testCases := []struct {
		vertices []string
	}{
		{[]string{"S", "B"}},
		{[]string{"A", "B"}},
		{[]string{"A", "D"}},
		{[]string{"A", "T"}},
		{[]string{"T", "A"}},
		{[]string{"D", "E"}},
		{[]string{"E", "D"}},
		{[]string{"C", "E"}},
		{[]string{"B", "E"}},
		{[]string{"D", "T"}},
		{[]string{"T", "D"}},
		{[]string{"F", "E"}},
		{[]string{"E", "F"}},
		{[]string{"E", "T"}},
		{[]string{"S", "C"}},
		{[]string{"S", "A"}},
	}
	for _, testCase := range testCases {
		wt := g.GetEdgeWeight(g.FindVertexByID(testCase.vertices[0]), g.FindVertexByID(testCase.vertices[1]))[0]
		g.AddFlow(g.FindVertexByID(testCase.vertices[0]), g.FindVertexByID(testCase.vertices[1]), wt)
	}
	for _, testCase := range testCases {
		if !g.IsFull(g.FindVertexByID(testCase.vertices[0]), g.FindVertexByID(testCase.vertices[1])) {
			test.Errorf("Should be full but %v %v", testCase.vertices[0], testCase.vertices[1])
		}
	}
}
