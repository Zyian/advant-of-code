package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

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

	var numInputs []int
	for _, n := range strings.Split(string(input), "\n") {
		num, err := strconv.ParseInt(n, 10, 64)
		if err != nil {
			fmt.Printf("cannot parse (%s): %v\n", n, err)
			continue
		}
		numInputs = append(numInputs, int(num))
	}

	fmt.Printf("Number of inputs: %d\n", len(numInputs))
	var i int
	increases := 0
	for i < len(numInputs)-1 {
		if numInputs[i] < numInputs[i+1] {
			increases++
		}
		i++
	}
	fmt.Printf("Number of part 1 increases: %d\n", increases)
	i = 0
	increases = 0
	for i < len(numInputs)-3 {
		if numInputs[i]+numInputs[i+1]+numInputs[i+2] < numInputs[i+1]+numInputs[i+2]+numInputs[i+3] {
			increases++
		}
		i++
	}
	fmt.Printf("Number of part 2 increases: %d\n", increases)
}
