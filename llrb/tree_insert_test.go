package llrb

import (
	"fmt"
	"testing"
)

func TestIsRed(t *testing.T) {
	if !isRed(NewNode(Int(5))) {
		t.Error("Expected Red")
	}
}

func TestInsert1(t *testing.T) {
	root := NewNode(String("S"))
	tr := New(root)
	tr.Insert(NewNode(String("E")))
	tr.Insert(NewNode(String("A")))
	nd := tr.Search(String("E"))
	if fmt.Sprintf("%v", nd.Key) != "E" {
		t.Errorf("Unexpected %v", nd.Key)
	}
	if fmt.Sprintf("%v", nd.Left.Key) != "A" {
		t.Errorf("Unexpected %v", nd.Left.Key)
	}
	if fmt.Sprintf("%v", nd.Right.Key) != "S" {
		t.Errorf("Unexpected %v", nd.Right.Key)
	}
	if !nd.Left.Black {
		t.Errorf("Left should be black but %v", nd.Left.Black)
	}
	if !nd.Right.Black {
		t.Errorf("Right should be black but %v", nd.Right.Black)
	}
}

func TestInsert2(t *testing.T) {
	root := NewNode(String("S"))
	tr := New(root)
	tr.Insert(NewNode(String("E")))
	tr.Insert(NewNode(String("A")))
	tr.Insert(NewNode(String("R")))
	tr.Insert(NewNode(String("C")))
	tr.Insert(NewNode(String("H")))
	tr.Insert(NewNode(String("X")))
	tr.Insert(NewNode(String("M")))
	tr.Insert(NewNode(String("P")))
	tr.Insert(NewNode(String("L")))
	nd := tr.Search(String("E"))
	if fmt.Sprintf("%v", nd.Key) != "E" {
		t.Errorf("Unexpected %v", nd.Key)
	}
	if fmt.Sprintf("%v", nd.Left.Key) != "C" {
		t.Errorf("Unexpected %v", nd.Left.Key)
	}
	if fmt.Sprintf("%v", nd.Right.Key) != "L" {
		t.Errorf("Unexpected %v", nd.Right.Key)
	}
	if !nd.Left.Black {
		t.Errorf("Left should be black but %v", nd.Left.Black)
	}
	if nd.Left.Left.Black {
		t.Errorf("Left, Left should be red but %v", nd.Left.Left.Black)
	}
	if !nd.Right.Black {
		t.Errorf("Right should be black but %v", nd.Right.Black)
	}
	if nd.Right.Left.Black {
		t.Errorf("Right, Left should be red but %v", nd.Right.Black)
	}
	if fmt.Sprintf("%v", nd.Right.Left.Key) != "H" {
		t.Errorf("Unexpected %v", nd.Right.Left.Key)
	}
	if fmt.Sprintf("%v", tr.Root.Key) != "M" {
		t.Errorf("Unexpected %v", tr.Root.Key)
	}
	if fmt.Sprintf("%v", tr.Root.Right.Right.Key) != "X" {
		t.Errorf("Unexpected %v", tr.Root.Right.Right.Key)
	}
	if fmt.Sprintf("%v", tr.Root.Right.Right.Left.Key) != "S" {
		t.Errorf("Unexpected %v", tr.Root.Right.Right.Left.Key)
	}
	if fmt.Sprintf("%v", tr.Root.Left.Key) != "E" {
		t.Errorf("Unexpected %v", tr.Root.Left.Key)
	}
	if fmt.Sprintf("%v", tr.Root.Right.Left.Key) != "P" {
		t.Errorf("Unexpected %v", tr.Root.Right.Left.Key)
	}
}

func TestBalanceInsert1(t *testing.T) {
	root := NewNode(Float64(1))
	tr := New(root)
	nums := []float64{3, 9, 13, 17, 20, 25, 39, 16, 15, 2, 2.5}
	for _, num := range nums {
		tr.Insert(NewNode(Float64(num)))
	}
}

