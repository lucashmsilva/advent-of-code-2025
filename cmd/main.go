package main

import (
	"flag"
	"fmt"
	"reflect"

	"github.com.br/lucashmsilva/advent-of-code-2025/solutions/day1"
)

type Solution interface {
	Part1(inputPath string)
	Part2(inputPath string)
}

type flags struct {
	day       int
	part      int
	inputPath string
}

func main() {
	solutionRegistry := map[int]Solution{
		1: day1.SolutionForDay{},
	}

	f := parseFlags()
	s, ok := solutionRegistry[f.day]

	if !ok {
		panic(fmt.Sprintf("solution %d not implemented", f.day))
	}

	runSolution(f, s)
}

func parseFlags() flags {
	dayPtr := flag.Int("day", 1, "day")
	partPtr := flag.Int("part", 1, "problem part")
	var inputPath string
	flag.StringVar(&inputPath, "input", "", "input file path for the problem")

	flag.Parse()

	if *dayPtr < 0 || *dayPtr > 12 {
		panic("day should be > 0 and <= 12")
	}

	if *partPtr < 0 || *partPtr > 2 {
		panic("part should be > 0 and <= 2")
	}

	if inputPath == "" {
		panic("input path should not be empty")
	}

	return flags{
		day:       *dayPtr,
		part:      *partPtr,
		inputPath: inputPath,
	}
}

func runSolution(f flags, s Solution) {
	st := reflect.TypeOf(s)
	m, _ := st.MethodByName(fmt.Sprintf("Part%d", f.part))

	m.Func.Call([]reflect.Value{
		reflect.ValueOf(s),
		reflect.ValueOf(f.inputPath),
	})
}
