package tsdfs

func main() {}

// Output differs every time it runs.
/*
import (
	"testing"

	"github.com/gyuho/goraph/graph/gs"
)

func TestTSDFS(t *testing.T) {
	g6 := gs.FromJSON("../../files/testgraph.json", "testgraph.006")
	g6s := TSDFS(g6)
	g6c := "E → D → C → B → A → F"
	if g6s != g6c {
		t.Errorf("Should be same but\n%v\n%v", g6s, g6c)
	}

	g7 := gs.FromJSON("../../files/testgraph.json", "testgraph.007")
	g7s := TSDFS(g7)
	g7c := "C → B → D → F → A → H → E → G"
	if g7s != g7c {
		t.Errorf("Should be same but\n%v\n%v", g7s, g7c)
	}
}
*/
