package functions

import (
	s "lem-in/datastruct"
)

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
var state *s.State
var original int

func Antsop(ants int, paths [][]string, ends string) [][]string {
	// Initialize state and set initial ant counts
	state := &s.State{Many: make(map[string]int)}
	original := ants
	var ant, apnum int

	// Adjust ant numbers if they exceed 400
	if ants > 400 {
		ant, apnum = modifyNumber(upperClosestDivisibleBy10(ants))
	} else {
		ant = ants
		apnum = ants
	}

	// Deep copy the paths and adjust their length based on ant count
	path := deepCopy(paths)
	if len(path) > ant {
		if len(path) >= 200 {
			path = path[:20]
		}
		if len(path) >= 60 {
			path = path[:12]
		}
		if len(path) >= ant {
			path = path[:ant]
		}
	}

	// Initialize the longest solution with the last element of path
	longestSolution := make([][]string, ant)
	for i := range longestSolution {
		longestSolution[i] = path[len(path)-1]
	}

	var soz [][]string
	// Modify the common room in the longest solution
	solution := modifyCommonRoom(deepCopy(longestSolution))
	number := len(longestSolution) - 1

	// Iterate through the paths in reverse to find the shortest and longest solutions
	for i := len(path) - 2; i >= 0; i-- {
		counter := len(longestSolution)
		shortestWithoutWait := deepCopy(longestSolution)
		longestSolutionWithWait := modifyCommonRoom(deepCopy(longestSolution))

		copy(shortestWithoutWait, reorderTallest(shortestWithoutWait, state))
		copy(shortestWithoutWait, appendNestedSlices(shortestWithoutWait, mapToNestedArray(state.Many)))
		copy(shortestWithoutWait, recur(shortestWithoutWait, path[i], number))

		shortestWithWait := modifyCommonRoom(deepCopy(shortestWithoutWait))

		if len(shortestWithWait[len(shortestWithWait)-1]) <= len(longestSolutionWithWait[len(shortestWithWait)-1]) {
			if len(shortestWithWait[len(shortestWithWait)-1]) < len(solution[len(shortestWithWait)-1]) {
				solution = deepCopy(shortestWithWait)
			}
			moh := len(longestSolution)
			for len(shortestWithWait[len(shortestWithWait)-1]) <= len(longestSolutionWithWait[len(shortestWithWait)-1]) {
				longestSolutionWithWait = deepCopy(shortestWithWait)
				copy(shortestWithoutWait, recur(shortestWithoutWait, path[i], number))
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
			longestSolution = deepCopy(shortestWithoutWait)
		} else {
			break
		}
		counter--
		if counter == 0 {
			break
		}
	}

	// If soz is empty, set it to the longest solution
	if len(soz) == 0 {
		soz = deepCopy(longestSolution)
	}

	// Return the modified common room solution
	return modifyCommonRoom((appendMultipleTimes(soz, apnum-1))[:original])
}
