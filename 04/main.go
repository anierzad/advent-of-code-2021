package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type BingoBoard struct {
	lines [5][5]int
	marks [5][5]bool
}

func (b *BingoBoard) MarkNumber(num int) {

	for l, numLine := range b.lines {
		for c, val := range numLine {
			
			// Does it match?
			if val == num {
				b.marks[l][c] = true
			}
		}
	}
}

func (b BingoBoard) Winner() bool {

	rows := []bool{true, true, true, true, true}
	cols := []bool{true, true, true, true, true}

	// Row loop.
	for r := 0; r < 5; r++ {

		// Column loop.
		for c :=0 ; c < 5; c++ {

			// If not marked.
			if !b.marks[r][c] {

				rows[r] = false
				cols[c] = false
			}
		}
	}

	// Check for row win.
	for _, row := range rows {
		if row {
			return true
		}
	}

	// Check for col win.
	for _, col := range cols {
		if col {
			return true
		}
	}

	return false
}

func (b BingoBoard) Score(num int) int {

	total := 0

	// Row loop.
	for r := 0; r < 5; r++ {

		// Column loop.
		for c :=0 ; c < 5; c++ {

			// If not marked.
			if !b.marks[r][c] {

				total = total + b.lines[r][c]
			}
		}
	}

	return total * num
}

func main() {

	// Read our input file.
	numbers := make([]int, 0)
	boards := make([]*BingoBoard, 0)

	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	// First line holds drawn numbers information.
	if scanner.Scan() {
		nl := scanner.Text()

		numbers = append(numbers, strToArray(nl)...)
	}

	// Skip blank line.
	scanner.Scan()

	// Build our bingo boards.
	lineIndex := 0;
	board := &BingoBoard{}
	for scanner.Scan() {

		l := scanner.Text()
		if len(l) == 0 {

			boards = append(boards, board)

			lineIndex = 0
			board = &BingoBoard{}

			continue
		}

		for i, v := range strToArray(l) {
			board.lines[lineIndex][i] = v
		}

		lineIndex++
	}

	// Play the game.
	winners := make([]*BingoBoard, 0)
	playedNumbers := make([]int, 0)

	for _, number := range numbers {

		for _, board := range boards {

			board.MarkNumber(number)

			if board.Winner() {
				winners = append(winners, board)
			}
		}

		playedNumbers = append(playedNumbers, number)

		if len(winners) > 0 {
			break
		}
	}

	lastNum := playedNumbers[len(playedNumbers) - 1]

	for i, winner := range winners {

		fmt.Printf("Winner #%d score: %d\n", i, winner.Score(lastNum))
	}
}

func strToArray(s string) []int {

	numbers := make([]int, 0)

	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, " ", ",")
	s = strings.ReplaceAll(s, ",,", ",")

	for _, n := range strings.Split(s, ",") {

		i, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal("Failed to convert number string to integer.")
		}
		numbers = append(numbers, i)
	}

	return numbers
}
