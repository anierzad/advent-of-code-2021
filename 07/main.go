package main

import (
	"bufio"
	"fmt"
	"log"
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

	start := sp.horizontal
	end := pos

	if start > end {
		start = pos
		end = sp.horizontal
	}

	reqFuel := 0
	currentBurn := 1

	for i := start; i < end; i++ {
		reqFuel += currentBurn
		currentBurn++
	}

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
	minPos := -1
	maxPos := -1
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

		// Max/min.
		if iv < minPos ||
			minPos < 0 {
			minPos = iv
		}
		if iv > maxPos ||
			maxPos < 0 {
			maxPos = iv
		}
	}

	// Work out cheapest.
	cheapest := -1

	for i := minPos; i <= maxPos; i++ {

		totalFuel := 0

		for _, iSp := range positions {

			// Add to fuel cost.
			totalFuel += iSp.FuelNeeded(i)
		}

		if totalFuel < cheapest ||
			cheapest < 0 {
			cheapest = totalFuel
		}
	}

	fmt.Println("Least fuel:", cheapest)
}
