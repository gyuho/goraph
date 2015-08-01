package llrb_test

import (
	"fmt"

	"github.com/gyuho/goraph/llrb"
)

func Example_rotateToLeft() {
	node3 := llrb.NewNode(llrb.Float64(3))
	node3.Black = true

	node1 := llrb.NewNode(llrb.Float64(1))
	node1.Black = true

	node13 := llrb.NewNode(llrb.Float64(13))
	node13.Black = false

	node9 := llrb.NewNode(llrb.Float64(9))
	node9.Black = true

	node17 := llrb.NewNode(llrb.Float64(17))
	node17.Black = true

	tr := llrb.New(node3)
	tr.Root.Right = node13
	tr.Root.Right.Left = node9
	tr.Root.Right.Right = node17
	tr.Root.Left = node1
	/*
	        3(B)
	      /      \
	   1(B)      13(R)
	            /   \
	         9(B)  17(B)
	*/
	fmt.Println("Before tr.Root = llrb.RotateToLeft(tr.Root)")
	fmt.Println(tr.Root.Left)
	fmt.Println(tr.Root)
	fmt.Println(tr.Root.Right)

	tr.Root = llrb.RotateToLeft(tr.Root)
	/*
			   	   13(B)
			   	  /     \
			   3(R)     17(B)
			  /   \
		   1(B)   9(B)
	*/

	fmt.Println("After tr.Root = llrb.RotateToLeft(tr.Root)")
	fmt.Println(tr.Root.Left)
	fmt.Println(tr.Root)
	fmt.Println(tr.Root.Right)
	// Output:
	// Before tr.Root = llrb.RotateToLeft(tr.Root)
	// [1(true)]
	// [[1(true)] 3(true) [[9(true)] 13(false) [17(true)]]]
	// [[9(true)] 13(false) [17(true)]]
	// After tr.Root = llrb.RotateToLeft(tr.Root)
	// [[1(true)] 3(false) [9(true)]]
	// [[[1(true)] 3(false) [9(true)]] 13(true) [17(true)]]
	// [17(true)]
}

func Example_rotateToRight() {
	node20 := llrb.NewNode(llrb.Float64(20))
	node20.Black = true

	node39 := llrb.NewNode(llrb.Float64(39))
	node39.Black = true

	node25 := llrb.NewNode(llrb.Float64(25))
	node25.Black = false

	node16 := llrb.NewNode(llrb.Float64(16))
	node16.Black = false

	node15 := llrb.NewNode(llrb.Float64(15))
	node15.Black = true

	node17 := llrb.NewNode(llrb.Float64(17))
	node17.Black = true

	tr := llrb.New(node20)
	tr.Root.Right = node39
	tr.Root.Right.Left = node25
	tr.Root.Left = node16
	tr.Root.Left.Left = node15
	tr.Root.Left.Right = node17
	/*
	             20(B)
	            /     \
	       16(R)     39(B)
	       /   \       /
	   15(B)  17(B)  25(R)
	*/
	fmt.Println("Before tr.Root = llrb.RotateToRight(tr.Root)")
	fmt.Println(tr.Root.Left)
	fmt.Println(tr.Root)
	fmt.Println(tr.Root.Right)

	tr.Root = llrb.RotateToRight(tr.Root)
	/*
	       16(B)
	      /     \
	   15(B)     20(R)
	            /    \
	        17(B)     39(B)
	                  /
	                25(R)
	*/

	fmt.Println("After tr.Root = llrb.RotateToRight(tr.Root)")
	fmt.Println(tr.Root.Left)
	fmt.Println(tr.Root)
	fmt.Println(tr.Root.Right)
	// Output:
	// Before tr.Root = llrb.RotateToRight(tr.Root)
	// [[15(true)] 16(false) [17(true)]]
	// [[[15(true)] 16(false) [17(true)]] 20(true) [[25(false)] 39(true)]]
	// [[25(false)] 39(true)]
	// After tr.Root = llrb.RotateToRight(tr.Root)
	// [15(true)]
	// [[15(true)] 16(true) [[17(true)] 20(false) [[25(false)] 39(true)]]]
	// [[17(true)] 20(false) [[25(false)] 39(true)]]
}

