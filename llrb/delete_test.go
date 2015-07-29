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
RotateToRight: 20
after nd = RotateToRight(nd) 16
nd.Right, deleted = tr.delete(nd.Right, key) at 16
calling delete on 20 for the key 39
nd.Right, deleted = tr.delete(nd.Right, key) at 20
calling delete on 39 for the key 39
RotateToRight: 39
after nd = RotateToRight(nd) 25
nd.Right, deleted = tr.delete(nd.Right, key) at 25
calling delete on 39 for the key 39
!nd.Key.Less(key) && nd.Right == nil when 39
FixUp 25
FixUp 20
FixUp 16
RotateToLeft: 16
FixUp 13
Deleted 39
3
13
20

calling delete on 13 for the key 20
nd.Right, deleted = tr.delete(nd.Right, key) at 13
calling delete on 20 for the key 20
RotateToRight: 20
after nd = RotateToRight(nd) 16
nd.Right, deleted = tr.delete(nd.Right, key) at 16
calling delete on 20 for the key 20
nd.Right != nil && !isRed(nd.Right) && !isRed(nd.Right.Left) when 20
MoveRedFromLeftToRight: 20
FlipColor: 20
DeleteMin 25
FixUp 25
FixUp 16
FixUp 13
Deleted 20
3
13
16

calling delete on 13 for the key 16
nd.Right != nil && !isRed(nd.Right) && !isRed(nd.Right.Left) when 13
MoveRedFromLeftToRight: 13
FlipColor: 13
MoveRedFromLeftToRight isRed(nd.Left.Left): 13
RotateToRight: 13
FlipColor: 3
nd.Right, deleted = tr.delete(nd.Right, key) at 3
calling delete on 13 for the key 16
nd.Right, deleted = tr.delete(nd.Right, key) at 13
calling delete on 16 for the key 16
DeleteMin 25
DeleteMin 17
FixUp 25
FixUp 17
FixUp 13
RotateToLeft: 13
FixUp 3
Deleted 16
2
3
[[9(true)] 13(false) [15(true)]]

calling delete on 3 for the key 9
nd.Right, deleted = tr.delete(nd.Right, key) at 3
calling delete on 17 for the key 9
calling delete on 13 for the key 9
MoveRedFromRightToLeft: 13
FlipColor: 13
calling delete on 9 for the key 9
!nd.Key.Less(key) && nd.Right == nil when 9
FixUp 13
RotateToLeft: 13
FixUp 17
FixUp 3
Deleted 9
2
3
[[[13(false)] 15(true)] 17(true) [25(true)]]

calling delete on 3 for the key 25
nd.Right != nil && !isRed(nd.Right) && !isRed(nd.Right.Left) when 3
MoveRedFromLeftToRight: 3
FlipColor: 3
nd.Right, deleted = tr.delete(nd.Right, key) at 3
calling delete on 17 for the key 25
nd.Right != nil && !isRed(nd.Right) && !isRed(nd.Right.Left) when 17
MoveRedFromLeftToRight: 17
FlipColor: 17
MoveRedFromLeftToRight isRed(nd.Left.Left): 17
RotateToRight: 17
FlipColor: 15
nd.Right, deleted = tr.delete(nd.Right, key) at 15
calling delete on 17 for the key 25
nd.Right, deleted = tr.delete(nd.Right, key) at 17
calling delete on 25 for the key 25
!nd.Key.Less(key) && nd.Right == nil when 25
FixUp 17
FixUp 15
FixUp 3
RotateToLeft: 3
RotateToRight: 15
FlipColor: 3
Deleted 25
2
3
[[13(true)] 15(true) [17(true)]]

calling delete on 3 for the key 2
MoveRedFromRightToLeft: 3
FlipColor: 3
calling delete on 2 for the key 2
nd.Right != nil && !isRed(nd.Right) && !isRed(nd.Right.Left) when 2
MoveRedFromLeftToRight: 2
FlipColor: 2
DeleteMin 2.5
FixUp 2.5
FixUp 3
RotateToLeft: 3
Deleted 2
[[[1(false)] 2.5(true)] 3(false) [13(true)]]
15
[17(true)]

calling delete on 15 for the key 3
calling delete on 3 for the key 3
nd.Right != nil && !isRed(nd.Right) && !isRed(nd.Right.Left) when 3
MoveRedFromLeftToRight: 3
FlipColor: 3
MoveRedFromLeftToRight isRed(nd.Left.Left): 3
RotateToRight: 3
FlipColor: 2.5
nd.Right, deleted = tr.delete(nd.Right, key) at 2.5
calling delete on 3 for the key 3
DeleteMin 13
FixUp 13
FixUp 2.5
FixUp 15
Deleted 3
2.5
15
17

*/
