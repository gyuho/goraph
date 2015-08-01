package llrb

import "testing"

func TestRotateToLeft(t *testing.T) {
	node3 := NewNode(Float64(3))
	node3.Black = true

	node1 := NewNode(Float64(1))
	node1.Black = true

	node13 := NewNode(Float64(13))
	node13.Black = false

	node9 := NewNode(Float64(9))
	node9.Black = true

	node17 := NewNode(Float64(17))
	node17.Black = true

	tr := New(node3)
	tr.Root.Right = node13
	tr.Root.Right.Left = node9
	tr.Root.Right.Right = node17
	tr.Root.Left = node1
	/*
	        3(B)
	      /      \
	   1(B)      13(R)
	            /   \
	         9(B)  17(B)
	*/
	tr.Root = RotateToLeft(tr.Root)
	/*
			   	   13(B)
			   	  /     \
			   3(R)     17(B)
			  /   \
		   1(B)   9(B)
	*/

	if tr.Root.Key != Float64(13) {
		t.Fatalf("Root should be 13 but got %v", tr.Root.Key)
	}
	if tr.Root.Left.Key != Float64(3) {
		t.Fatalf("tr.Root.Left should be 3 but got %v", tr.Root.Left.Key)
	}
	if tr.Root.Left.Black {
		t.Fatalf("3 must be red but %v", tr.Root.Left)
	}
	if tr.Search(Float64(3)).Black {
		t.Fatalf("3 must be red but %v", tr.Root.Left)
	}
	if tr.Search(Float64(3)).Left.Key != Float64(1) {
		t.Fatalf("3's Left must be 1 but %v", tr.Search(Float64(3)).Left)
	}
	if tr.Search(Float64(3)).Right.Key != Float64(9) {
		t.Fatalf("3's Right must be 9 but %v", tr.Search(Float64(3)).Right)
	}
	if tr.Search(Float64(13)).Right.Key != Float64(17) {
		t.Fatalf("13's Right must be 17 but %v", tr.Search(Float64(13)).Right)
	}
}

func TestRotateToRight(t *testing.T) {
	node20 := NewNode(Float64(20))
	node20.Black = true

	node39 := NewNode(Float64(39))
	node39.Black = true

	node25 := NewNode(Float64(25))
	node25.Black = false

	node16 := NewNode(Float64(16))
	node16.Black = false

	node15 := NewNode(Float64(15))
	node15.Black = true

	node17 := NewNode(Float64(17))
	node17.Black = true

	tr := New(node20)
	tr.Root.Right = node39
	tr.Root.Right.Left = node25
	tr.Root.Left = node16
	tr.Root.Left.Left = node15
	tr.Root.Left.Right = node17
	/*
	             20(B)
	            /     \
	       16(R)     39(B)
	       /   \       /
	   15(B)  17(B)  25(R)
	*/
	tr.Root = RotateToRight(tr.Root)
	/*
	       16(B)
	      /     \
	   15(B)     20(R)
	            /    \
	        17(B)     39(B)
	                  /
	                25(R)
	*/

	if tr.Root.Key != Float64(16) {
		t.Fatalf("Root should be 16 but got %v", tr.Root.Key)
	}
	if tr.Root.Left.Key != Float64(15) {
		t.Fatalf("tr.Root.Left should be 15 but got %v", tr.Root.Left.Key)
	}
	if !tr.Root.Left.Black {
		t.Fatalf("15 must be black but %v", tr.Root.Left)
	}
	if !tr.Search(Float64(16)).Black {
		t.Fatalf("16 must be black but %v", tr.Search(Float64(16)))
	}
	if !tr.Root.Black {
		t.Fatalf("Root must be black but %v", tr.Root)
	}
	if tr.Search(Float64(20)).Black {
		t.Fatalf("20 must be red but %v", tr.Search(Float64(20)))
	}
	if tr.Search(Float64(25)).Black {
		t.Fatalf("25 must be red but %v", tr.Search(Float64(25)))
	}
	if tr.Search(Float64(20)).Left.Key != Float64(17) {
		t.Fatalf("20's Left must be 17 but %v", tr.Search(Float64(20)).Left)
	}
	if tr.Search(Float64(20)).Right.Key != Float64(39) {
		t.Fatalf("20's Right must be 39 but %v", tr.Search(Float64(20)).Right)
	}
	if tr.Search(Float64(39)).Left.Key != Float64(25) {
		t.Fatalf("39's Left must be 25 but %v", tr.Search(Float64(39)).Left)
	}
	if tr.Search(Float64(16)).Right.Key != Float64(20) {
		t.Fatalf("16's Right must be 20 but %v", tr.Search(Float64(16)).Right)
	}
}

