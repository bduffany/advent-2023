package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if err := part2(); err != nil {
		log.Fatalf(err.Error())
	}
}

func part1() error {
	f, err := os.Open("input.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	sum := 0
	scanner := bufio.NewScanner(f)

	quantities := map[string]int{"red": 12, "green": 13, "blue": 14}

	for scanner.Scan() {
		line := scanner.Text()

		// Parse game ID
		colonIndex := strings.Index(line, ":")
		gamePrefix := line[:colonIndex]
		gamePrefixParts := strings.Fields(gamePrefix)
		gameIDString := gamePrefixParts[1]
		gameID, err := strconv.Atoi(gameIDString)
		if err != nil {
			return err
		}

		// Parse outcomes
		possible := true
		for _, round := range strings.Split(line[colonIndex+2:], "; ") {
			for _, part := range strings.Split(round, ", ") {
				parts := strings.Split(part, " ")
				countString := parts[0]
				count, err := strconv.Atoi(countString)
				if err != nil {
					return err
				}
				color := parts[1]
				if count > quantities[color] {
					possible = false
					break
				}
			}
			if !possible {
				break
			}
		}
		if possible {
			sum += gameID
		}
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

	sum := 0
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		colonIndex := strings.Index(line, ":")
		maxCounts := map[string]int{"red": 0, "green": 0, "blue": 0}
		for _, round := range strings.Split(line[colonIndex+2:], "; ") {
			for _, part := range strings.Split(round, ", ") {
				parts := strings.Split(part, " ")
				countString := parts[0]
				count, err := strconv.Atoi(countString)
				if err != nil {
					return err
				}
				color := parts[1]
				maxCounts[color] = max(maxCounts[color], count)
			}
		}
		power := 1
		for _, v := range maxCounts {
			power *= v
		}
		sum += power
	}
	fmt.Println(sum)
	return nil

}
