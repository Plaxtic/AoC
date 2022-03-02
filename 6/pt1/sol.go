package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Lanterns struct {
	Population []uint
	Day        uint
}

func (l *Lanterns) stepOneDay() {
	newLanterns := l.Population
	for i, lantern := range l.Population {
		if lantern == 0 {
			newLanterns = append(newLanterns, 8)
			newLanterns[i] = 6
		} else {
			newLanterns[i]--
		}
	}
	l.Population = newLanterns
	l.Day++
}

func (l *Lanterns) printPopulation() {
	fmt.Printf("Day %3d: ", l.Day)
	for _, lantern := range l.Population {
		fmt.Printf("%2d ", lantern)
	}
	fmt.Println("")
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
	var ages []uint
	for _, str := range input {
		age, err := strconv.Atoi(str)
		HandleErr(err)
		ages = append(ages, uint(age))
	}

	// create lanterns object
	lanterns := Lanterns{
		Population: ages,
		Day:        0,
	}

	// let 80 days pass
	for d := 0; d < 80; d++ {
		lanterns.stepOneDay()
		fmt.Printf("Day: %d\n", lanterns.Day)
	}
	fmt.Printf("Population size after 80 days %d\n", len(lanterns.Population))
}

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
