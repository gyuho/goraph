package main

import (
	"github.com/gyuho/gotree/tree/avl"
	"github.com/gyuho/gotree/tree/avlviz"
)

func main() {
	tr := avl.NewTree(4)
	tr.BalanceInsert(6)
	tr.BalanceInsert(5)
	avlviz.Show(tr, "avl-before.dot")

	tr.BalanceRL(5)
	avlviz.Show(tr, "avl-after.dot")
}
