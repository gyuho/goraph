package viz

import "testing"

func Test_Show(test *testing.T) {
	Show("testgraph.json", "testgraph.010", "testgraph.010.dot")
}
