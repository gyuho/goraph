package testgraph

// Graph contains test data.
type Graph struct {
	Name           string
	TotalNodeCount int
	TotalEdgeCount int
	IsDAG          bool
	EdgeToWeight   []EdgeToWeight
}

// EdgeToWeight maps Nodes to their weight between.
type EdgeToWeight struct {
	Nodes  []string
	Weight float64
}

// GraphSlice contains all test Graph data.
var GraphSlice = []Graph{
	Graph00,
	Graph01,
	Graph02,
	Graph03,
	Graph04,
	Graph05,
	Graph06,
	Graph07,
	Graph08,
	Graph09,
	Graph10,
	Graph11,
	Graph12,
	Graph13,
	Graph14,
	Graph15,
	Graph16,
}

// Graph00 represents graph_00.
var Graph00 = Graph{
	Name:           "graph_00",
	TotalNodeCount: 8,
	TotalEdgeCount: 30,
	IsDAG:          false,
	EdgeToWeight: []EdgeToWeight{
		{[]string{"S", "A"}, 100.0},
		{[]string{"S", "B"}, 14.0},
		{[]string{"S", "C"}, 200.0},

		{[]string{"A", "S"}, 15.0},
		{[]string{"A", "B"}, 5.0},
		{[]string{"A", "D"}, 20.0},
		{[]string{"A", "T"}, 44.0},

		{[]string{"B", "S"}, 14.0},
		{[]string{"B", "A"}, 5.0},
		{[]string{"B", "D"}, 30.0},
		{[]string{"B", "E"}, 18.0},

		{[]string{"C", "S"}, 9.0},
		{[]string{"C", "E"}, 24.0},

		{[]string{"D", "A"}, 20.0},
		{[]string{"D", "B"}, 30.0},
		{[]string{"D", "E"}, 2.0},
		{[]string{"D", "F"}, 11.0},
		{[]string{"D", "T"}, 16.0},

		{[]string{"E", "B"}, 18.0},
		{[]string{"E", "C"}, 24.0},
		{[]string{"E", "D"}, 2.0},
		{[]string{"E", "F"}, 6.0},
		{[]string{"E", "T"}, 19.0},

		{[]string{"F", "D"}, 11.0},
		{[]string{"F", "E"}, 6.0},
		{[]string{"F", "T"}, 6.0},

		{[]string{"T", "A"}, 44.0},
		{[]string{"T", "D"}, 16.0},
		{[]string{"T", "F"}, 6.0},
		{[]string{"T", "E"}, 19.0},
	},
}

// Graph01 represents graph_01.
var Graph01 = Graph{
	Name:           "graph_01",
	TotalNodeCount: 8,
	TotalEdgeCount: 24,
	IsDAG:          false,
	EdgeToWeight: []EdgeToWeight{
		{[]string{"S", "A"}, 100.0},
		{[]string{"S", "B"}, 14.0},
		{[]string{"S", "C"}, 200.0},

		{[]string{"A", "S"}, 15.0},
		{[]string{"A", "B"}, 5.0},
		{[]string{"A", "D"}, 20.0},
		{[]string{"A", "T"}, 44.0},

		{[]string{"B", "S"}, 14.0},
		{[]string{"B", "A"}, 5.0},
		{[]string{"B", "D"}, 30.0},
		{[]string{"B", "E"}, 18.0},

		{[]string{"C", "S"}, 9.0},
		{[]string{"C", "E"}, 24.0},

		{[]string{"E", "B"}, 18.0},
		{[]string{"E", "D"}, 2.0},
		{[]string{"E", "F"}, 6.0},
		{[]string{"E", "T"}, 19.0},

		{[]string{"F", "D"}, 11.0},
		{[]string{"F", "E"}, 6.0},
		{[]string{"F", "T"}, 6.0},

		{[]string{"T", "A"}, 44.0},
		{[]string{"T", "D"}, 16.0},
		{[]string{"T", "F"}, 6.0},
		{[]string{"T", "E"}, 19.0},
	},
}

