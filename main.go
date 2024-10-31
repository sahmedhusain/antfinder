package main

import (
    "fmt"                              // Importing the fmt package for formatted I/O
    structs "lem-in/datastruct"        // Importing custom data structures from the datastruct package
    antsMove "lem-in/logic/antsmoving" // Importing ant movement logic from the logic/antsmoving package
    fileOps "lem-in/operations/file"   // Importing file operations from the operations/file package
    strOps "lem-in/operations/strings" // Importing string operations from the operations/strings package
    "os"                               // Importing the os package for operating system functionality
)

var state *structs.State // Declaring a global variable to hold the state

func main() {
    if len(os.Args) > 2 { // Check if there are more than 2 command-line arguments
        fmt.Println("Error: Too many arguments provided.") // Print an error message
        fmt.Println("Usage: go run main.go <filename>")    // Print the correct usage
        os.Exit(1)                                         // Exit the program with an error code
    } else if len(os.Args) == 1 { // Check if there are no command-line arguments
        fmt.Println("Usage: go run main.go <filename>") // Print the correct usage
        os.Exit(1)                                      // Exit the program with an error code
    } else if len(os.Args) == 2 { // Check if there is exactly one command-line argument
        state = fileOps.ReadFile(os.Args[1]) // Read the file and initialize the state
        fmt.Printf("\nState after reading file: %+v\n", state) // Debug print the state

        solution := strOps.EqualizeSlices(antsMove.MoveLogic(state.Ants, state.Paths, state.End)) // Process the state and get the solution
        fmt.Printf("\nState after processing: %+v\n", state) // Debug print the state

        for i := 0; i < len(solution[0]); i++ { // Iterate over the length of the first solution slice
            for j := 0; j < state.Ants; j++ { // Iterate over the number of ants
                if i >= len(solution[j]) || solution[j][i] == "wait" { // Check if the current step is out of bounds or if the ant should wait
                    continue // Skip to the next iteration
                }
                fmt.Print("L")            // Print the ant move prefix
                fmt.Print(j + 1)          // Print the ant number (1-based index)
                fmt.Print("-")            // Print the separator
                fmt.Print(solution[j][i]) // Print the current position of the ant
                fmt.Print(" ")            // Print a space
            }
            fmt.Println() // Print a newline after each step
        }
    }
}
