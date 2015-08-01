package llrb

import "testing"

func TestTreeDeleteMin(t *testing.T) {
	node20 := NewNode(Float64(20))
	node20.Black = true

	node17 := NewNode(Float64(17))
	node17.Black = false

	node25 := NewNode(Float64(25))
	node25.Black = false

	tr := New(node20)
	tr.Root.Left = node17
	tr.Root.Right = node25
	/*
	        20(B)
	      /      \
	   17(R)     25(R)
	*/

	if tr.DeleteMin() != Float64(17) {
		t.Fatalf("Expected 17 but got %+v", tr)
	}
}

func TestDeleteMin(t *testing.T) {
	node20 := NewNode(Float64(20))
	node20.Black = true

	node17 := NewNode(Float64(17))
	node17.Black = false

	node25 := NewNode(Float64(25))
	node25.Black = false

	tr := New(node20)
	tr.Root.Left = node17
	tr.Root.Right = node25
	/*
	        20(B)
	      /      \
	   17(R)     25(R)
	*/
	// Deleting the Minimum value of Right Sub-Tree
	var subDeleted Interface
	tr.Root.Right, subDeleted = DeleteMin(tr.Root.Right)
	if subDeleted == nil {
		panic("Unexpected nil value")
	}
	_, tr.Root.Key = tr.Root.Key, subDeleted
	/*
	        25(B)
	      /
	   17(R)
	*/

	if tr.Root.Left.Key != Float64(17) {
		t.Fatalf("Expected 17 but got %+v", tr)
	}
	if tr.Root.Key != Float64(25) {
		t.Fatalf("Expected 25 but got %+v", tr)
	}
	if tr.Root.Right != nil {
		t.Fatalf("Expected nil but got %+v", tr)
	}
}

