package kruskal

import (
	"fmt"
	"testing"

	"github.com/gyuho/goraph/graph/gsd"
)

func Test_JSON_MST(test *testing.T) {
	g14 := gsd.JSONGraph("../../../testgraph/testgraph.json", "testgraph.014")
	result := MST(g14)
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
}
