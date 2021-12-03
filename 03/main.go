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
		zeros := 0

		for _, line := range dr.lines {

			if line[i] == '0' {
				zeros++

				if zeros > (len(dr.lines) / 2) {
					posRune = '0'
					break
				}
			}
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
		zeros := len(dr.lines)

		for _, line := range dr.lines {

			if line[i] == '1' {
				zeros--

				if zeros < (len(dr.lines) / 2) {
					posRune = '0'
					break
				}
			}
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

func (dr DiagnosticReport) Print() {

	fmt.Printf("Gamma Rate:..........%d\n", dr.GammaRate())
	fmt.Printf("Epsilon Rate:........%d\n", dr.EpsilonRate())
	fmt.Printf("Power Consumption:...%d\n", dr.PowerConsumption())
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
