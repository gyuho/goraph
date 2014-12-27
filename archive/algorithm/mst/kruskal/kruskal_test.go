package kruskal

import (
	"fmt"
	"testing"

	"github.com/gyuho/goraph/graph/gs"
)

func TestMST(t *testing.T) {
	g14 := gs.FromJSON("../../../files/testgraph.json", "testgraph.014")
	result, twgt := MST(g14)
	for _, edge := range result {
		fmt.Println(edge.Src.ID + " -- " + edge.Dst.ID)
	}
	/*
	   H -- G
	   G -- F
	   I -- C
	   A -- B
	   C -- F
	   D -- C
	   C -- B
	   D -- E
	*/
	if twgt != 37 {
		t.Errorf("expected 37 but %v", twgt)
	}
}
