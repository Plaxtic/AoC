package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	inputFile = "input"
)

func main() {

	// open and scan
	fd, err := os.Open(inputFile)
	HandleErr(err)
	scanner := bufio.NewScanner(fd)

	// initialise first window
	var prevWindow [3]int
	var currWindow [3]int
	scanner.Scan()
	currWindow[0], err = strconv.Atoi(scanner.Text())
	HandleErr(err)
	scanner.Scan()
	currWindow[1], err = strconv.Atoi(scanner.Text())
	HandleErr(err)
	scanner.Scan()
	currWindow[2], err = strconv.Atoi(scanner.Text())
	HandleErr(err)

	// iterate file
	total := 0
	for scanner.Scan() {

		// move window forward one and compare
		prevWindow[0] = currWindow[1]
		prevWindow[1] = currWindow[2]
		prevWindow[2], err = strconv.Atoi(scanner.Text())
		HandleErr(err)
		if sum(prevWindow) > sum(currWindow) {
			total++
		}

		// move window forward one and compare
		currWindow[0] = prevWindow[1]
		currWindow[1] = prevWindow[2]
		if !scanner.Scan() {
			break
		}
		currWindow[2], err = strconv.Atoi(scanner.Text())
		HandleErr(err)
		if sum(prevWindow) < sum(currWindow) {
			total++
		}
	}
	fmt.Printf("Total increases : %d\n", total)
}

func sum(window [3]int) (total int) {
	for _, i := range window {
		total += i
	}
	return total
}

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
