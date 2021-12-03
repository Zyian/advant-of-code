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

	inputs := strings.Split(string(input), "\n")

	gamma := ""
	epsilon := ""
	for i := 0; i < len(inputs[0]); i++ {
		zeros := 0
		ones := 0
		for j := 0; j < len(inputs); j++ {
			switch inputs[j][i] {
			case '0':
				zeros++
			case '1':
				ones++
			}
		}
		if zeros > ones {
			gamma += "0"
			epsilon += "1"
		} else {
			epsilon += "0"
			gamma += "1"
		}
	}
	fmt.Printf("Gamma: %s Epsilon: %s\n", gamma, epsilon)
	gammaNum, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonNum, _ := strconv.ParseInt(epsilon, 2, 64)
	fmt.Printf("Energy Consumption: %d\n", gammaNum*epsilonNum)

	oxygen := filter(inputs, true)
	fmt.Printf("Oxygen: %#v\n", oxygen)

	co2 := filter(inputs, false)
	fmt.Printf("CO2: %#v\n", co2)
	oxygenNum, _ := strconv.ParseInt(oxygen, 2, 64)
	co2Num, _ := strconv.ParseInt(co2, 2, 64)
	fmt.Printf("Life Support: %d\n", oxygenNum*co2Num)
}

func filter(input []string, mostCommon bool) string {
	r := input
	for i := 0; i < len(input[0]); i++ {
		if len(r) < 2 {
			break
		}
		r = filterOutCommonality(r, i, mostCommon)
	}
	return r[0]
}

func filterOutCommonality(input []string, position int, mostCommonly bool) []string {
	r := []string{}
	mostCommon := countMostCommon(input, position)
	for j := 0; j < len(input); j++ {
		if mostCommon == "2" {
			if string(input[j][position]) == getIntString(mostCommonly) {
				r = append(r, input[j])
			}
		} else {
			if string(input[j][position]) == mostCommon {
				if mostCommonly {
					r = append(r, input[j])
				}
			} else {
				if !mostCommonly {
					r = append(r, input[j])
				}
			}
		}
	}
	return r
}

func getIntString(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

func countMostCommon(input []string, position int) string {
	if position > len(input[0]) {
		return ""
	}
	zeros, ones := 0, 0
	for j := 0; j < len(input); j++ {
		switch input[j][position] {
		case '0':
			zeros++
		case '1':
			ones++
		}
	}
	mostCommon := "0"
	if ones > zeros {
		mostCommon = "1"
	}
	if ones == zeros {
		mostCommon = "2"
	}
	return mostCommon
}
