package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type State struct {
	Visited         map[string]bool
	Tunnels         map[string][]string
	Rooms           []string
	RoomsMapRoom    map[string]bool
	RoomsMapTunnels map[string]string
	StartEnd        map[string]int
	StartEndRooms   map[string]string
	Ants            int
	AntsHere        bool
	Solution        map[string]string
	RoomCoordinates map[string]int
	Paths           [][]string
	End             string
	StartAndEnd     []string
	Connect         map[string]int
	Many            map[string]int
}

var path [][]string
var ant int
var longestSolution [][]string
var end string
var longestSolutionWithWait [][]string
var shortestWithoutWait [][]string
var shortestWithWait [][]string
var shortest [][]string
var solution [][]string
var apnum int
var state *State
var original int
var soz [][]string
var counter int

// func Antsop(ants int, paths [][]string, ends string) [][]string {
// 	state = &State{
// 		Many: make(map[string]int),
// 	}
// 	original = ants
// 	if ants > 400 {
// 		ant, apnum = modifyNumber(upperClosestDivisibleBy10(ants))
// 	} else {
// 		ant = ants
// 		apnum = ants
// 	}
// 	end = ends
// 	path = deepCopy(paths)
// 	if len(path) > ant {
// 		if len(path) >= 200 {
// 			path = path[:20]
// 		}
// 		if len(path) >= 60 {
// 			path = path[:12]
// 		}
// 		if len(path) >= ant {
// 			path = path[:ant]
// 		}
// 	}
// 	for i := 1; i <= ant; i++ {
// 		longestSolution = append(longestSolution, path[len(path)-1])
// 	}
// 	var soz [][]string
// 	shortest = longestSolution
// 	number := len(longestSolution) - 1
// 	solution = modifyCommonRoom(deepCopy(longestSolution))
// 	for i := len(path) - 2; i >= 0; i-- {
// 		counter := len(longestSolution)
// 		longestSolution = deepCopy(shortest)
// 		shortestWithoutWait = deepCopy(shortest)
// 		longestSolutionWithWait = modifyCommonRoom(deepCopy(longestSolution))
// 		copy(shortestWithoutWait, reorderTallest((deepCopy(shortestWithoutWait)), state))
// 		copy(shortestWithoutWait, appendNestedSlices(shortestWithoutWait, mapToNestedArray(state.Many)))
// 		copy(shortestWithoutWait, recur(shortestWithoutWait, path[i], number))
// 		shortestWithWait = modifyCommonRoom(deepCopy(shortestWithoutWait))
// 		if len(shortestWithWait[len(shortestWithWait)-1]) <= len(longestSolutionWithWait[len(shortestWithWait)-1]) {
// 			if len(shortestWithWait[len(shortestWithWait)-1]) < len(solution[len(shortestWithWait)-1]) {
// 				solution = deepCopy(shortestWithWait)
// 			}
// 			moh := len(longestSolution)
// 			for len(shortestWithWait[len(shortestWithWait)-1]) <= len(longestSolutionWithWait[len(shortestWithWait)-1]) {
// 				longestSolutionWithWait = deepCopy(shortestWithWait)
// 				copy(shortestWithoutWait, recur(shortestWithoutWait, path[i], number))
// 				copy(shortestWithWait, modifyCommonRoom(deepCopy(shortestWithoutWait)))
// 				if len(shortestWithWait[len(shortestWithWait)-1]) < len(solution[len(shortestWithWait)-1]) {
// 					solution = deepCopy(shortestWithWait)
// 					soz = deepCopy(shortestWithoutWait)
// 				}
// 				moh--
// 				if moh == 0 {
// 					break
// 				}
// 			}
// 			shortest = deepCopy(shortestWithoutWait)
// 		} else {
// 			break
// 		}
// 		counter--
// 		if counter == 0 {
// 			break
// 		}
// 	}
// 	if len(soz) == 0 {
// 		soz = deepCopy(longestSolution)
// 	}
// 	shortest = (appendMultipleTimes(soz, apnum-1))[:original]
// 	longestSolution = (appendMultipleTimes(soz, apnum-1))[:original]
// 	solution = modifyCommonRoom(longestSolution)
// 	for i := len(path) - 2; i >= 0; i-- {
// 		counter := len(longestSolution)
// 		longestSolution = deepCopy(shortest)
// 		shortestWithoutWait = deepCopy(shortest)
// 		longestSolutionWithWait = modifyCommonRoom(deepCopy(longestSolution))
// 		copy(shortestWithoutWait, reorderTallest((deepCopy(shortestWithoutWait)), state))
// 		copy(shortestWithoutWait, appendNestedSlices(shortestWithoutWait, mapToNestedArray(state.Many)))
// 		copy(shortestWithoutWait, recur(shortestWithoutWait, path[i], number))
// 		shortestWithWait = modifyCommonRoom(deepCopy(shortestWithoutWait))
// 		if len(shortestWithWait[len(shortestWithWait)-1]) <= len(longestSolutionWithWait[len(shortestWithWait)-1]) {
// 			if len(shortestWithWait[len(shortestWithWait)-1]) < len(solution[len(shortestWithWait)-1]) {
// 				solution = deepCopy(shortestWithWait)
// 			}
// 			moh := len(longestSolution)
// 			for len(shortestWithWait[len(shortestWithWait)-1]) <= len(longestSolutionWithWait[len(shortestWithWait)-1]) {
// 				longestSolutionWithWait = deepCopy(shortestWithWait)
// 				copy(shortestWithoutWait, recur(shortestWithoutWait, path[i], number))
// 				copy(shortestWithWait, modifyCommonRoom(deepCopy(shortestWithoutWait)))
// 				if len(shortestWithWait[len(shortestWithWait)-1]) < len(solution[len(shortestWithWait)-1]) {
// 					solution = deepCopy(shortestWithWait)
// 					soz = deepCopy(shortestWithoutWait)
// 				}
// 				moh--
// 				if moh == 0 {
// 					break
// 				}
// 			}
// 			shortest = deepCopy(shortestWithoutWait)
// 		} else {
// 			break
// 		}
// 		counter--
// 		if counter == 0 {
// 			break
// 		}
// 	}
// 	if len(soz) == 0 {
// 		soz = deepCopy(longestSolution)
// 	}
// 	return modifyCommonRoom((appendMultipleTimes(soz, apnum-1))[:original])
//}

