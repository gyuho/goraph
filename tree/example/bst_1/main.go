package main

import (
	"fmt"
	"github.com/gyuho/gotree/tree/bst"
)

func main() {
	t1 := bst.NewTree(5)

	// Perm returns, as a slice of n ints
	// , a pseudo-random permutation of the integers [0,n).
	// func (r *Rand) Perm(n int) []int
	for i := 0; i < 10; i++ {
		t1 = t1.Insert(int64(i))
	}

	fmt.Println(t1.TreePrint())
	// ((0 (1 (2 (3 (4))))) 5 (6 (7 (8 (9)))))
	// 0 is nil

	ch1 := make(chan int64)
	fmt.Println(bst.StringInOrder(t1, ch1))
	// 0 1 2 3 4 5 6 7 8 9

	println()
	ch2 := make(chan int64)
	fmt.Println(bst.StringPreOrder(t1, ch2))
	// 5 0 1 2 3 4 6 7 8 9

	println()
	ch3 := make(chan int64)
	fmt.Println(bst.StringPostOrder(t1, ch3))
	// 4 3 2 1 0 9 8 7 6 5

	println()
	fmt.Println(bst.StringLevelOrder(t1))
	// 5 0 6 1 7 2 8 3 9 4
}
