package dotxd

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/gyuho/goraph/parsex"
)

// GetNodes returns the list of nodes in the specified graph.
func GetNodes(fpath string) []string {
	_, nmap := GetGraphMap(fpath)
	result := []string{}
	for k := range nmap {
		result = append(result, k)
	}
	return parsex.UniqElemStr(result)
}

// GetGraphMap reads lines from a DOT file.
func GetGraphMap(fpath string) (string, map[string]map[string][]float64) {
	file, err := os.OpenFile(fpath, os.O_RDONLY, 0444)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		// line S -> A [label=100
		if len(line) > 2 {
			lines = append(lines, strings.TrimSpace(line)[:len(line)-2])
		} else {
			lines = append(lines, line)
		}
	}

	if !strings.Contains(lines[0], "digraph") {
		log.Fatalln("Only for `digraph`")
	}

	gname := strings.TrimSpace(strings.Split(lines[0], "digraph")[1])
	gname = strings.Trim(gname, "{")

	nmap := make(map[string]map[string][]float64)
	for idx, line := range lines {
		if idx == 0 {
			continue
		}
		if idx == len(lines)-1 {
			break
		}

		wghv := parsex.StrToFloat64(strings.Split(line, "label=")[1])
		line = strings.TrimSpace(strings.Split(line, "[")[0])
		tmpls := strings.Split(line, "->")
		srcNodes := []string{}
		for _, v := range tmpls {
			srcNodes = append(srcNodes, strings.TrimSpace(v))
		}
		if len(srcNodes) < 2 {
			log.Println("Not enough nodes...")
			break
		}

		for k, node := range srcNodes {
			if k < len(srcNodes)-1 {
				nextNode := srcNodes[k+1]
				if _, ok := nmap[node]; ok {
					nmap[node][nextNode] = append(nmap[node][nextNode], wghv)
				} else {
					tmap := make(map[string][]float64)
					tmap[nextNode] = []float64{wghv}
					nmap[node] = tmap
				}
			}
		}
	}
	return gname, nmap
}
