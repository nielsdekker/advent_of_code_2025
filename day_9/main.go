package main

import (
	"bufio"
	"fmt"
	"iter"
	"math/big"
	"os"
	"strconv"
	"strings"
)

const CLEAR_LINE = "\033[G\033[K"

func main() {
	solvePartOne()
}

func solvePartOne() {
	points := []Point{}
	for p := range parseInput() {
		points = append(points, p)
	}

	buffer := createTileBuffer(points)

	biggestGrid := gridSize(points[0], points[1])
	biggestGridPointA := points[0]
	biggestGridPointB := points[1]

	biggestTileGrid := int(0)
	biggestTileGridPointA := points[0]
	biggestTileGridPointB := points[1]

	for i, p := range points {

		// If this point is to the bottom right of the TOP point then this will
		// always result in a smaller square and this value can be skipped. Same
		// if this point is to the top left of the BOTTOM point
		if !inGrid(p, biggestGridPointA, biggestGridPointB) {
			// Check if there is a bigger grid for this point
			for j, other := range points {
				// Skip ourselves
				if i == j {
					continue
				}

				g := gridSize(p, other)
				if g > biggestGrid {
					biggestGrid = g
					biggestGridPointA = p
					biggestGridPointB = other
				}
			}
		}

		if !inGrid(p, biggestTileGridPointA, biggestTileGridPointB) {
			// Check if there is a bigger tile grid
			for j, other := range points {
				fmt.Printf("%sChecking: %d/%d\tSubset %d", CLEAR_LINE, i, len(points), j)

				// Skip ourselves
				if i == j {
					continue
				}

				g := gridSize(p, other)
				if g > biggestTileGrid && buffer.isTileGrid(p, other) {
					biggestTileGrid = g
					biggestTileGridPointA = p
					biggestTileGridPointB = other
				}
			}
		}
	}

	fmt.Printf("%sSolution: %d\n", CLEAR_LINE, biggestGrid)
	fmt.Printf("For grid:\n\t%v\n\t%v\n", biggestGridPointA, biggestGridPointB)

	fmt.Printf("\n\nSolution part two: %d\n", biggestTileGrid)
	fmt.Printf("For grid:\n\t%v\n\t%v\n", biggestTileGridPointA, biggestTileGridPointB)
}

func inGrid(p Point, a Point, b Point) bool {
	inX := p.x < max(a.x, b.x) && p.x > min(a.x, b.x)
	inY := p.y < max(a.y, b.y) && p.y > min(a.y, b.y)

	return inX && inY
}

func gridSize(a Point, b Point) int {
	deltaX := a.x - b.x
	deltaY := a.y - b.y

	if deltaX < 0 {
		deltaX *= -1
	}
	if deltaY < 0 {
		deltaY *= -1
	}

	return (deltaX + 1) * (deltaY + 1)
}

func createTileBuffer(points []Point) *Buffer {
	// Get the largest x and y values
	buffer := NewBuffer(points)

	for pointIndex := range len(points) {
		from := points[pointIndex]

		toIndex := pointIndex + 1
		if len(points)-1 == pointIndex {
			toIndex = 0
		}
		to := points[toIndex]

		startX := min(from.x, to.x)
		startY := min(from.y, to.y)

		// We don't know if we move up or down so just loop twice and hope for
		for x := startX; x <= max(from.x, to.x); x++ {
			buffer.setBit(x, startY)
		}
		for y := startY; y <= max(from.y, to.y); y++ {
			buffer.setBit(startX, y)
		}
	}

	// Fill in the buffer
	for y := 0; y < buffer.height; y++ {
		fmt.Printf("%sFilling buffer %d/%d", CLEAR_LINE, y, buffer.height)
		// Use an intermediary array, we never want to fill after the last
		// intersection
		toFill := []int{}
		isFilling := false
		for x := 0; x < buffer.width; x++ {
			isSet := buffer.isSet(x, y)

			if isSet && !isFilling {
				// Toggle filling on
				isFilling = true
			} else if isSet && isFilling {
				// Add all to fill items
				for _, xToFill := range toFill {
					buffer.setBit(xToFill, y)
				}
				toFill = []int{}
			} else if !isSet && isFilling {
				toFill = append(toFill, x)
			}
		}
	}

	return buffer
}

func parseInput() iter.Seq[Point] {
	scanner := bufio.NewScanner(os.Stdin)

	parseOrPanic := func(val string) int64 {
		i, err := strconv.ParseInt(strings.TrimSpace(val), 10, 64)
		if err != nil {
			panic(err)
		}
		return i
	}

	return func(yield func(Point) bool) {
		for scanner.Scan() {
			splitted := strings.Split(scanner.Text(), ",")
			if !yield(Point{
				x: int(parseOrPanic(splitted[0])),
				y: int(parseOrPanic(splitted[1])),
			}) {
				return
			}
		}
	}
}

type Point struct {
	x int
	y int
}

type Buffer struct {
	data   []*big.Int
	width  int
	height int
}

func NewBuffer(points []Point) *Buffer {
	b := &Buffer{}

	maxX := points[0].x
	maxY := points[0].y
	for _, p := range points {
		maxX = max(p.x, maxX)
		maxY = max(p.y, maxY)
	}

	b.width = maxX + 1
	b.height = maxY + 1

	b.data = make([]*big.Int, b.height)
	for i := range len(b.data) {
		b.data[i] = big.NewInt(int64(0))
	}

	return b
}

func (buf *Buffer) setBit(x int, y int) {
	num := buf.data[y]
	num.SetBit(num, x, 1)
}

func (buf *Buffer) isSet(x int, y int) bool {
	b := buf.data[y].Bit(x)
	return b > 0
}

// HOLY SMOKES is this a slow and dumb way to do this :scream:
func (buf *Buffer) isTileGrid(a Point, b Point) bool {
	start := Point{
		x: min(a.x, b.x),
		y: min(a.y, b.y),
	}
	end := Point{
		x: max(a.x, b.x),
		y: max(a.y, b.y),
	}

	xMask := big.NewInt(int64(0))
	for x := start.x; x <= end.x; x++ {
		xMask.SetBit(xMask, x, 1)
	}

	for y := start.y; y <= end.y; y++ {
		line := buf.data[y]

		foo := new(big.Int).And(line, xMask)
		if foo.Cmp(xMask) < 0 {
			return false
		}
	}

	// No early return so all tiles are valid
	return true
}
