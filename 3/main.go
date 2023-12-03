package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type coordinate struct{ row, col int }

var directions = []coordinate{
	// Above
	{-1, -1},
	{-1, 0},
	{-1, 1},
	// Left and right
	{0, -1},
	{0, 1},
	// Below
	{1, -1},
	{1, 0},
	{1, 1},
}

func main() {
	if err := part2(); err != nil {
		log.Fatalf(err.Error())
	}
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isSymbol(c byte) bool {
	return c != '.' && !isDigit(c)
}

func containsNeighboringSymbol(lines []string, row, col int) bool {
	for _, d := range directions {
		c := col + d.col
		if c < 0 || c >= len(lines[row]) {
			continue
		}
		r := row + d.row
		if r < 0 || r >= len(lines) {
			continue
		}
		if isSymbol(lines[r][c]) {
			return true
		}
	}
	return false
}

func neighboringGearCoordinates(lines []string, row, col int) map[coordinate]struct{} {
	gears := map[coordinate]struct{}{}
	for _, d := range directions {
		c := col + d.col
		if c < 0 || c >= len(lines[row]) {
			continue
		}
		r := row + d.row
		if r < 0 || r >= len(lines) {
			continue
		}
		if lines[r][c] == '*' {
			gears[coordinate{r, c}] = struct{}{}
		}
	}
	return gears
}

func part1() error {
	f, err := os.Open("input.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		// Note: if memory were a constraint then we could buffer only 3 lines
		// instead of all of them.
		lines = append(lines, line)
	}

	sum := 0

	curNumber := 0
	isPartNumber := false

	// resetNumber is called whenever we see something that is not a number, as
	// well as after each row. It "consumes" the current part number if it exists,
	// adding it to the total, then resetting the current parsing state.
	resetNumber := func() {
		if isPartNumber {
			sum += curNumber
		}
		curNumber = 0
		isPartNumber = false
	}

	for row, line := range lines {
		for col, c := range []byte(line) {
			if isDigit(c) {
				curNumber = 10*curNumber + int(c-'0')
				if !isPartNumber && containsNeighboringSymbol(lines, row, col) {
					isPartNumber = true
				}
			} else {
				// Reset the number any time we see something that isn't a digit
				resetNumber()
			}
		}
		// Reset the number at the end of each row
		resetNumber()
	}
	fmt.Println(sum)
	return nil
}

func part2() error {
	f, err := os.Open("input.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	gearNeighboringPartNumbers := map[coordinate][]int{}

	curNumber := 0
	neighboringGears := map[coordinate]struct{}{}

	// resetNumber is called whenever we see something that is not a number, as
	// well as after each row. It "consumes" the current number if it exists,
	// adding it to the neighbor list for all of its neighboring gears, then
	// resetting the current parsing state.
	resetNumber := func() {
		for g := range neighboringGears {
			gearNeighboringPartNumbers[g] = append(gearNeighboringPartNumbers[g], curNumber)
		}
		curNumber = 0
		clear(neighboringGears)
	}

	for row, line := range lines {
		for col, c := range []byte(line) {
			if isDigit(c) {
				curNumber = 10*curNumber + int(c-'0')
				// Accumulate the set of neighboring gear coordinates for each
				// consecutive digit in the current part number
				for g := range neighboringGearCoordinates(lines, row, col) {
					neighboringGears[g] = struct{}{}
				}
			} else {
				// Reset the number any time we see something that isn't a digit
				resetNumber()
			}
		}
		// Reset the number at the end of each row
		resetNumber()
	}

	sum := 0
	for _, numbers := range gearNeighboringPartNumbers {
		if len(numbers) != 2 {
			continue
		}
		gearRatio := numbers[0] * numbers[1]
		sum += gearRatio
	}

	fmt.Println(sum)
	return nil
}