func Example_flipColor() {
	node3 := llrb.NewNode(llrb.Float64(3))
	node3.Black = true

	node1 := llrb.NewNode(llrb.Float64(1))
	node1.Black = true

	node13 := llrb.NewNode(llrb.Float64(13))
	node13.Black = true

	node9 := llrb.NewNode(llrb.Float64(9))
	node9.Black = true

	node17 := llrb.NewNode(llrb.Float64(17))
	node17.Black = true

	tr := llrb.New(node3)
	tr.Root.Right = node13
	tr.Root.Right.Left = node9
	tr.Root.Right.Right = node17
	tr.Root.Left = node1
	/*
	        3(B)
	      /      \
	   1(B)      13(B)
	            /   \
	         9(B)  17(B)
	*/
	fmt.Println("Before llrb.FlipColor(tr.Root.Right)")
	fmt.Println(tr.Root.Left)
	fmt.Println(tr.Root)
	fmt.Println(tr.Root.Right)

	llrb.FlipColor(tr.Root.Right)
	/*
	        3(B)
	      /      \
	   1(B)      13(R)
	            /   \
	         9(R)  17(R)
	*/

	fmt.Println("After llrb.FlipColor(tr.Root.Right)")
	fmt.Println(tr.Root.Left)
	fmt.Println(tr.Root)
	fmt.Println(tr.Root.Right)
	// Output:
	// Before llrb.FlipColor(tr.Root.Right)
	// [1(true)]
	// [[1(true)] 3(true) [[9(true)] 13(true) [17(true)]]]
	// [[9(true)] 13(true) [17(true)]]
	// After llrb.FlipColor(tr.Root.Right)
	// [1(true)]
	// [[1(true)] 3(true) [[9(false)] 13(false) [17(false)]]]
	// [[9(false)] 13(false) [17(false)]]
}

func Example_moveRedFromRightToLeft() {
	node3 := llrb.NewNode(llrb.Float64(3))
	node3.Black = true

	node2 := llrb.NewNode(llrb.Float64(2))
	node2.Black = true

	node15 := llrb.NewNode(llrb.Float64(15))
	node15.Black = true

	node1 := llrb.NewNode(llrb.Float64(1))
	node1.Black = true

	node2point5 := llrb.NewNode(llrb.Float64(2.5))
	node2point5.Black = true

	node13 := llrb.NewNode(llrb.Float64(13))
	node13.Black = false

	node17 := llrb.NewNode(llrb.Float64(17))
	node17.Black = true

	tr := llrb.New(node3)
	tr.Root.Right = node15
	tr.Root.Right.Left = node13
	tr.Root.Right.Right = node17
	tr.Root.Left = node2
	tr.Root.Left.Left = node1
	tr.Root.Left.Right = node2point5
	/*
	              3(B)
	            /      \
	        2(B)       15(B)
	       /   \       /   \
	   1(B)  2.5(B)  13(R) 17(B)
	*/
	fmt.Println("Before tr.Root = llrb.MoveRedFromRightToLeft(tr.Root)")
	fmt.Println(tr.Root.Left)
	fmt.Println(tr.Root)
	fmt.Println(tr.Root.Right)

	tr.Root = llrb.MoveRedFromRightToLeft(tr.Root)
	/*
	           13(B)
	          /     \
	       3(B)      15(B)
	       /          \
	     2(R)          17(B)
	     /   \
	   1(B)  2.5(B)
	*/

	fmt.Println("After tr.Root = llrb.MoveRedFromRightToLeft(tr.Root)")
	fmt.Println(tr.Root.Left)
	fmt.Println(tr.Root)
	fmt.Println(tr.Root.Right)
	// Output:
	// Before tr.Root = llrb.MoveRedFromRightToLeft(tr.Root)
	// [[1(true)] 2(true) [2.5(true)]]
	// [[[1(true)] 2(true) [2.5(true)]] 3(true) [[13(false)] 15(true) [17(true)]]]
	// [[13(false)] 15(true) [17(true)]]
	// After tr.Root = llrb.MoveRedFromRightToLeft(tr.Root)
	// [[[1(true)] 2(false) [2.5(true)]] 3(true)]
	// [[[[1(true)] 2(false) [2.5(true)]] 3(true)] 13(true) [15(true) [17(true)]]]
	// [15(true) [17(true)]]
}

