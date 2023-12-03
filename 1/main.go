package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	// if err := part1(); err != nil {
	// 	log.Fatalf(err.Error())
	// }
	if err := part2(); err != nil {
		log.Fatalf(err.Error())
	}
}

func parseDigit(d byte) int {
	return int(d) - '0'
}

func part1() error {
	f, err := os.Open("input.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		var firstDigit, lastDigit byte
		for _, c := range []byte(line) {
			if c < '1' || c > '9' {
				continue
			}
			if firstDigit == 0 {
				firstDigit = c
			}
			lastDigit = c
		}
		calibrationValue := parseDigit(firstDigit)*10 + parseDigit(lastDigit)
		sum += calibrationValue
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

	// Note: we can't use regexp.FindAllString, since for an input of
	// "oneight" it just returns "one", but should return "one" and "eight".
	digitAtStartRegex := regexp.MustCompile(`^([1-9]|one|two|three|four|five|six|seven|eight|nine)`)
	digitValue := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,

		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		var firstValue, lastValue int
		for i := 0; i < len(line); i++ {
			match := digitAtStartRegex.FindString(line[i:])
			if match == "" {
				continue
			}

			value := digitValue[match]
			if firstValue == 0 {
				firstValue = value
			}
			lastValue = value
		}
		calibrationValue := firstValue*10 + lastValue
		sum += calibrationValue
	}
	fmt.Println(sum)
	return nil
}
