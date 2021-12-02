package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	Direction string
	Units int
}

func main() {

	// Read input file in to commands.
	commands := make([]Command, 0)

	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		dir := parts[0]

		iv, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}

		nc := Command{
			Direction: dir,
			Units:     iv,
		}

		commands = append(commands, nc)
	}

	// Work out final position.
	horizontal := 0
	depth := 0

	for _, c := range commands {
		switch c.Direction {
		case "forward":
			horizontal = horizontal + c.Units
		case "down":
			depth = depth + c.Units
		case "up":
			depth = depth - c.Units
		}
	}

	// Print results.
	fmt.Printf("Horizontal: %d\n", horizontal)
	fmt.Printf("Depth: %d\n", depth)
	fmt.Printf("Multiplied: %d\n", horizontal * depth)
}
