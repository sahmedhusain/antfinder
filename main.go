package main

import (
	"fmt"
	structs "lem-in/datastruct"
	move "lem-in/logic/antsmoving"
	fileoperations "lem-in/operations/file"
	stringoperations "lem-in/operations/strings"
	"os"
)

var state *structs.State

func main() {

	if len(os.Args) > 2 {
		fmt.Println("Error: too many arguments")
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	} else if len(os.Args) == 1 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	} else if len(os.Args) == 2 {
		state = fileoperations.ReadFile(os.Args[1])
		solution := stringoperations.EqualizeSlices(move.MoveLogic(state.Ants, state.Paths, state.End))

		for i := 0; i < len(solution[0]); i++ {
			for j := 0; j < state.Ants; j++ {
				if i >= len(solution[j]) || solution[j][i] == "wait" {
					continue
				}
				fmt.Print("L")
				fmt.Print(j + 1)
				fmt.Print("-")
				fmt.Print(solution[j][i])
				fmt.Print(" ")
			}
			fmt.Println()
		}
	}
}
