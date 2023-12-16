package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Seed struct {
	Start  int
	Length int
}

func (s *Seed) SeedToRange() *Range {
	return &Range{
		Start:  s.Start,
		Length: s.Length,
	}
}

func (s *Seed) String() string {
	return fmt.Sprintf("Start of the range: %d, Length of the range: %d", s.Start, s.Length)
}

type SrcToDestMap struct {
	SourceString      string
	DestinationString string
	Map               []*SingleRangeMap
}

func ParseDestinationSourceRangeString(line string) (int, int, int) {
	line = strings.TrimSpace(line)
	nums := make([]int, 3)
	for i, numStr := range strings.Split(line, " ") {
		if numStr == "" {
			continue
		}
		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatalf(
				"Error while parsing destination range start or source range start or range length number to integer: %s\n",
				err,
			)
		}
		nums[i] = num
	}

	return nums[0], nums[1], nums[2]
}

func main() {
	start := time.Now()
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln("Error: while opening the input file, ", err)
	}

	scnr := bufio.NewScanner(f)

	// read seeds
	if !scnr.Scan() {
		log.Fatalln("Error: while scanning the first line for seeds, ", err)
	}
	seedsStr := scnr.Text()

	seedsNumsStr := strings.TrimPrefix(seedsStr, "seeds: ")
	seedsNumsStrSplit := strings.Split(seedsNumsStr, " ")

	pipelineFirstStageInputChannel := make(chan interface{})

	go func() {
		defer close(pipelineFirstStageInputChannel)
		// TODO: change with slog
		// log.Println("Seeds: ")
		for i := 0; i < len(seedsNumsStrSplit); i += 2 {
			seedStartRange, err := strconv.Atoi(seedsNumsStrSplit[i])
			if err != nil {
				log.Fatalln("Error: while parsing seed start range number string to integer, ", err)
			}
			seedLengthRange, err := strconv.Atoi(seedsNumsStrSplit[i+1])
			if err != nil {
				log.Fatalln("Error: while parsing seed length range number string to integer, ", err)
			}

			seed := &Range{
				Start:  seedStartRange,
				Length: seedLengthRange,
			}
			// TODO: change with slog
			// log.Println(seed)
			pipelineFirstStageInputChannel <- seed
		}
	}()

	scnr.Scan() // skip the second line

	scanning := true

	categories := make([]*SrcToDestMap, 0)

	for scanning {
		scanning = scnr.Scan()
		lines := make([]string, 0)
		line := scnr.Text()

		for line != "" {
			lines = append(lines, line)
			scnr.Scan()
			line = scnr.Text()
		}

		if len(lines) <= 0 {
			continue
		}

		srcToDestMap := &SrcToDestMap{
			Map: make([]*SingleRangeMap, 0),
		}

		for lineNum, lne := range lines {
			if lineNum == 0 {
				lne = strings.TrimSuffix(lne, " map:")
				lne = strings.TrimSpace(lne)
				srcToDestMap.SourceString = strings.Split(lne, "-to-")[0]
				srcToDestMap.DestinationString = strings.Split(lne, "-to-")[1]
			} else {
				destStart, srcStart, length := ParseDestinationSourceRangeString(lne)
				srcToDestMap.Map = append(srcToDestMap.Map, &SingleRangeMap{
					SourceRange:      &Range{srcStart, length},
					DestinationRange: &Range{destStart, length},
				})
			}
		}

		categories = append(categories, srcToDestMap)

		// TODO: change with slog
		// fmt.Println("Source: ", srcToDestMap.SourceString)
		// fmt.Println("Destination: ", srcToDestMap.DestinationString)
		// fmt.Println("Map: ")
		// for _, singleRangeMap := range srcToDestMap.Map {
		// 	fmt.Println(singleRangeMap)
		// }
	}

	builder := NewPipelineBuilder(context.Background())

	for i := 0; i < len(categories); i++ {
		builder.AddStage(&SrcToDestStage{
			srcToDestMap: categories[i],
		})
	}

	pipelineOutputChannel := builder.Build(pipelineFirstStageInputChannel)

	lowestLocation := math.MaxInt
	for v := range pipelineOutputChannel {
		locationRange := v.(*Range)

		if locationRange.Start < lowestLocation {
			lowestLocation = locationRange.Start
		}
	}

	fmt.Println("Lowest location: ", lowestLocation)
	fmt.Printf("Time taken: %s\n", time.Since(start))
}
