package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if err := part1(); err != nil {
		log.Fatalf(err.Error())
	}
}

func part1() error {
	f, err := os.Open("input.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	return nil
}

func part2() error {
	f, err := os.Open("input.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	return nil
}
