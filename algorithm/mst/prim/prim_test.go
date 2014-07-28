package prim

import (
	"container/heap"
	"testing"

	"github.com/gyuho/goraph/graph/gs"
)

func TestContains(t *testing.T) {
	g14 := gs.FromJSON("../../../files/testgraph.json", "testgraph.014")
	vertices := g14.GetVertices()
	minHeap := make(VertexSlice, 0)
	for _, vtx := range *vertices {
		if vtx == nil {
			continue
		}
		heap.Push(&minHeap, vtx.(*gs.Vertex))
	}
	if !minHeap.Contains(g14.FindVertexByID("A")) {
		t.Errorf("Should exists in the minHeap but %+v", minHeap.Contains(g14.FindVertexByID("A")))
	}
}

func TestMST(t *testing.T) {
	g14 := gs.FromJSON("../../../files/testgraph.json", "testgraph.014")
	twgt := MST(g14)
	if twgt != 37 {
		t.Errorf("Total weights should be 37 but %+v", twgt)
	}
	/*
		fmt.Println(g14.ShowPrev("A"))
		fmt.Println(g14.ShowPrev("B"))
		fmt.Println(g14.ShowPrev("C"))
		fmt.Println(g14.ShowPrev("D"))
		fmt.Println(g14.ShowPrev("E"))
		fmt.Println(g14.ShowPrev("F"))
		fmt.Println(g14.ShowPrev("G"))
		fmt.Println(g14.ShowPrev("H"))
		fmt.Println(g14.ShowPrev("I"))
	*/
}
