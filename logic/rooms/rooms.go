package rooms // Define the package name

import (
<<<<<<< Updated upstream
	stringoperations "lem-in/operations/strings"
)

=======
	"fmt"                              // Import the fmt package for formatted I/O
	structs "lem-in/datastruct"        // Import structs from a custom package
	strOps "lem-in/operations/strings" // Import string operations from a custom package
	"os"                               // Import the os package for operating system functionality
	"strconv"                          // Import the strconv package for string conversions
	"strings"                          // Import the strings package for string manipulation
)

// Function to map rooms based on the input string and update the state
func RoomMapping(s string, state *structs.State) {
	if s == "" { // Check if the input string is empty
		// Do nothing if the string is empty
	} else if strOps.IsNumber(s) && !state.AntsHere { // Check if the string is a number and ants are not yet set
		state.Ants, _ = strconv.Atoi(s) // Convert the string to an integer and set the number of ants
		if state.Ants < 0 {             // Check if the number of ants is negative
			fmt.Println("Error: Ants number is negative") // Print an error message
			os.Exit(1)                                    // Exit the program with an error code
		}
		state.AntsHere = true // Set the ants flag to true
		fmt.Printf("\nState after setting ants: %+v\n", state) // Debug print the state after setting ants
	} else if s == "##end" { // Check if the string is "##end"
		state.StartEnd["end"]++                                // Increment the end counter in the state
		state.StartAndEnd = append(state.StartAndEnd, "##end") // Append "##end" to the StartAndEnd slice
		fmt.Printf("\nState after setting end: %+v\n", state) // Debug print the state after setting end
	} else if s == "##start" { // Check if the string is "##start"
		state.StartEnd["start"]++                              // Increment the start counter in the state
		state.StartAndEnd = append(state.StartAndEnd, "##end") // Append "##end" to the StartAndEnd slice
		fmt.Printf("\nState after setting start: %+v\n", state) // Debug print the state after setting start
	} else if s[0] == '#' { // Check if the string starts with '#'
		// Do nothing for comments
	} else if len(strings.Fields(s)) == 3 { // Check if the string has three fields
		tru := true                               // Initialize a boolean flag
		for _, t := range strings.Fields(s)[1:] { // Iterate over the second and third fields
			if !strOps.IsNumber(t) { // Check if the field is not a number
				tru = false                          // Set the flag to false
				fmt.Println("invalid data format 1") // Print an error message
				os.Exit(1)                           // Exit the program with an error code
				break                                // Break the loop
			}
		}
		if tru { // If the flag is true
			state.Rooms = append(state.Rooms, strings.Fields(s)[0])                  // Append the room name to the Rooms slice
			state.RoomsMapRoom[strings.Fields(s)[0]] = true                          // Set the room name in the RoomsMapRoom map
			state.RoomCoordinates[strings.Join(strings.Fields(s)[1:], ",")]++        // Increment the room coordinates counter
			if state.RoomCoordinates[strings.Join(strings.Fields(s)[1:], ",")] > 1 { // Check if the coordinates are not unique
				fmt.Println("Error: room coordinates are not unique") // Print an error message
				os.Exit(1)                                            // Exit the program with an error code
			}
			fmt.Printf("\nState after adding room: %+v\n", state) // Debug print the state after adding room
			if state.StartEnd["start"] == 1 && state.StartEnd["end"] == 0 { // Check if start is set and end is not
				state.StartEndRooms["start"] = strings.Fields(s)[0] // Set the start room
				state.StartEnd["start"]--                           // Decrement the start counter
				fmt.Printf("\nState after setting start room: %+v\n", state) // Debug print the state after setting start room
			} else if state.StartEnd["end"] == 1 && state.StartEnd["start"] == 0 { // Check if end is set and start is not
				state.StartEndRooms["end"] = strings.Fields(s)[0] // Set the end room
				state.End = strings.Fields(s)[0]                  // Set the end room in the state
				state.StartEnd["end"]--                           // Decrement the end counter
				fmt.Printf("\nState after setting end room: %+v\n", state) // Debug print the state after setting end room
			}
		} else {
			fmt.Println("Error: Invalid data format.") // Print an error message
			os.Exit(1)                                 // Exit the program with an error code
		}
	} else if strings.ContainsRune(s, '-') { // Check if the string contains a hyphen
		state.Connect[s]++                                                                // Increment the connection counter
		state.Connect[strOps.ReverseHyphenatedString(s)]++                                // Increment the reversed connection counter
		if state.Connect[s] > 1 || state.Connect[strOps.ReverseHyphenatedString(s)] > 1 { // Check if the connection is not unique
			fmt.Println("Error: Tunnels are not unique.") // Print an error message
			os.Exit(1)                                    // Exit the program with an error code
		}
		tempar := strings.Split(s, "-") // Split the string by hyphen
		if len(tempar) == 2 {           // Check if the split result has two parts
			state.RoomsMapTunnels[tempar[0]] = tempar[0]                           // Set the first part in the RoomsMapTunnels map
			state.RoomsMapTunnels[tempar[1]] = tempar[1]                           // Set the second part in the RoomsMapTunnels map
			state.Tunnels[tempar[0]] = append(state.Tunnels[tempar[0]], tempar[1]) // Append the second part to the tunnels map
			state.Tunnels[tempar[1]] = append(state.Tunnels[tempar[1]], tempar[0]) // Append the first part to the tunnels map
			fmt.Printf("\nState after adding tunnel: %+v\n", state) // Debug print the state after adding tunnel
		} else {
			fmt.Println("Error: Invalid data format.") // Print an error message
			os.Exit(1)                                 // Exit the program with an error code
		}
	} else {
		fmt.Println("Error: Invalid data format.") // Print an error message
		os.Exit(1)                                 // Exit the program with an error code
	}
}

// Function to modify common rooms in sets of paths
>>>>>>> Stashed changes
func ModifyCommonRoom(sets [][]string, end string) [][]string {
	sets = strOps.SortByLength(sets)   // Sort the sets by length
	for i := 0; i < len(sets)-1; i++ { // Iterate over the sets
		for j := 0; j < len(sets[i]); j++ { // Iterate over the elements in the set
			for m := i + 1; m < len(sets); m++ { // Iterate over the remaining sets
				if j >= len(sets[m]) { // Check if the index is out of bounds
					continue // Continue to the next iteration
				}
				if sets[i][j] != "wait" && sets[i][j] == sets[m][j] { // Check if the elements are equal and not "wait"
					if sets[m][j] == end && !strOps.IdenticalSlices(sets[m], sets[i]) { // Check if the element is the end and the sets are not identical
						continue // Continue to the next iteration
					}
					before := sets[m][:j]                                           // Get the elements before the index
					after := sets[m][j:]                                            // Get the elements after the index
					sets[m] = append(before, append([]string{"wait"}, after...)...) // Insert "wait" at the index
				}
			}
		}
	}
	return strOps.SortByLength(sets) // Return the sorted sets
}

// Function to check if tunnels are valid for a room
func Roomtunels(tunel map[string]string, room []string) bool {
	for _, t := range tunel { // Iterate over the tunnels
		if strOps.Contains(room, t) { // Check if the room contains the tunnel
			continue // Continue to the next iteration
		} else if strOps.Contains(room, t) { // Check if the room contains the tunnel (redundant condition)
			return false // Return false
		} else {
			return true // Return true
		}
	}
	return true // Return true
}
