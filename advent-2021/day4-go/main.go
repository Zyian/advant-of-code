package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/Zyian/aocget"
)

func main() {
	aoc := aocget.NewClient(os.Getenv("SESSION_TOKEN"))
	input, err := aoc.DownloadInputString(2021, 4)
	if err != nil {
		fmt.Printf("unable to download input: %v\n", err)
		os.Exit(1)
	}

	input = strings.ReplaceAll(input, "  ", " ")
	r := regexp.MustCompile("(?m)^ ")

	input = r.ReplaceAllString(input, "")

	winners := strings.Split(strings.Split(input, "\n")[0], ",")
	fmt.Println(winners)

	boardsRaw := strings.Split(input, "\n\n")
	boards := [][][]string{}
	for _, b := range boardsRaw[1 : len(boardsRaw)-1] {
		board := blankGrid(5, 5)
		boardRows := strings.Split(b, "\n")
		for ir, br := range boardRows {
			boardCells := strings.Split(br, " ")
			for ic, c := range boardCells {
				board[ic][ir] = c
			}
		}
		boards = append(boards, board)
	}
	fmt.Printf("%#v\n", boards)
}

func blankGrid(w, h int) [][]string {
	grid := make([][]string, w)
	for i := 0; i < w; i++ {
		grid[i] = make([]string, h)
	}
	return grid
}
