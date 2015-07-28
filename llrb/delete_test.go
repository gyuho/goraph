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

	fmt.Println("Deleting", tr.Delete(Float64(39)))
	fmt.Println(tr)
	fmt.Println()

	fmt.Println("Deleting", tr.Delete(Float64(20)))
	fmt.Println(tr)
	fmt.Println()
}
