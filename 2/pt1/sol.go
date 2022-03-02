package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	inputFile = "input"
)

type Position struct {
	Depth  int
	HozPoz int
}

func (p *Position) down(meters int) {
	p.Depth += meters
}

func (p *Position) up(meters int) {
	if p.Depth > 0 {
		p.Depth -= meters
	}
}

func (p *Position) forward(meters int) {
	p.HozPoz += meters
}

func main() {

	// open and scan
	fd, err := os.Open(inputFile)
	HandleErr(err)
	scanner := bufio.NewScanner(fd)

	// create submarine
	sub := Position{
		Depth:  0,
		HozPoz: 0,
	}

	// iterate file
	for scanner.Scan() {
		command := strings.Split(scanner.Text(), " ")
		if len(command) != 2 {
			log.Panicf("Bad command")
		}

		distance, err := strconv.Atoi(command[1])
		HandleErr(err)
		switch command[0] {
		case "forward":
			sub.forward(distance)
		case "down":
			sub.down(distance)
		case "up":
			sub.up(distance)
		default:
			log.Panicf("Bad command")
		}

	}
	fmt.Printf("Total distance: %d\n", sub.HozPoz*sub.Depth)
}

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
