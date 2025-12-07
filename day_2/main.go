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
	fmt.Printf("solution: %d\n", partTwo())
}

func partOne() uint64 {
	invalidIdSum := uint64(0)

	for v := range puzzleIterator("./day_2/puzzle.txt") {
		for i := v.first; i <= v.last; i++ {
			asString := fmt.Sprintf("%d", i)
			l := len(asString)
			if l%2 != 0 {
				// Uneven so this will never match
				continue
			}

			if asString[0:l/2] == asString[l/2:] {
				invalidIdSum += i
			}
		}
	}

	return invalidIdSum
}

func partTwo() uint64 {
	invalidIdSum := uint64(0)
	for v := range puzzleIterator("./day_2/puzzle.txt") {
		for i := v.first; i <= v.last; i++ {
			s := fmt.Sprintf("%d", i)
			l := len(s)

			invalidIdFound := false
			for n := 2; n <= l && !invalidIdFound; n++ {
				if parts, ok := splitString(s, n); ok {
					// Check if all parts are equal
					prev := parts[0]
					allMatch := true
					for _, v := range parts {
						if v != prev {
							allMatch = false
							break
						}
					}

					if allMatch {
						fmt.Printf("Invalid id found: %s, [%d]\n", s, n)
						invalidIdFound = true
					}
				}
			}

			if invalidIdFound {
				invalidIdSum += i
			}
		}
	}

	return invalidIdSum
}

func splitString(s string, n int) ([]string, bool) {
	if len(s)%n != 0 {
		// Not split able
		return []string{}, false
	}

	res := make([]string, n)
	splitSize := len(s) / n

	for i := range n {
		res[i] = s[i*splitSize : (i+1)*splitSize]
	}

	return res, true
}

type Range struct {
	first uint64
	last  uint64
}

func puzzleIterator(path string) iter.Seq[Range] {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(func(data []byte, atEof bool) (int, []byte, error) {
		if atEof && len(data) == 0 {
			return 0, nil, nil
		}

		if i := strings.Index(string(data), ","); i >= 0 {
			return i + 1, data[0:i], nil
		} else if atEof {
			return len(data), data, nil
		} else {
			return 0, nil, nil
		}
	})

	return func(yield func(Range) bool) {
		for scanner.Scan() {
			data := scanner.Text()
			data = strings.TrimSpace(data)

			if len(data) == 0 {
				continue
			}

			splitted := strings.Split(data, "-")
			first, err := strconv.ParseUint(splitted[0], 10, 64)
			if err != nil {
				panic(err)
			}
			last, err := strconv.ParseUint(splitted[1], 10, 64)
			if err != nil {
				panic(err)
			}

			if !yield(Range{
				first: first,
				last:  last,
			}) {
				return
			}
		}
	}
}
