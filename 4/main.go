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

func parseNumberSet(s string) map[int]bool {
	out := map[int]bool{}
	for _, numStr := range strings.Fields(s) {
		n, err := strconv.Atoi(numStr)
		if err == nil {
			out[n] = true
		}
	}
	return out
}

func part1() error {
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
		dividerIndex := strings.Index(line, "|")
		winningNumbers := parseNumberSet(line[colonIndex+1 : dividerIndex])
		ourNumbers := parseNumberSet(line[dividerIndex+1:])
		score := 0
		for n := range ourNumbers {
			if !winningNumbers[n] {
				continue
			}
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}
		sum += score
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

	// Build an array of "scores" - the number of winning numbers in each card
	var cardScores []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		colonIndex := strings.Index(line, ":")
		dividerIndex := strings.Index(line, "|")
		winningNumbers := parseNumberSet(line[colonIndex+1 : dividerIndex])
		ourNumbers := parseNumberSet(line[dividerIndex+1:])
		cardScore := 0
		for n := range ourNumbers {
			if winningNumbers[n] {
				cardScore += 1
			}
		}
		cardScores = append(cardScores, cardScore)
	}

	// For each card, increment the number of copies for each successive card that
	// were won by this card plus any copies of the card we've won so far.
	//
	// Note: once we're done processing a card, no more copies of that card can be
	// won. This is because we're processing cards in order from first to last,
	// and because cards only win copies of subsequent cards. Therefore, once
	// we're done processing a card, we can add to the total.
	sum := 0
	copiesWon := make([]int, len(cardScores))
	for i, score := range cardScores {
		cardCount := 1 /*original*/ + copiesWon[i]
		for j := 0; j < score; j++ {
			copiesWon[i+1+j] += cardCount
		}
		sum += cardCount
	}

	fmt.Println(sum)
	return nil
}
