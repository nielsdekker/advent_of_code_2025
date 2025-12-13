package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	solvePartTwo()
}

func solvePartOne() {
	data := parsePuzzle()
	boxIndex := 0
	boxes := make([]int, len(data))

	prevLowest := float64(0)
	for range 1000 {
		lowestDistance := float64(math.MaxInt64)
		lowestFrom := 0
		lowestTo := 0

		for i := range len(data) - 1 {
			for n := range data[i+1:] {
				dist := distance(data[i], data[i+n+1])
				if dist < lowestDistance && dist > prevLowest {
					lowestFrom = i
					lowestTo = i + n + 1
					lowestDistance = dist
				}
			}
		}

		// Put these two points in a junction together
		boxFrom := boxes[lowestFrom]
		boxTo := boxes[lowestTo]
		if boxFrom > 0 && boxTo > 0 {
			// These are part of different junctions that are now connected.
			// Make sure to update any and all junctions to the same value
			for i := range boxes {
				if boxes[i] == boxFrom {
					boxes[i] = boxTo
				}
			}
			boxes[lowestFrom] = boxTo
			boxes[lowestTo] = boxTo
		} else if boxFrom > 0 || boxTo > 0 {
			// One is already part of a junction so combine them
			boxes[lowestFrom] = max(boxFrom, boxTo)
			boxes[lowestTo] = max(boxFrom, boxTo)
		} else {
			boxIndex++
			boxes[lowestFrom] = boxIndex
			boxes[lowestTo] = boxIndex
		}
		prevLowest = lowestDistance
	}

	// Now count the number of boxes per index
	count := make([]int64, boxIndex+1)
	for _, b := range boxes {
		if b == 0 {
			continue
		}
		count[b]++
	}
	slices.Sort(count)

	fmt.Printf("%d\n", boxes)

	solution := count[len(count)-1]
	solution *= count[len(count)-2]
	solution *= count[len(count)-3]
	fmt.Printf("Solution %d\n", solution)
}

func solvePartTwo() {
	data := parsePuzzle()
	boxIndex := 0
	boxes := make([]int, len(data))

	prevLowest := float64(0)
	for {
		lowestDistance := float64(math.MaxInt64)
		lowestFrom := 0
		lowestTo := 0

		for i := range len(data) - 1 {
			for n := range data[i+1:] {
				dist := distance(data[i], data[i+n+1])
				if dist < lowestDistance && dist > prevLowest {
					lowestFrom = i
					lowestTo = i + n + 1
					lowestDistance = dist
				}
			}
		}

		// Put these two points in a junction together
		boxFrom := boxes[lowestFrom]
		boxTo := boxes[lowestTo]
		if boxFrom > 0 && boxTo > 0 {
			// These are part of different junctions that are now connected.
			// Make sure to update any and all junctions to the same value
			for i := range boxes {
				if boxes[i] == boxFrom {
					boxes[i] = boxTo
				}
			}

			boxes[lowestFrom] = boxTo
			boxes[lowestTo] = boxTo
		} else if boxFrom > 0 || boxTo > 0 {
			// One is already part of a junction so combine them
			boxes[lowestFrom] = max(boxFrom, boxTo)
			boxes[lowestTo] = max(boxFrom, boxTo)
		} else {
			boxIndex++
			boxes[lowestFrom] = boxIndex
			boxes[lowestTo] = boxIndex
		}
		prevLowest = lowestDistance

		// When connecting a junction check if this is now solved
		solved := true
		for _, b := range boxes {
			if b != boxTo {
				solved = false
				break
			}
		}

		if solved {
			fmt.Printf("Last two points: %v %v\n", data[lowestFrom], data[lowestTo])
			fmt.Printf("Solution: %d\n", data[lowestFrom].x*data[lowestTo].x)
			break
		}
	}
}

type Point struct {
	x int64
	y int64
	z int64
}

func distance(from Point, to Point) float64 {
	return math.Sqrt(
		math.Pow(float64(from.x-to.x), 2) +
			math.Pow(float64(from.y-to.y), 2) +
			math.Pow(float64(from.z-to.z), 2),
	)
}

func parsePuzzle() []Point {
	scanner := bufio.NewScanner(os.Stdin)

	data := []Point{}
	for scanner.Scan() {
		splitted := strings.Split(scanner.Text(), ",")
		data = append(data, Point{
			x: parseOrPanic(splitted[0]),
			y: parseOrPanic(splitted[1]),
			z: parseOrPanic(splitted[2]),
		})
	}

	return data
}

func parseOrPanic(val string) int64 {
	res, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		panic(err)
	}

	return res
}
