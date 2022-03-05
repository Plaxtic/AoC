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
	boardSiz = 5
)

type Board struct{ Numbers [][]BingoNumber }
type BingoNumber struct {
	Value  int
	Strike bool
}

func checkLine(line []BingoNumber) bool {
	for _, n := range line {
		if !n.Strike {
			return false
		}
	}
	return true
}

func (b *Board) strikeBoard(num int) {
	for i, line := range b.Numbers {
		for j, bn := range line {
			if bn.Value == num {
				b.Numbers[i][j].Strike = true
			}
		}
	}
}

func (b *Board) hasWon() bool {
	for i, line := range b.Numbers {

		// check lines
		if checkLine(line) {
			return true
		}

		// check column
		winner := true
		for j := 0; j < boardSiz; j++ {
			if !b.Numbers[j][i].Strike {
				winner = false
			}
		}
		if winner {
			return true
		}
	}
	return false
}

func strikeAllBoards(boards []*Board, num int) *Board {
	for _, board := range boards {
		board.strikeBoard(num)
		if board.hasWon() {
			return board
		}
	}
	return nil
}

func getBoard(scanner *bufio.Scanner) *Board {

	// skip empty line
	if !scanner.Scan() {
		return nil
	}

	// get numbers
	board := &Board{}
	for i := 0; i < boardSiz && scanner.Scan(); i++ {

		var line []BingoNumber
		for _, n := range strings.Split(scanner.Text(), " ") {
			if len(n) != 0 {
				num, err := strconv.Atoi(n)
				HandleErr(err)
				bNum := BingoNumber{
					Value:  num,
					Strike: false,
				}
				line = append(line, bNum)
			}
		}
		if len(line) < boardSiz {
			log.Panic("Bad line")
		}
		board.Numbers = append(board.Numbers, line)
	}
	if len(board.Numbers) < boardSiz {
		log.Panic("Bad line")
		return nil
	}
	return board
}

func initGame(inputPath string) (inputs []int, players []*Board) {

	// open and scan file
	fd, err := os.Open(inputPath)
	HandleErr(err)
	scanner := bufio.NewScanner(fd)

	// get first line: input
	scanner.Scan()

	for _, n := range strings.Split(scanner.Text(), ",") {
		num, err := strconv.Atoi(n)
		HandleErr(err)
		inputs = append(inputs, num)
	}

	for {
		player := getBoard(scanner)
		if player == nil {
			return inputs, players
		}
		players = append(players, player)
	}
}

func main() {

	// initialize with input file
	var path string
	if len(os.Args) < 2 {
		path = "input"

	} else {
		path = os.Args[1]
	}
	inputs, boards := initGame(path)
	fmt.Printf("%d bingo boards loaded\n", len(boards))

	// iterate through input
	for round, input := range inputs {
		fmt.Printf("Round: %3d\r", round)

		winner := strikeAllBoards(boards, input)
		if winner != nil {

			// get score of board
			score := 0
			for _, line := range winner.Numbers {
				for _, bn := range line {
					if !bn.Strike {
						score += bn.Value
					}
				}
			}

			// multiply by final num
			score *= input
			fmt.Println("\nBINGO!")
			fmt.Printf("Winners score: %d\n", score)
			break
		}
	}
}

func (b *Board) printBoard() {
	for _, line := range b.Numbers {
		for _, n := range line {
			strike := " "
			if n.Strike {
				strike = "X"
			}
			fmt.Printf("%2d%s ", n.Value, strike)
		}
		fmt.Printf("\n")
	}
}

func printAll(boards []*Board) {
	for _, b := range boards {
		b.printBoard()
		fmt.Printf("\n")
	}
}

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