// Graph02 represents graph_02.
var Graph02 = Graph{
	Name:           "graph_02",
	TotalNodeCount: 8,
	TotalEdgeCount: 24,
	IsDAG:          false,
	EdgeToWeight: []EdgeToWeight{
		{[]string{"S", "A"}, 100.0},
		{[]string{"S", "B"}, 20.0},
		{[]string{"S", "C"}, 200.0},

		{[]string{"A", "S"}, 15.0},
		{[]string{"A", "B"}, 5.0},
		{[]string{"A", "D"}, 20.0},
		{[]string{"A", "T"}, 44.0},

		{[]string{"B", "S"}, 14.0},
		{[]string{"B", "A"}, 5.0},
		{[]string{"B", "D"}, 30.0},
		{[]string{"B", "E"}, 18.0},

		{[]string{"C", "S"}, 9.0},
		{[]string{"C", "E"}, 24.0},

		{[]string{"E", "B"}, 18.0},
		{[]string{"E", "D"}, 2.0},
		{[]string{"E", "F"}, 6.0},
		{[]string{"E", "T"}, 19.0},

		{[]string{"F", "D"}, 11.0},
		{[]string{"F", "E"}, 6.0},
		{[]string{"F", "T"}, 6.0},

		{[]string{"T", "A"}, 44.0},
		{[]string{"T", "D"}, 16.0},
		{[]string{"T", "F"}, 6.0},
		{[]string{"T", "E"}, 19.0},
	},
}

// Graph03 represents graph_03.
var Graph03 = Graph{
	Name:           "graph_03",
	TotalNodeCount: 8,
	TotalEdgeCount: 28,
	IsDAG:          false,
	EdgeToWeight: []EdgeToWeight{
		{[]string{"S", "B"}, 14.0},

		{[]string{"A", "S"}, 15.0},
		{[]string{"A", "B"}, 5.0},
		{[]string{"A", "D"}, 20.0},
		{[]string{"A", "T"}, 44.0},

		{[]string{"B", "S"}, 14.0},
		{[]string{"B", "A"}, 5.0},
		{[]string{"B", "D"}, 30.0},
		{[]string{"B", "E"}, 18.0},

		{[]string{"C", "S"}, 9.0},
		{[]string{"C", "E"}, 24.0},

		{[]string{"D", "A"}, 20.0},
		{[]string{"D", "B"}, 30.0},
		{[]string{"D", "E"}, 2.0},
		{[]string{"D", "F"}, 11.0},
		{[]string{"D", "T"}, 16.0},

		{[]string{"E", "B"}, 18.0},
		{[]string{"E", "C"}, 24.0},
		{[]string{"E", "D"}, 2.0},
		{[]string{"E", "F"}, 6.0},
		{[]string{"E", "T"}, 19.0},

		{[]string{"F", "D"}, 11.0},
		{[]string{"F", "E"}, 6.0},
		{[]string{"F", "T"}, 6.0},

		{[]string{"T", "A"}, 44.0},
		{[]string{"T", "D"}, 16.0},
		{[]string{"T", "F"}, 6.0},
		{[]string{"T", "E"}, 19.0},
	},
}

// Graph04 represents graph_04.
var Graph04 = Graph{
	Name:           "graph_04",
	TotalNodeCount: 6,
	TotalEdgeCount: 20,
	IsDAG:          false,
	EdgeToWeight: []EdgeToWeight{
		{[]string{"A", "B"}, 7.0},
		{[]string{"A", "C"}, 9.0},
		{[]string{"A", "F"}, 20.0},

		{[]string{"B", "A"}, 7.0},
		{[]string{"B", "C"}, 10.0},
		{[]string{"B", "D"}, 15.0},

		{[]string{"C", "A"}, 9.0},
		{[]string{"C", "B"}, 10.0},
		{[]string{"C", "D"}, 11.0},
		{[]string{"C", "E"}, 30.0},
		{[]string{"C", "F"}, 2.0},

		{[]string{"D", "B"}, 15.0},
		{[]string{"D", "C"}, 11.0},
		{[]string{"D", "E"}, 2.0},

		{[]string{"E", "F"}, 9.0},
		{[]string{"E", "C"}, 30.0},
		{[]string{"E", "D"}, 2.0},

		{[]string{"F", "A"}, 20.0},
		{[]string{"F", "C"}, 2.0},
		{[]string{"F", "E"}, 9.0},
	},
}

