package day4

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type SolutionForDay struct{}

func (SolutionForDay) Part1(inputPath string) {
	fmt.Println("Part 1")

	inputContent, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	var result int
	rollsMap := strings.Split(string(inputContent), "\n")
	for i := range rollsMap {
		for j := range rollsMap[0] {
			if rollsMap[i][j] == '@' && checkNeighbors(i, j, rollsMap) {
				result++
			}
		}
	}

	fmt.Println(result)
}

func (SolutionForDay) Part2(inputPath string) {
	fmt.Println("Part 2")

	inputContent, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	var result int
	rollsMap := strings.Split(string(inputContent), "\n")
	for previousResult := 1; previousResult != result; {
		previousResult = result
		for i := range rollsMap {
			for j := range rollsMap[0] {
				if rollsMap[i][j] == '@' && checkNeighbors(i, j, rollsMap) {
					result++
					rollsMap[i] = rollsMap[i][:j] + "." + rollsMap[i][j+1:]
				}
			}
		}
	}

	fmt.Println(result)
}

func checkNeighbors(x, y int, rollsMap []string) bool {
	var adjacentRolls int
	if x+1 < len(rollsMap) && rollsMap[x+1][y] == '@' {
		adjacentRolls++
	}

	if x-1 >= 0 && rollsMap[x-1][y] == '@' {
		adjacentRolls++
	}

	if y+1 < len(rollsMap[0]) && rollsMap[x][y+1] == '@' {
		adjacentRolls++
	}

	if y-1 >= 0 && rollsMap[x][y-1] == '@' {
		adjacentRolls++
	}

	if x+1 < len(rollsMap) && y+1 < len(rollsMap[0]) && rollsMap[x+1][y+1] == '@' {
		adjacentRolls++
	}

	if x+1 < len(rollsMap) && y-1 >= 0 && rollsMap[x+1][y-1] == '@' {
		adjacentRolls++
	}

	if x-1 >= 0 && y+1 < len(rollsMap[0]) && rollsMap[x-1][y+1] == '@' {
		adjacentRolls++
	}

	if y-1 >= 0 && x-1 >= 0 && rollsMap[x-1][y-1] == '@' {
		adjacentRolls++
	}

	return adjacentRolls < 4
}
