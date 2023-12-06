package main

import (
	"bufio"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"
)

const PUZZLE_INPUT_FILE = "input.txt"

func findMatchingNumbers(line string) int {
	numOfMatchingNumbers := 0
	winMap := make(map[int]bool)

	// slog.Debug("Current", "line", line)

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

	// slog.Debug("Numbers", numbers[0], numbers[1])

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
		// slog.Debug("number as string", "number", numStr)
		num, err := strconv.Atoi(numStr)
		if err != nil {
			slog.Error("Error while parsing one of the numbers you have to integer: %s", err)
			os.Exit(1)
		}
		if _, ok := winMap[num]; ok {
			// slog.Debug("Found winning number", "winning number", num)
			numOfMatchingNumbers += 1
		}
	}
	return numOfMatchingNumbers
}

func GetCardId(line string) string {
	return strings.Split(line, ":")[0]
}

func SerializeParams(a int, b []string, c int) string {
	str := fmt.Sprintf("%v:%p:%v", a, b, c)
	return str
}

func Memoize(f func(int, []string, int) int) func(int, []string, int) int {
	mem := make(map[string]int, 0)

	return func(ax int, ay []string, az int) int {
		if val, ok := mem[SerializeParams(ax, ay, az)]; ok {
			return val
		}

		vl := f(ax, ay, az)
		mem[SerializeParams(ax, ay, az)] = vl
		return vl
	}
}

var MemoizedWinScratchcard func(int, []string, int) int

func WinScratchcard(lcards int, cards []string, cardNum int) int {
	if cardNum >= lcards {
		return 0
	}

	numMatchingNum := findMatchingNumbers(cards[cardNum])

	if numMatchingNum == 0 {
		return 1
	}

	// original scratchcard
	total := 1
	for i := 1; i <= numMatchingNum; i++ {
		vl := MemoizedWinScratchcard(lcards, cards, cardNum+i)
		total += vl
	}

	return total
}

func main() {
	start := time.Now()
	MemoizedWinScratchcard = Memoize(WinScratchcard)

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

	lines := make([]string, 0)

	for scanner.Scan() {

		line := scanner.Text()
		if line == "" {
			continue
		}

		lines = append(lines, line)
	}

	scratchcards := 0
	lcards := len(lines)
	for cardNum := range lines {
		scratchcards += MemoizedWinScratchcard(lcards, lines, cardNum)
	}

	fmt.Printf("Number of won scratchcards is %d\n", scratchcards)

	elapsed := time.Since(start)
	fmt.Printf("elapsed time: %s\n", elapsed)
}
