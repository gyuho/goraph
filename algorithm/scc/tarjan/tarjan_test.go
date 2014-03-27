package tarjan

import (
	"fmt"
	"testing"

	"github.com/gyuho/goraph/graph/gsd"
)

func Test_JSON_SCC(test *testing.T) {
	g15 := gsd.JSONGraph("../../../testgraph/testgraph.json", "testgraph.015")
	fmt.Println(SCC(g15))
}
