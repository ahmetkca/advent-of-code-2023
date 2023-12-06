package main

import (
	"bufio"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

const PUZZLE_INPUT_FILE = "input.txt"

func main() {
	logLevelEnv, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		logLevelEnv = ""
	}

	var logLevel slog.Level
	switch logLevelEnv {
	case "DEBUG":
		logLevel = slog.LevelDebug
	default:
		logLevel = slog.LevelInfo
	}

	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: logLevel})
	slog.SetDefault(slog.New(handler))

	f, err := os.Open(PUZZLE_INPUT_FILE)
	if err != nil {
		log.Fatalf("Error reading the puzzle input file: %s", err)
	}

	scanner := bufio.NewScanner(f)

	totalPoints := 0

	for scanner.Scan() {
		points := 0

		winMap := make(map[int]bool)
		line := scanner.Text()
		if line == "" {
			continue
		}

		slog.Debug("Current", "line", line)

		_, afterLine, found := strings.Cut(line, ":")
		if !found {
			log.Fatalln("Expected card number seperated by ':'")
		}
		line = afterLine

		numbers := strings.Split(line, "|")
		if len(numbers) != 2 {
			log.Fatalln("Expected 2 list of numbers seperated by '|'")
		}

		numbers[0] = strings.TrimSpace(numbers[0])
		numbers[1] = strings.TrimSpace(numbers[1])

		slog.Debug("Numbers", numbers[0], numbers[1])

		for _, wNumStr := range strings.Split(numbers[0], " ") {
			if wNumStr == "" {
				continue
			}
			wNum, err := strconv.Atoi(wNumStr)
			if err != nil {
				slog.Error("Error while parsing winning number to integer: %s", err)
				os.Exit(1)
			}
			winMap[wNum] = true
		}

		for _, numStr := range strings.Split(numbers[1], " ") {
			if numStr == "" {
				continue
			}
			slog.Debug("number as string", "number", numStr)
			num, err := strconv.Atoi(numStr)
			if err != nil {
				slog.Error("Error while parsing one of the numbers you have to integer: %s", err)
				os.Exit(1)
			}
			if _, ok := winMap[num]; ok {
				slog.Debug("Found winning number", "winning number", num)
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
				slog.Debug("Currently accumulated points", "points", points)
			}
		}

		totalPoints += points
	}
	fmt.Printf("Elf's pile of scratchcards worth %d points\n", totalPoints)
}
