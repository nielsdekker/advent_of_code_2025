package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	solvePartTwo(readPuzzle("./day_4/puzzle.txt"))
}

func solvePartOne(data []rune, lineLength int) {
	itemAtIs := func(x int, y int) int {
		if x < 0 || x >= lineLength {
			return 0
		}
		if y < 0 || y*lineLength >= len(data) {
			return 0
		}

		if data[(y*lineLength)+x] == '@' {
			return 1
		} else {
			return 0
		}
	}

	solution := 0
	for i, r := range data {
		if r != '@' {
			continue
		}

		// Check all positions
		x := i % lineLength
		y := i / lineLength

		num_rolls := itemAtIs(x-1, y-1)
		num_rolls += itemAtIs(x, y-1)
		num_rolls += itemAtIs(x+1, y-1)
		num_rolls += itemAtIs(x-1, y)
		num_rolls += itemAtIs(x+1, y)
		num_rolls += itemAtIs(x-1, y+1)
		num_rolls += itemAtIs(x, y+1)
		num_rolls += itemAtIs(x+1, y+1)

		if num_rolls < 4 {
			solution++
		}
	}

	fmt.Printf("Solution %d\n", solution)
}

func solvePartTwo(data []rune, ll int) {
	itemAtIs := func(x int, y int) int {
		if x < 0 || x >= ll {
			return 0
		}
		if y < 0 || y*ll >= len(data) {
			return 0
		}

		if data[(y*ll)+x] == '@' {
			return 1
		} else {
			return 0
		}
	}

	solution := 0
	for {
		prevSolution := solution
		for i, r := range data {
			if r != '@' {
				continue
			}

			// Check all positions
			x := i % ll
			y := i / ll

			num_rolls := itemAtIs(x-1, y-1)
			num_rolls += itemAtIs(x, y-1)
			num_rolls += itemAtIs(x+1, y-1)
			num_rolls += itemAtIs(x-1, y)
			num_rolls += itemAtIs(x+1, y)
			num_rolls += itemAtIs(x-1, y+1)
			num_rolls += itemAtIs(x, y+1)
			num_rolls += itemAtIs(x+1, y+1)

			if num_rolls < 4 {
				solution++
				data[i] = 'X'
			}
		}

		if solution == prevSolution {
			// Nothing has been removed
			break
		}
	}

	fmt.Printf("Solution %d\n", solution)
}

func readPuzzle(path string) ([]rune, int) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	lineLength := 0
	puzzle := make([]rune, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lineLength = len(line)

		puzzle = append(puzzle, ([]rune(line))...)
	}

	return puzzle, lineLength
}
