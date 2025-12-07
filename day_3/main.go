package main

import (
	"bufio"
	"fmt"
	"iter"
	"math"
	"os"
)

func main() {
	fmt.Printf("Solution part one %d\n", solver(2))
	fmt.Printf("Solution part two %d\n", solver(12))
}

func solver(numBatteries int) uint64 {
	joltageSum := uint64(0)

	for bank := range bankIterator("./day_3/puzzle.txt") {
		largestValues := make([]int, numBatteries)

		for bankIndex, bankValue := range bank {
			// Update the largest values, starting at the largest value that
			// could be set moving forward. If we find a largest value reset all
			// values after end break

			itemsLeft := len(bank) - bankIndex
			startIndex := max(0, numBatteries-itemsLeft)
			valuesToUpdate := largestValues[startIndex:]

			for lIndex, lVal := range valuesToUpdate {
				if bankValue > lVal {
					for i := range valuesToUpdate[lIndex:] {
						valuesToUpdate[i+lIndex] = 0
					}
					valuesToUpdate[lIndex] = bankValue
					break
				}
			}
		}

		for i, v := range largestValues {
			joltageSum += uint64(math.Pow(10, float64(numBatteries-i-1)) * float64(v))
		}
	}

	return joltageSum
}

func bankIterator(path string) iter.Seq[[]int] {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	return func(yield func([]int) bool) {
		defer f.Close()
		for scanner.Scan() {
			line := scanner.Text()
			bank := make([]int, len(line))

			for i, r := range line {
				bank[i] = int(r) - 48
			}

			if !yield(bank) {
				return
			}
		}
	}
}
