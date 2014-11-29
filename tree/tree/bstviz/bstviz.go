// Package bstviz visualizes bst package.
package bstviz

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/gyuho/gotree/tree/bst"
)

// CheckStr returns true if the string exists in the slice.
func CheckStr(str string, slice []string) bool {
	for _, v := range slice {
		if str == v {
			return true
		}
	}
	return false
}

// Scan scans all nodes in the tree recursively.
func Scan(T *bst.Tree, slice *[]string) {
	if T == nil {
		return
	}
	if T.Left != nil {
		str := "\t" + strconv.FormatInt(T.Value, 10) + " -- " + strconv.FormatInt(T.Left.Value, 10)
		if !CheckStr(str, *slice) && strconv.FormatInt(T.Value, 10) != strconv.FormatInt(T.Left.Value, 10) {
			*slice = append(*slice, str)
		}
	}
	if T.Right != nil {
		str := "\t" + strconv.FormatInt(T.Value, 10) + " -- " + strconv.FormatInt(T.Right.Value, 10)
		if !CheckStr(str, *slice) && strconv.FormatInt(T.Value, 10) != strconv.FormatInt(T.Right.Value, 10) {
			*slice = append(*slice, str)
		}
	}
	Scan(T.Left, slice)
	Scan(T.Right, slice)
}

// Convert converts the tree into DOT format.
func Convert(T *bst.Tree, outputfile string) {
	file, err := os.Create(outputfile)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	result := "graph tree {" + "\n"
	slice := []string{}
	Scan(T, &slice)
	result += strings.Join(slice, "\n")
	result += "\n}"
	file.WriteString(result)
}

// Show visualizes the tree in DOT format.
func Show(T *bst.Tree, outputfile string) {
	Convert(T, outputfile)
	// func Command(name string, arg ...string) *Cmd
	// Command returns the Cmd struct to execute the named program with the given arguments.
	cmd := exec.Command("open", outputfile)
	er := cmd.Run()
	if er != nil {
		log.Fatal(er)
	}
}
