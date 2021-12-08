package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type DisplayDigit struct {
	output string
}

func (dd DisplayDigit) UniqueSegment() bool {
	if len(dd.output) == 2 ||
		len(dd.output) == 3 ||
		len(dd.output) == 4 ||
		len(dd.output) == 7 {
		return true
	}

	return false
}

func NewDisplayDigit(output string) *DisplayDigit {
	return &DisplayDigit{
		output: output,
	}
}

func main() {

	digits := make([]*DisplayDigit, 0)

	// Read our input file.
	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	// Read in.
	for scanner.Scan() {

		// Get text.
		line := scanner.Text()

		// Split display from inputs.
		splitLine := strings.Split(line, " | ")

		// Split display values.
		splitDisplay := strings.Split(splitLine[1], " ")

		for _, display := range splitDisplay {
			digit := NewDisplayDigit(display)

			digits = append(digits, digit)
		}
	}

	uniqueDigits := 0

	for _, digit := range digits {
		if digit.UniqueSegment() {
			uniqueDigits++
		}
	}

	fmt.Println("Unique digits:", uniqueDigits)
}
