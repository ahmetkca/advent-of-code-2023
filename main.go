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

func DoesItStartWithSpelledOutDigit(spelledOutDigit string) (int, bool) {
	for k := range digitSpelledOutToNumberMap {
		if strings.HasPrefix(k, spelledOutDigit) {
			digit, ok := digitSpelledOutToNumberMap[spelledOutDigit]
			if ok {
				return digit, true
			} else {
				return -1, true
			}
		}
	}
	return -1, false
}

func IsDigit(ch rune) (int, bool) {
	for i, digit := range digits {
		if ch == digit {
			return i, true
		}
	}
	return -1, false
}

func GetFirstAndLastDigitPartOne(input []byte) (int8, int8) {
	first, last := -1, -1
	for _, ch := range input {
		digit, isDigit := IsDigit(rune(ch))
		if isDigit {
			if first == -1 {
				first = int(digit)
			} else {
				last = int(digit)
			}
		}
	}
	if last == -1 {
		last = first
	}
	return int8(first), int8(last)
}

func GetFirstAndLastDigitSecondPart(input []byte) (int, int) {
	currStr := &strings.Builder{}

	first, last := -1, -1
	for _, ch := range input {
		// log.Printf("ch = %c\n", rune(ch))

		digit, isDigit := IsDigit(rune(ch))
		if isDigit {
			if first == -1 {
				first = int(digit)
			} else {
				last = int(digit)
			}
			currStr.Reset()
		} else {
			currStr.WriteByte(ch)
			// fmt.Printf("currStr = %s\n", currStr.String())
			dgt, doesItStartWithCurrStr := DoesItStartWithSpelledOutDigit(currStr.String())

			if dgt == -1 && doesItStartWithCurrStr == true {
				continue
			}

			if dgt != -1 && doesItStartWithCurrStr == true {
				// we found a digit and need to reset currStr to start over
				str := currStr.String()
				lastChar := str[len(str)-1]
				currStr.Reset()
				if len(str) > 1 {
					currStr.WriteByte(lastChar)
				}
				if first == -1 {
					first = dgt
				} else {
					last = dgt
				}
				continue
			}

			if dgt == -1 && doesItStartWithCurrStr == false {
				str := currStr.String()
				lastChar := str[len(str)-1]
				currStr.Reset()
				if len(str) > 1 {
					currStr.WriteByte(lastChar)
				}
			}
		}
	}

	if last == -1 {
		last = first
	}

	return first, last
}

// returns -1, -1 if not successful
func GetFirstAndLast1(input []byte) (int, int) {
	keys := make([]string, 0, len(digitSpelledOutToNumberMap))
	for k := range digitSpelledOutToNumberMap {
		keys = append(keys, k)
	}

	line := string(input)
	first, last := -1, -1
	for len(line) > 0 {
		ch := line[0]
		digit, isDigit := IsDigit(rune(ch))
		hasPrefix := false
		if isDigit {
			if first == -1 {
				first = int(digit)
			} else {
				last = int(digit)
			}
		} else {
			for _, k := range keys {
				fmt.Printf("%s, %s\n", line, k)
				if strings.HasPrefix(line, k) {
					line = strings.TrimPrefix(line, k)
					digit, ok := digitSpelledOutToNumberMap[k]
					fmt.Printf("Found: %d, %v\n", digit, ok)
					if ok {
						if first == -1 {
							first = digit
						} else {
							last = digit
						}
						hasPrefix = true
						break
					}
				}
			}
		}
		if !hasPrefix {
			line = line[1:]
		}
	}

	if last == -1 {
		last = first
	}

	return first, last
}

// find the last digit in the input
// last digit can be in either number opr spelled out format
// ex. eight or 8
func FindLast(input []byte) int {
	line := string(input)
	lastIndex := 1
	last := -1

	for spelledOutDigit, digitNum := range digitSpelledOutToNumberMap {
		index := strings.LastIndex(line, spelledOutDigit)
		if index != -1 {
			if last >= lastIndex {
				last = digitNum
				lastIndex = index
			}
		}
	}

	for i, ch := range line {
		if digitNum, isDigit := IsDigit(ch); isDigit {
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
		data = []byte("mkfone4ninefour")

		log.Printf("line = %s", string(data))

		first, last := GetFirstAndLast(data)

		calibrationValue := 10*int(first) + int(last)

		fmt.Printf("calibration value = %d\n", calibrationValue)

		sum += calibrationValue
		//////////////

		if err == io.EOF {
			running = false
		}

		log.Println(err)
	}

	fmt.Printf("Sum: %d\n", sum)
}