func TestFlipColor(t *testing.T) {
	node3 := NewNode(Float64(3))
	node3.Black = true

	node1 := NewNode(Float64(1))
	node1.Black = true

	node13 := NewNode(Float64(13))
	node13.Black = true

	node9 := NewNode(Float64(9))
	node9.Black = true

	node17 := NewNode(Float64(17))
	node17.Black = true

	tr := New(node3)
	tr.Root.Right = node13
	tr.Root.Right.Left = node9
	tr.Root.Right.Right = node17
	tr.Root.Left = node1
	/*
	        3(B)
	      /      \
	   1(B)      13(B)
	            /   \
	         9(B)  17(B)
	*/
	FlipColor(tr.Root.Right)
	/*
	        3(B)
	      /      \
	   1(B)      13(R)
	            /   \
	         9(R)  17(R)
	*/
	if tr.Search(Float64(13)).Left.Key != Float64(9) {
		t.Fatalf("13's Left must be 9 but %v", tr.Search(Float64(9)).Left)
	}
	if tr.Search(Float64(13)).Right.Key != Float64(17) {
		t.Fatalf("13's Right must be 17 but %v", tr.Search(Float64(17)).Right)
	}
	if tr.Search(Float64(13)).Black {
		t.Fatalf("13's Left must be red but %v", tr.Search(Float64(13)))
	}
	if tr.Search(Float64(9)).Black {
		t.Fatalf("9's Left must be red but %v", tr.Search(Float64(9)))
	}
	if tr.Search(Float64(17)).Black {
		t.Fatalf("17's Left must be red but %v", tr.Search(Float64(17)))
	}
}

func TestMoveRedFromRightToLeft(t *testing.T) {
	node3 := NewNode(Float64(3))
	node3.Black = true

	node2 := NewNode(Float64(2))
	node2.Black = true

	node15 := NewNode(Float64(15))
	node15.Black = true

	node1 := NewNode(Float64(1))
	node1.Black = true

	node2point5 := NewNode(Float64(2.5))
	node2point5.Black = true

	node13 := NewNode(Float64(13))
	node13.Black = false

	node17 := NewNode(Float64(17))
	node17.Black = true

	tr := New(node3)
	tr.Root.Right = node15
	tr.Root.Right.Left = node13
	tr.Root.Right.Right = node17
	tr.Root.Left = node2
	tr.Root.Left.Left = node1
	tr.Root.Left.Right = node2point5
	/*
	              3(B)
	            /      \
	        2(B)       15(B)
	       /   \       /   \
	   1(B)  2.5(B)  13(R) 17(B)
	*/
	tr.Root = MoveRedFromRightToLeft(tr.Root)
	/*
	           13(B)
	          /     \
	       3(B)      15(B)
	       /          \
	     2(R)          17(B)
	     /   \
	   1(B)  2.5(B)
	*/
	if tr.Search(Float64(13)).Left.Key != Float64(3) {
		t.Fatalf("13's Left must be 3 but %v", tr.Search(Float64(3)).Left)
	}
	if tr.Search(Float64(13)).Right.Key != Float64(15) {
		t.Fatalf("13's Right must be 15 but %v", tr.Search(Float64(15)).Right)
	}
	if !tr.Search(Float64(13)).Black {
		t.Fatalf("13's Left must be black but %v", tr.Search(Float64(13)))
	}
	if tr.Search(Float64(3)).Left.Key != Float64(2) {
		t.Fatalf("3's Left must be 2 but %v", tr.Search(Float64(3)).Left)
	}
	if tr.Search(Float64(2)).Left.Key != Float64(1) {
		t.Fatalf("2's Left must be 1 but %v", tr.Search(Float64(2)).Left)
	}
	if tr.Search(Float64(2)).Right.Key != Float64(2.5) {
		t.Fatalf("2's Right must be 2.5 but %v", tr.Search(Float64(2)).Right)
	}
	if tr.Search(Float64(15)).Right.Key != Float64(17) {
		t.Fatalf("15's Right must be 17 but %v", tr.Search(Float64(15)).Right)
	}
	if !tr.Search(Float64(3)).Black {
		t.Fatalf("3's Left must be black but %v", tr.Search(Float64(3)))
	}
	if tr.Search(Float64(2)).Black {
		t.Fatalf("2's Left must be red but %v", tr.Search(Float64(2)))
	}
	if !tr.Search(Float64(2.5)).Black {
		t.Fatalf("2.5's Left must be black but %v", tr.Search(Float64(2.5)))
	}
}

