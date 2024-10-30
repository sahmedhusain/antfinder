package fileoperations

import (
	"bufio"
	"fmt"
	structs "lem-in/datastruct"
	paths "lem-in/logic/paths"
	rooms "lem-in/logic/rooms"
	stringoperations "lem-in/operations/strings"
	"os"
)

func ReadFile(FilePath string) *structs.State {
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
	f, err := os.Open(FilePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		rooms.RoomMapping(line, state)
	}

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
			os.Exit(1)
		} else if state.Ants > 1000 {
			fmt.Println("error, Ants number is too large, Max Ants number is 1000")
			os.Exit(1)
		}
		state.Paths = stringoperations.SortByLength(stringoperations.GetUniqueStringSets(state.Paths))

		if state.Paths == nil {
			fmt.Println("error, no paths connect the start to the end")
			os.Exit(1)
		} else {
			return state
		}
	}
	fmt.Println("Error: start and end rooms are not defined or are duplicated")
	os.Exit(1)
	return nil
}
