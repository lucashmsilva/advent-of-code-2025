package day2

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type SolutionForDay struct{}

func (SolutionForDay) Part1(inputPath string) {
	fmt.Println("Part 1")

	inputContent, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	var invalidIdsSum int
	ranges := strings.SplitSeq(string(inputContent), ",")

	for r := range ranges {
		rangeLimits := strings.Split(r, "-")
		startId, _ := strconv.Atoi(rangeLimits[0])
		endId, _ := strconv.Atoi(rangeLimits[1])

		for i := startId; i <= endId; i++ {
			strId := strconv.Itoa(i)

			if len(strId)%2 != 0 {
				continue
			}

			idIsInvalid := true
			for j := 0; j < len(strId)/2; j++ {
				if strId[j] != strId[len(strId)/2+j] {
					idIsInvalid = false
					break
				}
			}

			if idIsInvalid {
				invalidIdsSum += i
			}
		}
	}

	fmt.Println(invalidIdsSum)
}

func (SolutionForDay) Part2(inputPath string) {
	fmt.Println("Part 2")
	inputContent, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	var invalidIdsSum int
	ranges := strings.SplitSeq(string(inputContent), ",")

	for r := range ranges {
		rangeLimits := strings.Split(r, "-")
		startId, _ := strconv.Atoi(rangeLimits[0])
		endId, _ := strconv.Atoi(rangeLimits[1])

		for i := startId; i <= endId; i++ {
			strId := strconv.Itoa(i)
			if sequenceIsRepeating(strId, getPivot(len(strId))) {
				invalidIdsSum += i
			}

		}

	}

	fmt.Println(invalidIdsSum)
}

func sequenceIsRepeating(strId string, pivot int) bool {
	if len(strId) == 1 {
		return false
	}

	subStrRepeats := true
	subStr := strId[:pivot]

	for i := len(subStr); i < len(strId); i += len(subStr) {
		if i+len(subStr) > len(strId) {
			subStrRepeats = false
			break
		}

		if subStr != strId[i:i+len(subStr)] {
			subStrRepeats = false
			break
		}
	}

	if !subStrRepeats {
		nextPivot := getPivot(pivot)
		if nextPivot == len(strId) || pivot == nextPivot {
			return false
		}

		return sequenceIsRepeating(strId, nextPivot)
	}

	return true
}

func getPivot(currentPos int) int {
	if currentPos%2 == 0 {
		return currentPos / 2
	}

	return currentPos/2 + 1
}
