package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
	"unicode"
)

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

func IsDigit(ch rune) (int, bool) {
	for i, digit := range digits {
		if ch == digit {
			return i, true
		}
	}
	return -1, false
}

const INPUT_PUZZLE_FILE = "input.txt"

var WIDTH int = -1

type Number struct {
	sbuilder strings.Builder
	Indexes  []int
}

func NewNumber() *Number {
	return &Number{
		sbuilder: strings.Builder{},
		Indexes:  make([]int, 0),
	}
}

func (number *Number) AddDigit(digit int, position int) {
	number.sbuilder.WriteString(strconv.Itoa(digit))
	number.Indexes = append(number.Indexes, position)
}

func (number *Number) String() string {
	return number.sbuilder.String()
}

func (number *Number) Value() (int, error) {
	num, err := strconv.Atoi(number.String())
	if err != nil {
		return -1, err
	}
	return num, nil
}

type Symbol struct {
	Position     int
	Value        rune
	adjacentNums []*Number
}

func NewSymbol(position int, value rune) *Symbol {
	return &Symbol{
		Position:     position,
		Value:        value,
		adjacentNums: make([]*Number, 0),
	}
}

func (symbol *Symbol) AddAdjacentNumber(number *Number) {
	symbol.adjacentNums = append(symbol.adjacentNums, number)
}

func (symbol *Symbol) String() string {
	return fmt.Sprintf("%c", symbol.Value)
}

var VALID_SYMBOLS = []rune{'/', '*', '%', '@', '#', '+', '&', '=', '-', '$'}

func (symbol *Symbol) IsValidSymbol() bool {
	return slices.Contains(VALID_SYMBOLS, symbol.Value)
}

type Numbers struct {
	Numbers []*Number
}

func (numbers *Numbers) AddNumber(number *Number) {
	numbers.Numbers = append(numbers.Numbers, number)
}

type Symbols struct {
	Symbols map[int]*Symbol
}

func (symbols *Symbols) AddSymbol(symbol *Symbol) {
	symbols.Symbols[symbol.Position] = symbol
}

type Grid struct {
	numbers *Numbers
	symbols *Symbols
}

func (number *Number) IsValidNumber(grid *Grid) bool {
	_, err := number.Value()
	if err != nil {
		log.Fatalln(err)
	}

	directions := []int{
		-1,         // west
		1,          // east
		-WIDTH,     // north
		WIDTH,      // south
		-WIDTH - 1, // north west
		-WIDTH + 1, // north east
		WIDTH - 1,  // south west
		WIDTH + 1,  // south east
	}

	for _, indx := range number.Indexes {
		for _, direction := range directions {
			if symbol, ok := grid.symbols.Symbols[indx+direction]; ok {
				symbol.AddAdjacentNumber(number)
				return true
			}
		}
	}

	return false
}

func (grid *Grid) FindValidNumbers() []int {
	nums := make([]int, 0)

	for _, num := range grid.numbers.Numbers {
		if num.IsValidNumber(grid) {
			nm, err := num.Value()
			if err != nil {
				log.Fatalln("Error getting the value of Number")
			}
			nums = append(nums, nm)
		}
	}

	return nums
}

func main() {
	start := time.Now()
	f, err := os.Open(INPUT_PUZZLE_FILE)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReaderSize(f, 1024*1024)

	numbers := &Numbers{
		Numbers: make([]*Number, 0),
	}

	symbols := &Symbols{
		Symbols: make(map[int]*Symbol),
	}

	grid := &Grid{
		numbers: numbers,
		symbols: symbols,
	}

	reading := true

	isFirstNewLine := true
	i := 0

	firstNum := true

	var currentNumber *Number
	var currentSymbol *Symbol

	for reading {
		i++
		ch, _, err := reader.ReadRune()
		if err == io.EOF {
			reading = false
			continue
		}
		if ch == '\n' || ch == unicode.ReplacementChar {
			if isFirstNewLine {
				WIDTH = i
				isFirstNewLine = false
			}
			continue
		}

		if digit, isDigit := IsDigit(ch); isDigit {
			if firstNum {
				firstNum = false
				if currentNumber == nil {
					currentNumber = NewNumber()
				}
				currentNumber.AddDigit(digit, i)
			} else {
				currentNumber.AddDigit(digit, i)
			}
		} else {
			// it could a valid symbol or '.'
			if currentSymbol == nil {
				currentSymbol = NewSymbol(i, ch)
				if currentSymbol.IsValidSymbol() {
					symbols.AddSymbol(currentSymbol)
					currentSymbol = nil
				} else {
					currentSymbol = nil
				}
			}
			if !firstNum {
				numbers.AddNumber(currentNumber)

				currentNumber = nil
				firstNum = true
			}
		}
	}

	fmt.Println("Width of the engineering schematic is ", WIDTH)

	validNums := grid.FindValidNumbers()

	sum2 := 0
	for _, symbol := range symbols.Symbols {
		if symbol.Value == '*' && len(symbol.adjacentNums) == 2 {
			num1, err := symbol.adjacentNums[0].Value()
			if err != nil {
				log.Fatalln(err)
			}
			num2, err := symbol.adjacentNums[1].Value()
			if err != nil {
				log.Fatalln(err)
			}
			sum2 += (num1 * num2)
		}
	}

	fmt.Printf("Sum of all of the gear ratios in the engine schematic is %d\n", sum2)

	sum := 0
	for _, vNum := range validNums {
		sum += vNum
	}

	fmt.Println("Sum of all the part numbers is ", sum)
	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %s\n", elapsed)
}
