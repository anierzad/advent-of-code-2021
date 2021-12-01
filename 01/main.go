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

		if p > -1 {

			if depths[i] > depths[p] {
				increaseCount++
			}
		}
	}

	// Print result.
	fmt.Printf("Depth increased %d times.\n", increaseCount)
}
