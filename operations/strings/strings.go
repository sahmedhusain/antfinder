package strOps

import (
<<<<<<< Updated upstream
	"fmt"
	structs "lem-in/datastruct"
	"os"
	"sort"
	"strconv"
	"strings"
=======
	"fmt"                       // Package for formatted I/O
	structs "lem-in/datastruct" // Importing custom data structures
	"sort"                      // Package for sorting slices
	"strconv"                   // Package for string conversions
	"strings"                   // Package for string manipulations
>>>>>>> Stashed changes
)

// SortByLength sorts a slice of string slices by their length in ascending order
func SortByLength(sets [][]string) [][]string {
	sort.Slice(sets, func(i, j int) bool { // Sort the slices based on their length
		return len(sets[i]) < len(sets[j])
	})
	return sets // Return the sorted slices
}

<<<<<<< Updated upstream
func Recur(a [][]string, b []string, number int) [][]string {
	if TestEq(a, b, number) {
		for i := number - 1; i >= 0; i-- {
			if !TestEq(a, b, i) {
				a[i] = b
				return a
=======
// Recur recursively updates a slice of string slices based on a condition
func Recur(a [][]string, b []string, number int) [][]string {
	if TestEq(a, b, number) { // Check if the slices are equal
		for i := number - 1; i >= 0; i-- { // Iterate backwards through the slices
			if !TestEq(a, b, i) { // If slices are not equal
				a[i] = b // Update the slice
				return a // Return the updated slices
>>>>>>> Stashed changes
			}
		}
	} else {
		a[number] = b // Update the slice at the given index
		return a      // Return the updated slices
	}
	return a // Return the slices if no updates were made
}

// TestEq checks if two slices of strings are equal
func TestEq(a [][]string, b []string, number int) bool {
	if len(a[number]) != len(b) { // Check if lengths are different
		return false // Return false if lengths are different
	}
	for i := range a[number] { // Iterate through the elements
		if a[number][i] != b[i] { // Check if elements are different
			return false // Return false if elements are different
		}
	}
	return true // Return true if all elements are equal
}

// DeepCopy creates a deep copy of a slice of string slices
func DeepCopy(sets [][]string) [][]string {
	copied := make([][]string, len(sets)) // Create a new slice with the same length
	for i := range sets {                 // Iterate through the original slices
		copied[i] = append([]string(nil), sets[i]...) // Copy each slice
	}
	return copied // Return the copied slices
}

// ReorderTallest reorders a slice of string slices based on their length and updates the state
func ReorderTallest(sets [][]string, state *structs.State) [][]string {
	if len(sets) == 0 { // Check if the input slice is empty
		return sets // Return the empty slice
	}
	pattern := sets                // Assign the input slice to a new variable
	number := len(pattern) - 1     // Get the last index of the slice
	length := len(pattern[number]) // Get the length of the last slice
	for i := number; i >= 0; i-- { // Iterate backwards through the slices
		if len(pattern[i]) == length { // Check if the lengths are equal
			state.Many[strings.Join(pattern[i], "&*")]++                         // Update the state
			fmt.Printf("\nState after updating Many at index %d: %+v\n", i, state) // Debug print the state
		} else {
			fmt.Printf("\nState before returning pattern at index %d: %+v\n", i, state) // Debug print the state
			return pattern[:number+1]                                                 // Return the reordered slices
		}
	}
	fmt.Printf("\nState before returning empty slice: %+v\n", state) // Debug print the state
	return [][]string{}                                            // Return an empty slice if no reordering was done
}

// MapToNestedArray converts a map of strings to integers into a nested slice of strings
func MapToNestedArray(m map[string]int) [][]string {
	var result [][]string           // Initialize the result slice
	tempMap := make(map[string]int) // Create a temporary map
	for k, v := range m {           // Copy the input map to the temporary map
		tempMap[k] = v
	}
	for len(tempMap) > 0 { // Iterate until the temporary map is empty
		for k, v := range tempMap { // Iterate through the temporary map
			if v > 0 { // Check if the value is greater than 0
				result = append(result, strings.Split(k, "&*")) // Split the key and append to the result
				tempMap[k] = v - 1                              // Decrement the value
				if tempMap[k] == 0 {                            // Check if the value is 0
					delete(tempMap, k) // Delete the key from the temporary map
				}
			}
		}
	}
	return result // Return the result slice
}

// AppendNestedSlices appends one nested slice of strings to another
func AppendNestedSlices(destination, source [][]string) [][]string {
	destination = append(destination, source...) // Append the source slice to the destination slice
	return destination                           // Return the updated destination slice
}

// IdenticalSlices checks if two slices of strings are identical
func IdenticalSlices(a, b []string) bool {
	if len(a) != len(b) { // Check if lengths are different
		return false // Return false if lengths are different
	}
	for i := range a { // Iterate through the elements
		if a[i] != b[i] { // Check if elements are different
			return false // Return false if elements are different
		}
	}
	return true // Return true if all elements are equal
}

// ModifyNumber modifies a number based on specific conditions and returns two integers
func ModifyNumber(n int) (j int, i int) {
	if n < 400 { // Check if the number is less than 400
		return n, 1 // Return the number and 1
	}
	if n >= 400 && n < 2000 { // Check if the number is between 400 and 2000
		for i := 10; i >= 2; i-- { // Iterate from 10 to 2
			if n%i == 0 { // Check if the number is divisible by i
				return n / i, i // Return the quotient and i
			}
		}
	}
	if n >= 2000 && n <= 5000 { // Check if the number is between 2000 and 5000
		for i := 50; i >= 2; i-- { // Iterate from 50 to 2
			if n%i == 0 { // Check if the number is divisible by i
				return n / i, i // Return the quotient and i
			}
		}
	}
	if n >= 5001 && n <= 20000 { // Check if the number is between 5001 and 20000
		for i := 100; i >= 2; i-- { // Iterate from 100 to 2
			if n%i == 0 { // Check if the number is divisible by i
				return n / i, i // Return the quotient and i
			}
		}
	}
	return n, 1 // Return the number and 1 if no conditions were met
}

// AppendMultipleTimes appends a nested slice of strings to itself multiple times
func AppendMultipleTimes(sets [][]string, n int) [][]string {
	result := DeepCopy(sets) // Create a deep copy of the input slice
	for j := 0; j < n; j++ { // Iterate n times
		result = append(result, sets...) // Append the input slice to the result
	}
	return result // Return the updated result slice
}

// UpperClosestDivisibleBy10 returns the closest number greater than or equal to n that is divisible by 10
func UpperClosestDivisibleBy10(n int) int {
	if n%10 == 0 { // Check if the number is already divisible by 10
		return n // Return the number
	}
	return ((n / 10) + 1) * 10 // Return the next number divisible by 10
}

<<<<<<< Updated upstream
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

=======
// IsNumber checks if a string can be converted to a number
>>>>>>> Stashed changes
func IsNumber(str string) bool {
	_, err := strconv.ParseFloat(str, 64) // Try to parse the string as a float
	return err == nil                     // Return true if no error occurred
}

// Contains checks if a slice contains a specific string
func Contains(slice []string, item string) bool {
	for _, i := range slice { // Iterate through the slice
		if i == item { // Check if the item is found
			return true // Return true if the item is found
		}
	}
	return false // Return false if the item is not found
}

// GetUniqueStringSets returns a slice of unique string slices
func GetUniqueStringSets(sets [][]string) [][]string {
	result := make([][]string, 0, len(sets)) // Initialize the result slice
	for _, set := range sets {               // Iterate through the input slices
		if IsUniqueStringSet(set) { // Check if the slice is unique
			result = append(result, set) // Append the unique slice to the result
		}
	}
	return result // Return the result slice
}

// IsUniqueStringSet checks if a slice of strings contains only unique elements
func IsUniqueStringSet(set []string) bool {
	seen := make(map[string]bool, len(set)) // Initialize a map to track seen elements
	for _, s := range set {                 // Iterate through the elements
		if seen[s] { // Check if the element has been seen before
			return false // Return false if the element is not unique
		}
		seen[s] = true // Mark the element as seen
	}
	return true // Return true if all elements are unique
}

// EqualizeSlices equalizes the length of slices by padding with "wait"
func EqualizeSlices(data [][]string) [][]string {
	maxLen := 0                // Initialize the maximum length
	for _, row := range data { // Iterate through the slices
		if len(row) > maxLen { // Check if the current slice is longer
			maxLen = len(row) // Update the maximum length
		}
	}
	result := make([][]string, len(data)) // Initialize the result slice
	for i, row := range data {            // Iterate through the input slices
		if len(row) < maxLen { // Check if the current slice is shorter
			result[i] = make([]string, maxLen)   // Create a new slice with the maximum length
			copy(result[i], row)                 // Copy the elements to the new slice
			for j := len(row); j < maxLen; j++ { // Pad the remaining elements with "wait"
				result[i][j] = "wait"
			}
		} else {
			result[i] = row // Use the original slice if it is already the maximum length
		}
	}
	return result // Return the result slice
}

// ReverseHyphenatedString reverses the parts of a hyphenated string
func ReverseHyphenatedString(str string) string {
	parts := strings.Split(str, "-")           // Split the string by hyphen
	before := ReverseString(parts[0])          // Reverse the first part
	after := ReverseString(parts[1])           // Reverse the second part
	return fmt.Sprintf("%s-%s", after, before) // Format and return the reversed string
}

// ReverseString reverses a string
func ReverseString(s string) string {
	runes := []rune(s)                                    // Convert the string to a slice of runes
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 { // Swap the runes
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes) // Convert the runes back to a string and return
}
