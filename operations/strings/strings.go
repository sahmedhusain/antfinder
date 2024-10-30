package stringoperations

import (
	"fmt"
	structs "lem-in/datastruct"
	"os"
	"sort"
	"strconv"
	"strings"
)

func SortByLength(sets [][]string) [][]string {
	sort.Slice(sets, func(i, j int) bool {
		return len(sets[i]) < len(sets[j])
	})
	return sets
}

func Recur(a [][]string, b []string, number int) [][]string {
	if TestEq(a, b, number) {
		for i := number - 1; i >= 0; i-- {
			if !TestEq(a, b, i) {
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

func TestEq(a [][]string, b []string, number int) bool {
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

func DeepCopy(sets [][]string) [][]string {
	copied := make([][]string, len(sets))
	for i := range sets {
		copied[i] = append([]string(nil), sets[i]...)
	}
	return copied
}

func ReorderTallest(sets [][]string, state *structs.State) [][]string {
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

func MapToNestedArray(m map[string]int) [][]string {
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

func AppendNestedSlices(destination, source [][]string) [][]string {
	destination = append(destination, source...)
	return destination
}

func IdenticalSlices(a, b []string) bool {
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

func ModifyNumber(n int) (j int, i int) {
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

func AppendMultipleTimes(sets [][]string, n int) [][]string {
	result := DeepCopy(sets)
	for j := 0; j < n; j++ {
		result = append(result, sets...)
	}
	return result
}

func UpperClosestDivisibleBy10(n int) int {
	if n%10 == 0 {
		return n
	}
	return ((n / 10) + 1) * 10
}

func VarState(s string, state *structs.State) {
	if s == "" {
	} else if IsNumber(s) && !state.AntsHere {
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
			if !IsNumber(t) {
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
		state.Connect[ReverseHyphenatedString(s)]++
		if state.Connect[s] > 1 || state.Connect[ReverseHyphenatedString(s)] > 1 {
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

func IsNumber(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

func Contains(slice []string, item string) bool {
	for _, i := range slice {
		if i == item {
			return true
		}
	}
	return false
}

func GetUniqueStringSets(sets [][]string) [][]string {
	result := make([][]string, 0, len(sets))
	for _, set := range sets {
		if IsUniqueStringSet(set) {
			result = append(result, set)
		}
	}
	return result
}

func IsUniqueStringSet(set []string) bool {
	seen := make(map[string]bool, len(set))
	for _, s := range set {
		if seen[s] {
			return false
		}
		seen[s] = true
	}
	return true
}

func EqualizeSlices(data [][]string) [][]string {
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

func ReverseHyphenatedString(str string) string {
	parts := strings.Split(str, "-")
	before := ReverseString(parts[0])
	after := ReverseString(parts[1])
	return fmt.Sprintf("%s-%s", after, before)
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
