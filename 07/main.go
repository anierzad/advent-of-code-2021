package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type ShipPosition struct {
	horizontal int
	count int
}

func (sp *ShipPosition) Add(count int) {
	sp.count += count
}

func (sp ShipPosition) FuelNeeded(pos int) int {
	reqFuel := int(math.Abs(float64(sp.horizontal - pos)))

	return reqFuel * sp.count
}

func NewShipPosition(horizontal, count int) *ShipPosition {
	return &ShipPosition{
		horizontal: horizontal,
		count: count,
	}
}

func main() {

	// Read our input file.
	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	// Get input line.
	scanner.Scan()
	line := scanner.Text()

	vals := strings.Split(line, ",")

	// Create positions.
	positions := make(map[int]*ShipPosition)

	for _, val := range vals {

		// Convert string to integer.
		iv, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal("Failed to convert string to int.")
		}

		// Have a position?
		position, exists := positions[iv]
		if !exists {

			// Create new position.
			position = NewShipPosition(iv, 0)

			// Add to map.
			positions[iv] = position
		}

		// Add one to count.
		position.Add(1)
	}

	// Work out cheapest.
	cheapest := -1

	for _, oSp := range positions {

		totalFuel := 0

		for _, iSp := range positions {

			// Skip if it's the same position.
			if iSp == oSp {
				continue
			}

			// Add to fuel cost.
			totalFuel += iSp.FuelNeeded(oSp.horizontal)
		}

		if totalFuel < cheapest ||
			cheapest < 0 {

			cheapest = totalFuel
		}
	}

	fmt.Println("Least fuel:", cheapest)
}
