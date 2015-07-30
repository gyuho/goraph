package llrb_test

import (
	"fmt"

	"github.com/gyuho/goraph/llrb"
)

func Example_Delete() {
	root := llrb.NewNode(llrb.Float64(1))
	tr := llrb.New(root)
	nums := []float64{3, 9, 13, 17, 20, 25, 39, 16, 15, 2, 2.5}
	for _, num := range nums {
		tr.Insert(llrb.NewNode(llrb.Float64(num)))
	}

	fmt.Println("Deleting start!")
	fmt.Println("Deleted", tr.Delete(llrb.Float64(39)))
	fmt.Println(tr.Root.Left.Key)
	fmt.Println(tr.Root.Key)
	fmt.Println(tr.Root.Right.Key)
	fmt.Println()

	fmt.Println("Deleted", tr.Delete(llrb.Float64(20)))
	fmt.Println(tr.Root.Left.Key)
	fmt.Println(tr.Root.Key)
	fmt.Println(tr.Root.Right.Key)
	fmt.Println()

	fmt.Println("Deleted", tr.Delete(llrb.Float64(16)))
	fmt.Println(tr.Root.Left.Key)
	fmt.Println(tr.Root.Key)
	fmt.Println(tr.Root.Right.Left)
	fmt.Println()

	fmt.Println("Deleted", tr.Delete(llrb.Float64(9)))
	fmt.Println(tr.Root.Left.Key)
	fmt.Println(tr.Root.Key)
	fmt.Println(tr.Root.Right)
	fmt.Println()

	fmt.Println("Deleted", tr.Delete(llrb.Float64(25)))
	fmt.Println(tr.Root.Left.Key)
	fmt.Println(tr.Root.Key)
	fmt.Println(tr.Root.Right)
	fmt.Println()

	fmt.Println("Deleted", tr.Delete(llrb.Float64(2)))
	fmt.Println(tr.Root.Left)
	fmt.Println(tr.Root.Key)
	fmt.Println(tr.Root.Right)
	fmt.Println()

	fmt.Println("Deleted", tr.Delete(llrb.Float64(3)))
	fmt.Println(tr.Root.Left.Key)
	fmt.Println(tr.Root.Key)
	fmt.Println(tr.Root.Right.Key)
	fmt.Println()
	// Output:
	// Deleting start!
	// Deleted 39
	// 3
	// 13
	// 20
	//
	// Deleted 20
	// 3
	// 13
	// 16
	//
	// Deleted 16
	// 2
	// 3
	// [[9(true)] 13(false) [15(true)]]
	//
	// Deleted 9
	// 2
	// 3
	// [[[13(false)] 15(true)] 17(true) [25(true)]]
	//
	// Deleted 25
	// 2
	// 3
	// [[13(true)] 15(true) [17(true)]]
	//
	// Deleted 2
	// [[[1(false)] 2.5(true)] 3(false) [13(true)]]
	// 15
	// [17(true)]
	//
	// Deleted 3
	// 2.5
	// 15
	// 17
}

func Example_DeleteMin() {
	node20 := llrb.NewNode(llrb.Float64(20))
	node20.Black = true

	node17 := llrb.NewNode(llrb.Float64(17))
	node17.Black = false

	node25 := llrb.NewNode(llrb.Float64(25))
	node25.Black = false

	tr := llrb.New(node20)
	tr.Root.Left = node17
	tr.Root.Right = node25
	/*
	        20(B)
	      /      \
	   17(R)     25(R)
	*/
	// Deleting the Minimum value of Right Sub-Tree
	var subDeleted llrb.Interface
	tr.Root.Right, subDeleted = llrb.DeleteMin(tr.Root.Right)
	if subDeleted == nil {
		panic("Unexpected nil value")
	}
	_, tr.Root.Key = tr.Root.Key, subDeleted
	/*
	        25(B)
	      /
	   17(R)
	*/

	fmt.Println("After tr.Root.Right, subDeleted = DeleteMin(tr.Root.Right), _, tr.Root.Key = tr.Root.Key, subDeleted")
	fmt.Println(tr.Root.Left)
	fmt.Println(tr.Root)
	fmt.Println(tr.Root.Right)
	// Output:
	// After tr.Root.Right, subDeleted = DeleteMin(tr.Root.Right), _, tr.Root.Key = tr.Root.Key, subDeleted
	// [17(false)]
	// [[17(false)] 25(true)]
	// []
}