func TestBalanceInsert2(t *testing.T) {
	root := NewNode(Float64(1))
	tr := New(root)
	nums := []float64{3, 9, 13, 17, 20, 25, 39, 16, 15, 2, 2.5}
	for _, num := range nums {
		t.Logf("Inserting: %f\n", num)
		tr.Insert(NewNode(Float64(num)))

		if tr.Search(Float64(num)) == nil {
			t.Fatal(num, "must not be nil")
		}
		if tr.Search(Float64(1000.12)) != nil {
			t.Fatal(1000.12, "must be nil")
		}

		switch num {
		case 3:
			if tr.Root.Key != Float64(3) {
				t.Fatalf("Root must be 3 but %v", tr.Root)
			}
			if tr.SearchParent(Float64(3)) != nil {
				t.Fatal("3's Parent must be nil")
			}
			if tr.Search(Float64(3)).Left.Key != Float64(1) {
				t.Fatal("3's Left child must be 1")
			}
			if tr.Search(Float64(3)).Right != nil {
				t.Fatal("3's Right child must be nil")
			}
			if tr.SearchParent(Float64(3)) != nil {
				t.Fatal("3's Parent must be nil")
			}
			if tr.SearchParent(Float64(1)).Key != Float64(3) {
				t.Fatal("1's Parent must be 3")
			}
			if tr.Search(Float64(1)).Right != nil {
				t.Fatal("1's Right child must be nil")
			}
			if tr.Search(Float64(1)).Left != nil {
				t.Fatal("1's Left child must be nil")
			}
			if isRed(tr.Search(Float64(3))) {
				t.Fatal("3 must be not red")
			}
			if !isRed(tr.Search(Float64(1))) {
				t.Fatal("1 must be red")
			}

		case 9:
			if tr.Root.Key != Float64(3) {
				t.Fatalf("Root must be 3 but %v", tr.Root)
			}
			if tr.SearchParent(Float64(3)) != nil {
				t.Fatal("3's Parent must be nil")
			}
			if tr.Search(Float64(3)).Left.Key != Float64(1) {
				t.Fatal("3's Left child must be 1")
			}
			if tr.Search(Float64(3)).Right.Key != Float64(9) {
				t.Fatal("3's Right child must be 9")
			}
			if tr.SearchParent(Float64(1)).Key != Float64(3) {
				t.Fatal("1's Parent must be 3")
			}
			if tr.SearchParent(Float64(9)).Key != Float64(3) {
				t.Fatal("9's Parent must be 3")
			}
			if tr.Search(Float64(1)).Right != nil {
				t.Fatal("1's Right child must be nil")
			}
			if tr.Search(Float64(1)).Left != nil {
				t.Fatal("1's Left child must be nil")
			}
			if tr.Search(Float64(9)).Right != nil {
				t.Fatal("9's Right child must be nil")
			}
			if tr.Search(Float64(9)).Left != nil {
				t.Fatal("9's Left child must be nil")
			}
			if isRed(tr.Search(Float64(3))) {
				t.Fatal("3 must be red (Root must be black always")
			}
			if isRed(tr.Search(Float64(1))) {
				t.Fatal("1 must be not red")
			}
			if isRed(tr.Search(Float64(9))) {
				t.Fatal("9 must be not red")
			}

		case 13:
			if tr.Root.Key != Float64(3) {
				t.Fatalf("Root must be 3 but %v", tr.Root)
			}
			if tr.SearchParent(Float64(3)) != nil {
				t.Fatal("3's Parent must be nil")
			}
			if tr.Search(Float64(3)).Left.Key != Float64(1) {
				t.Fatal("3's Left child must be 1")
			}
			if tr.Search(Float64(3)).Right.Key != Float64(13) {
				t.Fatal("3's Right child must be 13")
			}
			if tr.SearchParent(Float64(1)).Key != Float64(3) {
				t.Fatal("1's Parent must be 3")
			}
			if tr.SearchParent(Float64(13)).Key != Float64(3) {
				t.Fatal("13's Parent must be 3")
			}
			if tr.SearchParent(Float64(9)).Key != Float64(13) {
				t.Fatal("9's Parent must be 13")
			}
			if tr.Search(Float64(13)).Left.Key != Float64(9) {
				t.Fatal("13's Left child must be 9")
			}
			if tr.Search(Float64(13)).Right != nil {
				t.Fatal("13's Right child must be nil")
			}
			if tr.Search(Float64(1)).Right != nil {
				t.Fatal("1's Right child must be nil")
			}
			if tr.Search(Float64(1)).Left != nil {
				t.Fatal("1's Left child must be nil")
			}
			if tr.Search(Float64(9)).Right != nil {
				t.Fatal("9's Right child must be nil")
			}
			if tr.Search(Float64(9)).Left != nil {
				t.Fatal("9's Left child must be nil")
			}
			if isRed(tr.Search(Float64(3))) {
				t.Fatal("3 must be red (Root must be black always")
			}
			if isRed(tr.Search(Float64(1))) {
				t.Fatal("1 must be not red")
			}
			if isRed(tr.Search(Float64(13))) {
				t.Fatal("13 must be not red")
			}
			if !isRed(tr.Search(Float64(9))) {
				t.Fatal("9 must be red")
			}

		case 17:
			if tr.Root.Key != Float64(13) {
				t.Fatalf("Root must be 3 but %v", tr.Root)
			}
			if tr.SearchParent(Float64(13)) != nil {
				t.Fatal("13's Parent must be nil")
			}
			if tr.Search(Float64(13)).Left.Key != Float64(3) {
				t.Fatal("13's Left child must be 3")
			}
			if tr.Search(Float64(13)).Right.Key != Float64(17) {
				t.Fatal("13's Right child must be 17")
			}
			if tr.SearchParent(Float64(1)).Key != Float64(3) {
				t.Fatal("1's Parent must be 3")
			}
			if tr.SearchParent(Float64(13)) != nil {
				t.Fatal("13's Parent must be nil")
			}
			if tr.SearchParent(Float64(9)).Key != Float64(3) {
				t.Fatal("9's Parent must be 3")
			}
			if tr.Search(Float64(3)).Left.Key != Float64(1) {
				t.Fatal("3's Left child must be 1")
			}
			if tr.Search(Float64(3)).Right.Key != Float64(9) {
				t.Fatal("3's Right child must be 9")
			}
			if tr.SearchParent(Float64(17)).Key != Float64(13) {
				t.Fatal("17's Parent must be 13")
			}
			if tr.SearchParent(Float64(9)).Key != Float64(3) {
				t.Fatal("9's Parent must be 3")
			}
			if tr.Search(Float64(1)).Right != nil {
				t.Fatal("1's Right child must be nil")
			}
			if tr.Search(Float64(1)).Left != nil {
				t.Fatal("1's Left child must be nil")
			}
			if tr.Search(Float64(9)).Right != nil {
				t.Fatal("9's Right child must be nil")
			}
			if tr.Search(Float64(9)).Left != nil {
				t.Fatal("9's Left child must be nil")
			}
			if tr.Search(Float64(17)).Right != nil {
				t.Fatal("17's Right child must be nil")
			}
			if tr.Search(Float64(17)).Left != nil {
				t.Fatal("17's Left child must be nil")
			}
			if isRed(tr.Search(Float64(13))) {
				t.Fatal("13 must be red (Root must be black always")
			}
			if isRed(tr.Search(Float64(9))) {
				t.Fatal("9 must be not red")
			}
			if isRed(tr.Search(Float64(17))) {
				t.Fatal("17 must be not red")
			}
			if !isRed(tr.Search(Float64(3))) {
				t.Fatal("3 must be red")
			}

		case 20:
			if tr.Root.Key != Float64(13) {
				t.Fatalf("Root must be 3 but %v", tr.Root)
			}
			if tr.SearchParent(Float64(13)) != nil {
				t.Fatal("13's Parent must be nil")
			}
			if tr.Search(Float64(1)).Right != nil {
				t.Fatal("1's Right child must be nil")
			}
			if tr.Search(Float64(1)).Left != nil {
				t.Fatal("1's Left child must be nil")
			}
			if tr.Search(Float64(9)).Right != nil {
				t.Fatal("9's Right child must be nil")
			}
			if tr.Search(Float64(9)).Left != nil {
				t.Fatal("9's Left child must be nil")
			}
			if tr.Search(Float64(17)).Right != nil {
				t.Fatal("17's Right child must be nil")
			}
			if tr.Search(Float64(17)).Left != nil {
				t.Fatal("17's Left child must be nil")
			}
			if tr.SearchParent(Float64(3)).Key != Float64(13) {
				t.Fatal("3's Parent must be 13")
			}
			if tr.SearchParent(Float64(9)).Key != Float64(3) {
				t.Fatal("9's Parent must be 3")
			}
			if tr.SearchParent(Float64(17)).Key != Float64(20) {
				t.Fatal("17's Parent must be 20")
			}
			if tr.Search(Float64(13)).Left.Key != Float64(3) {
				t.Fatal("13's Left must be 3")
			}
			if tr.Search(Float64(13)).Right.Key != Float64(20) {
				t.Fatal("13's Right must be 20")
			}
			if tr.Search(Float64(20)).Left.Key != Float64(17) {
				t.Fatal("20's Left must be 17")
			}
			if tr.Search(Float64(3)).Left.Key != Float64(1) {
				t.Fatal("3's Left must be 1")
			}
			if tr.Search(Float64(3)).Right.Key != Float64(9) {
				t.Fatal("3's Right must be 9")
			}
			if isRed(tr.Search(Float64(13))) {
				t.Fatal("13 must be red (Root must be black always")
			}
			if isRed(tr.Search(Float64(1))) {
				t.Fatal("1 must be not red")
			}
			if isRed(tr.Search(Float64(9))) {
				t.Fatal("9 must be not red")
			}
			if isRed(tr.Search(Float64(20))) {
				t.Fatal("20 must be not red")
			}
			if !isRed(tr.Search(Float64(3))) {
				t.Fatal("3 must be red")
			}
			if !isRed(tr.Search(Float64(17))) {
				t.Fatal("17 must be red")
			}

		case 25:
			if tr.Root.Key != Float64(13) {
				t.Fatalf("Root must be 3 but %v", tr.Root)
			}
			if tr.SearchParent(Float64(13)) != nil {
				t.Fatal("13's Parent must be nil")
			}
			if tr.Search(Float64(1)).Right != nil {
				t.Fatal("1's Right child must be nil")
			}
			if tr.Search(Float64(1)).Left != nil {
				t.Fatal("1's Left child must be nil")
			}
			if tr.Search(Float64(9)).Right != nil {
				t.Fatal("9's Right child must be nil")
			}
			if tr.Search(Float64(9)).Left != nil {
				t.Fatal("9's Left child must be nil")
			}
			if tr.Search(Float64(17)).Right != nil {
				t.Fatal("17's Right child must be nil")
			}
			if tr.Search(Float64(17)).Left != nil {
				t.Fatal("17's Left child must be nil")
			}
			if tr.Search(Float64(25)).Right != nil {
				t.Fatal("25's Right child must be nil")
			}
			if tr.Search(Float64(25)).Left != nil {
				t.Fatal("25's Left child must be nil")
			}
			if tr.SearchParent(Float64(3)).Key != Float64(13) {
				t.Fatal("3's Parent must be 13")
			}
			if tr.SearchParent(Float64(9)).Key != Float64(3) {
				t.Fatal("9's Parent must be 3")
			}
			if tr.SearchParent(Float64(17)).Key != Float64(20) {
				t.Fatal("17's Parent must be 20")
			}
			if tr.Search(Float64(13)).Left.Key != Float64(3) {
				t.Fatal("13's Left must be 3")
			}
			if tr.Search(Float64(13)).Right.Key != Float64(20) {
				t.Fatal("13's Right must be 3")
			}
			if tr.Search(Float64(3)).Left.Key != Float64(1) {
				t.Fatal("3's Left must be 1")
			}
			if tr.Search(Float64(3)).Right.Key != Float64(9) {
				t.Fatal("3's Right must be 9")
			}
			if tr.Search(Float64(20)).Left.Key != Float64(17) {
				t.Fatal("20's Left must be 17")
			}
			if tr.Search(Float64(20)).Right.Key != Float64(25) {
				t.Fatal("20's Right must be 25")
			}
			if isRed(tr.Search(Float64(13))) {
				t.Fatal("13 must be red (Root must be black always")
			}
			if isRed(tr.Search(Float64(1))) {
				t.Fatal("1 must be not red")
			}
			if isRed(tr.Search(Float64(9))) {
				t.Fatal("9 must be not red")
			}
			if isRed(tr.Search(Float64(20))) {
				t.Fatal("20 must be not red")
			}
			if isRed(tr.Search(Float64(3))) {
				t.Fatal("3 must be not red")
			}
			if isRed(tr.Search(Float64(17))) {
				t.Fatal("17 must be not red")
			}
			if isRed(tr.Search(Float64(25))) {
				t.Fatal("25 must be red")
			}

		case 39:
			if tr.Root.Key != Float64(13) {
				t.Fatalf("Root must be 3 but %v", tr.Root)
			}
			if tr.SearchParent(Float64(13)) != nil {
				t.Fatal("13's Parent must be nil")
			}
			if tr.Search(Float64(1)).Right != nil {
				t.Fatal("1's Right child must be nil")
			}
			if tr.Search(Float64(1)).Left != nil {
				t.Fatal("1's Left child must be nil")
			}
			if tr.Search(Float64(9)).Right != nil {
				t.Fatal("9's Right child must be nil")
			}
			if tr.Search(Float64(9)).Left != nil {
				t.Fatal("9's Left child must be nil")
			}
			if tr.Search(Float64(17)).Right != nil {
				t.Fatal("17's Right child must be nil")
			}
			if tr.Search(Float64(17)).Left != nil {
				t.Fatal("17's Left child must be nil")
			}
			if tr.Search(Float64(25)).Right != nil {
				t.Fatal("25's Right child must be nil")
			}
			if tr.Search(Float64(25)).Left != nil {
				t.Fatal("25's Left child must be nil")
			}
			if tr.SearchParent(Float64(3)).Key != Float64(13) {
				t.Fatal("3's Parent must be 13")
			}
			if tr.SearchParent(Float64(9)).Key != Float64(3) {
				t.Fatal("9's Parent must be 3")
			}
			if tr.SearchParent(Float64(17)).Key != Float64(20) {
				t.Fatal("17's Parent must be 20")
			}
			if tr.SearchParent(Float64(25)).Key != Float64(39) {
				t.Fatal("25's Parent must be 39")
			}
			if tr.Search(Float64(13)).Left.Key != Float64(3) {
				t.Fatal("13's Left must be 3")
			}
			if tr.Search(Float64(13)).Right.Key != Float64(20) {
				t.Fatal("13's Right must be 3")
			}
			if tr.Search(Float64(3)).Left.Key != Float64(1) {
				t.Fatal("3's Left must be 1")
			}
			if tr.Search(Float64(3)).Right.Key != Float64(9) {
				t.Fatal("3's Right must be 9")
			}
			if tr.Search(Float64(20)).Left.Key != Float64(17) {
				t.Fatal("20's Left must be 17")
			}
			if tr.Search(Float64(20)).Right.Key != Float64(39) {
				t.Fatal("20's Right must be 39")
			}
			if tr.Search(Float64(39)).Left.Key != Float64(25) {
				t.Fatal("39's Left must be 25")
			}
			if isRed(tr.Search(Float64(13))) {
				t.Fatal("13 must be red (Root must be black always")
			}
			if isRed(tr.Search(Float64(1))) {
				t.Fatal("1 must be not red")
			}
			if isRed(tr.Search(Float64(9))) {
				t.Fatal("9 must be not red")
			}
			if isRed(tr.Search(Float64(20))) {
				t.Fatal("20 must be not red")
			}
			if isRed(tr.Search(Float64(3))) {
				t.Fatal("3 must be not red")
			}
			if isRed(tr.Search(Float64(17))) {
				t.Fatal("17 must be not red")
			}
			if isRed(tr.Search(Float64(39))) {
				t.Fatal("39 must be not red")
			}
			if !isRed(tr.Search(Float64(25))) {
				t.Fatal("25 must be red")
			}

		case 16:
			if tr.Root.Key != Float64(13) {
				t.Fatalf("Root must be 3 but %v", tr.Root)
			}
			if tr.SearchParent(Float64(13)) != nil {
				t.Fatal("13's Parent must be nil")
			}
			if tr.Search(Float64(1)).Right != nil {
				t.Fatal("1's Right child must be nil")
			}
			if tr.Search(Float64(1)).Left != nil {
				t.Fatal("1's Left child must be nil")
			}
			if tr.Search(Float64(9)).Right != nil {
				t.Fatal("9's Right child must be nil")
			}
			if tr.Search(Float64(9)).Left != nil {
				t.Fatal("9's Left child must be nil")
			}
			if tr.Search(Float64(16)).Right != nil {
				t.Fatal("16's Right child must be nil")
			}
			if tr.Search(Float64(16)).Left != nil {
				t.Fatal("16's Left child must be nil")
			}
			if tr.Search(Float64(25)).Right != nil {
				t.Fatal("25's Right child must be nil")
			}
			if tr.Search(Float64(25)).Left != nil {
				t.Fatal("25's Left child must be nil")
			}
			if tr.SearchParent(Float64(3)).Key != Float64(13) {
				t.Fatal("3's Parent must be 13")
			}
			if tr.SearchParent(Float64(9)).Key != Float64(3) {
				t.Fatal("9's Parent must be 3")
			}
			if tr.SearchParent(Float64(17)).Key != Float64(20) {
				t.Fatal("17's Parent must be 20")
			}
			if tr.SearchParent(Float64(25)).Key != Float64(39) {
				t.Fatal("25's Parent must be 39")
			}
			if tr.Search(Float64(13)).Left.Key != Float64(3) {
				t.Fatal("13's Left must be 3")
			}
			if tr.Search(Float64(13)).Right.Key != Float64(20) {
				t.Fatal("13's Right must be 3")
			}
			if tr.Search(Float64(3)).Left.Key != Float64(1) {
				t.Fatal("3's Left must be 1")
			}
			if tr.Search(Float64(3)).Right.Key != Float64(9) {
				t.Fatal("3's Right must be 9")
			}
			if tr.Search(Float64(20)).Left.Key != Float64(17) {
				t.Fatal("20's Left must be 17")
			}
			if tr.Search(Float64(20)).Right.Key != Float64(39) {
				t.Fatal("20's Right must be 39")
			}
			if tr.Search(Float64(39)).Left.Key != Float64(25) {
				t.Fatal("39's Left must be 25")
			}
			if tr.Search(Float64(17)).Left.Key != Float64(16) {
				t.Fatal("17's Left must be 16")
			}
			if isRed(tr.Search(Float64(13))) {
				t.Fatal("13 must be red (Root must be black always")
			}
			if isRed(tr.Search(Float64(1))) {
				t.Fatal("1 must be not red")
			}
			if isRed(tr.Search(Float64(9))) {
				t.Fatal("9 must be not red")
			}
			if isRed(tr.Search(Float64(20))) {
				t.Fatal("20 must be not red")
			}
			if isRed(tr.Search(Float64(3))) {
				t.Fatal("3 must be not red")
			}
			if isRed(tr.Search(Float64(17))) {
				t.Fatal("17 must be not red")
			}
			if isRed(tr.Search(Float64(39))) {
				t.Fatal("39 must be not red")
			}
			if !isRed(tr.Search(Float64(25))) {
				t.Fatal("25 must be red")
			}
			if !isRed(tr.Search(Float64(16))) {
				t.Fatal("16 must be red")
			}

		case 15:
			if tr.Root.Key != Float64(13) {
				t.Fatalf("Root must be 3 but %v", tr.Root)
			}
			if tr.SearchParent(Float64(13)) != nil {
				t.Fatal("13's Parent must be nil")
			}
			if tr.Search(Float64(1)).Right != nil {
				t.Fatal("1's Right child must be nil")
			}
			if tr.Search(Float64(1)).Left != nil {
				t.Fatal("1's Left child must be nil")
			}
			if tr.Search(Float64(9)).Right != nil {
				t.Fatal("9's Right child must be nil")
			}
			if tr.Search(Float64(9)).Left != nil {
				t.Fatal("9's Left child must be nil")
			}
			if tr.Search(Float64(15)).Right != nil {
				t.Fatal("15's Right child must be nil")
			}
			if tr.Search(Float64(15)).Left != nil {
				t.Fatal("15's Left child must be nil")
			}
			if tr.Search(Float64(17)).Right != nil {
				t.Fatal("17's Right child must be nil")
			}
			if tr.Search(Float64(17)).Left != nil {
				t.Fatal("17's Left child must be nil")
			}
			if tr.Search(Float64(25)).Right != nil {
				t.Fatal("25's Right child must be nil")
			}
			if tr.Search(Float64(25)).Left != nil {
				t.Fatal("25's Left child must be nil")
			}
			if tr.SearchParent(Float64(3)).Key != Float64(13) {
				t.Fatal("3's Parent must be 13")
			}
			if tr.SearchParent(Float64(9)).Key != Float64(3) {
				t.Fatal("9's Parent must be 3")
			}
			if tr.SearchParent(Float64(15)).Key != Float64(16) {
				t.Fatal("15's Parent must be 16")
			}
			if tr.SearchParent(Float64(17)).Key != Float64(16) {
				t.Fatal("17's Parent must be 16")
			}
			if tr.SearchParent(Float64(25)).Key != Float64(39) {
				t.Fatal("25's Parent must be 39")
			}
			if tr.Search(Float64(13)).Left.Key != Float64(3) {
				t.Fatal("13's Left must be 3")
			}
			if tr.Search(Float64(13)).Right.Key != Float64(20) {
				t.Fatal("13's Right must be 3")
			}
			if tr.Search(Float64(3)).Left.Key != Float64(1) {
				t.Fatal("3's Left must be 1")
			}
			if tr.Search(Float64(3)).Right.Key != Float64(9) {
				t.Fatal("3's Right must be 9")
			}
			if tr.Search(Float64(20)).Left.Key != Float64(16) {
				t.Fatal("20's Left must be 16")
			}
			if tr.Search(Float64(20)).Right.Key != Float64(39) {
				t.Fatal("20's Right must be 39")
			}
			if tr.Search(Float64(39)).Left.Key != Float64(25) {
				t.Fatal("39's Left must be 25")
			}
			if tr.Search(Float64(16)).Left.Key != Float64(15) {
				t.Fatal("16's Left must be 15")
			}
			if tr.Search(Float64(16)).Right.Key != Float64(17) {
				t.Fatal("16's Right must be 17")
			}
			if isRed(tr.Search(Float64(13))) {
				t.Fatal("13 must be red (Root must be black always")
			}
			if isRed(tr.Search(Float64(1))) {
				t.Fatal("1 must be not red")
			}
			if isRed(tr.Search(Float64(9))) {
				t.Fatal("9 must be not red")
			}
			if isRed(tr.Search(Float64(20))) {
				t.Fatal("20 must be not red")
			}
			if isRed(tr.Search(Float64(3))) {
				t.Fatal("3 must be not red")
			}
			if isRed(tr.Search(Float64(17))) {
				t.Fatal("17 must be not red")
			}
			if isRed(tr.Search(Float64(39))) {
				t.Fatal("39 must be not red")
			}
			if isRed(tr.Search(Float64(15))) {
				t.Fatal("15 must be not red")
			}
			if !isRed(tr.Search(Float64(25))) {
				t.Fatal("25 must be red")
			}
			if !isRed(tr.Search(Float64(16))) {
				t.Fatal("16 must be red")
			}

		case 2:
			if tr.Root.Key != Float64(13) {
				t.Fatalf("Root must be 3 but %v", tr.Root)
			}
			if tr.SearchParent(Float64(13)) != nil {
				t.Fatal("13's Parent must be nil")
			}
			if tr.Search(Float64(1)).Right != nil {
				t.Fatal("1's Right child must be nil")
			}
			if tr.Search(Float64(1)).Left != nil {
				t.Fatal("1's Left child must be nil")
			}
			if tr.Search(Float64(9)).Right != nil {
				t.Fatal("9's Right child must be nil")
			}
			if tr.Search(Float64(9)).Left != nil {
				t.Fatal("9's Left child must be nil")
			}
			if tr.Search(Float64(15)).Right != nil {
				t.Fatal("15's Right child must be nil")
			}
			if tr.Search(Float64(15)).Left != nil {
				t.Fatal("15's Left child must be nil")
			}
			if tr.Search(Float64(17)).Right != nil {
				t.Fatal("17's Right child must be nil")
			}
			if tr.Search(Float64(17)).Left != nil {
				t.Fatal("17's Left child must be nil")
			}
			if tr.Search(Float64(25)).Right != nil {
				t.Fatal("25's Right child must be nil")
			}
			if tr.Search(Float64(25)).Left != nil {
				t.Fatal("25's Left child must be nil")
			}
			if tr.SearchParent(Float64(3)).Key != Float64(13) {
				t.Fatal("3's Parent must be 13")
			}
			if tr.SearchParent(Float64(9)).Key != Float64(3) {
				t.Fatal("9's Parent must be 3")
			}
			if tr.SearchParent(Float64(15)).Key != Float64(16) {
				t.Fatal("15's Parent must be 16")
			}
			if tr.SearchParent(Float64(17)).Key != Float64(16) {
				t.Fatal("17's Parent must be 16")
			}
			if tr.SearchParent(Float64(1)).Key != Float64(2) {
				t.Fatal("1's Parent must be 2")
			}
			if tr.SearchParent(Float64(25)).Key != Float64(39) {
				t.Fatal("25's Parent must be 39")
			}
			if tr.Search(Float64(13)).Left.Key != Float64(3) {
				t.Fatal("13's Left must be 3")
			}
			if tr.Search(Float64(13)).Right.Key != Float64(20) {
				t.Fatal("13's Right must be 3")
			}
			if tr.Search(Float64(3)).Left.Key != Float64(2) {
				t.Fatal("3's Left must be 2")
			}
			if tr.Search(Float64(3)).Right.Key != Float64(9) {
				t.Fatal("3's Right must be 9")
			}
			if tr.Search(Float64(20)).Left.Key != Float64(16) {
				t.Fatal("20's Left must be 16")
			}
			if tr.Search(Float64(20)).Right.Key != Float64(39) {
				t.Fatal("20's Right must be 39")
			}
			if tr.Search(Float64(39)).Left.Key != Float64(25) {
				t.Fatal("39's Left must be 25")
			}
			if tr.Search(Float64(16)).Left.Key != Float64(15) {
				t.Fatal("16's Left must be 15")
			}
			if tr.Search(Float64(16)).Right.Key != Float64(17) {
				t.Fatal("16's Right must be 17")
			}
			if tr.Search(Float64(2)).Left.Key != Float64(1) {
				t.Fatal("2's Left must be 1")
			}
			if isRed(tr.Search(Float64(13))) {
				t.Fatal("13 must be red (Root must be black always")
			}
			if !isRed(tr.Search(Float64(1))) {
				t.Fatal("1 must be red")
			}
			if isRed(tr.Search(Float64(9))) {
				t.Fatal("9 must be not red")
			}
			if isRed(tr.Search(Float64(20))) {
				t.Fatal("20 must be not red")
			}
			if isRed(tr.Search(Float64(3))) {
				t.Fatal("3 must be not red")
			}
			if isRed(tr.Search(Float64(17))) {
				t.Fatal("17 must be not red")
			}
			if isRed(tr.Search(Float64(39))) {
				t.Fatal("39 must be not red")
			}
			if isRed(tr.Search(Float64(15))) {
				t.Fatal("15 must be not red")
			}
			if isRed(tr.Search(Float64(2))) {
				t.Fatal("2 must be not red")
			}
			if !isRed(tr.Search(Float64(25))) {
				t.Fatal("25 must be red")
			}
			if !isRed(tr.Search(Float64(16))) {
				t.Fatal("16 must be red")
			}

		case 2.5:
			if tr.Root.Key != Float64(13) {
				t.Fatalf("Root must be 3 but %v", tr.Root)
			}
			if tr.SearchParent(Float64(13)) != nil {
				t.Fatal("13's Parent must be nil")
			}
			if tr.Search(Float64(1)).Right != nil {
				t.Fatal("1's Right child must be nil")
			}
			if tr.Search(Float64(1)).Left != nil {
				t.Fatal("1's Left child must be nil")
			}
			if tr.Search(Float64(2.5)).Right != nil {
				t.Fatal("2.5's Right child must be nil")
			}
			if tr.Search(Float64(2.5)).Left != nil {
				t.Fatal("2.5's Left child must be nil")
			}
			if tr.Search(Float64(9)).Right != nil {
				t.Fatal("9's Right child must be nil")
			}
			if tr.Search(Float64(9)).Left != nil {
				t.Fatal("9's Left child must be nil")
			}
			if tr.Search(Float64(15)).Right != nil {
				t.Fatal("15's Right child must be nil")
			}
			if tr.Search(Float64(15)).Left != nil {
				t.Fatal("15's Left child must be nil")
			}
			if tr.Search(Float64(17)).Right != nil {
				t.Fatal("17's Right child must be nil")
			}
			if tr.Search(Float64(17)).Left != nil {
				t.Fatal("17's Left child must be nil")
			}
			if tr.Search(Float64(25)).Right != nil {
				t.Fatal("25's Right child must be nil")
			}
			if tr.Search(Float64(25)).Left != nil {
				t.Fatal("25's Left child must be nil")
			}
			if tr.Search(Float64(39)).Right != nil {
				t.Fatal("39's Right child must be nil")
			}
			if tr.SearchParent(Float64(3)).Key != Float64(13) {
				t.Fatal("3's Parent must be 13")
			}
			if tr.SearchParent(Float64(9)).Key != Float64(3) {
				t.Fatal("9's Parent must be 3")
			}
			if tr.SearchParent(Float64(15)).Key != Float64(16) {
				t.Fatal("15's Parent must be 16")
			}
			if tr.SearchParent(Float64(17)).Key != Float64(16) {
				t.Fatal("17's Parent must be 16")
			}
			if tr.SearchParent(Float64(1)).Key != Float64(2) {
				t.Fatal("1's Parent must be 2")
			}
			if tr.SearchParent(Float64(25)).Key != Float64(39) {
				t.Fatal("25's Parent must be 39")
			}
			if tr.SearchParent(Float64(2.5)).Key != Float64(2) {
				t.Fatal("2.5's Parent must be 2")
			}
			if tr.SearchParent(Float64(2.1)) != nil {
				t.Fatal("2.1's Parent must be nil")
			}
			if tr.Search(Float64(13)).Left.Key != Float64(3) {
				t.Fatal("13's Left must be 3")
			}
			if tr.Search(Float64(13)).Right.Key != Float64(20) {
				t.Fatal("13's Right must be 3")
			}
			if tr.Search(Float64(3)).Left.Key != Float64(2) {
				t.Fatal("3's Left must be 2")
			}
			if tr.Search(Float64(3)).Right.Key != Float64(9) {
				t.Fatal("3's Right must be 9")
			}
			if tr.Search(Float64(20)).Left.Key != Float64(16) {
				t.Fatal("20's Left must be 16")
			}
			if tr.Search(Float64(20)).Right.Key != Float64(39) {
				t.Fatal("20's Right must be 39")
			}
			if tr.Search(Float64(39)).Left.Key != Float64(25) {
				t.Fatal("39's Left must be 25")
			}
			if tr.Search(Float64(16)).Left.Key != Float64(15) {
				t.Fatal("16's Left must be 15")
			}
			if tr.Search(Float64(16)).Right.Key != Float64(17) {
				t.Fatal("16's Right must be 17")
			}
			if tr.Search(Float64(2)).Left.Key != Float64(1) {
				t.Fatal("2's Left must be 1")
			}
			if tr.Search(Float64(2)).Right.Key != Float64(2.5) {
				t.Fatal("2's Right must be 2.5")
			}
			if isRed(tr.Search(Float64(13))) {
				t.Fatal("13 must be red (Root must be black always")
			}
			if isRed(tr.Search(Float64(1))) {
				t.Fatal("1 must be not red")
			}
			if isRed(tr.Search(Float64(9))) {
				t.Fatal("9 must be not red")
			}
			if isRed(tr.Search(Float64(2.5))) {
				t.Fatal("2.5 must be not red")
			}
			if isRed(tr.Search(Float64(20))) {
				t.Fatal("20 must be not red")
			}
			if isRed(tr.Search(Float64(3))) {
				t.Fatal("3 must be not red")
			}
			if isRed(tr.Search(Float64(17))) {
				t.Fatal("17 must be not red")
			}
			if isRed(tr.Search(Float64(39))) {
				t.Fatal("39 must be not red")
			}
			if isRed(tr.Search(Float64(15))) {
				t.Fatal("15 must be not red")
			}
			if !isRed(tr.Search(Float64(2))) {
				t.Fatal("2 must be red")
			}
			if !isRed(tr.Search(Float64(25))) {
				t.Fatal("25 must be red")
			}
			if !isRed(tr.Search(Float64(16))) {
				t.Fatal("16 must be red")
			}

		default:
			t.Fatalf("%f shouldn't be there...", num)
		}
	}
}
