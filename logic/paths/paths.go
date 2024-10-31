package paths

import (
    "fmt"                       // Importing the fmt package for formatted I/O
    structs "lem-in/datastruct" // Importing custom data structures from the datastruct package
)

// Start initiates the pathfinding process
func Start(path []string, end string, route string, state *structs.State) {
    if route == end { // Check if the current route is the end
        path = append(path, route)                                     // Append the end route to the path
        state.Paths = append(state.Paths, append([]string{}, path...)) // Add the completed path to the state's paths
        fmt.Printf("\nState after adding completed path: %+v\n", state)  // Debug print the state after adding completed path
    } else {
        path = append(path, route)                       // Append the current route to the path
        fmt.Printf("\nState after appending current route: %+v\n", state) // Debug print the state after appending current route
        for _, nextRoute := range state.Tunnels[route] { // Iterate over the tunnels connected to the current route
            if !state.Visited[nextRoute] { // Check if the next route has not been visited
                state.Visited[nextRoute] = true // Mark the next route as visited
                fmt.Printf("\nState after marking next route as visited: %+v\n", state) // Debug print the state after marking next route as visited
                if nextRoute == end {           // Check if the next route is the end
                    path = append(path, nextRoute)                                 // Append the end route to the path
                    state.Paths = append(state.Paths, append([]string{}, path...)) // Add the completed path to the state's paths
                    fmt.Printf("\nState after adding completed path: %+v\n", state)  // Debug print the state after adding completed path
                } else {
                    Start(path, end, nextRoute, state) // Recursively start the pathfinding process for the next route
                }
                state.Visited[nextRoute] = false // Mark the next route as unvisited (backtracking)
                fmt.Printf("\nState after backtracking: %+v\n", state) // Debug print the state after backtracking
            } else {
                continue // Continue to the next iteration if the route has already been visited
            }
        }
    }
}
