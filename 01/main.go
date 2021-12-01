package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	// Read our input file.
	depths := []int{}

	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		iv, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		depths = append(depths, iv)
	}

	// See how many times it increases.
	increaseCount := 0

	for i := range depths {

		p := i - 1

		// Set up indexes.
		iOne := i
		iTwo := i + 1
		iThree := i + 2

		pOne := p
		pTwo := p + 1
		pThree := p + 2

		if pOne > -1 && iThree < len(depths) {

			// Sum windows.
			iVal := depths[iOne] + depths[iTwo] + depths[iThree]
			pVal := depths[pOne] + depths[pTwo] + depths[pThree]

			if iVal > pVal {
				increaseCount++
			}
		}
	}

	// Print result.
	fmt.Printf("Depth increased %d times.\n", increaseCount)
}
