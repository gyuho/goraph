package avlviz

import (
	"testing"

	"github.com/gyuho/gotree/tree/avl"
)

// Show works as expected,
// but it is commented out
// just for the sake of Travis.org testing
func Test_Show1(test *testing.T) {
	tr := avl.NewTree(4)
	tr.Inserts(3, 2)
	// Show(tr, "avl01.dot")
}

func Test_Show2(test *testing.T) {
	tr := avl.NewTree(5)
	tr.Inserts(7, 8, 4, 2, 1, 3)
	// Show(tr, "avl02.dot")
}

func Test_Show3(test *testing.T) {
	tr := avl.NewTree(5)
	tr.Inserts(7, 8, 4, 2, 1, 6, 3, 9, 10)
	// Show(tr, "avl03.dot")
}

func Test_Show4(test *testing.T) {
	tr := avl.NewTree(4)
	tr.BalanceInsert(7)
	tr.BalanceInsert(5)
	// Show(tr, "avl04.dot")
}

func Test_Show5(test *testing.T) {
	tr := avl.NewTree(4)
	tr.BalanceInsert(3)
	tr.BalanceInsert(2)
	tr.BalanceLL(2)
	// Show(tr, "avl05.dot")
}

func Test_Show6(test *testing.T) {
	tr := avl.NewTree(4)
	tr.BalanceInsert(2)
	tr.BalanceInsert(3)
	// Show(tr, "avl06.dot")

	tr.BalanceLR(3)
	// Show(tr, "avl06.dot")
}

func Test_Show7(test *testing.T) {
	tr := avl.NewTree(4)
	tr.BalanceInsert(6)
	tr.BalanceInsert(5)
	// Show(tr, "avl07.dot")

	tr.BalanceRL(5)
	// Show(tr, "avl07.dot")
}

func Test_Show8(test *testing.T) {
	tr := avl.NewTree(10)
	tr.BalanceInserts(13, 17, 5, 4, 7)
	// Show(tr, "avl08.dot")
}

func Test_Show9(test *testing.T) {
	tr := avl.NewTree(10)
	tr.BalanceInserts(13, 17)
	// fmt.Println(tr.Right.Value)
	// 17
	// Show(tr, "avl09.dot")
}

func Test_Show10(test *testing.T) {
	tr := avl.NewTree(13)
	tr.Inserts(5, 17, 4, 10)
	// Show(tr, "avl10.dot")
}

func Test_Show11(test *testing.T) {
	tr := avl.NewTree(13)
	tr.Inserts(5, 17, 4, 10)
	// Show(tr, "avl11.dot")
}

func Test_Show12(test *testing.T) {
	tr := avl.NewTree(13)
	tr.Inserts(5, 17, 4, 10)
	tr.RotateRight(13)
	// Show(tr, "avl12.dot")
}

func Test_Show13(test *testing.T) {
	tr := avl.NewTree(5)
	tr.Inserts(4, 13, 10, 17)
	// Show(tr, "avl13.dot")
}

func Test_Show14(test *testing.T) {
	tr := avl.NewTree(5)
	tr.Inserts(4, 13, 10, 17)
	tr.RotateLeft(5)
	// Show(tr, "avl14.dot")
}

// Left Left Case
func Test_Show15(test *testing.T) {
	tr := avl.NewTree(13)
	tr.Inserts(5, 17, 3, 10, 4, 2)
	// fmt.Println(tr.Height(13))
	// 2
	// Show(tr, "avl15.dot")
}

// Left Left Case
func Test_Show16(test *testing.T) {
	tr := avl.NewTree(13)
	tr.Inserts(5, 17, 3, 10, 4, 2)
	tr.RotateRight(13)
	// Show(tr, "avl16.dot")
}

// Left Right Case
func Test_Show17(test *testing.T) {
	tr := avl.NewTree(13)
	tr.Inserts(5, 17, 3, 10, 12, 9)
	// fmt.Println(tr.Height(13))
	// 2
	// Show(tr, "avl17.dot")
}

// Left Right Case
func Test_Show18(test *testing.T) {
	tr := avl.NewTree(13)
	tr.Inserts(5, 17, 3, 10, 12, 9)
	tr.RotateLeft(5)
	// Show(tr, "avl18.dot")
}

// Left Right Case
func Test_Show19(test *testing.T) {
	tr := avl.NewTree(13)
	tr.Inserts(5, 17, 3, 10, 12, 9)
	tr.RotateLeft(5)
	tr.RotateRight(13)
	// Show(tr, "avl19.dot")
}

// Right Right Case
func Test_Show20(test *testing.T) {
	tr := avl.NewTree(7)
	tr.Inserts(4, 12, 8, 15, 17, 13)
	// fmt.Println(tr.Height(7))
	// -2
	// Show(tr, "avl20.dot")
}

// Right Right Case
func Test_Show21(test *testing.T) {
	tr := avl.NewTree(7)
	tr.Inserts(4, 12, 8, 15, 17, 13)
	tr.RotateLeft(7)
	// Show(tr, "avl21.dot")
}

// Right Left Case
func Test_Show22(test *testing.T) {
	tr := avl.NewTree(7)
	tr.Inserts(4, 12, 9, 15, 8, 10)
	// fmt.Println(tr.Height(7))
	// -2
	// Show(tr, "avl22.dot")
}

// Right Left Case
func Test_Show23(test *testing.T) {
	tr := avl.NewTree(7)
	tr.Inserts(4, 12, 9, 15, 8, 10)
	tr.RotateRight(12)
	// Show(tr, "avl23.dot")
}

// Right Left Case
func Test_Show24(test *testing.T) {
	tr := avl.NewTree(7)
	tr.Inserts(4, 12, 9, 15, 8, 10)
	tr.RotateRight(12)
	tr.RotateLeft(7)
	// Show(tr, "avl24.dot")
}

func Test_Show25(test *testing.T) {
	// Left Left Case
	tr1 := avl.NewTree(13)
	tr1.TreeInserts(5, 17, 3, 10, 4, 2)
	// Show(tr1, "avl_balanced_25.dot")

	// Left Right Case
	tr2 := avl.NewTree(13)
	tr2.TreeInserts(5, 17, 3, 10, 12, 9)
	// Show(tr2, "avl_balanced_26.dot")

	// Right Right Case
	tr3 := avl.NewTree(7)
	tr3.TreeInserts(4, 12, 8, 15, 17, 13)
	// Show(tr3, "avl_balanced_27.dot")

	// Right Left Case
	tr4 := avl.NewTree(7)
	tr4.TreeInserts(4, 12, 9, 15, 8, 10)
	// Show(tr4, "avl_balanced_28.dot")
}
