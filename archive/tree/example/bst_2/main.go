package main

import (
	"fmt"

	"github.com/gyuho/gotree/tree/bst"
	"github.com/gyuho/gotree/tree/bstviz"
)

func main() {
	t1 := bst.NewTree(5)

	// Perm returns, as a slice of n ints
	// , a pseudo-random permutation of the integers [0,n).
	// func (r *Rand) Perm(n int) []int
	for i := 0; i < 10; i++ {
		t1 = t1.Insert(int64(i))
	}

	fmt.Println(t1.Find(6))
	// &{<nil> 6 0x210323140 1}

	bstviz.Show(t1, "bstviz.dot")
}
