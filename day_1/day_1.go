package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//f, err := os.Open("./day_1/data_example.txt")
	// Answer is 6671
	f, err := os.Open("./day_1/data_puzzle_1.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	current := 50
	solution := 0
	for scanner.Scan() {
		line := scanner.Text()

		prefix := line[0]
		num64, err := strconv.ParseInt(line[1:], 10, 32)
		if err != nil {
			panic(err)
		}
		num := int(num64)

		switch prefix {
		case 'L':
			// This is stupid but was struggling with of by one errors
			for range num {
				current--

				switch current {
				case 0:
					solution++
				case -1:
					current = 99
				}
			}
		case 'R':
			current += num
			solution += current / 100
			current = current % 100
		}
	}

	fmt.Printf("Solution: %d\n", solution)
}
