package day1

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type SolutionForDay struct{}

const nDials = 100
const startingPos = 50

func (SolutionForDay) Part1(inputPath string) {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	position := float64(startingPos)
	clicksAtZero := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		direction := line[0]
		distance, _ := strconv.ParseFloat(line[1:], 64)

		if direction == 'R' {
			clicksAtZero += (int(position) + int(distance))
			position = float64(int(position+distance) % nDials)
		} else if direction == 'L' {
			distance > position

			position = float64(((int(distance-position) % 100) + 100) % 100)
		}

		position = float64((int((position + distance)) % nDials))
		if position < 0 {
			position += nDials
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
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	position := float64(startingPos)
	zeroPositions := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		direction := line[0]
		distance, _ := strconv.ParseFloat(line[1:], 64)

		if direction == 'L' {
			distance *= -1
		}

		previousPos := position
		movement := position + distance

		position = float64((int((movement)) % nDials))
		if position < 0 {
			position += nDials
		}

		fmt.Printf("The dial is rotated %s to point at %d", line, int(position))

		if math.Abs(movement) > nDials {
			timesClicked := int(math.Abs(movement)) / nDials
			zeroPositions += timesClicked
			fmt.Printf("; during this rotation, it points at 0 %d times", timesClicked)
		} else if movement < 0 && previousPos != 0 && position != 0 {
			zeroPositions++
			fmt.Printf("; during this rotation, it points at 0 1 times")
		}

		if position == 0 {
			zeroPositions++
		}
		fmt.Printf("; we clicked zero %d times so far", zeroPositions)
		fmt.Print(".\n")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(zeroPositions)
}
