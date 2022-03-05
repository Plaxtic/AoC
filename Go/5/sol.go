package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Grid struct {
	MaxX, MaxY int
	Points     [probMaxSiz][probMaxSiz]int
}

type Vector struct {
	X1, Y1 int
	X2, Y2 int
}

const (
	probMaxSiz = 1000 // sacrifice elegance for speed
)

// PARSE FILE
func loadVectors(path string) (vectors []Vector) {

	// open and scan
	fd, err := os.Open(path)
	HandleErr(err)
	scanner := bufio.NewScanner(fd)

	// parse vector strings
	for scanner.Scan() {

		// split by direction
		coordinates := strings.Split(scanner.Text(), " -> ")
		if len(coordinates) < 2 {
			log.Panic("Bad Line")
		}

		// split into x/y
		XY1 := strings.Split(coordinates[0], ",")
		XY2 := strings.Split(coordinates[1], ",")
		if len(XY1) < 2 || len(XY2) < 2 {
			log.Panic("Bad Line")
		}

		// convert  all to integer
		X1, err := strconv.Atoi(XY1[0])
		HandleErr(err)
		Y1, err := strconv.Atoi(XY1[1])
		HandleErr(err)
		X2, err := strconv.Atoi(XY2[0])
		HandleErr(err)
		Y2, err := strconv.Atoi(XY2[1])
		HandleErr(err)

		// create vector and append to vector array
		vector := Vector{
			X1: X1,
			Y1: Y1,
			X2: X2,
			Y2: Y2,
		}
		vectors = append(vectors, vector)
	}
	return vectors

}

// GRID FUNCTIONS
func newGrid(vecs []Vector, part int) *Grid {
	grid := &Grid{}

	for _, vec := range vecs {
		if vec.X1 == vec.X2 { // vertical vector
			grid.addVertVec(vec)
		} else if vec.Y1 == vec.Y2 { // horizontal vector
			grid.addHozVec(vec)
		} else if part == 1 { // stop if running part one
			continue
		} else if vec.X1-vec.X2 == vec.Y1-vec.Y2 { // left leaning diagonal
			grid.addDiagonalVecR(vec)
		} else if vec.X2-vec.X1 == vec.Y1-vec.Y2 { // right leaning diagonal
			grid.addDiagonalVecL(vec)
		}
	}
	return grid
}

func (g *Grid) addHozVec(vec Vector) {

	// get and update max/min values
	hX, lX, hY, _ := vec.getMaxMinXY()
	g.updateMaxes(hX, hY)

	// add points
	for x := lX; x <= hX; x++ {
		g.Points[hY][x]++
	}
}

func (g *Grid) addVertVec(vec Vector) {

	// get and update max/min values
	hX, _, hY, lY := vec.getMaxMinXY()
	g.updateMaxes(hX, hY)

	// add points
	for y := lY; y <= hY; y++ {
		g.Points[y][hX]++
	}
}

func (g *Grid) addDiagonalVecL(vec Vector) {

	// get and update max/min values
	hX, lX, hY, lY := vec.getMaxMinXY()
	g.updateMaxes(hX, hY)

	// add points
	x := lX
	for y := lY; y <= hY; y++ {
		g.Points[y][x]++
		x++
	}
}

func (g *Grid) addDiagonalVecR(vec Vector) {

	// get and update max/min values
	hX, _, hY, lY := vec.getMaxMinXY()
	g.updateMaxes(hX, hY)

	// add points
	x := hX
	for y := lY; y <= hY; y++ {
		g.Points[y][x]++
		x--
	}
}

func (g *Grid) updateMaxes(x, y int) {

	// check in bounds
	if x > probMaxSiz || y > probMaxSiz {
		log.Panic("Too big (increase buffer)")
	}

	// check if new maxes
	if y > g.MaxY {
		g.MaxY = y
	}
	if x > g.MaxX {
		g.MaxX = x
	}
}

// cacluate intersecting
func (g *Grid) numIntersections() (n int) {
	for y := 0; y <= g.MaxY; y++ {
		for x := 0; x <= g.MaxX; x++ {
			intsec := g.Points[y][x]
			if intsec > 1 {
				n++
			}
		}
	}
	return n
}

func (v *Vector) getMaxMinXY() (maxX, minX, maxY, minY int) {

	// get max X coordinate
	minX = v.X1
	maxX = v.X2
	if v.X2 < minX {
		minX = v.X2
		maxX = v.X1
	}

	// get max Y coordinate
	minY = v.Y1
	maxY = v.Y2
	if v.Y2 < minY {
		minY = v.Y2
		maxY = v.Y1
	}

	return maxX, minX, maxY, minY
}

// DEBUG FUNCTIONS
func (g *Grid) printGrid() {
	for y := 0; y <= g.MaxY; y++ {
		fmt.Printf("%d| ", y)
		for x := 0; x <= g.MaxX; x++ {

			// check intersections
			intsec := g.Points[y][x]
			if intsec > 0 {
				fmt.Printf("%d", intsec)
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println("")
	}
}

func (v *Vector) printVec() {
	fmt.Printf("(%d, %d) -> (%d, %d)\n", v.X1, v.Y1, v.X2, v.Y2)
}

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// driver code
func main() {

	// allow file selection
	path := "input"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	// parse file into vectors
	vecs := loadVectors(path)

	// map vectors onto grid
	grid1 := newGrid(vecs, 1)
	grid2 := newGrid(vecs, 2)

	// results
	fmt.Printf("Mx: %d, My: %d\n", grid2.MaxX, grid2.MaxY)
	fmt.Printf("Lines overlapping (part 1) : %d\n", grid1.numIntersections())
	fmt.Printf("Lines overlapping (part 2) : %d\n", grid2.numIntersections())
}