func initializeState() *State {
	return &State{
		Many: make(map[string]int),
	}
}

func adjustAnts(ants int) (int, int) {
	if ants > 400 {
		return modifyNumber(upperClosestDivisibleBy10(ants))
	}
	return ants, ants
}

func adjustPaths(paths [][]string, ant int) [][]string {
	if len(paths) > ant {
		if len(paths) >= 200 {
			paths = paths[:20]
		}
		if len(paths) >= 60 {
			paths = paths[:12]
		}
		if len(paths) >= ant {
			paths = paths[:ant]
		}
	}
	return paths
}

func initializeLongestSolution(paths [][]string, ant int) [][]string {
	var longestSolution [][]string
	for i := 1; i <= ant; i++ {
		longestSolution = append(longestSolution, paths[len(paths)-1])
	}
	return longestSolution
}

func calculateSolutions(paths [][]string, ant int, state *State, end string) [][]string {
	//var soz [][]string
	shortest := initializeLongestSolution(paths, ant)
	number := len(shortest) - 1
	solution := modifyCommonRoom(deepCopy(shortest))

	for i := len(paths) - 2; i >= 0; i-- {
		counter = len(shortest)
		longestSolution := deepCopy(shortest)
		shortestWithoutWait := deepCopy(shortest)
		longestSolutionWithWait := modifyCommonRoom(deepCopy(longestSolution))
		copy(shortestWithoutWait, reorderTallest(deepCopy(shortestWithoutWait), state))
		copy(shortestWithoutWait, appendNestedSlices(shortestWithoutWait, mapToNestedArray(state.Many)))
		copy(shortestWithoutWait, recur(shortestWithoutWait, paths[i], number))
		shortestWithWait := modifyCommonRoom(deepCopy(shortestWithoutWait))

		if len(shortestWithWait[len(shortestWithWait)-1]) <= len(longestSolutionWithWait[len(shortestWithWait)-1]) {
			if len(shortestWithWait[len(shortestWithWait)-1]) < len(solution[len(shortestWithWait)-1]) {
				solution = deepCopy(shortestWithWait)
			}
			moh := len(longestSolution)
			for len(shortestWithWait[len(shortestWithWait)-1]) <= len(longestSolutionWithWait[len(shortestWithWait)-1]) {
				longestSolutionWithWait = deepCopy(shortestWithWait)
				copy(shortestWithoutWait, recur(shortestWithoutWait, paths[i], number))
				copy(shortestWithWait, modifyCommonRoom(deepCopy(shortestWithoutWait)))
				if len(shortestWithWait[len(shortestWithWait)-1]) < len(solution[len(shortestWithWait)-1]) {
					solution = deepCopy(shortestWithWait)
					soz = deepCopy(shortestWithoutWait)
				}
				moh--
				if moh == 0 {
					break
				}
			}
			shortest = deepCopy(shortestWithoutWait)
		} else {
			break
		}
	}
	return solution
}

