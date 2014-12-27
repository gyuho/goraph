package helper

import (
	"os"
	"testing"
)

func TestWriteToFile(t *testing.T) {
	WriteToFile("./tmp.txt", "A\nB\nC")
	if CountLines("./tmp.txt") != 3 {
		t.Fatalf("expected 3 but %v", CountLines("./tmp.txt"))
	}
	os.RemoveAll("./tmp.txt")
}
