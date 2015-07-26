package bst_test

import (
	"fmt"

	"github.com/gyuho/goraph/bst"
)

// This example inserts float values to a binary search tree
// and deletes some of nodes in a tree.
func Example_delete() {
	root := bst.NewNode(bst.Float64(1))
	tr := bst.New(root)

	slice := []float64{3, 9, 13, 17, 20, 25, 39, 16, 15, 2, 2.5}
	for _, num := range slice {
		tr.Insert(bst.NewNode(bst.Float64(num)))
	}

	deletes := []float64{13, 17, 3, 15, 1, 2.5}
	for _, num := range deletes {
		fmt.Println("Deleting", num)
		tr.Delete(tr.Search(bst.Float64(num)))
		fmt.Println("Deleted", num, ":", tr)
		fmt.Println()
	}
	// Output:
	// Deleting 13
	// Deleted 13 : [1 [[2 [2.5]] 3 [9 [[[15] 16] 17 [20 [25 [39]]]]]]]
	//
	// Deleting 17
	// Deleted 17 : [1 [[2 [2.5]] 3 [9 [[15] 16 [20 [25 [39]]]]]]]
	//
	// Deleting 3
	// Deleted 3 : [1 [[2] 2.5 [9 [[15] 16 [20 [25 [39]]]]]]]
	//
	// Deleting 15
	// Deleted 15 : [1 [[2] 2.5 [9 [16 [20 [25 [39]]]]]]]
	//
	// Deleting 1
	// Deleted 1 : [[2] 2.5 [9 [16 [20 [25 [39]]]]]]
	//
	// Deleting 2.5
	// Deleted 2.5 : [2 [9 [16 [20 [25 [39]]]]]]
}
