package llrb_test

import (
	"fmt"

	"github.com/gyuho/goraph/llrb"
)

func Example_RotateToRight() {
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
	tr.Root = llrb.RotateToRight(tr.Root)

	fmt.Println(tr.Root.Left)
	fmt.Println(tr.Root)
	fmt.Println(tr.Root.Right)
	// Output:
	// RotateToRight: 20
	// [15(true)]
	// [[15(true)] 16(true) [[17(true)] 20(false) [[25(false)] 39(true)]]]
	// [[17(true)] 20(false) [[25(false)] 39(true)]]
}
