package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type DiagnosticReport struct {
	lines []string
}

func (dr DiagnosticReport) GammaRate() int {

	gamma := ""

	for i := 0; i < 12; i++ {

		posRune := '1'
		zeros := bitCount(dr.lines, i)

		if zeros > (len(dr.lines) / 2) {
			posRune = '0'
		}

		gamma = gamma + string(posRune)
	}

	// Convert to integer.
	iv, err := strconv.ParseInt(gamma, 2, 0)
	if err != nil {
		log.Fatal(err)
	}

	return int(iv)
}

func (dr DiagnosticReport) EpsilonRate() int {

	epsilon := ""

	for i := 0; i < 12; i++ {

		posRune := '1'
		zeros := bitCount(dr.lines, i)

		if zeros < (len(dr.lines) / 2) {
			posRune = '0'
		}

		epsilon = epsilon + string(posRune)
	}

	// Convert to integer.
	iv, err := strconv.ParseInt(epsilon, 2, 0)
	if err != nil {
		log.Fatal(err)
	}

	return int(iv)
}

func (dr DiagnosticReport) PowerConsumption() int {

	ga := dr.GammaRate()
	ep := dr.EpsilonRate()

	return ga * ep
}

func (dr DiagnosticReport) OxygenGeneratorRating() int {

	filtered := dr.lines

	for i := 0; i < 12; i++ {

		if len(filtered) < 2 {
			break
		}

		tb := '1'

		if bitCount(filtered, i) > (len(filtered) / 2) {
			tb = '0'
		}

		filtered = bitFilter(filtered, i, tb)
	}

	// Convert to integer.
	iv, err := strconv.ParseInt(filtered[0], 2, 0)
	if err != nil {
		log.Fatal(err)
	}

	return int(iv)
}

func (dr DiagnosticReport) CO2ScrubberRating() int {

	filtered := dr.lines

	for i := 0; i < 12; i++ {

		if len(filtered) < 2 {
			break
		}

		tb := '1'

		if bitCount(filtered, i) <= (len(filtered) / 2) {
			tb = '0'
		}

		filtered = bitFilter(filtered, i, tb)
	}

	// Convert to integer.
	iv, err := strconv.ParseInt(filtered[0], 2, 0)
	if err != nil {
		log.Fatal(err)
	}

	return int(iv)
}

func (dr DiagnosticReport) LifeSupportRating() int {

	og := dr.OxygenGeneratorRating()
	cs := dr.CO2ScrubberRating()

	return og * cs
}

func (dr DiagnosticReport) Print() {

	fmt.Printf("Gamma Rate:................%d\n", dr.GammaRate())
	fmt.Printf("Epsilon Rate:..............%d\n", dr.EpsilonRate())
	fmt.Printf("Power Consumption:.........%d\n", dr.PowerConsumption())
	fmt.Printf("Oxygen Generator Rating:...%d\n", dr.OxygenGeneratorRating())
	fmt.Printf("CO2 Scrubber Rating:.......%d\n", dr.CO2ScrubberRating())
	fmt.Printf("Life Support Rating:.......%d\n", dr.LifeSupportRating())
}

func bitCount(lines []string, bit int) int {

	zeros := 0

	for _, line := range lines {
		if line[bit] == '0' {
			zeros++
		}
	}

	return zeros
}

func bitFilter(lines []string, bit int, tb rune) []string {

	filtered := []string{}

	for _, line := range lines {

		if rune(line[bit]) == tb {
			filtered = append(filtered, line)
		}
	}

	return filtered
}

func main() {

	// Read our inut file.
	diagnostics := DiagnosticReport{}

	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		diagnostics.lines = append(diagnostics.lines, scanner.Text())
	}

	diagnostics.Print()
}
