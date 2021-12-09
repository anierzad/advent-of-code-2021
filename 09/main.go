package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		if p.IsLow(pm.adjacentTo(p)) {
			risk += p.Risk()
		}
	}

	return risk
}

func (pm PointMap) adjacentTo(p *Point) (adjPoints []*Point) {

	//adjPoints
	adjacent := make([]Coordinate, 0)

	// Get above.
	if p.coordinate.y > 0 {
		adjacent = append(adjacent, NewCoordinate(p.coordinate.x, (p.coordinate.y - 1)))
	}

	// Get Below.
	if p.coordinate.y < pm.yMax {
		adjacent = append(adjacent, NewCoordinate(p.coordinate.x, (p.coordinate.y + 1)))
	}

	// Get left.
	if p.coordinate.x > 0 {
		adjacent = append(adjacent, NewCoordinate((p.coordinate.x - 1), p.coordinate.y))
	}

	// Get Below.
	if p.coordinate.x < pm.xMax {
		adjacent = append(adjacent, NewCoordinate((p.coordinate.x + 1), p.coordinate.y))
	}

	// Get points at those coordinates.
	for _, ac := range adjacent {
		ap, exists := pm.points[ac]
		if !exists {
			fmt.Printf("Unable to find point at %+v\n", ac)
		}

		adjPoints = append(adjPoints, ap)
	}

	return adjPoints
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

func (p Point) Risk() int {
	return 1 + p.height
}

func (p Point) IsLow(points []*Point) bool {

	for _, ap := range points {
		if ap.height <= p.height {
			return false
		}
	}

	return true
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
}
