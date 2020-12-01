package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		logrus.WithField("err", err).Fatalf("unable to read input")
	}

	intInputs := []int{}
	for _, v := range strings.Split(string(input), "\n") {
		if v == "" {
			continue
		}
		i, err := strconv.Atoi(v)
		if err != nil {
			logrus.WithField("err", err).Errorf("not convertible value")
		}
		intInputs = append(intInputs, i)
	}
	logrus.Infof("Using %d inputs", len(intInputs))
Outer:
	for i := 0; i < len(intInputs)-1; i++ {
		for j := 1; j < len(intInputs); j++ {
			if intInputs[i]+intInputs[j] == 2020 {
				logrus.Infof("%d+%d=2020; %d*%d = %d", intInputs[i], intInputs[j], intInputs[i], intInputs[j], intInputs[i]*intInputs[j])
				break Outer
			}
		}
	}

	for i := 0; i < len(intInputs)-2; i++ {
		for j := 1; j < len(intInputs)-1; j++ {
			for k := 2; k < len(intInputs); k++ {
				if intInputs[i]+intInputs[j]+intInputs[k] == 2020 {
					logrus.Infof("%d+%d+%d=2020; %d*%d*%d = %d", intInputs[i], intInputs[j], intInputs[k], intInputs[i], intInputs[j], intInputs[k], intInputs[i]*intInputs[j]*intInputs[k])
					os.Exit(0)
				}
			}
		}
	}

}