func TestDelete(t *testing.T) {
	root := NewNode(Float64(1))
	tr := New(root)
	nums := []float64{3, 9, 13, 17, 20, 25, 39, 16, 15, 2, 2.5}
	for _, num := range nums {
		tr.Insert(NewNode(Float64(num)))
	}

	tr.Delete(Float64(39))
	if tr.Search(Float64(13)).Left.Key != Float64(3) {
		t.Fatal("13's Left child must be 3")
	}
	if tr.Search(Float64(3)).Left.Key != Float64(2) {
		t.Fatal("3's Left child must be 2")
	}
	if tr.Search(Float64(3)).Left.Black != false {
		t.Fatal("3's Left child must be red")
	}
	if tr.Search(Float64(20)).Right.Key != Float64(25) {
		t.Fatal("20's Right child must be 25")
	}
	if tr.SearchParent(Float64(25)).Key != Float64(20) {
		t.Fatal("25's Parent must be 20")
	}
	if tr.SearchParent(Float64(16)).Key != Float64(20) {
		t.Fatal("16's Parent must be 20")
	}
	if tr.Search(Float64(25)).Right != nil {
		t.Fatal("25's Right child must be nil")
	}
	if tr.Search(Float64(25)).Left != nil {
		t.Fatal("25's Left child must be nil")
	}
	if tr.Search(Float64(39)) != nil {
		t.Fatal("39 must be nil")
	}

	nt20 := tr.Delete(Float64(20))
	if nt20 != Float64(20) {
		t.Fatal("tr.Delete(Float64(20)) must return 20 but", nt20)
	}
	if tr.Search(Float64(16)).Left.Key != Float64(15) {
		t.Fatal("16's Left child must be 15")
	}
	if tr.Search(Float64(16)).Right.Key != Float64(25) {
		t.Fatal("16's Right child must be 25")
	}
	if tr.Search(Float64(25)).Left.Key != Float64(17) {
		t.Fatal("25's Left child must be 17")
	}
	if tr.Search(Float64(25)).Left.Black != false {
		t.Fatal("25's Left child must be red")
	}
	if tr.Search(Float64(17)).Black != false {
		t.Fatal("17 must be red")
	}
	if tr.SearchParent(Float64(25)).Key != Float64(16) {
		t.Fatal("25's Parent must be 16")
	}
	if tr.SearchParent(Float64(25)).Key != Float64(16) {
		t.Fatal("25's Parent must be 16")
	}
	if tr.SearchParent(Float64(17)).Key != Float64(25) {
		t.Fatal("17's Parent must be 25")
	}
	if tr.SearchParent(Float64(16)).Key != Float64(13) {
		t.Fatal("16's Parent must be 13")
	}

	nt16 := tr.Delete(Float64(16))
	if nt16 != Float64(16) {
		t.Fatal("tr.Delete(Float64(16)) must return 16 but", nt16)
	}
	if tr.SearchParent(Float64(13)).Key != Float64(17) {
		t.Fatal("13's Parent must be 17")
	}
	if tr.SearchParent(Float64(25)).Key != Float64(17) {
		t.Fatal("25's Parent must be 17")
	}
	if tr.Search(Float64(13)).Black != false {
		t.Fatal("13 must be red")
	}
	if tr.Search(Float64(17)).Left.Black != false {
		t.Fatal("17's Left child must be red")
	}
	if tr.Search(Float64(17)).Left.Key != Float64(13) {
		t.Fatal("17's Left child must be 13")
	}
	if tr.Search(Float64(17)).Right.Key != Float64(25) {
		t.Fatal("17's Right child must be 25")
	}
	if tr.Search(Float64(2)).Left.Key != Float64(1) {
		t.Fatal("2's Left child must be 1")
	}
	if tr.Search(Float64(2)).Right.Key != Float64(2.5) {
		t.Fatal("2's Right child must be 2.5")
	}
	if tr.Search(Float64(13)).Left.Key != Float64(9) {
		t.Fatal("13's Left child must be 9")
	}
	if tr.Search(Float64(13)).Right.Key != Float64(15) {
		t.Fatal("13's Right child must be 15")
	}

	nt9 := tr.Delete(Float64(9))
	if nt9 != Float64(9) {
		t.Fatal("tr.Delete(Float64(9)) must return 9 but", nt9)
	}
	if tr.Search(Float64(3)).Left.Key != Float64(2) {
		t.Fatal("3's Left child must be 2")
	}
	if tr.Search(Float64(3)).Right.Key != Float64(17) {
		t.Fatal("3's Right child must be 17")
	}
	if tr.Search(Float64(2)).Left.Key != Float64(1) {
		t.Fatal("2's Left child must be 1")
	}
	if tr.Search(Float64(2)).Right.Key != Float64(2.5) {
		t.Fatal("2's Right child must be 2.5")
	}
	if tr.Search(Float64(17)).Left.Key != Float64(15) {
		t.Fatal("17's Left child must be 15")
	}
	if tr.Search(Float64(17)).Right.Key != Float64(25) {
		t.Fatal("17's Right child must be 25")
	}
	if tr.Search(Float64(15)).Left.Key != Float64(13) {
		t.Fatal("15's Left child must be 13")
	}
	if tr.SearchParent(Float64(13)).Key != Float64(15) {
		t.Fatal("13's Parent must be 15")
	}
	if tr.SearchParent(Float64(15)).Key != Float64(17) {
		t.Fatal("15's Parent must be 17")
	}
	if tr.SearchParent(Float64(25)).Key != Float64(17) {
		t.Fatal("25's Parent must be 17")
	}
	if tr.SearchParent(Float64(2)).Key != Float64(3) {
		t.Fatal("2's Parent must be 3")
	}
	if tr.SearchParent(Float64(17)).Key != Float64(3) {
		t.Fatal("17's Parent must be 3")
	}
	if tr.Search(Float64(13)).Black != false {
		t.Fatal("13 must be red")
	}
	if tr.Search(Float64(15)).Black == false {
		t.Fatal("15 must be black")
	}

	nt25 := tr.Delete(Float64(25))
	if nt25 != Float64(25) {
		t.Fatal("tr.Delete(Float64(25)) must return 25 but", nt25)
	}
	if tr.Root.Key != Float64(3) {
		t.Fatal("Root must be 3")
	}
	if tr.Search(Float64(3)).Left.Key != Float64(2) {
		t.Fatal("3's Left child must be 2")
	}
	if tr.Search(Float64(3)).Right.Key != Float64(15) {
		t.Fatal("3's Right child must be 15")
	}
	if tr.Search(Float64(2)).Left.Key != Float64(1) {
		t.Fatal("2's Left child must be 1")
	}
	if tr.Search(Float64(2)).Right.Key != Float64(2.5) {
		t.Fatal("2's Right child must be 2.5")
	}
	if tr.Search(Float64(15)).Left.Key != Float64(13) {
		t.Fatal("15's Left child must be 13")
	}
	if tr.Search(Float64(15)).Right.Key != Float64(17) {
		t.Fatal("15's Right child must be 17")
	}
	if tr.SearchParent(Float64(2)).Key != Float64(3) {
		t.Fatal("2's Parent must be 3")
	}
	if tr.SearchParent(Float64(15)).Key != Float64(3) {
		t.Fatal("15's Parent must be 3")
	}
	if tr.SearchParent(Float64(1)).Key != Float64(2) {
		t.Fatal("1's Parent must be 2")
	}
	if tr.SearchParent(Float64(2.5)).Key != Float64(2) {
		t.Fatal("2.5's Parent must be 2")
	}
	if tr.SearchParent(Float64(13)).Key != Float64(15) {
		t.Fatal("13's Parent must be 15")
	}
	if tr.SearchParent(Float64(17)).Key != Float64(15) {
		t.Fatal("17's Parent must be 15")
	}
	if tr.Search(Float64(17)).Left != nil {
		t.Fatal("17's Left child must be nil")
	}
	if tr.Search(Float64(17)).Right != nil {
		t.Fatal("17's Right child must be nil")
	}
	if tr.Search(Float64(3)).Black != true {
		t.Fatal("3 must be black")
	}
	if tr.Search(Float64(2)).Black != true {
		t.Fatal("2 must be black")
	}
	if tr.Search(Float64(15)).Black != true {
		t.Fatal("15 must be black")
	}
	if tr.Search(Float64(1)).Black != true {
		t.Fatal("1 must be black")
	}
	if tr.Search(Float64(2.5)).Black != true {
		t.Fatal("2.5 must be black")
	}
	if tr.Search(Float64(13)).Black != true {
		t.Fatal("13 must be black")
	}
	if tr.Search(Float64(17)).Black != true {
		t.Fatal("17 must be black")
	}

	nt2 := tr.Delete(Float64(2))
	if nt2 != Float64(2) {
		t.Fatal("tr.Delete(Float64(2)) must return 2 but", nt2)
	}
	if tr.Root.Key != Float64(15) {
		t.Fatal("Root must be 15")
	}
	if tr.Search(Float64(15)).Left.Key != Float64(3) {
		t.Fatal("15's Left child must be 3")
	}
	if tr.Search(Float64(15)).Right.Key != Float64(17) {
		t.Fatal("15's Right child must be 17")
	}
	if tr.Search(Float64(3)).Left.Key != Float64(2.5) {
		t.Fatal("3's Left child must be 2.5")
	}
	if tr.Search(Float64(3)).Right.Key != Float64(13) {
		t.Fatal("3's Right child must be 13")
	}
	if tr.Search(Float64(2.5)).Left.Key != Float64(1) {
		t.Fatal("2.5's Left child must be 1")
	}
	if tr.SearchParent(Float64(2.5)).Key != Float64(3) {
		t.Fatal("2.5's Parent must be 3")
	}
	if tr.SearchParent(Float64(13)).Key != Float64(3) {
		t.Fatal("13's Parent must be 3")
	}
	if tr.SearchParent(Float64(1)).Key != Float64(2.5) {
		t.Fatal("1's Parent must be 2.5")
	}
	if tr.Search(Float64(3)).Black != false {
		t.Fatal("3 must be red")
	}
	if tr.Search(Float64(15)).Black != true {
		t.Fatal("15 must be black")
	}
	if tr.Search(Float64(1)).Black != false {
		t.Fatal("1 must be red")
	}
	if tr.Search(Float64(2.5)).Black != true {
		t.Fatal("2.5 must be black")
	}
	if tr.Search(Float64(13)).Black != true {
		t.Fatal("13 must be black")
	}
	if tr.Search(Float64(17)).Black != true {
		t.Fatal("17 must be black")
	}

	nt3 := tr.Delete(Float64(3))
	if nt3 != Float64(3) {
		t.Fatal("tr.Delete(Float64(3)) must return 3 but", nt3)
	}
	if tr.Root.Key != Float64(15) {
		t.Fatal("Root must be 15")
	}
	if tr.Search(Float64(15)).Left.Key != Float64(2.5) {
		t.Fatal("15's Left child must be 2.5")
	}
	if tr.Search(Float64(15)).Right.Key != Float64(17) {
		t.Fatal("15's Right child must be 17")
	}
	if tr.Search(Float64(2.5)).Left.Key != Float64(1) {
		t.Fatal("2.5's Left child must be 1")
	}
	if tr.SearchParent(Float64(2.5)).Key != Float64(15) {
		t.Fatal("2.5's Parent must be 15")
	}
	if tr.SearchParent(Float64(17)).Key != Float64(15) {
		t.Fatal("17's Parent must be 15")
	}
	if tr.SearchParent(Float64(13)).Key != Float64(2.5) {
		t.Fatal("13's Parent must be 2.5")
	}
	if tr.SearchParent(Float64(1)).Key != Float64(2.5) {
		t.Fatal("1's Parent must be 2.5")
	}
	if tr.SearchParent(Float64(17)).Key != Float64(15) {
		t.Fatal("17's Parent must be 15")
	}
	if tr.Search(Float64(15)).Black != true {
		t.Fatal("15 must be black")
	}
	if tr.Search(Float64(1)).Black != true {
		t.Fatal("1 must be black")
	}
	if tr.Search(Float64(2.5)).Black != false {
		t.Fatal("2.5 must be red")
	}
	if tr.Search(Float64(13)).Black != true {
		t.Fatal("13 must be black")
	}
	if tr.Search(Float64(17)).Black != true {
		t.Fatal("17 must be black")
	}

	//
	//
	//
	//
	// fmt.Println("Deleting start!")
	// fmt.Println("Deleted", tr.Delete(Float64(39)))
	// fmt.Println(tr.Root.Left.Key)
	// fmt.Println(tr.Root.Key)
	// fmt.Println(tr.Root.Right.Key)
	// fmt.Println()

	// fmt.Println("Deleted", tr.Delete(Float64(20)))
	// fmt.Println(tr.Root.Left.Key)
	// fmt.Println(tr.Root.Key)
	// fmt.Println(tr.Root.Right.Key)
	// fmt.Println()

	// fmt.Println("Deleted", tr.Delete(Float64(16)))
	// fmt.Println(tr.Root.Left.Key)
	// fmt.Println(tr.Root.Key)
	// fmt.Println(tr.Root.Right.Left)
	// fmt.Println()

	// fmt.Println("Deleted", tr.Delete(Float64(9)))
	// fmt.Println(tr.Root.Left.Key)
	// fmt.Println(tr.Root.Key)
	// fmt.Println(tr.Root.Right)
	// fmt.Println()

	// fmt.Println("Deleted", tr.Delete(Float64(25)))
	// fmt.Println(tr.Root.Left.Key)
	// fmt.Println(tr.Root.Key)
	// fmt.Println(tr.Root.Right)
	// fmt.Println()

	// fmt.Println("Deleted", tr.Delete(Float64(2)))
	// fmt.Println(tr.Root.Left)
	// fmt.Println(tr.Root.Key)
	// fmt.Println(tr.Root.Right)
	// fmt.Println()

	// fmt.Println("Deleted", tr.Delete(Float64(3)))
	// fmt.Println(tr.Root.Left.Key)
	// fmt.Println(tr.Root.Key)
	// fmt.Println(tr.Root.Right.Key)
	// fmt.Println()
}

