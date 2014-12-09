package main

import (
	"fmt"

	"github.com/gyuho/gotree/tree/avl"
	"github.com/gyuho/gotree/tree/avlviz"
)

func main() {
	tr := avl.NewTree(10)
	// tr.BalanceInserts(13, 17, 5, 4)
	tr.BalanceInsert(13)
	tr.BalanceInsert(17)
	fmt.Println("10's right", tr.Right.Value)
	tr.BalanceRR(17)
	fmt.Println("10's left", tr.Left.Value)
	tr.BalanceInsert(5)
	tr.BalanceInsert(4)
	avlviz.Show(tr, "avl09.dot")

	fmt.Println("5's parent", tr.Parent(5).Value)
	fmt.Println("4's parent", tr.Parent(4).Value)
	fmt.Println(10, tr.IsBalanced(10), tr.GetSize(10))
	fmt.Println(13, tr.IsBalanced(13), tr.GetSize(13))
	fmt.Println(17, tr.IsBalanced(17), tr.Find(17))
	fmt.Println(5, tr.IsBalanced(5), tr.GetSize(5))
	fmt.Println(4, tr.IsBalanced(4), tr.GetSize(4))
}
