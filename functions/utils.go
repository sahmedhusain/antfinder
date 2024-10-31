package functions

import (
	"bufio"
	"fmt"
	s "lem-in/datastruct"
	"os"
	"sort"
	"strconv"
	"strings"
)

func modifyNumber(n int) (j int, i int) {
	if n < 400 {
		return n, 1
	}
	if n >= 400 && n < 2000 {
		for i := 10; i >= 2; i-- {
			if n%i == 0 {
				return n / i, i
			}
		}
	}
	if n >= 2000 && n <= 5000 {
		for i := 50; i >= 2; i-- {
			if n%i == 0 {
				return n / i, i
			}
		}
	}
	if n >= 5001 && n <= 20000 {
		for i := 100; i >= 2; i-- {
			if n%i == 0 {
				return n / i, i
			}
		}
	}
	return n, 1
}

func appendMultipleTimes(sets [][]string, n int) [][]string {
	result := deepCopy(sets)
	for j := 0; j < n; j++ {
		result = append(result, sets...)
	}
	return result
}

func deepCopy(sets [][]string) [][]string {
	copied := make([][]string, len(sets))
	for i := range sets {
		copied[i] = append([]string(nil), sets[i]...)
	}
	return copied
}

func upperClosestDivisibleBy10(n int) int {
	if n%10 == 0 {
		return n
	}
	return ((n / 10) + 1) * 10
}

func modifyCommonRoom(sets [][]string) [][]string {
	sets = sortByLength(sets)
	for i := 0; i < len(sets)-1; i++ {
		for j := 0; j < len(sets[i]); j++ {
			for m := i + 1; m < len(sets); m++ {
				if j >= len(sets[m]) {
					continue
				}
				if sets[i][j] != "wait" && sets[i][j] == sets[m][j] {
					if sets[m][j] == end && !identicalSlices(sets[m], sets[i]) {
						continue
					}
					before := sets[m][:j]
					after := sets[m][j:]
					sets[m] = append(before, append([]string{"wait"}, after...)...)
				}
			}
		}
	}
	return sortByLength(sets)
}

func recur(a [][]string, b []string, number int) [][]string {
	if testEq(a, b, number) {
		for i := number - 1; i >= 0; i-- {
			if !testEq(a, b, i) {
				a[i] = b
				return a
			}
		}
	} else {
		a[number] = b
		return a
	}
	return a
}

func reorderTallest(sets [][]string, state *s.State) [][]string {
	if len(sets) == 0 {
		return sets
	}
	pattern := sets
	number := len(pattern) - 1
	length := len(pattern[number])
	for i := number; i >= 0; i-- {
		if len(pattern[i]) == length {
			state.Many[strings.Join(pattern[i], "&*")]++
		} else {
			return pattern[:number+1]
		}
	}
	return [][]string{}
}

func appendNestedSlices(destination, source [][]string) [][]string {
	destination = append(destination, source...)
	return destination
}

func mapToNestedArray(m map[string]int) [][]string {
	var result [][]string
	tempMap := make(map[string]int)
	for k, v := range m {
		tempMap[k] = v
	}
	for len(tempMap) > 0 {
		for k, v := range tempMap {
			if v > 0 {
				result = append(result, strings.Split(k, "&*"))
				tempMap[k] = v - 1
				if tempMap[k] == 0 {
					delete(tempMap, k)
				}
			}
		}
	}
	return result
}

func variableop(s string, state *s.State) {
	if s == "" {
	} else if isNumber(s) && !state.AntsHere {
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
			if !isNumber(t) {
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
		state.Connect[reverseHyphenatedString(s)]++
		if state.Connect[s] > 1 || state.Connect[reverseHyphenatedString(s)] > 1 {
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

func roomtunels(tunel map[string]string, room []string) bool {
	for _, t := range tunel {
		if contains(room, t) {
			continue
		} else if !contains(room, t) {
			return false
		} else {
			return true
		}
	}
	return true
}

func startPathOpt(path []string, end string, route string, state *s.State) {
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
					startPathOpt(path, end, nextRoute, state)
				}
				state.Visited[nextRoute] = false
			} else {
				continue
			}
		}
	}
}

//

//

func getUniqueStringSets(sets [][]string) [][]string {
	result := make([][]string, 0, len(sets))
	for _, set := range sets {
		if isUniqueStringSet(set) {
			result = append(result, set)
		}
	}
	return result
}

func isUniqueStringSet(set []string) bool {
	seen := make(map[string]bool, len(set))
	for _, s := range set {
		if seen[s] {
			return false
		}
		seen[s] = true
	}
	return true
}

func equalizeSlices(data [][]string) [][]string {
	maxLen := 0
	for _, row := range data {
		if len(row) > maxLen {
			maxLen = len(row)
		}
	}
	result := make([][]string, len(data))
	for i, row := range data {
		if len(row) < maxLen {
			result[i] = make([]string, maxLen)
			copy(result[i], row)
			for j := len(row); j < maxLen; j++ {
				result[i][j] = "wait"
			}
		} else {
			result[i] = row
		}
	}
	return result
}

func reverseHyphenatedString(str string) string {
	parts := strings.Split(str, "-")
	before := reverseString(parts[0])
	after := reverseString(parts[1])
	return fmt.Sprintf("%s-%s", after, before)
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func contains(slice []string, item string) bool {
	for _, i := range slice {
		if i == item {
			return true
		}
	}
	return false
}

func isNumber(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

func sortByLength(sets [][]string) [][]string {
	sort.Slice(sets, func(i, j int) bool {
		return len(sets[i]) < len(sets[j])
	})
	return sets
}

func testEq(a [][]string, b []string, number int) bool {
	if len(a[number]) != len(b) {
		return false
	}
	for i := range a[number] {
		if a[number][i] != b[i] {
			return false
		}
	}
	return true
}

func identicalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func ProcessFile(fileName string) {
	state := &s.State{
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

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		variableop(line, state)
	}

	if !roomtunels(state.RoomsMapTunnels, state.Rooms) {
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
			startPathOpt(path, endedroom, route, state)
		}
		state.Paths = sortByLength(getUniqueStringSets(state.Paths))
		solution := equalizeSlices(Antsop(state.Ants, state.Paths, state.End))

		if state.Ants == 0 {
			fmt.Println("Error: no Ants found")
		} else if state.Ants > 10000 {
			fmt.Println("Error: Ants number is too large")
		} else {
			if state.Paths == nil {
				fmt.Println("Error: no paths connect the start to the end")
			} else {
				fmt.Println(len(solution[state.Ants-1]))
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
