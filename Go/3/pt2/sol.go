package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	inputFile   = "input"
	bits        = 12
	numeralBase = 2
)

func filterArray(bitArray []string, idx int, most bool) []string {

	// get most common bit
	total := 0.0
	for _, b := range bitArray {
		if b[idx] == '1' {
			total++
		}
	}

	// least if not most common
	bit := '0'
	if (total/float64(len(bitArray)) >= 0.5) == most {
		bit = '1'
	}

	// create reduced array
	var newArray []string
	for _, b := range bitArray {
		if rune(b[idx]) == bit {
			newArray = append(newArray, b)
		}
	}
	return newArray
}

func main() {

	// get entire file
	bins, err := os.ReadFile(inputFile)
	HandleErr(err)

	// split into binary strings
	OXArray := strings.Split(string(bins), "\n")
	C02Array := strings.Split(string(bins), "\n")

	// remove null entry
	OXArray = OXArray[:len(OXArray)-1]
	C02Array = C02Array[:len(C02Array)-1]

	// iterate
	for i := 0; i < bits; i++ {

		// reduce oxygen array
		if len(OXArray) > 1 {
			OXArray = filterArray(OXArray, i, true)
		}

		// reduce c02 array
		if len(C02Array) > 1 {
			C02Array = filterArray(C02Array, i, false)
		}
	}

	// convert binary (base 2) to int
	C02, err := strconv.ParseInt(C02Array[0], 2, 64)
	HandleErr(err)
	OX, err := strconv.ParseInt(OXArray[0], 2, 64)
	HandleErr(err)

	// done
	fmt.Println("Solution:")
	fmt.Printf("Oxygen generator rating = %0*b = %d\n", bits, OX, OX)
	fmt.Printf("CO2 scrubber rating     = %0*b = %d\n", bits, C02, C02)
	fmt.Printf("power consumption = %d\n", OX*C02)
}

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
