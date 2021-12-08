package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var definitions = map[int][]rune{
	0: {'a','b','c','e','f','g'},
	1: {'c','f'},
	2: {'a','c','d','e','g'},
	3: {'a','c','d','f','g'},
	4: {'b','c','d','f'},
	5: {'a','b','d','f','g'},
	6: {'a','b','d','e','f','g'},
	7: {'a','c','f'},
	8: {'a','b','c','d','e','f','g'},
	9: {'a','b','c','d','f','g'},
}

type DisplayDigit struct {
	output string
	segments map[rune]bool
}

func (dd DisplayDigit) Value() int {
	longest := 0
	value := -1

	for v, definition := range definitions {
		match := true

		for _, r := range definition {
			if !dd.segments[r] {
				match = false
				break
			}
		}

		if match {
			length := len(definition)
			if length > longest {
				longest = len(definition)
				value = v
			}
		}
	}

	return value
}

func (dd DisplayDigit) Print() [][]rune {
	lines := make([][]rune, 7)

	for i := range lines {
		line := make([]rune, 6)

		for p := range line {
			line[p] = ' '
		}

		switch i {
		case 0:
			if dd.segments['a'] {
				line[1] = 'a'
				line[2] = 'a'
				line[3] = 'a'
				line[4] = 'a'
			}
		case 1, 2:
			if dd.segments['b'] {
				line[0] = 'b'
			}
			if dd.segments['c'] {
				line[5] = 'c'
			}
		case 3:
			if dd.segments['d'] {
				line[1] = 'd'
				line[2] = 'd'
				line[3] = 'd'
				line[4] = 'd'
			}
		case 4, 5:
			if dd.segments['e'] {
				line[0] = 'e'
			}
			if dd.segments['f'] {
				line[5] = 'f'
			}
		case 6:
			if dd.segments['g'] {
				line[1] = 'g'
				line[2] = 'g'
				line[3] = 'g'
				line[4] = 'g'
			}
		}

		lines[i] = line
	}

	return lines
}

func NewDisplayDigit(output string) *DisplayDigit {
	dd := &DisplayDigit{
		output: output,
		segments: make(map[rune]bool),
	}

	for _, r := range output {
		dd.segments[r] = true
	}

	return dd
}

type Display struct {
	digits []*DisplayDigit
	signals []string
}

func (d Display) Print() {

	lines := make([]string, 7)

	for _, digit := range d.digits {
		for ln, line := range digit.Print() {
			if lines[ln] == "" {
				lines[ln] = string(line)
				continue
			}
			lines[ln] = fmt.Sprintf("%s %s", lines[ln], string(line))
		}
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}

func (d Display) Value() int {
	str := ""

	for _, digit := range d.digits {
		str = fmt.Sprintf("%s%d", str, digit.Value())
	}

	value, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal("Failed to convert string to int.")
	}

	return value
}

func NewDisplay(signals, outputs []string) *Display {

	d := &Display{
		digits:  make([]*DisplayDigit, 0),
		signals: signals,
	}

	// We need to work out what segments map to which signals.
	sigOne := stringWithLength(signals, len(definitions[1]))
	sigFour := stringWithLength(signals, len(definitions[4]))
	sigSeven := stringWithLength(signals, len(definitions[7]))
	sigEight := stringWithLength(signals, len(definitions[8]))

	// Make a mapping.
	mapping := make(map[rune]rune)

	// Process one.
	for _, r := range sigOne {
		count := runeOccourances(signals, r)

		if count == 8 {
			mapping[r] = 'c'
		}

		if count == 9 {
			mapping[r] = 'f'
		}
	}

	// Process four.
	for _, r := range sigFour {

		// Skip what we've already got.
		if _, exists := mapping[r]; exists {
			continue
		}

		count := runeOccourances(signals, r)

		if count == 6 {
			mapping[r] = 'b'
		}

		if count == 7 {
			mapping[r] = 'd'
		}
	}

	// Process seven.
	for _, r := range sigSeven {

		// Skip what we've already got.
		if _, exists := mapping[r]; exists {
			continue
		}

		count := runeOccourances(signals, r)

		if count == 8 {
			mapping[r] = 'a'
		}
	}

	// Process eight.
	for _, r := range sigEight {

		// Skip what we've already got.
		if _, exists := mapping[r]; exists {
			continue
		}

		count := runeOccourances(signals, r)

		if count == 4 {
			mapping[r] = 'e'
		}

		if count == 7 {
			mapping[r] = 'g'
		}
	}

	// Create display.
	for _, output := range outputs {

		// Remap!
		remapped := ""

		for _, r := range output {
			remapped = fmt.Sprintf("%s%s", remapped, string(mapping[r]))
		}

		digit := NewDisplayDigit(remapped)

		d.digits = append(d.digits, digit)
	}

	return d
}

func runeOccourances(signals []string, r rune) (count int) {
	for _, signal := range signals {
		if strings.Contains(signal, string(r)) {
			count++
		}
	}

	return count
}

func stringWithLength(strs []string, length int) string {
	for _, s := range strs {
		if len(s) == length {
			return s
		}
	}

	return ""
}

func main() {

	displays := make([]*Display, 0)

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

		signals := strings.Split(splitLine[0], " ")
		outputs := strings.Split(splitLine[1], " ")

		display := NewDisplay(signals, outputs)

		displays = append(displays, display)
	}

	total := 0

	for _, display := range displays {
		total += display.Value()
	}

	fmt.Println("Total:", total)
}