func Example_moveRedFromLeftToRight() {
	node13 := llrb.NewNode(llrb.Float64(13))
	node13.Black = true

	node3 := llrb.NewNode(llrb.Float64(3))
	node3.Black = true

	node16 := llrb.NewNode(llrb.Float64(16))
	node16.Black = true

	node2 := llrb.NewNode(llrb.Float64(2))
	node2.Black = false

	node9 := llrb.NewNode(llrb.Float64(9))
	node9.Black = true

	node15 := llrb.NewNode(llrb.Float64(15))
	node15.Black = true

	node25 := llrb.NewNode(llrb.Float64(25))
	node25.Black = true

	node1 := llrb.NewNode(llrb.Float64(1))
	node1.Black = true

	node2point5 := llrb.NewNode(llrb.Float64(2.5))
	node2point5.Black = true

	node17 := llrb.NewNode(llrb.Float64(17))
	node17.Black = false

	tr := llrb.New(node13)
	tr.Root.Right = node16
	tr.Root.Right.Left = node15
	tr.Root.Right.Right = node25
	tr.Root.Right.Right.Left = node17
	tr.Root.Left = node3
	tr.Root.Left.Left = node2
	tr.Root.Left.Right = node9
	tr.Root.Left.Left.Left = node1
	tr.Root.Left.Left.Right = node2point5
	/*
		             13(B)
		            /      \
		        3(B)       16(B)
		       /   \       /   \
		   2(R)    9(B) 15(B)  25(B)
		   /  \                /
		1(B) 2.5(B)         17(R)
	*/
	fmt.Println("Before tr.Root = llrb.MoveRedFromLeftToRight(tr.Root)")
	fmt.Println(tr.Root.Left)
	fmt.Println(tr.Root)
	fmt.Println(tr.Root.Right)

	tr.Root = llrb.MoveRedFromLeftToRight(tr.Root)
	/*
	              3(B)
	            /      \
	        2(B)       13(B)
	       /   \       /    \
	   1(B)   2.5(B) 9(B)  16(R)
	                        /  \
	                    15(B)   25(B)
	                            /
	                          17(R)
	*/

	fmt.Println("After tr.Root = llrb.MoveRedFromRightToLeft(tr.Root)")
	fmt.Println(tr.Root.Left)
	fmt.Println(tr.Root)
	fmt.Println(tr.Root.Right)
	// Output:
	// Before tr.Root = llrb.MoveRedFromLeftToRight(tr.Root)
	// [[[1(true)] 2(false) [2.5(true)]] 3(true) [9(true)]]
	// [[[[1(true)] 2(false) [2.5(true)]] 3(true) [9(true)]] 13(true) [[15(true)] 16(true) [[17(false)] 25(true)]]]
	// [[15(true)] 16(true) [[17(false)] 25(true)]]
	// After tr.Root = llrb.MoveRedFromRightToLeft(tr.Root)
	// [[1(true)] 2(true) [2.5(true)]]
	// [[[1(true)] 2(true) [2.5(true)]] 3(true) [[9(true)] 13(true) [[15(true)] 16(false) [[17(false)] 25(true)]]]]
	// [[9(true)] 13(true) [[15(true)] 16(false) [[17(false)] 25(true)]]]
}

func Example_balance() {
	node13 := llrb.NewNode(llrb.Float64(13))
	node13.Black = true

	node3 := llrb.NewNode(llrb.Float64(3))
	node3.Black = false

	node2 := llrb.NewNode(llrb.Float64(2))
	node2.Black = false

	tr1 := llrb.New(node13)
	tr1.Root.Left = node3
	tr1.Root.Left.Left = node2
	/*
	        13(B)
	       /
	      3(R)
	     /
	   2(R)
	*/
	tr1.Root = llrb.Balance(tr1.Root)
	/*
	      3(R)
	     /   \
	   2(B)   13(B)
	*/
	fmt.Println("After tr1.Root = llrb.Balance(tr1.Root):", tr1)
	// Output:
	// After tr1.Root = llrb.Balance(tr1.Root): [[2(true)] 3(false) [13(true)]]

	// Don't do this:
	//
	// tr2 := llrb.New(node13)
	// tr2.Root.Left = node3
	// tr2.Root.Left.Left = node2
	// llrb.Balance(tr2.Root) // (X) this will cut off the tree
	// fmt.Println(tr2)       // (X) this will [stack growth] error
}
