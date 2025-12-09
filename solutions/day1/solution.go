package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type SolutionForDay struct{}

const nDials = 100
const startingPos = 50

func (SolutionForDay) Part1(inputPath string) {
	fmt.Println("Part 1")

	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var position int = startingPos
	var zeroPositions int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		direction := line[0]
		distance, _ := strconv.Atoi(line[1:])

		if direction == 'R' {
			position = (position + distance) % nDials
		} else {
			position = ((position-distance)%nDials + nDials) % nDials
		}

		if position == 0 {
			zeroPositions++
		}

		// fmt.Printf("The dial is rotated %s to point at %d.\n", line, int(position))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(zeroPositions)
}

func (SolutionForDay) Part2(inputPath string) {
	fmt.Println("Part 2")

	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var position int = startingPos
	var totalZeroClicks int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var clickedZero int
		line := scanner.Text()

		direction := line[0]
		distance, _ := strconv.Atoi(line[1:])

		if direction == 'R' {
			clickedZero += (position + distance) / nDials
			position = (position + distance) % nDials
		} else {
			if position == 0 {
				clickedZero += distance / 100
			} else if distance >= position {
				clickedZero += 1 + (distance-position)/nDials
			}
			position = ((position-distance)%nDials + nDials) % nDials
		}

		fmt.Printf("The dial is rotated %s to point at %d", line, int(position))
		if clickedZero > 0 {
			fmt.Printf("; during this rotation, it points at 0 %d times", clickedZero)
			totalZeroClicks += clickedZero
		}
		fmt.Printf(".\n")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(totalZeroClicks)
}
