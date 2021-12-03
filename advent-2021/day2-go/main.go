package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Directions struct {
	Direction string
	Distance  int
}

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("unexpected err opening file: ", err)
		os.Exit(1)
	}
	input, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("err reading file: ", err)
		os.Exit(1)
	}

	var numInputs []Directions
	for _, n := range strings.Split(string(input), "\n") {
		parts := strings.Split(n, " ")
		if len(parts) != 2 {
			continue
		}
		num, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			fmt.Printf("cannot parse (%s): %v\n", n, err)
			continue
		}
		numInputs = append(numInputs, Directions{parts[0], int(num)})
	}

	horizontal, vertical := 0, 0
	for _, d := range numInputs {
		switch d.Direction {
		case "forward":
			horizontal += d.Distance
		case "up":
			vertical -= d.Distance
		case "down":
			vertical += d.Distance
		}
	}
	fmt.Printf("Part 1: Horizontal * Vertical: %d\n", horizontal*vertical)

	horizontal, vertical, aim := 0, 0, 0
	for _, d := range numInputs {
		switch d.Direction {
		case "forward":
			horizontal += d.Distance
			vertical += aim * d.Distance
		case "up":
			aim -= d.Distance
		case "down":
			aim += d.Distance
		}
	}
	fmt.Printf("Part 2: Horizontal * Vertical: %d\n", horizontal*vertical)
}
