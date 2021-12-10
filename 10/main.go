package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

var (
	syntax = map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}
	scores = map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
)

type ChunkLine struct {
	chunks []*Chunk
}

func (cl ChunkLine) AutoComplete() []rune {

	corrections := make([]rune, 0)

	for _, chunk := range cl.chunks {
		ccs := chunk.AutoComplete()

		corrections = append(corrections, ccs...)
	}

	return corrections
}

func (cl ChunkLine) AutoCompleteScore() int {

	total := 0

	for _, correction := range cl.AutoComplete() {
		total *= 5
		total += scores[correction]
	}

	return total
}

func NewChunkLine(line string) (*ChunkLine, int, bool) {

	cl := &ChunkLine{
		chunks: make([]*Chunk, 0),
	}

	// Start iterating from our start position.
	for i := 0; i < len(line); i++ {

		c, ePos, invalid := NewChunk(line, i, 0)
		if invalid {
			return cl, ePos, invalid
		}

		cl.chunks = append(cl.chunks, c)

		// Move past what's already been processed.
		i = ePos
	}

	return cl, len(line), false
}

type Chunk struct {
	open rune
	close rune
	chunks []*Chunk
}

func (c Chunk) AutoComplete() []rune {

	corrections := make([]rune, 0)

	for _, chunk := range c.chunks {
		ccs := chunk.AutoComplete()

		corrections = append(corrections, ccs...)
	}

	if c.close == 0 {
		corrections = append(corrections, syntax[c.open])
	}

	return corrections
}

func NewChunk(line string, position, depth int) (*Chunk, int, bool) {

	// Make our new chunk as we're always expecting one.
	c := &Chunk{
		open: 0,
		close: 0,
		chunks: make([]*Chunk, 0),
	}

	// Start iterating from our start position.
	for i := position; i < len(line); i++ {

		// Get token at this position.
		r := rune(line[i])

		// We need to grab a start token for ourselves.
		if c.open == 0 {

			// Check it's a start token.
			if isOpen(r, syntax) {
				c.open = r
				continue
			}

			// If we make it here then we're invalid.
			return c, i, true
		}

		// We also need an end token for ourselves.
		if c.close == 0 {

			// Check it's an end token and it's ours.
			if isClose(r, syntax) && r == syntax[c.open] {
				c.close = r
				return c, i, false
			}
		}

		cc, ePos, invalid := NewChunk(line, i, depth + 1)
		if invalid {
			return c, ePos, invalid
		}

		c.chunks = append(c.chunks, cc)

		// Move past what's already been processed.
		i = ePos
	}

	return c, len(line), false
}

func isOpen(r rune, syntax map[rune]rune) bool {
	for k := range syntax {
		if r == k {
			return true
		}
	}
	return false
}

func isClose(r rune, syntax map[rune]rune) bool {
	for _, v := range syntax {
		if r == v {
			return true
		}
	}
	return false
}

func main() {

	chunkLines := make([]*ChunkLine, 0)

	// Read our input file.
	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		text := scanner.Text()

		cl, _, invalid := NewChunkLine(text)
		if !invalid {
			chunkLines = append(chunkLines, cl)
		}
	}

	acScores := make([]int, 0)

	for _, cl := range chunkLines {
		acScores = append(acScores, cl.AutoCompleteScore())
	}

	sort.Ints(acScores)

	middle := int(math.Round(float64(len(acScores) / 2)))

	fmt.Println("Auto complete score:", acScores[middle])
}
