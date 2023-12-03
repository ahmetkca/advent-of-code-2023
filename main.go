package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
)

const PUZZLE_INPUT_FILE = "input.txt"

var digits = [...]rune{
	rune('0'),
	rune('1'),
	rune('2'),
	rune('3'),
	rune('4'),
	rune('5'),
	rune('6'),
	rune('7'),
	rune('8'),
	rune('9'),
}

var digitSpelledOutToNumberMap = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func IsDigit(ch rune) (int, bool) {
	for i, digit := range digits {
		if ch == digit {
			return i, true
		}
	}
	return -1, false
}

// find the last digit in the input
// last digit can be in either number opr spelled out format
// ex. eight or 8
func FindLast(input []byte) int {
	line := string(input)
	lastIndex := -1
	last := -1

	for spelledOutDigit, digitNum := range digitSpelledOutToNumberMap {
		index := strings.LastIndex(line, spelledOutDigit)
		fmt.Printf("index: %d, lastIndex: %d, spelledOutDigit: %s, last: %d\n", index, lastIndex, spelledOutDigit, last)
		if index != -1 {
			if index >= lastIndex {
				last = digitNum
				lastIndex = index
			}
		}
	}

	for i, ch := range line {
		if digitNum, isDigit := IsDigit(ch); isDigit {
			fmt.Printf("index: %d, lastIndex: %d, last: %d", i, lastIndex, last)
			if i > lastIndex {
				last = digitNum
				lastIndex = i
			}
		}
	}

	return last
}

// Find the first digit in the input
// digit can be in number or spelled out format
// ex. one or 1
func FindFirst(input []byte) int {
	line := string(input)
	firstIndex := math.MaxInt
	first := -1

	for spelledOutDigit, digitNum := range digitSpelledOutToNumberMap {
		index := strings.Index(line, spelledOutDigit)
		if index != -1 {
			if index <= firstIndex {
				first = digitNum
				firstIndex = index
			}
		}
	}

	for i, ch := range line {
		if digitNum, isDigit := IsDigit(ch); isDigit {
			if i < firstIndex {
				first = digitNum
				firstIndex = i
			}
		}
	}

	return first
}

func GetFirstAndLast(input []byte) (int, int) {
	first, last := FindFirst(input), FindLast(input)

	if last == -1 {
		return first, first
	}
	return first, last
}

func main() {
	f, err := os.Open(PUZZLE_INPUT_FILE)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(f)

	sum := 0
	running := true

	for running {
		data, err := reader.ReadBytes('\n')

		log.Printf("line = %s", string(data))

		first, last := GetFirstAndLast(data)

		calibrationValue := 10*int(first) + int(last)

		fmt.Printf("calibration value = %d\n", calibrationValue)

		sum += calibrationValue

		if err == io.EOF {
			running = false
		}

		log.Println(err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
