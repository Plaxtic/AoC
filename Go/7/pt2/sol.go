package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func calcFuel(amount int) (total int) { // I know there must be some formula for this
	for amount > 0 {
		total += amount
		amount--
	}
	return total
}

func totalFuelRequired(crabs map[int]int, position int) (total int) {
	for pos, nCrabs := range crabs {
		if pos > position {
			total += calcFuel(pos-position) * nCrabs
		} else if pos < position {
			total += calcFuel(position-pos) * nCrabs
		}
	}
	return total
}

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {

	// allow alternative file selection
	path := "input"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	// open and scan
	data, err := os.ReadFile(path)
	HandleErr(err)

	// remove newline
	dataStr := strings.TrimSuffix(string(data), "\n")

	// parse to list
	input := strings.Split(dataStr, ",")

	// convert to integers
	maxPos := 0
	crabs := make(map[int]int)
	for _, pStr := range input {
		pos, err := strconv.Atoi(pStr)
		HandleErr(err)
		crabs[pos]++

		if pos > maxPos {
			maxPos = pos
		}
	}

	// brute force
	minFuel := math.MaxInt
	bestPos := 0
	for p := 0; p <= maxPos; p++ {
		fuel := totalFuelRequired(crabs, p)
		if fuel < minFuel {
			minFuel = fuel
			bestPos = p
		}
	}
	fmt.Printf("Min fuel: %d at pos %d\n", minFuel, bestPos)
}