/*
Deleting start!
calling delete on 13 for the key 39
nd.Right, deleted = tr.delete(nd.Right, key) at 13
calling delete on 20 for the key 39
RotateToRight 20
after nd = RotateToRight(nd) 16
nd.Right, deleted = tr.delete(nd.Right, key) at 16
calling delete on 20 for the key 39
nd.Right, deleted = tr.delete(nd.Right, key) at 20
calling delete on 39 for the key 39
RotateToRight 39
after nd = RotateToRight(nd) 25
nd.Right, deleted = tr.delete(nd.Right, key) at 25
calling delete on 39 for the key 39
!nd.Key.Less(key) && nd.Right == nil when 39
FixUp 25
FixUp 20
FixUp 16
RotateToLeft 16
after FixUp nd = RotateToLeft(nd) 20
FixUp 13
Deleted 39
3
13
20

calling delete on 13 for the key 20
nd.Right, deleted = tr.delete(nd.Right, key) at 13
calling delete on 20 for the key 20
RotateToRight 20
after nd = RotateToRight(nd) 16
nd.Right, deleted = tr.delete(nd.Right, key) at 16
calling delete on 20 for the key 20
nd.Right != nil && !isRed(nd.Right) && !isRed(nd.Right.Left) when [[17(true)] 20(false) [25(true)]]
MoveRedFromLeftToRight 20
FlipColor 20
DeleteMin 25
after deleted, nd.Key = nd.Key, subDeleted [[17(false)] 25(true)]
after deleted, nd.Key = nd.Key, subDeleted 25
FixUp 25
FixUp 16
FixUp 13
Deleted 20
3
13
16

calling delete on 13 for the key 16
nd.Right != nil && !isRed(nd.Right) && !isRed(nd.Right.Left) when [[[[1(true)] 2(false) [2.5(true)]] 3(true) [9(true)]] 13(true) [[15(true)] 16(true) [[17(false)] 25(true)]]]
MoveRedFromLeftToRight 13
FlipColor 13


RotateToRight 13
FlipColor 3
nd.Right, deleted = tr.delete(nd.Right, key) at 3
calling delete on 13 for the key 16
nd.Right, deleted = tr.delete(nd.Right, key) at 13
calling delete on 16 for the key 16
DeleteMin 25
DeleteMin 17
FixUp 25
after deleted, nd.Key = nd.Key, subDeleted [[15(true)] 17(false) [25(true)]]
after deleted, nd.Key = nd.Key, subDeleted 17
FixUp 17
FixUp 13
RotateToLeft 13
after FixUp nd = RotateToLeft(nd) 17
FixUp 3
Deleted 16
2
3
[[9(true)] 13(false) [15(true)]]

calling delete on 3 for the key 9
nd.Right, deleted = tr.delete(nd.Right, key) at 3
calling delete on 17 for the key 9
calling delete on 13 for the key 9
MoveRedFromRightToLeft 13
FlipColor 13
returning 13
calling delete on 9 for the key 9
!nd.Key.Less(key) && nd.Right == nil when 9
FixUp 13
RotateToLeft 13
after FixUp nd = RotateToLeft(nd) 15
FixUp 17
FixUp 3
Deleted 9
2
3
[[[13(false)] 15(true)] 17(true) [25(true)]]

calling delete on 3 for the key 25
nd.Right != nil && !isRed(nd.Right) && !isRed(nd.Right.Left) when [[[1(true)] 2(true) [2.5(true)]] 3(true) [[[13(false)] 15(true)] 17(true) [25(true)]]]
MoveRedFromLeftToRight 3
FlipColor 3
nd.Right, deleted = tr.delete(nd.Right, key) at 3
calling delete on 17 for the key 25
nd.Right != nil && !isRed(nd.Right) && !isRed(nd.Right.Left) when [[[13(false)] 15(true)] 17(false) [25(true)]]
MoveRedFromLeftToRight 17
FlipColor 17
RotateToRight 17
FlipColor 15
nd.Right, deleted = tr.delete(nd.Right, key) at 15
calling delete on 17 for the key 25
nd.Right, deleted = tr.delete(nd.Right, key) at 17
calling delete on 25 for the key 25
!nd.Key.Less(key) && nd.Right == nil when 25
FixUp 17
FixUp 15
FixUp 3
RotateToLeft 3
after FixUp nd = RotateToLeft(nd) 15
RotateToRight 15
FlipColor 3
Deleted 25
2
3
[[13(true)] 15(true) [17(true)]]

calling delete on 3 for the key 2
MoveRedFromRightToLeft 3
FlipColor 3
returning 3
calling delete on 2 for the key 2
nd.Right != nil && !isRed(nd.Right) && !isRed(nd.Right.Left) when [[1(true)] 2(false) [2.5(true)]]
MoveRedFromLeftToRight 2
FlipColor 2
DeleteMin 2.5
after deleted, nd.Key = nd.Key, subDeleted [[1(false)] 2.5(true)]
after deleted, nd.Key = nd.Key, subDeleted 2.5
FixUp 2.5
FixUp 3
RotateToLeft 3
after FixUp nd = RotateToLeft(nd) 15
Deleted 2
[[[1(false)] 2.5(true)] 3(false) [13(true)]]
15
[17(true)]

calling delete on 15 for the key 3
calling delete on 3 for the key 3
nd.Right != nil && !isRed(nd.Right) && !isRed(nd.Right.Left) when [[[1(false)] 2.5(true)] 3(false) [13(true)]]
MoveRedFromLeftToRight 3
FlipColor 3
RotateToRight 3
FlipColor 2.5
nd.Right, deleted = tr.delete(nd.Right, key) at 2.5
calling delete on 3 for the key 3
DeleteMin 13
after deleted, nd.Key = nd.Key, subDeleted [13(true)]
after deleted, nd.Key = nd.Key, subDeleted 13
FixUp 13
FixUp 2.5
FixUp 15
Deleted 3
2.5
15
17



*/
