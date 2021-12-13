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

func (cs CaveSystem) AllPaths(initial, end string) []*Route {

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

	routes := sc.AllPaths(NewRoute(), ec, 0)

	for _, r := range routes {
		r.Print()
	}

	return routes
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

func (c *Cave) AllPaths(route *Route, target *Cave, depth int) []*Route {

	// Add myself to the route.
	route = route.AddCave(c)

	allRoutes := make([]*Route, 0)

	// End of path?
	if c == target {

		allRoutes = append(allRoutes, route)

		return allRoutes
	}

	// Pass the route to all my paths and store returned routes.
	for _, p := range c.paths {

		// See if we can visit that cave again.
		if p.canVisit(route) {
			retRoutes := p.AllPaths(route, target, depth + 1)
			allRoutes = append(allRoutes, retRoutes...)
		}
	}

	return allRoutes
}

func (c *Cave) canVisit(route *Route) bool {

	// Big caves can always be visited.
	if c.isBig() {
		return true
	}

	if route.Contains(c) {
		return false
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

type Route struct {
	locations []*Cave
}

func (r Route) AddCave(cave *Cave) *Route {

	nr := &Route{
		locations: make([]*Cave, 0),
	}

	for _, l := range r.locations {
		nr.locations = append(nr.locations, l)
	}

	nr.locations = append(nr.locations, cave)

	return nr
}

func (r Route) Contains(cave *Cave) bool {
	for _, loc := range r.locations {
		if loc.name == cave.name {
			return true
		}
	}
	return false
}

func (r *Route) Print() {

	routeStr := ""

	for _, l := range r.locations {

		if routeStr == "" {
			routeStr = l.name
			continue
		}

		routeStr = fmt.Sprintf("%s -> %s", routeStr, l.name)
	}

	fmt.Printf("address: %v | %v\n", &r, routeStr)
}

func NewRoute() *Route {
	return &Route{
		locations: make([]*Cave, 0),
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
