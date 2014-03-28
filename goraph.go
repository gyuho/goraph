// goraph provides graph visualizing tools and algorithm implementations.
package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gyuho/goraph/algorithm/bfs"
	"github.com/gyuho/goraph/algorithm/dfs"
	"github.com/gyuho/goraph/algorithm/mst/kruskal"
	"github.com/gyuho/goraph/algorithm/spd"
	"github.com/gyuho/goraph/graph/gsd"
	"github.com/gyuho/goraph/viz"
)

func main() {
	println()
	gstr := `   ______                                    __  
  / ____/  ____    _____  ____ _    ____    / /_ 
 / / __   / __ \  / ___/ / __  /   / __ \  / __ \
/ /_/ /  / /_/ / / /    / /_/ /   / /_/ / / / / /
\____/   \____/ /_/     \__._/   / .___/ /_/ /_/ 
                                /_/                       `
	fmt.Println("===================================================")
	fmt.Println(gstr)
	fmt.Println("===================================================")

	fmt.Println(`
Goraph provides:
	- graph visualizing tool
	- graph algorithm implementations

This is a personal project, non-committal on a timeline.
Everything is still in heavy development.

Type "help" for help.
Type "Ctrl + c" or "q" to exit.`)

	for {
		fmt.Println(`
[1] Visualize a Graph
[2] Algorithm Visualization : Breadth First Search
[3] Algorithm Visualization : Depth First Search
[4] Algorithm Visualization : Shortest Path
[5] Algorithm Visualization : Minimum Spanning Tree
`)
		mm := make(map[string]string)
		mm["1"] = "Visualize a Graph"
		mm["2"] = "Algorithm Visualization : Breadth First Search"
		mm["3"] = "Algorithm Visualization : Depth First Search"
		mm["4"] = "Algorithm Visualization : Shortest Path"
		mm["5"] = "Algorithm Visualization : Minimum Spanning Tree"

		fmt.Print("Select the menu: ")
		var choice string
		fmt.Scanf("%s", &choice)
		if choice == "q" || choice == "Q" || choice == "quit" {
			fmt.Print("Are you sure? [y/n]: ")
			var qas string
			fmt.Scanf("%s", &qas)
			if strings.ToLower(qas) == "y" {
				break
			} else if strings.ToLower(qas) == "n" {
				continue
			}
		} else {
			fmt.Println("You've chosen \"" + mm[choice] + "\"")
		}

		switch choice {
		case "1":
			fmt.Print("✔ Type or Drag the JSON graph file: ")
			var path string
			fmt.Scanf("%s", &path)
			fmt.Print("✔ Type the name of graph to visualize: ")
			var name string
			fmt.Scanf("%s", &name)
			viz.Show(path, name, name+"-output.dot")
		case "2":
			fmt.Print("✔ Type or Drag the JSON graph file: ")
			var path string
			fmt.Scanf("%s", &path)
			fmt.Print("✔ Type the name of graph to visualize: ")
			var name string
			fmt.Scanf("%s", &name)
			g := gsd.JSONGraph(path, name)
			fmt.Print("✔ Type the name of node to start: ")
			var node string
			fmt.Scanf("%s", &node)
			bfs.ShowBFS(g, g.FindVertexByID(node), name+"-output.dot")
		case "3":
			fmt.Print("✔ Type or Drag the JSON graph file: ")
			var path string
			fmt.Scanf("%s", &path)
			fmt.Print("✔ Type the name of graph to visualize: ")
			var name string
			fmt.Scanf("%s", &name)
			g := gsd.JSONGraph(path, name)
			dfs.ShowDFS(g, name+"-outputdfs.dot")
		case "4":
			fmt.Print("✔ Type or Drag the JSON graph file: ")
			var path string
			fmt.Scanf("%s", &path)
			fmt.Print("✔ Type the name of graph to visualize: ")
			var name string
			fmt.Scanf("%s", &name)
			g := gsd.JSONGraph(path, name)
			fmt.Print("✔ Start Node: ")
			var start string
			fmt.Scanf("%s", &start)
			fmt.Print("✔ End Node: ")
			var end string
			fmt.Scanf("%s", &end)
			spd.ShowSPD(g, start, end, name+"-output.dot")
		case "5":
			fmt.Print("✔ Type or Drag the JSON graph file: ")
			var path string
			fmt.Scanf("%s", &path)
			fmt.Print("✔ Type the name of graph to visualize: ")
			var name string
			fmt.Scanf("%s", &name)
			g := gsd.JSONGraph(path, name)
			kruskal.ShowMST(g, name+"-output.dot")
		default:
			fmt.Println("Done")
		}

		time.Sleep(time.Second)
	}
}
