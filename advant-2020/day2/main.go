package main

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {
	i := readInput()
	validPwds := 0
	validPwds2 := 0
	for _, l := range strings.Split(string(i), "\n") {
		lparts := strings.Split(l, " ")
		minmax := strings.Split(lparts[0], "-")
		min, err := strconv.Atoi(minmax[0])
		if err != nil {
			logrus.WithField("err", err).Errorf("Unable to parse min")
			continue
		}
		max, err := strconv.Atoi(minmax[1])
		if err != nil {
			logrus.WithField("err", err).Errorf("Unable to parse max")
			continue
		}
		char := strings.TrimSuffix(lparts[1], ":")
		countSet := strings.Count(lparts[2], char)
		if countSet >= min && countSet <= max {
			validPwds++
		}
		if xorCheck(string(lparts[2][min-1]) == char, string(lparts[2][max-1]) == char) {
			validPwds2++
		}
	}
	logrus.Infof("There are %d passwords for part 1\n", validPwds)
	logrus.Infof("There are %d passwords for part 2\n", validPwds2)
}

func readInput() []byte {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		logrus.WithField("err", err).Fatalf("unable to read input")
	}
	return input
}

func xorCheck(X, Y bool) bool {
	return (X || Y) && !(X && Y)
}
