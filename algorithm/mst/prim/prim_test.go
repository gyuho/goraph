package prim

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/gyuho/goraph/graph/gsd"
)

func Test_JSON_Contains(test *testing.T) {
	g14 := gsd.JSONGraph("../../../example_files/testgraph.json", "testgraph.014")
	vertices := g14.GetVertices()
	minHeap := make(VertexSlice, 0)
	for _, vtx := range *vertices {
		if vtx == nil {
			continue
		}
		heap.Push(&minHeap, vtx.(*gsd.Vertex))
	}
	if !minHeap.Contains(g14.FindVertexByID("A")) {
		test.Errorf("Should exists in the minHeap but %+v", minHeap.Contains(g14.FindVertexByID("A")))
	}
}

func Test_JSON_MST(test *testing.T) {
	g14 := gsd.JSONGraph("../../../example_files/testgraph.json", "testgraph.014")
	result, twgt := MST(g14)
	fmt.Println(result)
	/*
	   graph PrimMST {
	   	A -- B
	   	H -- A
	   	C -- F
	   	D -- C
	   	F -- G
	   	E -- D
	   	G -- H
	   	I -- C
	   }
	*/
	if twgt != 37 {
		test.Errorf("Total weights should be 37 but %+v", twgt)
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
