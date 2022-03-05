package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	inputFile   = "input"
	bits        = 12
	numeralBase = 2
)

func fileIter(path string) *bufio.Scanner {

	// open and scan
	fd, err := os.Open(path)
	HandleErr(err)
	scanner := bufio.NewScanner(fd)

	return scanner
}

func main() {

	// get file scanner
	scanner := fileIter(inputFile)

	// get most common bits
	var bitCount [bits][numeralBase]float64
	entries := 0.0
	for scanner.Scan() {
		binNum := scanner.Text()

		for i, c := range binNum {
			if c == '0' {
				bitCount[i][0]++
			} else {
				bitCount[i][1]++
			}
		}
		entries++
	}

	// calculate gamma and epsilon
	var gamma uint
	var epsilon uint
	for i, b := range bitCount {
		if b[0] > b[1] {
			epsilon += (1 << (bits - i - 1)) // remember we are iterating "backwards"
		} else if b[0] < b[1] {
			gamma += (1 << (bits - i - 1))
		} else {
			log.Panic("Equal bit frequency, implement rounding")
		}
	}
	fmt.Println("Solution:")
	fmt.Printf("gamma rate   = %0*b = %d\n", bits, gamma, gamma)
	fmt.Printf("epsilon rate = %0*b = %d\n", bits, epsilon, epsilon)
	fmt.Printf("power consumption = %d\n", epsilon*gamma)
}

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
