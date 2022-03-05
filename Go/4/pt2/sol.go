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

type Game struct {
	Inputs  []int
	Players []*Board
}

type Board struct {
	Numbers [][]BingoNumber
	Win     bool
}

type BingoNumber struct {
	Value  int
	Strike bool
}

// board methods
func getBoard(scanner *bufio.Scanner) *Board {

	// skip empty line
	if !scanner.Scan() {
		return nil
	}

	// init board
	board := &Board{}
	board.Win = false

	// get numbers
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

// whole game methods
func newGame(inputPath string) (game Game) {

	// open and scan file
	fd, err := os.Open(inputPath)
	HandleErr(err)
	scanner := bufio.NewScanner(fd)

	// get first line: input
	scanner.Scan()

	for _, n := range strings.Split(scanner.Text(), ",") {
		num, err := strconv.Atoi(n)
		HandleErr(err)
		game.Inputs = append(game.Inputs, num)
	}

	for {
		player := getBoard(scanner)
		if player == nil {
			return game
		}
		game.Players = append(game.Players, player)
	}
}
func (g *Game) strikeAllBoards(num int) {
	for _, board := range g.Players {
		if !board.Win {
			board.strikeBoard(num)

			if board.hasWon() {
				board.Win = true
			}
		}
	}
}

func (g *Game) playersRemaining() []*Board {
	var remaining []*Board

	for _, board := range g.Players {
		if !board.Win {
			remaining = append(remaining, board)
		}
	}
	return remaining
}

func (g *Game) printAll() {
	for _, b := range g.Players {
		b.printBoard()
		fmt.Printf("\n")
	}
}

// generic
func checkLine(line []BingoNumber) bool {
	for _, n := range line {
		if !n.Strike {
			return false
		}
	}
	return true
}

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// MAIN
func main() {

	// initialize with input file
	var path string
	if len(os.Args) < 2 {
		path = "input"

	} else {
		path = os.Args[1]
	}
	game := newGame(path)
	numBoards := len(game.Players)
	fmt.Printf("%d bingo boards loaded\n", numBoards)

	// play game
	var remaining []*Board
	for round, input := range game.Inputs {
		fmt.Printf("Round: %3d, remaining: %3d\r", round, len(remaining))

		game.strikeAllBoards(input)

		remaining = game.playersRemaining()
		if len(remaining) == 1 {
			looser := remaining[0]

			// play till looser wins
			for !looser.hasWon() {
				round++
				input = game.Inputs[round]
				looser.strikeBoard(input)
			}

			// get score of board
			score := 0
			for _, line := range looser.Numbers {
				for _, bn := range line {
					if !bn.Strike {
						score += bn.Value
					}
				}
			}

			// multiply by final num
			fmt.Printf("\nLoosers score: %d * %d = %d\n", score, input, score*input)
			break
		}
	}
}
