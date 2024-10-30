package rooms

import (
	stringoperations "lem-in/operations/strings"
)

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
