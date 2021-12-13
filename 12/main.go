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

	// Get all routes.
	routes := sc.AllPaths(NewRoute(), sc, ec, false, false)

	// Dedupe routes.
	deDupe := make(map[string]*Route)
	for _, r := range routes {
		str := r.ToString()
		_, exists := deDupe[str]
		if !exists {
			deDupe[str] = r
		}
	}

	// Make back in to a slice.
	routes = make([]*Route, 0)
	for _, v := range deDupe {
		routes = append(routes, v)
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

func (c *Cave) AllPaths(route *Route, start, target *Cave, revisit, rUsed bool) []*Route {

	// Hold routes.
	allRoutes := make([]*Route, 0)

	// If we're the start and we've been here before return early.
	if c == start && !c.canVisit(route) {
		return allRoutes
	}

	// Add myself to the route.
	route = route.AddCave(c)

	// End of path?
	if c == target {

		allRoutes = append(allRoutes, route)

		return allRoutes
	}

	// Pass the route to all my paths and store returned routes.
	for _, p := range c.paths {

		// See if we can visit that cave again.
		can := p.canVisit(route)
		if can || revisit {

			if !can {
				rUsed = true
			}

			if !rUsed {
				retRoutes := p.AllPaths(route, start, target, true, rUsed)
				allRoutes = append(allRoutes, retRoutes...)
			}

			retRoutes := p.AllPaths(route, start, target, false, rUsed)
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

func (r Route) ToString() string {
	routeStr := ""

	for _, l := range r.locations {

		if routeStr == "" {
			routeStr = l.name
			continue
		}

		routeStr = fmt.Sprintf("%s -> %s", routeStr, l.name)
	}

	return routeStr
}

func (r *Route) Print() {
	fmt.Printf("address: %v | %v\n", &r, r.ToString())
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
