package llrb

import (
	"fmt"
	"testing"
)

func TestDelete(t *testing.T) {
	root := NewNode(Float64(1))
	tr := New(root)
	nums := []float64{3, 9, 13, 17, 20, 25, 39, 16, 15, 2, 2.5}
	for _, num := range nums {
		tr.Insert(NewNode(Float64(num)))
	}

	fmt.Println("Deleted", tr.Delete(Float64(39)))
	fmt.Println(tr)
	fmt.Println()

	fmt.Println("Deleted", tr.Delete(Float64(20)))
	fmt.Println(tr.Root.Left.Key)
	fmt.Println(tr.Root.Key)
	fmt.Println(tr.Root.Right.Key)
	fmt.Println()

	fmt.Println("Deleted", tr.Delete(Float64(16)))
	fmt.Println(tr.Root.Left.Key)
	fmt.Println(tr.Root.Key)
	fmt.Println(tr.Root.Right.Left)
	fmt.Println()

	fmt.Println("Deleted", tr.Delete(Float64(9)))
	fmt.Println(tr.Root.Left.Key)
	fmt.Println(tr.Root.Key)
	fmt.Println(tr.Root.Right)
	fmt.Println()

	fmt.Println("Deleted", tr.Delete(Float64(25)))
	fmt.Println(tr.Root.Left.Key)
	fmt.Println(tr.Root.Key)
	fmt.Println(tr.Root.Right)
	fmt.Println()

	fmt.Println("Deleted", tr.Delete(Float64(2)))
	fmt.Println(tr.Root.Left)
	fmt.Println(tr.Root.Key)
	fmt.Println(tr.Root.Right)
	fmt.Println()

	fmt.Println("Deleted", tr.Delete(Float64(3)))
	fmt.Println(tr.Root.Left.Key)
	fmt.Println(tr.Root.Key)
	fmt.Println(tr.Root.Right.Key)
	fmt.Println()
}
