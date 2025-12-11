package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	beams := make(map[int]int)

	scanner.Scan()
	for i, r := range scanner.Text() {
		if r == 'S' {
			beams[i] = 1
		}
	}

	// For this we basically only need to know how many ways we can get to a
	// certain spot. Then we can just sum the number of ways at the end and get
	// the answer.
	splits := 0
	for scanner.Scan() {
		for i, r := range scanner.Text() {
			switch r {
			case '^':
				if val, ok := beams[i]; ok {
					// There is a beam above, remove this
					delete(beams, i)

					// Also create two new beams, if needed
					beams[i-1] = val + beams[i-1]
					beams[i+1] = val + beams[i+1]
					splits++
				}
			}
		}
	}

	// Sum all beam values
	sum := 0
	for _, v := range beams {
		sum += v
	}
	fmt.Printf("Splits: %d, quantum %d\n", splits, sum)
}
