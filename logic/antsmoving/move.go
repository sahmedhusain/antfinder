package move

import (
	structs "lem-in/datastruct"
	rooms "lem-in/logic/rooms"
	stringoperations "lem-in/operations/strings"
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
var state *structs.State
var original int

func MoveLogic(ants int, paths [][]string, ends string) [][]string {
	state = &structs.State{
		Many: make(map[string]int),
	}
	original = ants
	if ants > 400 {
		ant, apnum = stringoperations.ModifyNumber(stringoperations.UpperClosestDivisibleBy10(ants))
	} else {
		ant = ants
		apnum = ants
	}
	end = ends
	path = stringoperations.DeepCopy(paths)
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
	for i := 1; i <= ant; i++ {
		longestSolution = append(longestSolution, path[len(path)-1])
	}
	var soz [][]string
	shortest = longestSolution
	number := len(longestSolution) - 1
	solution = rooms.ModifyCommonRoom(stringoperations.DeepCopy(longestSolution), end)
	for i := len(path) - 2; i >= 0; i-- {
		counter := len(longestSolution)
		longestSolution = stringoperations.DeepCopy(shortest)
		shortestWithoutWait = stringoperations.DeepCopy(shortest)
		longestSolutionWithWait = rooms.ModifyCommonRoom(stringoperations.DeepCopy(longestSolution), end)
		copy(shortestWithoutWait, stringoperations.ReorderTallest((stringoperations.DeepCopy(shortestWithoutWait)), state))
		copy(shortestWithoutWait, stringoperations.AppendNestedSlices(shortestWithoutWait, stringoperations.MapToNestedArray(state.Many)))
		copy(shortestWithoutWait, stringoperations.Recur(shortestWithoutWait, path[i], number))
		shortestWithWait = rooms.ModifyCommonRoom(stringoperations.DeepCopy(shortestWithoutWait), end)
		if len(shortestWithWait[len(shortestWithWait)-1]) <= len(longestSolutionWithWait[len(shortestWithWait)-1]) {
			if len(shortestWithWait[len(shortestWithWait)-1]) < len(solution[len(shortestWithWait)-1]) {
				solution = stringoperations.DeepCopy(shortestWithWait)
			}
			moh := len(longestSolution)
			for len(shortestWithWait[len(shortestWithWait)-1]) <= len(longestSolutionWithWait[len(shortestWithWait)-1]) {
				longestSolutionWithWait = stringoperations.DeepCopy(shortestWithWait)
				copy(shortestWithoutWait, stringoperations.Recur(shortestWithoutWait, path[i], number))
				copy(shortestWithWait, rooms.ModifyCommonRoom(stringoperations.DeepCopy(shortestWithoutWait), end))
				if len(shortestWithWait[len(shortestWithWait)-1]) < len(solution[len(shortestWithWait)-1]) {
					solution = stringoperations.DeepCopy(shortestWithWait)
					soz = stringoperations.DeepCopy(shortestWithoutWait)
				}
				moh--
				if moh == 0 {
					break
				}
			}
			shortest = stringoperations.DeepCopy(shortestWithoutWait)
		} else {
			break
		}
		counter--
		if counter == 0 {
			break
		}
	}
	if len(soz) == 0 {
		soz = stringoperations.DeepCopy(longestSolution)
	}
	shortest = (stringoperations.AppendMultipleTimes(soz, apnum-1))[:original]
	longestSolution = (stringoperations.AppendMultipleTimes(soz, apnum-1))[:original]
	solution = rooms.ModifyCommonRoom(longestSolution, end)
	for i := len(path) - 2; i >= 0; i-- {
		counter := len(longestSolution)
		longestSolution = stringoperations.DeepCopy(shortest)
		shortestWithoutWait = stringoperations.DeepCopy(shortest)
		longestSolutionWithWait = rooms.ModifyCommonRoom(stringoperations.DeepCopy(longestSolution), end)
		copy(shortestWithoutWait, stringoperations.ReorderTallest((stringoperations.DeepCopy(shortestWithoutWait)), state))
		copy(shortestWithoutWait, stringoperations.AppendNestedSlices(shortestWithoutWait, stringoperations.MapToNestedArray(state.Many)))
		copy(shortestWithoutWait, stringoperations.Recur(shortestWithoutWait, path[i], number))
		shortestWithWait = rooms.ModifyCommonRoom(stringoperations.DeepCopy(shortestWithoutWait), end)
		if len(shortestWithWait[len(shortestWithWait)-1]) <= len(longestSolutionWithWait[len(shortestWithWait)-1]) {
			if len(shortestWithWait[len(shortestWithWait)-1]) < len(solution[len(shortestWithWait)-1]) {
				solution = stringoperations.DeepCopy(shortestWithWait)
			}
			moh := len(longestSolution)
			for len(shortestWithWait[len(shortestWithWait)-1]) <= len(longestSolutionWithWait[len(shortestWithWait)-1]) {
				longestSolutionWithWait = stringoperations.DeepCopy(shortestWithWait)
				copy(shortestWithoutWait, stringoperations.Recur(shortestWithoutWait, path[i], number))
				copy(shortestWithWait, rooms.ModifyCommonRoom(stringoperations.DeepCopy(shortestWithoutWait), end))
				if len(shortestWithWait[len(shortestWithWait)-1]) < len(solution[len(shortestWithWait)-1]) {
					solution = stringoperations.DeepCopy(shortestWithWait)
					soz = stringoperations.DeepCopy(shortestWithoutWait)
				}
				moh--
				if moh == 0 {
					break
				}
			}
			shortest = stringoperations.DeepCopy(shortestWithoutWait)
		} else {
			break
		}
		counter--
		if counter == 0 {
			break
		}
	}
	if len(soz) == 0 {
		soz = stringoperations.DeepCopy(longestSolution)
	}
	return rooms.ModifyCommonRoom((stringoperations.AppendMultipleTimes(soz, apnum-1))[:original], end)
}
