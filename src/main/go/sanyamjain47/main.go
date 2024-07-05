package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Stats struct {
	min   float64
	max   float64
	avg   float64
	sum   float64
	count uint64
}

// A map to store the stats for each string
var statsMap map[string]*Stats

func main() {
	// Take the first argument as the file name to open
	fileName := os.Args[1]
	statsMap = make(map[string]*Stats, 2048)
	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Read the current line
		line := scanner.Text()

		// Split the line by ';'
		parts := strings.Split(line, ";")

		// Extract the values
		location := parts[0]
		valueStr := parts[1]
		// Convert value to float64
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			panic(err)
		}

		if _, ok := statsMap[location]; !ok {
			statsMap[location] = &Stats{
				min:   0,
				max:   0,
				avg:   0,
				sum:   0,
				count: 0,
			}
		}
		if statsMap[location].min > value {
			statsMap[location].min = value
		}
		if statsMap[location].max < value {
			statsMap[location].max = value
		}
		statsMap[location].sum += value
		statsMap[location].count++
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	stations := make([]string, 0, len(statsMap))
	for k := range statsMap {
		stations = append(stations, k)
	}
	sort.Strings(stations)

	fmt.Print("{")
	for i, station := range stations {
		if i > 0 {
			fmt.Print(", ")
		}
		s := statsMap[station]
		mean := s.sum / float64(s.count)
		fmt.Printf("%s=%.1f/%.1f/%.1f", station, s.min, mean, s.max)
	}
	fmt.Println("}")

}
