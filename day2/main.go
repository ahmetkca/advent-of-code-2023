package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const INPUT_PUZZLE_FILE = "input.txt"

// R, G, B
var constraints = [3]int{
	12,
	13,
	14,
}

var colorToConstraint = map[string]int{
	"red":   0,
	"green": 1,
	"blue":  2,
}

func main() {
	f, err := os.Open(INPUT_PUZZLE_FILE)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	gameNum := 1
	sum := 0

	sum2 := 0
	for scanner.Scan() {
		fmt.Printf("Game #%d\n", gameNum)
		line := scanner.Text()

		line = strings.SplitAfter(line, ":")[1]

		sets := strings.Split(line, ";")

		isItPossible := true
		fewestCubesForEachColor := map[string]int{
			"red":   -1,
			"green": -1,
			"blue":  -1,
		}
		for i, set := range sets {
			fmt.Printf("\tSet #%d\n", i+1)
			cubes := strings.Split(set, ",")
			for _, cubesWithColor := range cubes {
				cubesWithColor = strings.TrimSpace(cubesWithColor)
				cwc := strings.Split(cubesWithColor, " ")
				numOfCubes, err := strconv.Atoi(cwc[0])
				if err != nil {
					panic(err)
				}
				color := cwc[1]
				fmt.Printf("\t\tColor: %s, Num of cubes: %d\n", color, numOfCubes)
				if numOfCubes > constraints[colorToConstraint[color]] {
					fmt.Printf("\t\t\tOver constraint: %d > %d\n", numOfCubes, constraints[colorToConstraint[color]])
					isItPossible = false
				}

				if fewestCubesForEachColor[color] == -1 || numOfCubes > fewestCubesForEachColor[color] {
					fewestCubesForEachColor[color] = numOfCubes
				}
			}
		}

		powerOfCubes := 1
		for color, numOfCubes := range fewestCubesForEachColor {
			fmt.Printf("\tGame #%d, Fewest cubes for color %s: %d\n", gameNum, color, numOfCubes)
			powerOfCubes *= numOfCubes
		}
		fmt.Printf("\tGame #%d, Power of cubes: %d\n", gameNum, powerOfCubes)
		sum2 += powerOfCubes

		if isItPossible {
			fmt.Printf("\tGame #%d is possible\n", gameNum)
			sum += gameNum
		}

		gameNum++
	}

	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Sum of power of cubes: %d\n", sum2)
}