// Graph05 represents graph_05.
var Graph05 = Graph{
	Name:           "graph_05",
	TotalNodeCount: 6,
	TotalEdgeCount: 6,
	IsDAG:          true,
	EdgeToWeight: []EdgeToWeight{
		{[]string{"A", "F"}, 1.0},

		{[]string{"B", "A"}, 1.0},

		{[]string{"D", "B"}, 1.0},
		{[]string{"D", "C"}, 1.0},

		{[]string{"E", "C"}, 1.0},
		{[]string{"E", "F"}, 1.0},
	},
}

// Graph06 represents graph_06.
var Graph06 = Graph{
	Name:           "graph_06",
	TotalNodeCount: 8,
	TotalEdgeCount: 9,
	IsDAG:          true,
	EdgeToWeight: []EdgeToWeight{
		{[]string{"A", "E"}, 1.0},
		{[]string{"A", "H"}, 1.0},

		{[]string{"B", "D"}, 1.0},

		{[]string{"C", "D"}, 1.0},
		{[]string{"C", "E"}, 1.0},

		{[]string{"D", "F"}, 1.0},
		{[]string{"D", "G"}, 1.0},
		{[]string{"D", "H"}, 1.0},

		{[]string{"E", "G"}, 1.0},
	},
}

// Graph07 represents graph_07.
var Graph07 = Graph{
	Name:           "graph_07",
	TotalNodeCount: 6,
	TotalEdgeCount: 8,
	IsDAG:          false,
	EdgeToWeight: []EdgeToWeight{
		{[]string{"A", "E"}, 1.0},
		{[]string{"A", "F"}, 1.0},

		{[]string{"B", "A"}, 1.0},

		{[]string{"D", "B"}, 1.0},
		{[]string{"D", "C"}, 1.0},

		{[]string{"E", "D"}, 1.0},
		{[]string{"E", "C"}, 1.0},
		{[]string{"E", "F"}, 1.0},
	},
}

// Graph08 represents graph_08.
var Graph08 = Graph{
	Name:           "graph_08",
	TotalNodeCount: 8,
	TotalEdgeCount: 15,
	IsDAG:          false,
	EdgeToWeight: []EdgeToWeight{
		{[]string{"A", "B"}, 1.0},
		{[]string{"A", "E"}, 1.0},
		{[]string{"A", "H"}, 1.0},

		{[]string{"B", "C"}, 1.0},
		{[]string{"B", "D"}, 1.0},

		{[]string{"C", "D"}, 1.0},
		{[]string{"C", "E"}, 1.0},

		{[]string{"D", "F"}, 1.0},
		{[]string{"D", "G"}, 1.0},
		{[]string{"D", "H"}, 1.0},

		{[]string{"E", "A"}, 1.0},
		{[]string{"E", "G"}, 1.0},

		{[]string{"F", "E"}, 1.0},

		{[]string{"G", "H"}, 1.0},

		{[]string{"H", "F"}, 1.0},
	},
}

// Graph09 represents graph_09.
var Graph09 = Graph{
	Name:           "graph_09",
	TotalNodeCount: 6,
	TotalEdgeCount: 12,
	IsDAG:          false,
	EdgeToWeight: []EdgeToWeight{
		{[]string{"A", "C"}, 9.0},
		{[]string{"A", "F"}, 20.0},

		{[]string{"B", "A"}, 1.0},
		{[]string{"B", "D"}, 15.0},

		{[]string{"C", "B"}, 10.0},
		{[]string{"C", "E"}, 30.0},

		{[]string{"D", "C"}, 11.0},
		{[]string{"D", "E"}, 2.0},

		{[]string{"E", "C"}, 30.0},
		{[]string{"E", "F"}, 9.0},

		{[]string{"F", "A"}, 20.0},
		{[]string{"F", "C"}, 2.0},
	},
}

