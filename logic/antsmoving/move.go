package antsMove

import (
    "fmt"                              // Importing the fmt package for formatted I/O
    structs "lem-in/datastruct"        // Importing custom data structures
    rooms "lem-in/logic/rooms"         // Importing custom room logic
    strOps "lem-in/operations/strings" // Importing custom string operations
)

var path [][]string                    // Variable to store paths
var ant int                            // Variable to store the number of ants
var longestSolution [][]string         // Variable to store the longest solution
var end string                         // Variable to store the end room
var longestSolutionWithWait [][]string // Variable to store the longest solution with wait times
var shortestWithoutWait [][]string     // Variable to store the shortest path without wait times
var shortestWithWait [][]string        // Variable to store the shortest path with wait times
var shortest [][]string                // Variable to store the shortest path
var solution [][]string                // Variable to store the final solution
var apnum int                          // Variable to store the adjusted number of paths
var state *structs.State               // Variable to store the state
var original int                       // Variable to store the original number of ants

// MoveLogic calculates the optimal paths for the given number of ants and paths
func MoveLogic(ants int, paths [][]string, ends string) [][]string {
    state = &structs.State{ // Initialize the state with default values
        Many: make(map[string]int), // Map to track occurrences of paths
    }
    fmt.Printf("\nInitial state: %+v\n", state) // Debug print the initial state

    original = ants // Store the original number of ants
    if ants > 400 { // Check if the number of ants is greater than 400
        ant, apnum = strOps.ModifyNumber(strOps.UpperClosestDivisibleBy10(ants)) // Adjust the number of ants and paths
    } else {
        ant = ants   // Use the original number of ants
        apnum = ants // Use the original number of paths
    }
    fmt.Printf("\nState after adjusting ants and paths: %+v\n", state) // Debug print the state after adjusting ants and paths

    end = ends                    // Store the end room
    path = strOps.DeepCopy(paths) // Create a deep copy of the paths
    fmt.Printf("\nState after copying paths: %+v\n", state) // Debug print the state after copying paths

    if len(path) > ant { // Check if the number of paths is greater than the number of ants
        if len(path) >= 200 { // Check if the number of paths is greater than or equal to 200
            path = path[:20] // Limit the paths to 20
        }
        if len(path) >= 60 { // Check if the number of paths is greater than or equal to 60
            path = path[:12] // Limit the paths to 12
        }
        if len(path) >= ant { // Check if the number of paths is greater than or equal to the number of ants
            path = path[:ant] // Limit the paths to the number of ants
        }
    }
    fmt.Printf("\nState after limiting paths: %+v\n", state) // Debug print the state after limiting paths

    for i := 1; i <= ant; i++ { // Iterate through the number of ants
        longestSolution = append(longestSolution, path[len(path)-1]) // Append the longest path to the longest solution
    }
    fmt.Printf("\nState after appending longest paths: %+v\n", state) // Debug print the state after appending longest paths

    var soz [][]string                                                       // Variable to store the temporary solution
    shortest = longestSolution                                               // Initialize the shortest path with the longest solution
    number := len(longestSolution) - 1                                       // Get the index of the last path in the longest solution
    solution = rooms.ModifyCommonRoom(strOps.DeepCopy(longestSolution), end) // Modify the longest solution to include the end room
    fmt.Printf("\nState after modifying common room: %+v\n", state) // Debug print the state after modifying common room

    for i := len(path) - 2; i >= 0; i-- {                                    // Iterate through the paths in reverse order
        counter := len(longestSolution)                                                                                // Initialize the counter with the length of the longest solution
        longestSolution = strOps.DeepCopy(shortest)                                                                    // Create a deep copy of the shortest path
        shortestWithoutWait = strOps.DeepCopy(shortest)                                                                // Create a deep copy of the shortest path without wait times
        longestSolutionWithWait = rooms.ModifyCommonRoom(strOps.DeepCopy(longestSolution), end)                        // Modify the longest solution to include wait times
        copy(shortestWithoutWait, strOps.ReorderTallest((strOps.DeepCopy(shortestWithoutWait)), state))                // Reorder the shortest path without wait times
        copy(shortestWithoutWait, strOps.AppendNestedSlices(shortestWithoutWait, strOps.MapToNestedArray(state.Many))) // Append nested slices to the shortest path without wait times
        copy(shortestWithoutWait, strOps.Recur(shortestWithoutWait, path[i], number))                                  // Recursively update the shortest path without wait times
        shortestWithWait = rooms.ModifyCommonRoom(strOps.DeepCopy(shortestWithoutWait), end)                           // Modify the shortest path to include wait times
        fmt.Printf("\nState after modifying shortest path: %+v\n", state) // Debug print the state after modifying shortest path

        if len(shortestWithWait[len(shortestWithWait)-1]) <= len(longestSolutionWithWait[len(shortestWithWait)-1]) {   // Check if the length of the shortest path with wait times is less than or equal to the length of the longest solution with wait times
            if len(shortestWithWait[len(shortestWithWait)-1]) < len(solution[len(shortestWithWait)-1]) { // Check if the length of the shortest path with wait times is less than the length of the solution
                solution = strOps.DeepCopy(shortestWithWait) // Update the solution with the shortest path with wait times
            }
            moh := len(longestSolution)                                                                                   // Initialize the variable to track the length of the longest solution
            for len(shortestWithWait[len(shortestWithWait)-1]) <= len(longestSolutionWithWait[len(shortestWithWait)-1]) { // Iterate while the length of the shortest path with wait times is less than or equal to the length of the longest solution with wait times
                longestSolutionWithWait = strOps.DeepCopy(shortestWithWait)                                  // Update the longest solution with wait times
                copy(shortestWithoutWait, strOps.Recur(shortestWithoutWait, path[i], number))                // Recursively update the shortest path without wait times
                copy(shortestWithWait, rooms.ModifyCommonRoom(strOps.DeepCopy(shortestWithoutWait), end))    // Modify the shortest path to include wait times
                if len(shortestWithWait[len(shortestWithWait)-1]) < len(solution[len(shortestWithWait)-1]) { // Check if the length of the shortest path with wait times is less than the length of the solution
                    solution = strOps.DeepCopy(shortestWithWait) // Update the solution with the shortest path with wait times
                    soz = strOps.DeepCopy(shortestWithoutWait)   // Update the temporary solution with the shortest path without wait times
                }
                moh--         // Decrement the length of the longest solution
                if moh == 0 { // Check if the length of the longest solution is zero
                    break // Exit the loop
                }
            }
            shortest = strOps.DeepCopy(shortestWithoutWait) // Update the shortest path with the shortest path without wait times
        } else {
            break // Exit the loop if the length of the shortest path with wait times is greater than the length of the longest solution with wait times
        }
        counter--         // Decrement the counter
        if counter == 0 { // Check if the counter is zero
            break // Exit the loop
        }
        fmt.Printf("\nState after iteration: %+v\n", state) // Debug print the state after each iteration
    }
    if len(soz) == 0 { // Check if the temporary solution is empty
        soz = strOps.DeepCopy(longestSolution) // Update the temporary solution with the longest solution
    }
    fmt.Printf("\nState after updating temporary solution: %+v\n", state) // Debug print the state after updating temporary solution

    shortest = (strOps.AppendMultipleTimes(soz, apnum-1))[:original]        // Append the temporary solution multiple times to the shortest path
    longestSolution = (strOps.AppendMultipleTimes(soz, apnum-1))[:original] // Append the temporary solution multiple times to the longest solution
    solution = rooms.ModifyCommonRoom(longestSolution, end)                 // Modify the longest solution to include the end room
    fmt.Printf("\nState after final modifications: %+v\n", state) // Debug print the state after final modifications

    for i := len(path) - 2; i >= 0; i-- {                                   // Iterate through the paths in reverse order
        counter := len(longestSolution)                                                                                // Initialize the counter with the length of the longest solution
        longestSolution = strOps.DeepCopy(shortest)                                                                    // Create a deep copy of the shortest path
        shortestWithoutWait = strOps.DeepCopy(shortest)                                                                // Create a deep copy of the shortest path without wait times
        longestSolutionWithWait = rooms.ModifyCommonRoom(strOps.DeepCopy(longestSolution), end)                        // Modify the longest solution to include wait times
        copy(shortestWithoutWait, strOps.ReorderTallest((strOps.DeepCopy(shortestWithoutWait)), state))                // Reorder the shortest path without wait times
        copy(shortestWithoutWait, strOps.AppendNestedSlices(shortestWithoutWait, strOps.MapToNestedArray(state.Many))) // Append nested slices to the shortest path without wait times
        copy(shortestWithoutWait, strOps.Recur(shortestWithoutWait, path[i], number))                                  // Recursively update the shortest path without wait times
        shortestWithWait = rooms.ModifyCommonRoom(strOps.DeepCopy(shortestWithoutWait), end)                           // Modify the shortest path to include wait times
        fmt.Printf("\nState after modifying shortest path: %+v\n", state) // Debug print the state after modifying shortest path

        if len(shortestWithWait[len(shortestWithWait)-1]) <= len(longestSolutionWithWait[len(shortestWithWait)-1]) {   // Check if the length of the shortest path with wait times is less than or equal to the length of the longest solution with wait times
            if len(shortestWithWait[len(shortestWithWait)-1]) < len(solution[len(shortestWithWait)-1]) { // Check if the length of the shortest path with wait times is less than the length of the solution
                solution = strOps.DeepCopy(shortestWithWait) // Update the solution with the shortest path with wait times
            }
            moh := len(longestSolution)                                                                                   // Initialize the variable to track the length of the longest solution
            for len(shortestWithWait[len(shortestWithWait)-1]) <= len(longestSolutionWithWait[len(shortestWithWait)-1]) { // Iterate while the length of the shortest path with wait times is less than or equal to the length of the longest solution with wait times
                longestSolutionWithWait = strOps.DeepCopy(shortestWithWait)                                  // Update the longest solution with wait times
                copy(shortestWithoutWait, strOps.Recur(shortestWithoutWait, path[i], number))                // Recursively update the shortest path without wait times
                copy(shortestWithWait, rooms.ModifyCommonRoom(strOps.DeepCopy(shortestWithoutWait), end))    // Modify the shortest path to include wait times
                if len(shortestWithWait[len(shortestWithWait)-1]) < len(solution[len(shortestWithWait)-1]) { // Check if the length of the shortest path with wait times is less than the length of the solution
                    solution = strOps.DeepCopy(shortestWithWait) // Update the solution with the shortest path with wait times
                    soz = strOps.DeepCopy(shortestWithoutWait)   // Update the temporary solution with the shortest path without wait times
                }
                moh--         // Decrement the length of the longest solution
                if moh == 0 { // Check if the length of the longest solution is zero
                    break // Exit the loop
                }
            }
            shortest = strOps.DeepCopy(shortestWithoutWait) // Update the shortest path with the shortest path without wait times
        } else {
            break // Exit the loop if the length of the shortest path with wait times is greater than the length of the longest solution with wait times
        }
        counter--         // Decrement the counter
        if counter == 0 { // Check if the counter is zero
            break // Exit the loop
        }
        fmt.Printf("\nState after iteration: %+v\n", state) // Debug print the state after each iteration
    }
    if len(soz) == 0 { // Check if the temporary solution is empty
        soz = strOps.DeepCopy(longestSolution) // Update the temporary solution with the longest solution
    }
    fmt.Printf("\nState after updating temporary solution: %+v\n", state) // Debug print the state after updating temporary solution

    return rooms.ModifyCommonRoom((strOps.AppendMultipleTimes(soz, apnum-1))[:original], end) // Return the final solution
}
