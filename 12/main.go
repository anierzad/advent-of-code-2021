package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type CaveSystem struct {
	caves map[string]*Cave
}

func (cs *CaveSystem) AddPath(locations []string) {
	for _, start := range locations {
		for _, end := range locations {

			// Only if start is not the same as end.
			if start != end {

				// Does this start exist?
				sc, exists := cs.caves[start]
				if !exists {
					sc = NewCave(start)
					cs.caves[start] = sc
				}

				// Does this end exist?
				se , exists := cs.caves[end]
				if !exists {
					se = NewCave(end)
					cs.caves[end] = se
				}

				// Add end to start.
				sc.AddPath(se)
			}
		}
	}
}

func (cs CaveSystem) AllPaths(initial, end string) [][]*Cave {

	// Get initial location.
	sc, exists := cs.caves[initial]
	if !exists {
		log.Fatal("Cave not found.")
	}

	// Get initial location.
	ec, exists := cs.caves[end]
	if !exists {
		log.Fatal("Cave not found.")
	}

	paths := sc.AllPaths(make([]*Cave, 0), ec, 0)

	for _, p := range paths {
		for _, r := range p {
			fmt.Printf("%s -> ", r.name)
		}
		fmt.Println()
	}

	return paths
}

func NewCaveSystem() *CaveSystem {
	return &CaveSystem{
		caves: make(map[string]*Cave, 0),
	}
}

type Cave struct {
	name string
	paths []*Cave
}

func (c *Cave) AddPath(p *Cave) {
	c.paths = append(c.paths, p)
}

func (c *Cave) AllPaths(route []*Cave, target *Cave, depth int) [][]*Cave {

	// Add myself to the route.
	route = append(route, c)

	//fmt.Println("cave", c.name)

	// for _, r := range route {
	// 	fmt.Printf("%s -> ", r.name)
	// }
	// fmt.Println()

	allRoutes := make([][]*Cave, 0)

	// End of path?
	if c == target {

		allRoutes = append(allRoutes, route)

		fmt.Println("appending:", c.name, "depth:", depth)

		for _, ar := range allRoutes {
			for _, r := range ar {
				fmt.Printf("%s -> ", r.name)
			}
			fmt.Println()
		}

		fmt.Println()
		fmt.Println()
		fmt.Scanln()

		return allRoutes
	}

	// Pass the route to all my paths and store returned routes.
	for _, p := range c.paths {

		// See if we can visit that cave again.
		if p.canVisit(route) {

			retRoutes := p.AllPaths(route, target, depth + 1)

			for _, rr := range retRoutes {
				allRoutes = append(allRoutes, rr)
			}


			fmt.Println("cave:", c.name, "to:", p.name, "depth:", depth)

			for _, ar := range allRoutes {
				for _, r := range ar {
					fmt.Printf("%s -> ", r.name)
				}
				fmt.Println()
			}

			fmt.Println()
			fmt.Println()
			fmt.Scanln()
		}
	}

	return allRoutes
}

func (c *Cave) canVisit(route []*Cave) bool {

	// Big caves can always be visited.
	if c.isBig() {
		return true
	}

	// Is it already in the route?
	for _, rp := range route {
		if rp == c {
			return false
		}
	}

	return true
}

func (c *Cave) isBig() bool {
	r := string(c.name[0])

	return strings.ToUpper(r) == r
}

func NewCave(name string) *Cave {
	return &Cave{
		name: name,
		paths: make([]*Cave, 0),
	}
}

func main() {

	caveSystem := NewCaveSystem()

	// Read our input file.
	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, "-")

		caveSystem.AddPath(split)
	}

	fmt.Println("Total routes:", len(caveSystem.AllPaths("start", "end")))
}
