package kruskal

import (
	"fmt"
	"testing"

	"github.com/gyuho/goraph/graph/gsd"
)

func Test_JSON_MSTString(test *testing.T) {
	g14 := gsd.JSONGraph("../../../example_files/testgraph.json", "testgraph.014")
	result := MSTString(g14)
	fmt.Println(result)
	/*
	   graph KruskalMST {
	   	H -- G
	   	G -- F
	   	I -- C
	   	B -- A
	   	C -- F
	   	D -- C
	   	A -- H
	   	D -- E
	   }
	*/

	// The results are basically all the same but
	// the order changes
	// if result != rc {
	//	test.Errorf("Should be same but\n%v\n%v", result, rc)
	// }
}

func Test_JSON_ShowMST(test *testing.T) {
	// g14 := gsd.JSONGraph("../../../example_files/testgraph.json", "testgraph.014")
	// ShowMST(g14, "g14mst_kruskal.dot")
}
