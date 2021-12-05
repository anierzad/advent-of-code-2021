package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Plottable interface {
	AllPoints() []Point
}

type Point struct {
	x int
	y int
}

type Line struct {
	start Point
	end Point
}

func (l Line) AllPoints() []Point {

	// We need to ignore diagonal lines.
	if l.start.x != l.end.x &&
		l.start.y != l.end.y {
		return make([]Point, 0)
	}

	points := make([]Point, 0)

	// Decide direction.
	horizontal := true

	if l.start.y != l.end.y {
		horizontal = false
	}

	if horizontal {

		// X logic, Y remains constant.
		for _, num := range NumbersBetween(l.start.x, l.end.x) {
			p := Point{
				x: num,
				y: l.start.y,
			}
			points = append(points, p)
		}

	} else {

		// Y logic, X remains constant.
		for _, num := range NumbersBetween(l.start.y, l.end.y) {
			p := Point{
				x: l.start.x,
				y: num,
			}
			points = append(points, p)
		}

	}

	return points
}

type VentMap struct {
	points map[Point]int
}

func (v *VentMap) Plot(p Plottable) {
	for _, point := range p.AllPoints() {
		v.points[point]++
	}
}

func (v VentMap) Overlaps() int {
	total := 0
	for _, v := range v.points {
		if v > 1 {
			total++
		}
	}
	return total
}

func NewVentMap(xMin, xMax, yMin, yMax int) *VentMap {

	vm := &VentMap{
		points: make(map[Point]int),
	}

	for xi := xMin; xi < xMax; xi++ {
		for yi := yMin; yi < yMax; yi++ {
			p := Point{
				x: xi,
				y: yi,
			}

			vm.points[p] = 0
		}
	}

	return vm
}

func NumbersBetween(s, e int) []int {

	// Firstly put the numbers in order.
	start := s
	end := e

	if s > e {
		end = s
		start = e
	}

	nums := make([]int, 0)

	for i := start; i <= end; i++ {
		nums = append(nums, i)
	}

	return nums
}

func main() {

	// Read our input file.
	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	xMax := 0
	xMin := 9999
	yMax := 0
	yMin := 9999
	lines := make([]*Line, 0)

	for scanner.Scan() {
		l := scanner.Text()

		// Split start and end.
		linePoints := strings.Split(l, " -> ")

		// Split start and end in to x and y.
		lineStart := strings.Split(linePoints[0], ",")
		lineEnd := strings.Split(linePoints[1], ",")

		// Convert to integers.
		sX, err := strconv.Atoi(lineStart[0])
		if err != nil {
			log.Fatal("Can't convert x string to int.")
		}
		sY, err := strconv.Atoi(lineStart[1])
		if err != nil {
			log.Fatal("Can't convert y string to int.")
		}
		eX, err := strconv.Atoi(lineEnd[0])
		if err != nil {
			log.Fatal("Can't convert x string to int.")
		}
		eY, err := strconv.Atoi(lineEnd[1])
		if err != nil {
			log.Fatal("Can't convert y string to int.")
		}

		// Make line with points.
		line := &Line{
			start: Point{
				x: sX,
				y: sY,
			},
			end: Point{
				x: eX,
				y: eY,
			},
		}

		lines = append(lines, line)

		// Update best scores.
		if sX > xMax {
			xMax = sX
		}
		if eX > xMax {
			xMax = eX
		}
		if sY > yMax {
			yMax = sY
		}
		if eY > yMax {
			yMax = eY
		}

		// Update worst scores.
		if sX < xMin {
			xMin = sX
		}
		if eX < xMin {
			xMin = eX
		}
		if sY < yMin {
			yMin = sY
		}
		if eY < yMin {
			yMin = eY
		}
	}

	// Create vent map.
	vm := NewVentMap(xMin, xMax, yMin, yMax)

	// Plot all lines.
	for _, line := range lines {
		vm.Plot(line)
	}

	fmt.Println("Overlaps:", vm.Overlaps())
}
