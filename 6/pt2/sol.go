package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Lanterns struct {
	Population map[int]int
	Day        uint
}

const (
	ageRange = 9
)

func copyMap(dest map[int]int, source map[int]int) map[int]int {
	for k, v := range source {
		dest[k] = v
	}
	return dest
}

// has to be faster as previous function was exponential, the speed of this is constant
func (l *Lanterns) stepOneDay() {

	// make temporary map
	newLanterns := make(map[int]int)
	copyMap(newLanterns, l.Population)

	// age all lanterns by 1 day
	for i := 0; i < 8; i++ {
		newLanterns[i] += l.Population[i+1]
		newLanterns[i] -= l.Population[i]
	}

	// birth new lanterns with 8 days till mitosis
	newLanterns[6] += l.Population[0]
	newLanterns[8] += l.Population[0]
	newLanterns[8] -= l.Population[8]

	// update population, next day
	copyMap(l.Population, newLanterns)
	l.Day++
}

func (l *Lanterns) totalPopulation() (total int) {
	for _, num := range l.Population {
		total += num
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

	// create lanterns object
	lanterns := Lanterns{
		Population: make(map[int]int),
		Day:        0,
	}

	// fill Population map with ages
	for _, str := range input {
		age, err := strconv.Atoi(str)
		HandleErr(err)
		lanterns.Population[age]++
	}

	// let 256 days pass
	for d := 0; d < 256; d++ {
		lanterns.stepOneDay()
	}
	fmt.Printf("Population size after 256 days %d\n", lanterns.totalPopulation())
}
