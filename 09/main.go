package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type PointMap struct {
	points map[Coordinate]*Point
	xMax int
	yMax int
}

func (pm *PointMap) AddPoint(x, y, height int) {
	p := NewPoint(x, y ,height)

	pm.points[p.coordinate] = p

	for c := range pm.points {
		if c.x > pm.xMax {
			pm.xMax = c.x
		}
		if c.y > pm.yMax {
			pm.yMax = c.y
		}
	}
}

func (pm PointMap) Risk() int {

	risk := 0

	for _, p := range pm.points {
		if p.IsLow(pm.points) {
			risk += p.Risk()
		}
	}

	return risk
}

func (pm PointMap) BasinScore() int {

	scores := make([]int, 0)

	for _, p := range pm.points {

		if p.IsLow(pm.points) {
			scores = append(scores, p.BasinSize(pm.points))
		}
	}

	if len(scores) < 3 {
		log.Fatal("Not enough basins!")
	}

	sort.Ints(scores)
	total := scores[len(scores)-1] * scores[len(scores)-2] * scores[len(scores)-3]

	return total
}

func NewPointMap() *PointMap {
	return &PointMap{
		points: make(map[Coordinate]*Point),
	}
}

type Point struct {
	coordinate Coordinate
	height int
}

func (p Point) adjacent(points map[Coordinate]*Point) (adjPoints []*Point) {

	adjacent := []Coordinate{
		NewCoordinate(p.coordinate.x, (p.coordinate.y - 1)),
		NewCoordinate(p.coordinate.x, (p.coordinate.y + 1)),
		NewCoordinate((p.coordinate.x - 1), p.coordinate.y),
		NewCoordinate((p.coordinate.x + 1), p.coordinate.y),
	}

	// Get points at those coordinates.
	for _, ac := range adjacent {
		ap, exists := points[ac]
		if !exists {
			continue
		}

		adjPoints = append(adjPoints, ap)
	}

	return adjPoints
}

func (p Point) Risk() int {
	return 1 + p.height
}

func (p Point) IsLow(points map[Coordinate]*Point) bool {

	adjacent := p.adjacent(points)

	for _, ap := range adjacent {
		if ap.height <= p.height {
			return false
		}
	}

	return true
}

func (p *Point) BasinSize(points map[Coordinate]*Point) int {
	all := p.identifyBasin(points, make([]*Point, 0))

	return len(all)
}

func (p *Point) identifyBasin(points map[Coordinate]*Point, stack []*Point) []*Point {

	// Heights of 9 or above don't count.
	if p.height > 8 {
		return stack
	}

	// Add myself to the stack.
	stack = append(stack, p)

	// Get adjacent.
	adjacentPoints := p.adjacent(points)


	for _, ap := range adjacentPoints {

		// Skip if already in stack.
		inStack := false

		for _, sp := range stack {
			if ap == sp {
				inStack = true
				break
			}
		}

		if inStack {
			continue
		}

		if ap.height > p.height {
			stack = ap.identifyBasin(points, stack)
		}
	}

	return stack
}

func NewPoint(x, y, height int) *Point {
	return &Point{
		coordinate: NewCoordinate(x, y),
		height: height,
	}
}

type Coordinate struct {
	x int
	y int
}

func NewCoordinate(x, y int) Coordinate{
	return Coordinate{
		x: x,
		y: y,
	}
}

func main() {

	pointMap := NewPointMap()

	// Read our input file.
	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	yPos := 0

	for scanner.Scan() {

		for xPos, r := range scanner.Text() {
			
			height, err := strconv.Atoi(string(r))
			if err != nil {
				log.Fatal("Failed to convert string to int.")
			}

			pointMap.AddPoint(xPos, yPos, height)
		}

		yPos++
	}

	fmt.Println("Risk:", pointMap.Risk())
	fmt.Println("Basin score:", pointMap.BasinScore())
}
