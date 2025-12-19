package main

import (
	"flag"
	"fmt"

	"github.com.br/lucashmsilva/advent-of-code-2025/solutions/day1"
	"github.com.br/lucashmsilva/advent-of-code-2025/solutions/day2"
	"github.com.br/lucashmsilva/advent-of-code-2025/solutions/day3"
	"github.com.br/lucashmsilva/advent-of-code-2025/solutions/day4"
	"github.com.br/lucashmsilva/advent-of-code-2025/solutions/day5"
)

type Solution interface {
	Part1(inputPath string)
	Part2(inputPath string)
}

type flags struct {
	day       int
	inputPath string
}

func main() {
	solutionRegistry := map[int]Solution{
		1: day1.SolutionForDay{},
		2: day2.SolutionForDay{},
		3: day3.SolutionForDay{},
		4: day4.SolutionForDay{},
		5: day5.SolutionForDay{},
	}

	f := parseFlags()
	s, ok := solutionRegistry[f.day]

	if !ok {
		panic(fmt.Sprintf("solution %d not implemented", f.day))
	}

	s.Part1(f.inputPath)
	s.Part2(f.inputPath)
}

func parseFlags() flags {
	dayPtr := flag.Int("day", 1, "day")
	var inputPath string
	flag.StringVar(&inputPath, "input", "", "input file path for the problem")

	flag.Parse()

	if *dayPtr < 0 || *dayPtr > 12 {
		panic("day should be > 0 and <= 12")
	}

	if inputPath == "" {
		panic("input path should not be empty")
	}

	return flags{
		day:       *dayPtr,
		inputPath: inputPath,
	}
}
