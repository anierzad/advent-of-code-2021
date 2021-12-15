package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Point struct {
	x int8
	y int8
}

func NewPoint(x, y int8) Point {
	return Point{
		x: x,
		y: y,
	}
}

type Route struct {
	points []Point
}

func (r Route) AddPoint(point Point) Route {
	newPoints := make([]Point, len(r.points))

	for i, p := range r.points {
		newPoints[i] = p
	}

	newPoints = append(newPoints, point)

	return Route{
		points: newPoints,
	}
}

func (r Route) Contains(point Point) bool {
	for _, p := range r.points {
		if p == point {
			return true
		}
	}
	return false
}

func (r Route) Risk(risks map[Point]int8) int {
	risk := 0

	for _, p := range r.points {
		risk += int(risks[p])
	}

	return risk
}

func NewRoute() Route {
	return Route{
		points: []Point{},
	}
}

type Cavern struct {
	points map[Point]int8
	xMax int8
	yMax int8
}

func (c *Cavern) AddPoint(point Point, risk int8) {
	c.points[point] = risk

	if point.x > c.xMax {
		c.xMax = point.x
	}

	if point.y > c.yMax {
		c.yMax = point.y
	}
}

func (c *Cavern) LeastRisk() int {

	// Start and finish.
	start := NewPoint(0, 0)
	finish := NewPoint(c.xMax, c.yMax)

	visited := make(map[Point]bool)
	distances := make(map[Point]int)
	for p := range c.points {
		distances[p] = -1
	}
	distances[start] = 0

	current := start

	for {
		for _, ap := range c.adjacentTo(current) {
			vis, exists := visited[ap]
			if exists && vis {
				continue
			}

			// Calculate tentative distance.
			tentative := distances[current] + int(c.points[ap])

			if distances[ap] == -1 || tentative < distances[ap] {
				distances[ap] = tentative
			}
		}

		visited[current] = true

		// Destination node marked as visited?
		vis, exists := visited[finish]
		if exists && vis {
			return distances[finish]
		}

		lowestDistance := -1
		lowestDistancePoint := Point{}

		for p, dist := range distances {
			vis, exists := visited[p]
			if exists && vis {
				continue
			}

			if lowestDistance < 0 || (dist > 0 && dist < lowestDistance) {
				lowestDistance = dist
				lowestDistancePoint = p
			}
		}

		// No connection.
		if lowestDistance < 0 {
			break
		}

		current = lowestDistancePoint
	}

	return -1
}

func printDistances(distances map[Point]int, xMax, yMax int8) {

	fmt.Println()
	for y := int8(0); y <= yMax; y++ {
		for x := int8(0); x <= xMax; x++ {
			p := NewPoint(x, y)

			val := distances[p]

			fmt.Printf("%d", val)
		}
		fmt.Println()
	}
	fmt.Println()
}


func (c Cavern) adjacentTo(point Point) []Point {

	adjacent := make([]Point, 0)

	// Right.
	if point.x < c.xMax {
		adjacent = append(adjacent, NewPoint(point.x + 1, point.y))
	}

	// Below.
	if point.y < c.yMax {
		adjacent = append(adjacent, NewPoint(point.x, point.y + 1))
	}

	// Left.
	if point.x > 0 {
		adjacent = append(adjacent, NewPoint(point.x - 1, point.y))
	}

	// Above.
	if point.y > 0 {
		adjacent = append(adjacent, NewPoint(point.x, point.y - 1))
	}

	return adjacent
}

func NewCavern() *Cavern {
	return &Cavern{
		points: make(map[Point]int8),
		xMax: 0,
		yMax: 0,
	}
}

func main() {

	cavern := NewCavern()

	// Read our input file.
	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for y := 0; scanner.Scan(); y++ {
		for x, r := range scanner.Text() {
			val, err := strconv.Atoi(string(r))
			if err != nil {
				log.Fatal("Failed to convert string to int.")
			}

			p := NewPoint(int8(y), int8(x))

			cavern.AddPoint(p, int8(val))
		}
	}

	leastRisk := cavern.LeastRisk()

	fmt.Println("Least risk:", leastRisk)
}
