package paths

import (
	structs "lem-in/datastruct"
)

func Start(path []string, end string, route string, state *structs.State) {
	if route == end {
		path = append(path, route)
		state.Paths = append(state.Paths, append([]string{}, path...))
	} else {
		path = append(path, route)
		for _, nextRoute := range state.Tunnels[route] {
			if !state.Visited[nextRoute] {
				state.Visited[nextRoute] = true
				if nextRoute == end {
					path = append(path, nextRoute)
					state.Paths = append(state.Paths, append([]string{}, path...))
				} else {
					Start(path, end, nextRoute, state)
				}
				state.Visited[nextRoute] = false
			} else {
				continue
			}
		}
	}
}
