package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type SolutionForDay struct{}

func (SolutionForDay) Part1(inputPath string) {
	fmt.Println("Part 1")

	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var totalOutput int64
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		bank := scanner.Text()

		var largestSoFar int64

		joltages := strings.Split(bank, "")
		for i := 0; i < len(joltages)-1; i++ {
			for j := i + 1; j < len(joltages); j++ {
				combination, _ := strconv.ParseInt(joltages[i]+joltages[j], 10, 64)
				if combination > largestSoFar {
					largestSoFar = combination
				}

			}
		}

		totalOutput += largestSoFar
	}

	fmt.Println(totalOutput)
}

func (SolutionForDay) Part2(inputPath string) {
	fmt.Println("Part 2")

	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var totalOutput int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		bank := scanner.Text()
		joltages := strings.Split(bank, "")

		maxJoltage, _ := strconv.Atoi(findMaxJoltage(joltages, 12))

		totalOutput += maxJoltage
	}

	fmt.Println(totalOutput)
}

func findMaxJoltage(joltages []string, size int) string {
	if size == 1 {
		return slices.Max(joltages)
	}

	searchSpace := joltages[:len(joltages)-size+1]
	largest := slices.Max(searchSpace)
	maxJoltageIdx := slices.Index(searchSpace, largest)
	nextDigits := findMaxJoltage(joltages[maxJoltageIdx+1:], size-1)

	return largest + nextDigits
}
