package fileoperations

import (
	"bufio"
	"fmt"
	structs "lem-in/datastruct"
	stringoperations "lem-in/operations/strings"
	"os"
)

func ReadFile(FilePath string, state *structs.State) {

	f, err := os.Open(FilePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		stringoperations.VarState(line, state)
	}
}
