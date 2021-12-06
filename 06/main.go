package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type LanternFishSchool struct {
	current uint8
	count int
}

func (f *LanternFishSchool) PassDay() int {

	// Time to spawn?
	if f.current < 1 {

		// Reset current and return spawn count.
		f.current = 6

		return f.count
	}

	f.current--

	return 0
}

func (f *LanternFishSchool) Add(c int) {
	f.count += c
}

func (f LanternFishSchool) Total() int {
	return f.count
}

func NewLanternFishSchool(current uint8, count int) *LanternFishSchool {
	return &LanternFishSchool{
		current: current,
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

	// Create fish in groups.
	schools := make(map[string]*LanternFishSchool)

	for _, val := range vals {

		// Convert string to integer.
		iv, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal("Failed to convert string to int.")
		}

		// Construct a key.
		schoolKey := fmt.Sprintf("%d-%d", 0, iv)

		// Does the school exist?
		school, exists := schools[schoolKey]
		if !exists {

			// Create school.
			school = NewLanternFishSchool(uint8(iv), 0)

			// Add to map.
			schools[schoolKey] = school
		}

		// Add one to count.
		school.Add(1)
	}

	// Simulate days.
	for i := 0; i < 256; i++ {

		// Keep a track of the number of fish we need in our new school.
		newFishToday := 0

		for _, school := range(schools) {
			newFish := school.PassDay()
			newFishToday += newFish
		}

		// Create a new school if we need to.
		if newFishToday > 0 {

			// Construct a key.
			schoolKey := fmt.Sprintf("%d-%d", i, 8)

			ns := NewLanternFishSchool(8, newFishToday)

			schools[schoolKey] = ns
		}
	}

	totalFish := 0
	for _, school := range schools {
		totalFish += school.Total()
	}

	// Total fish.
	fmt.Println("Fish Count:", totalFish)
}
