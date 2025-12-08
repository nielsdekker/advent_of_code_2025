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
	solvePartOne(readPuzzle("./day_5/puzzle.txt"))
	solvePartTwo(readPuzzle("./day_5/puzzle.txt"))
}

func solvePartOne(root *Range, ids iter.Seq[uint64]) {
	solution := 0
	for id := range ids {
		if root.contains(id) {
			solution++
		}
	}

	fmt.Printf("Solution %d\n", solution)
}

func solvePartTwo(root *Range, _ iter.Seq[uint64]) {
	// Skip the 0,0 range at the start
	current := root.next
	solution := uint64(0)

	current.print()

	for current != nil {
		solution += ((current.to + 1) - current.from)

		current = current.next
	}

	fmt.Printf("Solution %d\n", solution)
}

func readPuzzle(path string) (*Range, iter.Seq[uint64]) {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	root := &Range{}
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			break
		}

		splitted := strings.Split(line, "-")
		f, err := strconv.ParseInt(splitted[0], 10, 64)
		if err != nil {
			panic(err)
		}
		s, err := strconv.ParseInt(splitted[1], 10, 64)
		if err != nil {
			panic(err)
		}

		root.insert(&Range{
			from: uint64(f),
			to:   uint64(s),
		})
	}

	return root, func(yield func(uint64) bool) {
		for scanner.Scan() {
			line := scanner.Text()
			id, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				panic(err)
			}

			if !yield(uint64(id)) {
				return
			}
		}
	}
}

type Range struct {
	from uint64
	to   uint64
	next *Range
}

func (r *Range) insert(toAdd *Range) {
	current := r
	for current != nil {
		// Check for overlap
		if current.overlaps(toAdd) {
			// Update the current because this overlaps
			current.from = min(current.from, toAdd.from)
			current.to = max(current.to, toAdd.to)

			// It could be that now we overlap with the next item. This would
			// only change the to value
			if current.overlaps(current.next) {
				current.to = current.next.to
				current.next = current.next.next
			}
			return
		}

		// Check if this item is before the next (or the next item is after the
		// other)
		if current.next.isAfter(toAdd) {
			toAdd.next = current.next
			current.next = toAdd
			return
		}

		// No items after this so this should be add the end
		if current.next == nil {
			current.next = toAdd
			return
		}

		// Move to the next item if we didn't return yet
		current = current.next
	}

}

func (r *Range) overlaps(other *Range) bool {
	if r == nil || other == nil {
		return false
	} else if other.from >= r.from && other.from <= r.to {
		return true
	} else if other.to >= r.from && other.to <= r.to {
		return true
	} else if other.from < r.from && other.to > r.to {
		return true
	} else {
		return false
	}
}

func (r *Range) isAfter(other *Range) bool {
	if r == nil {
		return false
	}

	return r.from > other.to
}

func (r *Range) contains(index uint64) bool {
	current := r
	for current != nil {
		if current.from <= index && current.to >= index {
			return true
		}
		current = current.next
	}
	return false
}

func (r *Range) print() {
	fmt.Printf("All ranges:\n")
	current := r
	for current != nil {
		fmt.Printf("\t%d -> %d\n", current.from, current.to)
		current = current.next
	}
}