// Graph10 represents graph_10.
var Graph10 = Graph{
	Name:           "graph_10",
	TotalNodeCount: 8,
	TotalEdgeCount: 25,
	IsDAG:          false,
	EdgeToWeight: []EdgeToWeight{
		{[]string{"S", "A"}, 11.0},
		{[]string{"S", "B"}, 17.0},
		{[]string{"S", "C"}, 9.0},

		{[]string{"A", "S"}, 11.0},
		{[]string{"A", "B"}, 5.0},
		{[]string{"A", "D"}, 50.0},
		{[]string{"A", "T"}, 500.0},

		{[]string{"B", "S"}, 17.0},
		{[]string{"B", "D"}, 30.0},

		{[]string{"C", "S"}, 9.0},

		{[]string{"D", "A"}, 50.0},
		{[]string{"D", "B"}, 30.0},
		{[]string{"D", "E"}, 3.0},
		{[]string{"D", "F"}, 11.0},

		{[]string{"E", "B"}, 18.0},
		{[]string{"E", "D"}, 2.0},
		{[]string{"E", "F"}, 6.0},
		{[]string{"E", "T"}, 19.0},

		{[]string{"F", "D"}, 11.0},
		{[]string{"F", "E"}, 6.0},
		{[]string{"F", "T"}, 77.0},

		{[]string{"T", "A"}, 500.0},
		{[]string{"T", "D"}, 10.0},
		{[]string{"T", "F"}, 77.0},
		{[]string{"T", "E"}, 19.0},
	},
}

// Graph11 represents graph_11.
var Graph11 = Graph{
	Name:           "graph_11",
	TotalNodeCount: 5,
	TotalEdgeCount: 10,
	IsDAG:          false,
	EdgeToWeight: []EdgeToWeight{
		{[]string{"S", "A"}, 7.0},
		{[]string{"S", "B"}, 6.0},

		{[]string{"A", "C"}, -3.0},
		{[]string{"A", "T"}, 9.0},

		{[]string{"B", "A"}, 8.0},
		{[]string{"B", "C"}, 5.0},
		{[]string{"B", "T"}, -4.0},

		{[]string{"C", "B"}, -2.0},

		{[]string{"T", "C"}, 7.0},
		{[]string{"T", "S"}, 2.0},
	},
}

// Graph12 represents graph_12.
var Graph12 = Graph{
	Name:           "graph_12",
	TotalNodeCount: 5,
	TotalEdgeCount: 10,
	IsDAG:          false,
	EdgeToWeight: []EdgeToWeight{
		{[]string{"S", "A"}, 7.0},
		{[]string{"S", "B"}, 6.0},

		{[]string{"A", "C"}, -3.0},
		{[]string{"A", "T"}, 9.0},

		{[]string{"B", "A"}, -8.0},
		{[]string{"B", "C"}, 5.0},
		{[]string{"B", "T"}, -4.0},

		{[]string{"C", "B"}, -2.0},

		{[]string{"T", "S"}, 2.0},
		{[]string{"T", "C"}, 7.0},
	},
}

