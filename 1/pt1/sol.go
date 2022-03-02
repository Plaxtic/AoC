package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	total := 0
	prev := math.MaxInt

	// iterate file
	for scanner.Scan() {

		line, err := strconv.Atoi(scanner.Text())
		HandleErr(err)
		if line > prev {
			total++
		}

		if !scanner.Scan() {
			break
		}
		prev, err = strconv.Atoi(scanner.Text())
		HandleErr(err)
		if line < prev {
			total++
		}
	}
	fmt.Printf("Total increases : %d\n", total)
}

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
