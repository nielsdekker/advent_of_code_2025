package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type OPERATION int

const (
	ADD OPERATION = iota
	SUB
	MUL
	DIV
)

func main() {
	data := parsePuzzle()

	solvePartTwo(data)
}

func solvePartOne(data []Data) {
	totalSum := int64(0)

	for _, d := range data {
		totalSum += applyOperation(d.nums, d.operation)
	}

	fmt.Printf("Solution %d\n", totalSum)
}

func solvePartTwo(data []Data) {
	totalSum := int64(0)

	for _, d := range data {
		cephalodNums := make([]int64, 0)

		for i := range len(d.raw[0]) {
			acc := ""
			for _, r := range d.raw {
				rune := r[i]
				if rune == ' ' {
					continue
				} else {
					acc += string(rune)
				}
			}

			if len(acc) > 0 {
				asInt, err := strconv.ParseInt(acc, 10, 64)
				if err != nil {
					panic(err)
				}

				cephalodNums = append(cephalodNums, asInt)
			}
		}

		totalSum += applyOperation(cephalodNums, d.operation)
	}

	fmt.Printf("Solution %d\n", totalSum)
}

func applyOperation(nums []int64, operation OPERATION) int64 {
	answer := nums[0]

	switch operation {
	case ADD:
		for i := 1; i < len(nums); i++ {
			answer += nums[i]
		}
	case SUB:
		for i := 1; i < len(nums); i++ {
			answer -= nums[i]
		}
	case DIV:
		for i := 1; i < len(nums); i++ {
			answer /= nums[i]
		}
	case MUL:
		for i := 1; i < len(nums); i++ {
			answer *= nums[i]
		}
	}

	return answer
}

func parsePuzzle() []Data {
	scanner := bufio.NewScanner(os.Stdin)

	// Get all the lines
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// The last line contains the operations and also determines how the data is
	// splitted
	lastLine := lines[len(lines)-1]
	prevOperation := ADD
	data := []Data{}
	startIndex := 0
	for i, r := range lastLine {
		if r != ' ' {
			if i != startIndex {
				raw := make([]string, len(lines)-1)
				for n := 0; n < len(lines)-1; n++ {
					raw[n] = lines[n][startIndex:i]
				}

				data = append(data, Data{
					raw:       raw,
					operation: prevOperation,
				})

				startIndex = i
			}
			prevOperation = getOperation(r)
		}
	}

	// Also add the remainder
	raw := make([]string, len(lines)-1)
	for n := 0; n < len(lines)-1; n++ {
		raw[n] = lines[n][startIndex:]
	}

	data = append(data, Data{
		raw:       raw,
		operation: prevOperation,
	})

	// Parse the nums
	for i := range data {
		d := &data[i]

		d.nums = make([]int64, len(d.raw))
		for i, r := range d.raw {
			parsed, err := strconv.ParseInt(strings.TrimSpace(r), 10, 64)
			if err != nil {
				panic(err)
			}

			d.nums[i] = parsed
		}
	}

	return data
}

func getOperation(r rune) OPERATION {
	switch r {
	case '+':
		return ADD
	case '-':
		return SUB
	case '*':
		return MUL
	default:
		return DIV
	}
}

type Data struct {
	raw       []string
	nums      []int64
	operation OPERATION
}