// Graph13 represents graph_13.
var Graph13 = Graph{
	Name:           "graph_13",
	TotalNodeCount: 9,
	TotalEdgeCount: 28,
	IsDAG:          false,
	EdgeToWeight: []EdgeToWeight{
		{[]string{"A", "B"}, 4.0},
		{[]string{"A", "H"}, 8.0},

		{[]string{"B", "A"}, 4.0},
		{[]string{"B", "H"}, 11.0},
		{[]string{"B", "C"}, 8.0},

		{[]string{"C", "B"}, 8.0},
		{[]string{"C", "I"}, 2.0},
		{[]string{"C", "F"}, 4.0},
		{[]string{"C", "D"}, 7.0},

		{[]string{"D", "C"}, 7.0},
		{[]string{"D", "F"}, 14.0},
		{[]string{"D", "E"}, 9.0},

		{[]string{"E", "D"}, 9.0},
		{[]string{"E", "F"}, 10.0},

		{[]string{"F", "G"}, 2.0},
		{[]string{"F", "C"}, 4.0},
		{[]string{"F", "D"}, 14.0},
		{[]string{"F", "E"}, 10.0},

		{[]string{"G", "H"}, 1.0},
		{[]string{"G", "I"}, 6.0},
		{[]string{"G", "F"}, 2.0},

		{[]string{"H", "A"}, 8.0},
		{[]string{"H", "B"}, 11.0},
		{[]string{"H", "I"}, 7.0},
		{[]string{"H", "G"}, 1.0},

		{[]string{"I", "H"}, 7.0},
		{[]string{"I", "G"}, 6.0},
		{[]string{"I", "C"}, 2.0},
	},
}

// Graph14 represents graph_14.
var Graph14 = Graph{
	Name:           "graph_14",
	TotalNodeCount: 8,
	TotalEdgeCount: 14,
	IsDAG:          false,
	EdgeToWeight: []EdgeToWeight{
		{[]string{"A", "B"}, 1.0},

		{[]string{"B", "E"}, 1.0},
		{[]string{"B", "F"}, 1.0},
		{[]string{"B", "C"}, 1.0},

		{[]string{"C", "D"}, 1.0},
		{[]string{"C", "G"}, 1.0},

		{[]string{"D", "C"}, 1.0},
		{[]string{"D", "H"}, 1.0},

		{[]string{"E", "A"}, 1.0},
		{[]string{"E", "F"}, 1.0},

		{[]string{"F", "G"}, 1.0},

		{[]string{"G", "F"}, 1.0},
		{[]string{"G", "H"}, 1.0},

		{[]string{"H", "H"}, 1.0},
	},
}

// Graph15 represents graph_15.
var Graph15 = Graph{
	Name:           "graph_15",
	TotalNodeCount: 10,
	TotalEdgeCount: 17,
	IsDAG:          false,
	EdgeToWeight: []EdgeToWeight{
		{[]string{"A", "B"}, 1.0},
		{[]string{"A", "F"}, 1.0},

		{[]string{"B", "C"}, 1.0},
		{[]string{"B", "G"}, 1.0},

		{[]string{"C", "D"}, 1.0},
		{[]string{"C", "H"}, 1.0},

		{[]string{"D", "H"}, 1.0},
		{[]string{"D", "I"}, 1.0},
		{[]string{"D", "J"}, 1.0},
		{[]string{"D", "E"}, 1.0},

		{[]string{"E", "J"}, 1.0},

		{[]string{"F", "G"}, 1.0},

		{[]string{"G", "A"}, 1.0},
		{[]string{"G", "H"}, 1.0},

		{[]string{"H", "C"}, 1.0},

		{[]string{"I", "J"}, 1.0},

		{[]string{"J", "E"}, 1.0},
	},
}

// Graph16 represents graph_16.
var Graph16 = Graph{
	Name:           "graph_16",
	TotalNodeCount: 8,
	TotalEdgeCount: 15,
	IsDAG:          false,
	EdgeToWeight: []EdgeToWeight{
		{[]string{"S", "A"}, 10.0},
		{[]string{"S", "B"}, 5.0},
		{[]string{"S", "C"}, 15.0},

		{[]string{"A", "B"}, 4.0},
		{[]string{"A", "D"}, 9.0},
		{[]string{"A", "E"}, 15.0},

		{[]string{"B", "C"}, 4.0},
		{[]string{"B", "E"}, 8.0},

		{[]string{"C", "F"}, 16.0},

		{[]string{"D", "E"}, 15.0},
		{[]string{"D", "T"}, 10.0},

		{[]string{"E", "T"}, 10.0},
		{[]string{"E", "F"}, 15.0},

		{[]string{"F", "T"}, 10.0},
		{[]string{"F", "B"}, 6.0},
	},
}
