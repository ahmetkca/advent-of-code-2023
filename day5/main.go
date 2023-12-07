package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

const PUZZLE_INPUT_FILE = "input.txt"

type Range struct {
	Begin  int
	Length int
}

func (rng *Range) End() int {
	return rng.Begin + rng.Length - 1
}

func (rng *Range) IsInRange(num int) bool {
	if num >= rng.Begin && num <= rng.End() {
		return true
	}
	return false
}

type RangeMap struct {
	SrcRange  *Range
	DestRange *Range
}

func (rngMp *RangeMap) GetMappedNumber(num int) (int, error) {
	if !rngMp.SrcRange.IsInRange(num) {
		return -1, fmt.Errorf("The number (%d) is not in the range\n", num)
	}

	diff := rngMp.SrcRange.End() - num
	return rngMp.DestRange.End() - diff, nil
}

type DestToSrcMap struct {
	Destination string
	Source      string
	RangeMaps   []*RangeMap
}

func (dtosMap *DestToSrcMap) FindSrcToDest(src int) int {
	for _, rngMp := range dtosMap.RangeMaps {
		val, err := rngMp.GetMappedNumber(src)
		if err == nil {
			return val
		}
	}

	return src
}

func (dtosMap *DestToSrcMap) String() string {
	strBuilder := strings.Builder{}
	strBuilder.WriteString(
		fmt.Sprintf("Mapping from %s to %s:\n", dtosMap.Source, dtosMap.Destination),
	)

	for _, rngMp := range dtosMap.RangeMaps {
		strBuilder.WriteString(
			fmt.Sprintf("\t%d\t%d\n", rngMp.SrcRange.Begin, rngMp.DestRange.Begin),
		)
		strBuilder.WriteString("\t.\t.\n")
		strBuilder.WriteString("\t.\t.\n")
		strBuilder.WriteString("\t.\t.\n")
		strBuilder.WriteString(
			fmt.Sprintf("\t%d\t%d\n", rngMp.SrcRange.End(), rngMp.DestRange.End()),
		)
	}
	return strBuilder.String()
}

type AllMappings struct {
	SeedToSoil            *DestToSrcMap
	SoilToFertilizer      *DestToSrcMap
	FertilizerToWater     *DestToSrcMap
	WaterToLight          *DestToSrcMap
	LightToTemperature    *DestToSrcMap
	TemperatureToHumidity *DestToSrcMap
	HumidityToLocation    *DestToSrcMap
}

func (allMappings *AllMappings) FindSeedToLocation(seedNum int) (location int) {
	location = allMappings.SeedToSoil.FindSrcToDest(seedNum)
	location = allMappings.SoilToFertilizer.FindSrcToDest(location)
	location = allMappings.FertilizerToWater.FindSrcToDest(location)
	location = allMappings.WaterToLight.FindSrcToDest(location)
	location = allMappings.LightToTemperature.FindSrcToDest(location)
	location = allMappings.TemperatureToHumidity.FindSrcToDest(location)
	location = allMappings.HumidityToLocation.FindSrcToDest(location)
	return
}

func ParseDestinationToSourceMapNameString(line string) (src string, dest string) {
	line = strings.TrimSuffix(line, " map:")
	line = strings.TrimSpace(line)
	destinationAndSourceNames := strings.Split(line, "-to-")
	src = strings.TrimSpace(
		destinationAndSourceNames[0],
	)
	dest = strings.TrimSpace(
		destinationAndSourceNames[1],
	)
	return
}

func NewDestToSrcMap(destination string, source string) *DestToSrcMap {
	return &DestToSrcMap{
		Destination: destination,
		Source:      source,
		RangeMaps:   make([]*RangeMap, 0),
	}
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

func ProcessDestinationToSourceMap(lines []string, dtosMapChnl chan<- *DestToSrcMap) {
	source, destination := ParseDestinationToSourceMapNameString(lines[0])
	dtosMap := NewDestToSrcMap(destination, source)

	// process destionation range start, source range start and range length lines //
	dtosLines := lines[1:]
	for _, dtosLine := range dtosLines {
		destRngStrt, srcRngStrt, rngLngth := ParseDestinationSourceRangeString(dtosLine)

		dtosMap.RangeMaps = append(dtosMap.RangeMaps, &RangeMap{
			SrcRange: &Range{
				Begin:  srcRngStrt,
				Length: rngLngth,
			},
			DestRange: &Range{
				Begin:  destRngStrt,
				Length: rngLngth,
			},
		})
	}
	/////////////////////////////////////////////////////////////////////////////////

	dtosMapChnl <- dtosMap
}

func ParseSeeds(line string) []int {
	seeds := make([]int, 0)
	line = strings.TrimPrefix(line, "seeds: ")
	seedsAsStr := strings.Split(line, " ")
	for _, seedAsStr := range seedsAsStr {
		if seedAsStr != "" {
			seed, err := strconv.Atoi(seedAsStr)
			if err != nil {
				log.Fatalf("Error while parsing seed number: %s\n", err)
			}
			seeds = append(seeds, seed)
		}
	}
	return seeds
}

const NUMBER_OF_DEST_TO_SRC_MAPS = 7

func main() {
	start := time.Now()
	f, err := os.Open(PUZZLE_INPUT_FILE)
	if err != nil {
		log.Fatalf("Error while opening the puzzle input file: %s\n", err)
	}

	scanner := bufio.NewScanner(f)

	// scan the first line for seeds.
	ok := scanner.Scan()
	if !ok {
		log.Fatalf("Error while scanning the first line for seeds: %s\n", err)
	}
	firstLine := scanner.Text()
	seeds := ParseSeeds(firstLine)
	// skip new line
	scanner.Scan()
	scanner.Text()

	fmt.Printf("seeds: %v\n", seeds)
	scanning := true

	dtosMaps := make([]*DestToSrcMap, NUMBER_OF_DEST_TO_SRC_MAPS)

	dtosMapChnl := make(chan *DestToSrcMap)

	for scanning {
		scanning = scanner.Scan()
		lines := make([]string, 0)
		line := scanner.Text()

		for line != "" {
			lines = append(lines, line)
			scanner.Scan()
			line = scanner.Text()
		}

		if len(lines) <= 0 {
			continue
		}
		go ProcessDestinationToSourceMap(lines, dtosMapChnl)

	}

	for i := 0; i < NUMBER_OF_DEST_TO_SRC_MAPS; i++ {
		fmt.Println("Waiting for Source to Destination Mappings...")
		dtosMap := <-dtosMapChnl
		dtosMaps[i] = dtosMap
		fmt.Println(dtosMaps[i])
	}

	allMappings := &AllMappings{}

	for _, dtosMp := range dtosMaps {
		switch dtosMp.Source {
		case "seed":
			allMappings.SeedToSoil = dtosMp
		case "soil":
			allMappings.SoilToFertilizer = dtosMp
		case "fertilizer":
			allMappings.FertilizerToWater = dtosMp
		case "water":
			allMappings.WaterToLight = dtosMp
		case "light":
			allMappings.LightToTemperature = dtosMp
		case "temperature":
			allMappings.TemperatureToHumidity = dtosMp
		case "humidity":
			allMappings.HumidityToLocation = dtosMp
		}
	}

	min := math.MaxInt
	for _, seed := range seeds {
		locationNum := allMappings.FindSeedToLocation(seed)
		if locationNum < min {
			min = locationNum
		}
		fmt.Printf("Seed: %d, Location: %d\n", seed, locationNum)
	}
	fmt.Printf("Lowest location number is %d\n", min)
	fmt.Printf("elapsed time: %s\n", time.Since(start))
}
