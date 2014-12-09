package helper

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

// DeleteNonAlnum removes all alphanumeric characters.
func DeleteNonAlnum(str string) string {
	// alphanumeric (== [0-9A-Za-z])
	// \s is a white space character
	validID := regexp.MustCompile(`[^[:alnum:]\s]`)
	return validID.ReplaceAllString(str, "")
}

// OpenToOverwrite creates or opens a file for overwriting.
// Make sure to close the file.
func OpenToOverwrite(fpath string) *os.File {
	file, err := os.OpenFile(fpath, os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		// log.Fatal(err)
		file, err = os.Create(fpath)
		if err != nil {
			log.Fatal(err)
		}
	}
	return file
}

// WriteToFile writes the input string slice into a text file.
func WriteToFile(fpath, str string) {
	file := OpenToOverwrite(fpath)
	defer file.Close()
	txt := bufio.NewWriter(file)
	_, err := txt.WriteString(str)
	if err != nil {
		log.Fatal(err)
	}
	defer txt.Flush()
}

// ReadLines reads lines from a specified file,
// and returns them in string slice format.
func ReadLines(fpath string) []string {
	file, err := os.OpenFile(fpath, os.O_RDONLY, 0444)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// CountLines returns the number of lines.
func CountLines(fpath string) int {
	return len(ReadLines(fpath))
}