func TestMoveRedFromLeftToRight(t *testing.T) {
	node13 := NewNode(Float64(13))
	node13.Black = true

	node3 := NewNode(Float64(3))
	node3.Black = true

	node16 := NewNode(Float64(16))
	node16.Black = true

	node2 := NewNode(Float64(2))
	node2.Black = false

	node9 := NewNode(Float64(9))
	node9.Black = true

	node15 := NewNode(Float64(15))
	node15.Black = true

	node25 := NewNode(Float64(25))
	node25.Black = true

	node1 := NewNode(Float64(1))
	node1.Black = true

	node2point5 := NewNode(Float64(2.5))
	node2point5.Black = true

	node17 := NewNode(Float64(17))
	node17.Black = false

	tr := New(node13)
	tr.Root.Right = node16
	tr.Root.Right.Left = node15
	tr.Root.Right.Right = node25
	tr.Root.Right.Right.Left = node17
	tr.Root.Left = node3
	tr.Root.Left.Left = node2
	tr.Root.Left.Right = node9
	tr.Root.Left.Left.Left = node1
	tr.Root.Left.Left.Right = node2point5
	/*
		             13(B)
		            /      \
		        3(B)       16(B)
		       /   \       /   \
		   2(R)    9(B) 15(B)  25(B)
		   /  \                /
		1(B) 2.5(B)         17(R)
	*/
	tr.Root = MoveRedFromLeftToRight(tr.Root)
	/*
	              3(B)
	            /      \
	        2(B)       13(B)
	       /   \       /    \
	   1(B)   2.5(B) 9(B)  16(R)
	                        /  \
	                    15(B)   25(B)
	                            /
	                          17(R)
	*/

	if tr.Search(Float64(3)).Left.Key != Float64(2) {
		t.Fatalf("3's Left must be 2 but %v", tr.Search(Float64(3)).Left)
	}
	if tr.Search(Float64(3)).Right.Key != Float64(13) {
		t.Fatalf("3's Right must be 13 but %v", tr.Search(Float64(3)).Right)
	}
	if tr.Search(Float64(13)).Left.Key != Float64(9) {
		t.Fatalf("13's Left must be 9 but %v", tr.Search(Float64(3)).Left)
	}
	if tr.Search(Float64(13)).Right.Key != Float64(16) {
		t.Fatalf("13's Right must be 16 but %v", tr.Search(Float64(13)).Right)
	}
	if !tr.Search(Float64(13)).Black {
		t.Fatalf("13's Left must be black but %v", tr.Search(Float64(13)))
	}
	if tr.Search(Float64(16)).Black {
		t.Fatalf("16's Left must be red but %v", tr.Search(Float64(16)))
	}
	if tr.Search(Float64(3)).Left.Key != Float64(2) {
		t.Fatalf("3's Left must be 2 but %v", tr.Search(Float64(3)).Left)
	}
	if tr.Search(Float64(2)).Left.Key != Float64(1) {
		t.Fatalf("2's Left must be 1 but %v", tr.Search(Float64(2)).Left)
	}
	if tr.Search(Float64(2)).Right.Key != Float64(2.5) {
		t.Fatalf("2's Right must be 2.5 but %v", tr.Search(Float64(2)).Right)
	}
	if tr.Search(Float64(15)).Right != nil {
		t.Fatalf("15's Right must be nil but %v", tr.Search(Float64(15)).Right)
	}
	if !tr.Search(Float64(16)).Left.Black {
		t.Fatalf("16's Left must be black but %v", tr.Search(Float64(16)).Left)
	}
	if tr.Search(Float64(17)).Black {
		t.Fatalf("17's Left must be red but %v", tr.Search(Float64(17)))
	}
	if !tr.Search(Float64(2.5)).Black {
		t.Fatalf("2.5's Left must be black but %v", tr.Search(Float64(2.5)))
	}
}

func TestBalance(t *testing.T) {
	node13 := NewNode(Float64(13))
	node13.Black = true

	node3 := NewNode(Float64(3))
	node3.Black = false

	node2 := NewNode(Float64(2))
	node2.Black = false

	tr := New(node13)
	tr.Root.Left = node3
	tr.Root.Left.Left = node2
	/*
	        13(B)
	       /
	      3(R)
	     /
	   2(R)
	*/
	tr.Root = Balance(tr.Root)
	/*
	      3(R)
	     /   \
	   2(B)   13(B)
	*/

	if tr.Search(Float64(3)).Left.Key != Float64(2) {
		t.Fatalf("3's Left must be 2 but %v", tr.Search(Float64(3)).Left)
	}
	if tr.Search(Float64(3)).Right.Key != Float64(13) {
		t.Fatalf("3's Right must be 13 but %v", tr.Search(Float64(3)).Right)
	}
	if !tr.Search(Float64(2)).Black {
		t.Fatalf("2's must be black but %v", tr.Search(Float64(2)))
	}
	if !tr.Search(Float64(13)).Black {
		t.Fatalf("13's must be black but %v", tr.Search(Float64(13)))
	}
}

func TestFixUp(t *testing.T) {
	node13 := NewNode(Float64(13))
	node13.Black = true

	node3 := NewNode(Float64(3))
	node3.Black = false

	node2 := NewNode(Float64(2))
	node2.Black = false

	tr := New(node13)
	tr.Root.Left = node3
	tr.Root.Left.Left = node2
	/*
	        13(B)
	       /
	      3(R)
	     /
	   2(R)
	*/
	tr.Root = FixUp(tr.Root)
	/*
	      3(R)
	     /   \
	   2(B)   13(B)
	*/

	if tr.Search(Float64(3)).Left.Key != Float64(2) {
		t.Fatalf("3's Left must be 2 but %v", tr.Search(Float64(3)).Left)
	}
	if tr.Search(Float64(3)).Right.Key != Float64(13) {
		t.Fatalf("3's Right must be 13 but %v", tr.Search(Float64(3)).Right)
	}
	if !tr.Search(Float64(2)).Black {
		t.Fatalf("2's must be black but %v", tr.Search(Float64(2)))
	}
	if !tr.Search(Float64(13)).Black {
		t.Fatalf("13's must be black but %v", tr.Search(Float64(13)))
	}
}
