package rooms

import (
	"fmt"
	stringoperations "lem-in/operations/strings"
	structs "lem-in/datastruct"
	"os"
	"strconv"
	"strings"
)

func RoomMapping(s string, state *structs.State) {
	if s == "" {
	} else if stringoperations.IsNumber(s) && !state.AntsHere {
		state.Ants, _ = strconv.Atoi(s)
		if state.Ants < 0 {
			fmt.Println("Error: Ants number is negative")
			os.Exit(1)
		}
		state.AntsHere = true
	} else if s == "##end" {
		state.StartEnd["end"]++
		state.StartAndEnd = append(state.StartAndEnd, "##end")
	} else if s == "##start" {
		state.StartEnd["start"]++
		state.StartAndEnd = append(state.StartAndEnd, "##end")
	} else if s[0] == '#' {
	} else if len(strings.Fields(s)) == 3 {
		tru := true
		for _, t := range strings.Fields(s)[1:] {
			if !stringoperations.IsNumber(t) {
				tru = false
				fmt.Println("invalid data format 1")
				os.Exit(1)
				break
			}
		}
		if tru {
			state.Rooms = append(state.Rooms, strings.Fields(s)[0])
			state.RoomsMapRoom[strings.Fields(s)[0]] = true
			state.RoomCoordinates[strings.Join(strings.Fields(s)[1:], ",")]++
			if state.RoomCoordinates[strings.Join(strings.Fields(s)[1:], ",")] > 1 {
				fmt.Println("Error: room coordinates are not unique")
				os.Exit(1)
			}
			if state.StartEnd["start"] == 1 && state.StartEnd["end"] == 0 {
				state.StartEndRooms["start"] = strings.Fields(s)[0]
				state.StartEnd["start"]--
			} else if state.StartEnd["end"] == 1 && state.StartEnd["start"] == 0 {
				state.StartEndRooms["end"] = strings.Fields(s)[0]
				state.End = strings.Fields(s)[0]
				state.StartEnd["end"]--
			}
		} else {
			fmt.Println("invalid data format 2")
			os.Exit(1)
		}
	} else if strings.ContainsRune(s, '-') {
		state.Connect[s]++
		state.Connect[stringoperations.ReverseHyphenatedString(s)]++
		if state.Connect[s] > 1 || state.Connect[stringoperations.ReverseHyphenatedString(s)] > 1 {
			fmt.Println("Error: tunnels are not unique")
			os.Exit(1)
		}
		tempar := strings.Split(s, "-")
		if len(tempar) == 2 {
			state.RoomsMapTunnels[tempar[0]] = tempar[0]
			state.RoomsMapTunnels[tempar[1]] = tempar[1]
			state.Tunnels[tempar[0]] = append(state.Tunnels[tempar[0]], tempar[1])
			state.Tunnels[tempar[1]] = append(state.Tunnels[tempar[1]], tempar[0])
		} else {
			fmt.Println("invalid data format")
			os.Exit(1)
		}
	} else {
		fmt.Println("invalid data format")
		os.Exit(1)
	}
}

func ModifyCommonRoom(sets [][]string, end string) [][]string {
	sets = stringoperations.SortByLength(sets)
	for i := 0; i < len(sets)-1; i++ {
		for j := 0; j < len(sets[i]); j++ {
			for m := i + 1; m < len(sets); m++ {
				if j >= len(sets[m]) {
					continue
				}
				if sets[i][j] != "wait" && sets[i][j] == sets[m][j] {
					if sets[m][j] == end && !stringoperations.IdenticalSlices(sets[m], sets[i]) {
						continue
					}
					before := sets[m][:j]
					after := sets[m][j:]
					sets[m] = append(before, append([]string{"wait"}, after...)...)
				}
			}
		}
	}
	return stringoperations.SortByLength(sets)
}

func Roomtunels(tunel map[string]string, room []string) bool {
	for _, t := range tunel {
		if stringoperations.Contains(room, t) {
			continue
		} else if stringoperations.Contains(room, t) {
			return false
		} else {
			return true
		}
	}
	return true
}
