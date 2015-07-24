package tree_test

import (
	"fmt"

	"github.com/gyuho/goraph/tree"
)

func Example() {
	root := tree.NewNode(tree.Float64(1))
	data := tree.New(root)

	slice := []float64{3, 9, 13, 17, 20, 25, 39, 16, 15, 2, 2.5}
	for _, num := range slice {
		data.Insert(tree.NewNode(tree.Float64(num)))
	}

	deletes := []float64{13, 17, 3, 15, 1}
	for _, num := range deletes {
		fmt.Println("Deleting", num)
		data.Delete(data.Search(tree.Float64(num)))
		fmt.Println("After deleting", num, ":", data)
		fmt.Println()
	}

	// Output:
	// Deleting 13
	// After deleting 13 : [1 [[2 [2.5]] 3 [9 [[[15] 16] 17 [20 [25 [39]]]]]]]
	//
	// Deleting 17
	// After deleting 17 : [1 [[2 [2.5]] 3 [9 [[15] 16 [20 [25 [39]]]]]]]
	//
	// Deleting 3
	// After deleting 3 : [1 [[2] 2.5 [9 [[15] 16 [20 [25 [39]]]]]]]
	//
	// Deleting 15
	// After deleting 15 : [1 [[2] 2.5 [9 [[<nil>] 16 [20 [25 [39]]]]]]]
	//
	// Deleting 1
	// After deleting 1 : [[2] 2.5 [9 [[<nil>] 16 [20 [25 [39]]]]]]
}
