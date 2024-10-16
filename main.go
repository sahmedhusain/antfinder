package main

import (
	"fmt"
	structs "lem-in/datastruct"
	handler "lem-in/handlers"
)

func main() {
	var filePath string
	fmt.Print("Enter the path to the file: ")
	fmt.Scanln(&filePath)

	lines, err := handler.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	connections := structs.ParseLines(lines)
	graph := structs.BuildGraph(connections)

	// Print the graph to verify
	for name, room := range graph.Rooms {
		fmt.Printf("Room: %s, Connections: ", name)
		for conn := range room.Connections {
			fmt.Printf("%s ", conn)
		}
		fmt.Println()
	}
}
