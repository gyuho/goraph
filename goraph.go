// goraph provides graph visualizing tools and graph algorithm implementations.
package main

import (
	"fmt"
	"strings"
	"time"
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
Hello World! This is Gyu-Ho (gyuho.cs@gmail.com).

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
[5] Algorithm Visualization : Topological Sort
[6] Algorithm Visualization : Minimum Spanning Tree
[7] Algorithm Visualization : Network Flow
`)
		mm := make(map[string]string)
		mm["1"] = "Visualize a Graph"
		mm["2"] = "Algorithm Visualization : Breadth First Search"
		mm["3"] = "Algorithm Visualization : Depth First Search"
		mm["4"] = "Algorithm Visualization : Shortest Path"
		mm["5"] = "Algorithm Visualization : Topological Sort"
		mm["6"] = "Algorithm Visualization : Minimum Spanning Tree"
		mm["7"] = "Algorithm Visualization : Network Flow"

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
			fmt.Println("1")
		case "2":
			fmt.Println("2")
		case "3":
			fmt.Println("3")
		case "4":
			fmt.Println("4")
		case "5":
			fmt.Println("5")
		case "6":
			fmt.Println("6")
		case "7":
			fmt.Println("7")
		default:
			fmt.Println("Done")
		}

		time.Sleep(time.Second)
	}
}
