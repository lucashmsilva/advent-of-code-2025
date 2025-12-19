package day5

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type SolutionForDay struct{}

type Range struct {
	start, end int
}

type Ranges []Range

func (a Ranges) Len() int           { return len(a) }
func (a Ranges) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Ranges) Less(i, j int) bool { return a[i].start < a[j].start }
func (a Ranges) Find(itemId int) bool {
	for _, r := range a {
		if itemId >= r.start && itemId <= r.end {
			return true
		}
	}

	return false
}

func (SolutionForDay) Part1(inputPath string) {
	fmt.Println("Part 1")

	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var ranges Ranges
	var freshIngredients int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		rangeStr := strings.Split(line, "-")
		start, _ := strconv.Atoi(rangeStr[0])
		end, _ := strconv.Atoi(rangeStr[1])
		ranges = append(ranges, Range{start, end})
	}

	sort.Sort(ranges)

	for scanner.Scan() {
		line := scanner.Text()
		itemId, _ := strconv.Atoi(line)

		found := ranges.Find(itemId)

		if found {
			freshIngredients++
		}
	}

	fmt.Println(freshIngredients)
}

func (SolutionForDay) Part2(inputPath string) {
	fmt.Println("Part 2")

	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var ranges Ranges
	var compactedRanges Ranges
	var freshIngredients int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		rangeStr := strings.Split(line, "-")
		start, _ := strconv.Atoi(rangeStr[0])
		end, _ := strconv.Atoi(rangeStr[1])
		ranges = append(ranges, Range{start, end})
	}

	sort.Sort(ranges)
	compactedRanges = Ranges{ranges[0]}
	for i := 1; i < ranges.Len(); i++ {
		j := compactedRanges.Len() - 1
		if ranges[i].start <= compactedRanges[j].end {
			compactedRanges[j].end = int(math.Max(float64(ranges[i].end), float64(compactedRanges[j].end)))
		} else {
			compactedRanges = append(compactedRanges, ranges[i])
		}
	}

	for _, r := range compactedRanges {
		freshIngredients += r.end - r.start + 1
	}

	fmt.Println(freshIngredients)
}
