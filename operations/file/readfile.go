package fileOps

import (
	"bufio"                            // Package for buffered I/O
	"fmt"                              // Package for formatted I/O
	structs "lem-in/datastruct"        // Importing custom data structures
	paths "lem-in/logic/paths"         // Importing custom path logic
	rooms "lem-in/logic/rooms"         // Importing custom room logic
	strOps "lem-in/operations/strings" // Importing custom string operations
	"os"                               // Package for OS functions
)

// ReadFile reads the file at the given FilePath and returns a pointer to the State struct
func ReadFile(FilePath string) *structs.State {
	// Initialize the State struct with default values
	state := &structs.State{
		Visited:         make(map[string]bool),     // Map to track visited rooms
		Tunnels:         make(map[string][]string), // Map of tunnels between rooms
		RoomsMapRoom:    make(map[string]bool),     // Map of rooms
		RoomsMapTunnels: make(map[string]string),   // Map of room tunnels
		StartEnd:        make(map[string]int),      // Map to track start and end rooms
		StartEndRooms:   make(map[string]string),   // Map of start and end room names
		RoomCoordinates: make(map[string]int),      // Map of room coordinates
		Solution:        make(map[string]string),   // Map to store the solution
		Connect:         make(map[string]int),      // Map to track connections
	}
	fmt.Printf("\nInitial state: %+v\n", state) // Debug print the initial state

	// Open the file at the given FilePath
	f, err := os.Open(FilePath)
	if err != nil {
		// Print error message and exit if file cannot be opened
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer f.Close() // Ensure the file is closed when the function exits

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Read each line from the file
		line := scanner.Text()
<<<<<<< Updated upstream
		stringoperations.VarState(line, state)
=======
		// Map the room information from the line to the state
		rooms.RoomMapping(line, state)
		fmt.Printf("\nState after RoomMapping: %+v\n", state) // Debug print the state after RoomMapping
>>>>>>> Stashed changes
	}

	// Check if the tunnels connect to valid rooms
	if !rooms.Roomtunels(state.RoomsMapTunnels, state.Rooms) {
		fmt.Println("Error: tunnels do not connect to a valid room")
		os.Exit(1)
	}
	fmt.Printf("\nState after Roomtunels check: %+v\n", state) // Debug print the state after Roomtunels check

	// Check if the start and end rooms are connected
	if state.RoomsMapTunnels[state.End] == "" {
		fmt.Println("Error: start and end rooms do not connect")
		os.Exit(1)
	}
	fmt.Printf("\nState after start and end rooms connection check: %+v\n", state) // Debug print the state after start and end rooms connection check

	// Check if start and end rooms are defined and not duplicated
	if state.StartEndRooms["start"] != "" && state.StartEndRooms["end"] != "" && len(state.StartAndEnd) == 2 {
		startedroom := state.StartEndRooms["start"] // Get the start room
		endedroom := state.StartEndRooms["end"]     // Get the end room
		// Iterate through the tunnels from the start room
		for _, route := range state.Tunnels[startedroom] {
			path := []string{}                // Initialize an empty path
			state.Visited[startedroom] = true // Mark the start room as visited
			fmt.Printf("\nState after marking start room as visited: %+v\n", state) // Debug print the state after marking start room as visited
			// Start finding paths from the start to the end room
			paths.Start(path, endedroom, route, state)
			fmt.Printf("\nState after paths.Start: %+v\n", state) // Debug print the state after paths.Start
		}
<<<<<<< Updated upstream
		if state.Ants == 0 {
			fmt.Println("error, no Ants found")
			os.Exit(1)
		} else if state.Ants > 1000 {
			fmt.Println("error, Ants number is too large, Max Ants number is 1000")
=======
		maxants := 1000 // Define the maximum number of ants
		// Check if the number of ants is valid
		if state.Ants == 0 {
			fmt.Println("error, no Ants found")
			os.Exit(1)
		} else if state.Ants > maxants {
			fmt.Printf("error, Ants number is too large, Max Ants number is %v\n", maxants)
>>>>>>> Stashed changes
			os.Exit(1)
		}
		fmt.Printf("\nState after checking ants: %+v\n", state) // Debug print the state after checking ants

		// Sort and get unique paths
		state.Paths = strOps.SortByLength(strOps.GetUniqueStringSets(state.Paths))
		fmt.Printf("\nState after sorting and getting unique paths: %+v\n", state) // Debug print the state after sorting and getting unique paths

		// Check if there are valid paths from start to end
		if state.Paths == nil {
			fmt.Println("error, no paths connect the start to the end")
			os.Exit(1)
		} else {
			fmt.Printf("\nFinal state before returning: %+v\n", state) // Debug print the final state before returning
			return state // Return the state if everything is valid
		}
	}
	// Print error if start and end rooms are not defined or are duplicated
	fmt.Println("Error: start and end rooms are not defined or are duplicated")
	os.Exit(1)
	return nil // Return nil if there is an error
}
