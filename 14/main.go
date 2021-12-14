package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type PolymerTemplate struct {
	current string
	rules map[string]string
}

func (pt *PolymerTemplate) AddRule(pair, insert string) {
	pt.rules[pair] = insert
}

func (pt *PolymerTemplate) Step() {
	pt.current = pt.recStep(pt.current, 0)
}

func (pt *PolymerTemplate) recStep(cur string, start int) string {

	for p1, p2 := start, start + 1; p2 < len(cur); p1, p2 = p1 + 1, p2 + 1 {
		part := cur[p1:p2+1]

		rule, exists := pt.rules[part]
		if exists {

			newCur :=  cur[:p1+1] + rule + cur[p2:]

			return pt.recStep(newCur, p2 + 1)
		}
	}

	return cur
}

func (pt *PolymerTemplate) MostCommon() (rune, int) {
	counts := make(map[rune]int)

	for _, r := range pt.current {
		count, exists := counts[r]
		if !exists {
			count = 0
		}

		count++

		counts[r] = count
	}

	var best rune
	bestCount := -1

	for k, v := range counts {

		if bestCount < 0 || v > bestCount {
			bestCount = v
			best = k
		}
	}

	return best, bestCount
}

func (pt *PolymerTemplate) LeastCommon() (rune, int) {
	counts := make(map[rune]int)

	for _, r := range pt.current {
		count, exists := counts[r]
		if !exists {
			count = 0
		}

		count++

		counts[r] = count
	}

	var worst rune
	worstCount := -1

	for k, v := range counts {
		if worstCount < 0 || v < worstCount {
			worstCount = v
			worst = k
		}
	}

	return worst, worstCount
}

func NewPolymerTemplate(template string) *PolymerTemplate {
	return &PolymerTemplate{
		current: template,
		rules: make(map[string]string),
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

	scanner.Scan()

	template := NewPolymerTemplate(scanner.Text())

	scanner.Scan()

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " -> ")

		template.AddRule(split[0], split[1])
	}

	for i := 0; i < 10; i++ {
		template.Step()
	}

	_, best := template.MostCommon()
	_, worst := template.LeastCommon()

	difference := best - worst

	fmt.Println("Difference:", difference)
}