func Antsop(ants int, paths [][]string, ends string) [][]string {
	state := initializeState()
	original = ants
	ant, apnum = adjustAnts(ants)
	end := ends
	path := deepCopy(paths)
	path = adjustPaths(path, ant)
	longestSolution = initializeLongestSolution(path, ant)
	return calculateSolutions(path, ant, state, end)
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

func sortByLength(sets [][]string) [][]string {
	sort.Slice(sets, func(i, j int) bool {
		return len(sets[i]) < len(sets[j])
	})
	return sets
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

func deepCopy(sets [][]string) [][]string {
	copied := make([][]string, len(sets))
	for i := range sets {
		copied[i] = append([]string(nil), sets[i]...)
	}
	return copied
}

func reorderTallest(sets [][]string, state *State) [][]string {
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

func appendNestedSlices(destination, source [][]string) [][]string {
	destination = append(destination, source...)
	return destination
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

func upperClosestDivisibleBy10(n int) int {
	if n%10 == 0 {
		return n
	}
	return ((n / 10) + 1) * 10
}

func handleFileInput(args []string) (*os.File, error) {
	if len(args) == 2 {
		f, err := os.Open(args[1])
		if err != nil {
			return nil, fmt.Errorf("error opening file: %w", err)
		}
		return f, nil
	}
	return nil, fmt.Errorf("invalid number of arguments")
}

func initial() *State {
	return &State{
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
}

func scanFile(file *os.File, state *State) error {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		variableop(line, state)
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error scanning file: %w", err)
	}
	return nil
}

func calculatePaths(state *State) error {
	startedroom := state.StartEndRooms["start"]
	endedroom := state.StartEndRooms["end"]
	for _, route := range state.Tunnels[startedroom] {
		path := []string{}
		state.Visited[startedroom] = true
		startPathOpt(path, endedroom, route, state)
	}

	state.Paths = sortByLength(getUniqueStringSets(state.Paths))
	solution = equalizeSlices(Antsop(state.Ants, state.Paths, state.End))
	if state.Ants == 0 {
		return fmt.Errorf("error, no Ants found")
	} else if state.Ants > 10000 {
		return fmt.Errorf("error, Ants number is too large")
	}
	// Further processing of the solution can be added here
	return nil
}

func validateTunnels(state *State) error {
	if !roomtunels(state.RoomsMapTunnels, state.Rooms) {
		return fmt.Errorf("tunnels do not connect to a valid room")
	}
	if state.RoomsMapTunnels[state.End] == "" {
		return fmt.Errorf("start and end rooms do not connect")
	}
	if state.StartEndRooms["start"] == "" || state.StartEndRooms["end"] == "" || len(state.StartAndEnd) != 2 {
		return fmt.Errorf("invalid start or end rooms")
	}
	return nil
}

func main() {

	state := initializeState()

	file, err := handleFileInput(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Scan the file and process each line
	if err := scanFile(file, state); err != nil {
		fmt.Println(err)
		return
	}

	// Validate the tunnels and connections
	if err := validateTunnels(state); err != nil {
		fmt.Println(err)
		return
	}

	// Calculate the paths and solutions
	if err := calculatePaths(state); err != nil {
		fmt.Println(err)
		return
	}
}

func variableop(s string, state *State) {
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

func isNumber(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

func contains(slice []string, item string) bool {
	for _, i := range slice {
		if i == item {
			return true
		}
	}
	return false
}

func startPathOpt(path []string, end string, route string, state *State) {
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
