package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type LanternFish struct {
	spawnRate int
	current int
}

func (f *LanternFish) PassDay() (*LanternFish, bool) {

	// Time to spawn?
	if f.current < 1 {

		// Reset.
		f.current = f.spawnRate - 1

		// Create new fish.
		fish := NewLanternFish()

		return fish, true
	}

	f.current--

	return nil, false
}

func NewLanternFish() *LanternFish {
	return &LanternFish{
		spawnRate: 7,
		current: 8,
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

	// Create fish.
	school := make([]*LanternFish, 0)

	for _, val := range vals {

		// Convert string to integer.
		iv, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal("Failed to convert string to int.")
		}

		fish := NewLanternFish()
		fish.current = iv

		school = append(school, fish)
	}

	// Simulate days.
	for i := 0; i < 80; i++ {

		newFish := make([]*LanternFish, 0)

		for _, fish := range school {

			f, s := fish.PassDay()
			if s {
				newFish = append(newFish, f)
			}
		}

		school = append(school, newFish...)
	}

	// Total fish.
	fmt.Println("Fish Count:", len(school))
}
