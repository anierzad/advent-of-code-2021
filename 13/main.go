package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func NewPoint(x, y int) Point {
	return Point{
		x: x,
		y: y,
	}
}

type Paper struct {
	dots map[Point]bool
	xSize int
	ySize int
}

func (p *Paper) FoldX(index int) {
	for y := 0; y <= p.ySize; y++ {

	writeX := index - 1

		for x := index + 1; x <= p.xSize; x++ {

			sp := NewPoint(x, y)
			dp := NewPoint(writeX, y)

			sv := p.dots[sp]
			dv := p.dots[dp]

			p.dots[dp] = sv || dv

			writeX--
		}
	}

	p.xSize = index - 1
}

func (p *Paper) FoldY(index int) {
	writeY := index - 1

	for y := index + 1; y <= p.ySize; y++ {
		for x := 0; x <= p.xSize; x++ {

			sp := NewPoint(x, y)
			dp := NewPoint(x, writeY)

			sv := p.dots[sp]
			dv := p.dots[dp]

			p.dots[dp] = sv || dv
		}
		writeY--
	}

	p.ySize = index - 1
}

func (p *Paper) AddDot(point Point) {
	p.dots[point] = true
}

func (p *Paper) DotCount() int {
	count := 0

	for y := 0; y <= p.ySize; y++ {
		for x := 0; x <= p.xSize; x++ {
			val := p.dots[NewPoint(x, y)]
			if val {
				count++
			}
		}
	}

	return count
}

func NewPaper(xSize, ySize int) *Paper {
	p := &Paper{
		dots: make(map[Point]bool),
		xSize: xSize,
		ySize: ySize,
	}

	for y := 0; y <= ySize; y++ {
		for x := 0; x <= xSize; x++ {
			p.dots[NewPoint(x, y)] = false
		}
	}

	return p
}

func main() {

	// Read our input file.
	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	dotPoints := make([]Point, 0)
	xMax := 0
	yMax := 0

	for scanner.Scan() {

		// Check we've not got to the folds yet.
		if len(scanner.Text()) == 0 {
			break
		}

		split := strings.Split(scanner.Text(), ",")

		x, err := strconv.Atoi(split[0])
		if err != nil {
			log.Fatal("Failed to convert string to int.")
		}
		y, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal("Failed to convert string to int.")
		}

		dotPoints = append(dotPoints, NewPoint(x, y))

		if x > xMax {
			xMax = x
		}
		if y > yMax {
			yMax = y
		}
	}

	paper := NewPaper(xMax, yMax)

	for _, dp := range dotPoints {
		paper.AddDot(dp)
	}

	for scanner.Scan() {

		split := strings.Split(scanner.Text(), " ")
		split = strings.Split(split[2], "=")

		index, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal("Failed to convert string to int.")
		}

		if split[0] == "x"{
			paper.FoldX(index)
		}

		if split[0] == "y"{
			paper.FoldY(index)
		}
		break
	}

	fmt.Println("Dot count:", paper.DotCount())
}
