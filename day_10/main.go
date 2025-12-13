package main

import (
	"bufio"
	"fmt"
	"iter"
	"os"
	"strconv"
	"strings"
)

func main() {
	solution := 0
	for p := range parsePuzzle() {
		solution += p.solve()
	}

	fmt.Printf("Solution: %d\n", solution)
}

func parsePuzzle() iter.Seq[PuzzleInput] {
	scanner := bufio.NewScanner(os.Stdin)

	return func(yield func(PuzzleInput) bool) {
		for scanner.Scan() {
			input := PuzzleInput{}

			// The first value is the target input
			// Second value to second from last are the buttons, we skip the
			// joltage for now
			splitted := strings.Split(scanner.Text(), " ")

			target := strings.Trim(splitted[0], "[]")
			for i, r := range target {
				if r == '#' {
					input.target += 1 << i
				}
			}

			input.buttons = []int16{}
			input.buttonsRaw = [][]int{}
			for _, l := range splitted[1 : len(splitted)-1] {
				numbers := strings.Split(strings.Trim(l, "()"), ",")
				button := int16(0)
				buttonsRaw := []int{}

				for _, s := range numbers {
					n, err := strconv.ParseInt(s, 10, 16)
					if err != nil {
						panic(err)
					}

					button += 1 << n
					buttonsRaw = append(buttonsRaw, int(n))
				}

				input.buttons = append(input.buttons, button)

				input.buttonsRaw = append(input.buttonsRaw, buttonsRaw)
			}

			input.joltage = []int{}
			for s := range strings.SplitSeq(strings.Trim(splitted[len(splitted)-1], "{}"), ",") {
				n, err := strconv.ParseInt(s, 10, 16)
				if err != nil {
					panic(err)
				}
				input.joltage = append(input.joltage, int(n))
			}

			if !yield(input) {
				return
			}
		}
	}
}

type PuzzleInput struct {
	// This is the target we want, f.e. [..#..]
	target int16

	// Contains the button inputs, each input defines the bits it toggles
	buttons []int16

	// The raw data instead of an bitmask
	buttonsRaw [][]int

	// Joltage requirements
	joltage []int
}

func (p PuzzleInput) solve() int {
	statesSeen := make(map[int16]struct{})
	iteration := 1
	for _, b := range p.buttons {
		statesSeen[b] = struct{}{}

		// One press solved
		if b == p.target {
			return iteration
		}
	}

	for {
		statesToAdd := []int16{}
		iteration++

		for k := range statesSeen {
			for _, b := range p.buttons {
				nextState := b ^ k
				if _, ok := statesSeen[nextState]; ok {
					// We've already seen this state with less presses, no need
					// to add it
					continue
				}

				// This is a new state
				statesToAdd = append(statesToAdd, nextState)

				if nextState == p.target {
					return iteration
				}
			}
		}

		if len(statesToAdd) == 0 {
			panic("No solution")
		} else {
			for _, v := range statesToAdd {
				statesSeen[v] = struct{}{}
			}
		}
	}
}
