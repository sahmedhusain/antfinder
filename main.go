package main

import (
	"fmt"
	structs "lem-in/datastruct"
	move "lem-in/logic/antsmoving"
	paths "lem-in/logic/paths"
	rooms "lem-in/logic/rooms"
	fileoperations "lem-in/operations/file"
	stringoperations "lem-in/operations/strings"
	"os"
)

func main() {
	state := &structs.State{
		Visited:         make(map[string]bool),
		Tunnels:         make(map[string][]string),
		RoomsMapRoom:    make(map[string]bool),
		RoomsMapTunnels: make(map[string]string),
		StartEnd:        make(map[string]int),
		StartEndRooms:   make(map[string]string),
		RoomCoordinates: make(map[string]int),
		Solution:        make(map[string]string),
		Connect:         make(map[string]int),
	}
	if len(os.Args) > 2 {
		fmt.Println("Error: too many arguments")
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	} else if len(os.Args) == 1 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	} else if len(os.Args) == 2 {
		fileoperations.ReadFile(os.Args[1], state)

		if !rooms.Roomtunels(state.RoomsMapTunnels, state.Rooms) {
			fmt.Println("Error: tunnels do not connect to a valid room")
			os.Exit(1)
		}
		if state.RoomsMapTunnels[state.End] == "" {
			fmt.Println("Error: start and end rooms do not connect")
			os.Exit(1)
		}
		if state.StartEndRooms["start"] != "" && state.StartEndRooms["end"] != "" && len(state.StartAndEnd) == 2 {
			startedroom := state.StartEndRooms["start"]
			endedroom := state.StartEndRooms["end"]
			for _, route := range state.Tunnels[startedroom] {
				path := []string{}
				state.Visited[startedroom] = true
				paths.Start(path, endedroom, route, state)
			}
			if state.Ants == 0 {
				fmt.Println("error, no Ants found")
				return
			} else if state.Ants > 1000 {
				fmt.Println("error, Ants number is too large, Max Ants number is 1000")
				return
			}
			state.Paths = stringoperations.SortByLength(stringoperations.GetUniqueStringSets(state.Paths))
			solution := stringoperations.EqualizeSlices(move.MoveLogic(state.Ants, state.Paths, state.End))

			if state.Paths == nil {
				fmt.Println("error, no paths connect the start to the end")
				return
			} else {
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
	} else {
		fmt.Println("Error: start and end rooms are not defined or are duplicated")
		os.Exit(1)
	}
}
