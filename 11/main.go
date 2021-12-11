package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type OctopusGrid struct {
	gridPoints []GridPoint
	octopi map[GridPoint]*Octopus
	tick int
	flashes int
	xMax int
	yMax int
}

func (og *OctopusGrid) AddOctopus(x, y, power int) {
	gp := NewGridPoint(x, y)
	og.gridPoints = append(og.gridPoints, gp)
	og.octopi[gp] = NewOctopus(gp, power)

	if x > og.xMax {
		og.xMax = x
	}
	if y > og.yMax {
		og.yMax = y
	}
}

func (og *OctopusGrid) Tick() {

	// Tick all octupi.
	for _, gp := range og.gridPoints {
		o := og.octopi[gp]
		o.Tick()
	}

	// Initialise flash cascade.
	og.flashCascade()

	og.tick++
}

func (og OctopusGrid) Flashes() int {
	return og.flashes
}

func (og OctopusGrid) AllFlash() bool {

	allFlash := true

	for _, gp := range og.gridPoints {
		o := og.octopi[gp]
		if !o.Flashed {
			allFlash = false
			break
		}
	}

	return allFlash
}

func (og *OctopusGrid) flashCascade() {

	recurse := false

	// Make them flash.
	for _, gp := range og.gridPoints {
		o := og.octopi[gp]
		if o.Flash(og.tick) {

			// Flag we need to do another pass and increment flash count.
			recurse = true
			og.flashes++

			// Identify adjacent octupi and increase their power.
			for _, agp := range og.adjacentPoints(o.point) {
				ao := og.octopi[agp]
				ao.IncreasePower()
			}
		}
	}

	if recurse {
		og.flashCascade()
	}
}

func (og OctopusGrid) adjacentPoints(gp GridPoint) []GridPoint {

	gps := make([]GridPoint, 0)

	for y := gp.y - 1; y <= gp.y + 1; y++ {

		if y < 0 || y > og.yMax {
			continue
		}

		for x := gp.x - 1; x <= gp.x + 1; x++ {
			
			if x < 0 || x > og.xMax {
				continue
			}

			agp := NewGridPoint(x, y)

			if agp == gp {
				continue
			}

			gps = append(gps, agp)
		}
	}

	return gps
}

func NewOctopusGrid() *OctopusGrid {
	return &OctopusGrid{
		gridPoints: make([]GridPoint, 0),
		octopi: make(map[GridPoint]*Octopus),
		tick: 0,
		flashes: 0,
	}
}

type Octopus struct {
	point GridPoint
	power int
	Flashed bool
}

func (o *Octopus) Tick() {

	// If we previously flashed we should reset.
	if o.Flashed {
		o.Reset()
	}

	o.Flashed = false
	o.IncreasePower()
}

func (o *Octopus) IncreasePower() {
	o.power++
}

func (o *Octopus) Reset() {
	o.power = 0
}

func (o *Octopus) Flash(tick int) bool {

	// Ready to flash?
	if o.power > 9 {

		// Only flash if we've not flashed this tick.
		if !o.Flashed {
			o.Flashed = true
			o.Reset()
			return true
		}
	}
	return false
}

func NewOctopus(gp GridPoint, power int) *Octopus {
	return &Octopus{
		point: gp,
		power: power,
		Flashed: false,
	}
}

type GridPoint struct {
	x int
	y int
}

func NewGridPoint(x, y int) GridPoint {
	return GridPoint{
		x: x,
		y: y,
	}
}

func main() {

	grid := NewOctopusGrid()

	// Read our input file.
	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for y := 0; scanner.Scan(); y++ {
		for x, r := range scanner.Text() {
			power, err := strconv.Atoi(string(r))
			if err != nil {
				log.Fatal("Failed to convert string to int.")
			}

			grid.AddOctopus(x, y, power)
		}
	}

	for i := 1;; i++ {
		grid.Tick()

		if grid.AllFlash() {
			fmt.Println("First all flash:", i)
			break
		}
	}

	fmt.Println("Flashes:", grid.Flashes())
}
