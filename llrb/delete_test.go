package llrb

import (
	"fmt"
	"testing"
)

func TestDeleteExample(t *testing.T) {
	root := NewNode(Float64(1))
	tr := New(root)
	nums := []float64{3, 9, 13, 17, 20, 25, 39, 16, 15, 2, 2.5}
	for _, num := range nums {
		tr.Insert(NewNode(Float64(num)))
	}

	fmt.Println()
	fmt.Println("Deleting start!")
	fmt.Println("Deleted", tr.Delete(Float64(39)))
	fmt.Println(tr.Root.Left.Key)
	fmt.Println(tr.Root.Key)
	fmt.Println(tr.Root.Right.Key)
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

/*
Deleting start!
calling delete on 13 for the key 39
nd.Right, deleted = tr.delete(nd.Right, key) at 13
calling delete on 20 for the key 39
rotateToRight: 20
nd.Right, deleted = tr.delete(nd.Right, key) at 16
calling delete on 20 for the key 39
nd.Right, deleted = tr.delete(nd.Right, key) at 20
calling delete on 39 for the key 39
rotateToRight: 39
nd.Right, deleted = tr.delete(nd.Right, key) at 25
calling delete on 39 for the key 39
!nd.Key.Less(key) && nd.Right == nil when 39
fixUp 25
fixUp 20
fixUp 16
rotateToLeft: 16
fixUp 13
Deleted 39
3
13
20

calling delete on 13 for the key 20
nd.Right, deleted = tr.delete(nd.Right, key) at 13
calling delete on 20 for the key 20
rotateToRight: 20
nd.Right, deleted = tr.delete(nd.Right, key) at 16
calling delete on 20 for the key 20
nd.Right != nil && !isRed(nd.Right) && !isRed(nd.Right.Left) when 20
moveRedToRight: 20
flipColor: 20
deleteMin 25
fixUp 25
fixUp 16
fixUp 13
Deleted 20
3
13
16

calling delete on 13 for the key 16
nd.Right != nil && !isRed(nd.Right) && !isRed(nd.Right.Left) when 13
moveRedToRight: 13
flipColor: 13
moveRedToRight isRed(nd.Left.Left): 13
rotateToRight: 13
flipColor: 3
nd.Right, deleted = tr.delete(nd.Right, key) at 3
calling delete on 13 for the key 16
nd.Right, deleted = tr.delete(nd.Right, key) at 13
calling delete on 16 for the key 16
deleteMin 25
deleteMin 17
fixUp 25
fixUp 17
fixUp 13
rotateToLeft: 13
fixUp 3
Deleted 16
2
3
[[9(true)] 13(false) [15(true)]]

calling delete on 3 for the key 9
nd.Right, deleted = tr.delete(nd.Right, key) at 3
calling delete on 17 for the key 9
calling delete on 13 for the key 9
moveRedToLeft: 13
flipColor: 13
calling delete on 9 for the key 9
!nd.Key.Less(key) && nd.Right == nil when 9
fixUp 13
rotateToLeft: 13
fixUp 17
fixUp 3
Deleted 9
2
3
[[[13(false)] 15(true)] 17(true) [25(true)]]

calling delete on 3 for the key 25
nd.Right != nil && !isRed(nd.Right) && !isRed(nd.Right.Left) when 3
moveRedToRight: 3
flipColor: 3
nd.Right, deleted = tr.delete(nd.Right, key) at 3
calling delete on 17 for the key 25
nd.Right != nil && !isRed(nd.Right) && !isRed(nd.Right.Left) when 17
moveRedToRight: 17
flipColor: 17
moveRedToRight isRed(nd.Left.Left): 17
rotateToRight: 17
flipColor: 15
nd.Right, deleted = tr.delete(nd.Right, key) at 15
calling delete on 17 for the key 25
nd.Right, deleted = tr.delete(nd.Right, key) at 17
calling delete on 25 for the key 25
!nd.Key.Less(key) && nd.Right == nil when 25
fixUp 17
fixUp 15
fixUp 3
rotateToLeft: 3
rotateToRight: 15
flipColor: 3
Deleted 25
2
3
[[13(true)] 15(true) [17(true)]]

calling delete on 3 for the key 2
moveRedToLeft: 3
flipColor: 3
calling delete on 2 for the key 2
nd.Right != nil && !isRed(nd.Right) && !isRed(nd.Right.Left) when 2
moveRedToRight: 2
flipColor: 2
deleteMin 2.5
fixUp 2.5
fixUp 3
rotateToLeft: 3
Deleted 2
[[[1(false)] 2.5(true)] 3(false) [13(true)]]
15
[17(true)]

calling delete on 15 for the key 3
calling delete on 3 for the key 3
nd.Right != nil && !isRed(nd.Right) && !isRed(nd.Right.Left) when 3
moveRedToRight: 3
flipColor: 3
moveRedToRight isRed(nd.Left.Left): 3
rotateToRight: 3
flipColor: 2.5
nd.Right, deleted = tr.delete(nd.Right, key) at 2.5
calling delete on 3 for the key 3
deleteMin 13
fixUp 13
fixUp 2.5
fixUp 15
Deleted 3
2.5
15
17

*/
