package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type Digits struct {
	Zero, One, Two, Three, Four,
	Five, Six, Seven, Eight, Nine string
	A,
	B, C,
	D,
	E, F,
	G rune
	Segments map[rune]rune
}

const (
	SegMarkers = "abcdefg"

	_ int = iota
	One
	Seven
	Four
	TwoThreeFive
	ZeroSixNine
	Eight
)

func (d *Digits) inferFromSix() error {
	if len(d.One) == 0 ||
		len(d.Six) == 0 ||
		len(d.Seven) == 0 {
		return errors.New("Require One, Six, and Seven")
	}

	// get C and F (because six does not contain C)
	if strings.Contains(d.Six, string(d.One[0])) {
		d.C = rune(d.One[1])
		d.F = rune(d.One[0])
	} else {
		d.C = rune(d.One[0])
		d.F = rune(d.One[1])
	}

	// get A (because its the only extra in seven)
	for _, c := range d.Seven {
		if !strings.Contains(d.One, string(c)) {
			d.A = c
			break
		}
	}
	d.Segments['a'] = d.A
	d.Segments['c'] = d.C
	d.Segments['f'] = d.F

	return nil
}

func (d *Digits) inferFromFive() error {
	if len(d.Five) != TwoThreeFive || d.C == 0 {
		fmt.Println(d.C)
		return errors.New("Require Five and C")
	}

	// get E (because only E and C are not in seven, and we know C by now)
	for seg := range SegMarkers {
		if !strings.Contains(d.Five, fmt.Sprint(seg)) &&
			!strings.Contains(d.Five, string(d.C)) {
			d.E = rune(seg)
			break
		}
	}
	d.Segments['e'] = d.E

	return nil
}

func (d *Digits) inferFromTwo() error {
	if len(d.Two) != TwoThreeFive || d.F == 0 {
		return errors.New("Require Two and F")
	}

	// get B (because only B and F are not in seven, and we know F by now)
	for seg := range SegMarkers {
		if !strings.Contains(d.Two, fmt.Sprint(seg)) &&
			!strings.Contains(d.Two, string(d.F)) {
			d.B = rune(seg)
			break
		}
	}
	d.Segments['b'] = d.B

	return nil
}

func (d *Digits) inferFromFour() error {
	if d.B == 0 || d.C == 0 || d.F == 0 || len(d.Four) != Four {
		return errors.New("Require B, C, F, and four")
	}

	// get B (because we know B, C, and F, so the remaining is D)
	for _, c := range d.Four {
		if c != d.B && c != d.C && c != d.F {
			d.D = c
			break
		}
	}
	d.Segments['d'] = d.D

	return nil
}

func (d *Digits) inferSegemtents(signalPattern []string) {

	// get single digits 1, 4, 7, 8
	for i, pat := range signalPattern {
		switch len(pat) {
		case One:
			d.One = pat
			remove(signalPattern, i)
		case Seven:
			d.Seven = pat
			remove(signalPattern, i)
		case Four:
			d.Four = pat
			remove(signalPattern, i)
		case Eight:
			d.Eight = pat
			remove(signalPattern, i)
		}
	}

	// get three and six
	for i, pat := range signalPattern {
		if len(pat) == TwoThreeFive {
			if strings.Contains(pat, string(d.One[0])) &&
				strings.Contains(pat, string(d.One[1])) {
				d.Three = pat

				remove(signalPattern, i)
			} else if strings.Contains(pat, string(d.One[0])) ||
				strings.Contains(pat, string(d.One[1])) {

				d.Six = pat
				HandleErr(d.inferFromSix()) // we now have C, F, and A

				remove(signalPattern, i)

				if len(d.Three) == TwoThreeFive {
					break
				}
			}
		}
	}

	// get five and two
	for i, pat := range signalPattern {
		if len(pat) == TwoThreeFive {
			if !strings.Contains(pat, string(d.C)) {
				d.Five = pat
				remove(signalPattern, i)

				if len(d.Two) == TwoThreeFive {
					break
				}
			} else {
				d.Two = pat
				remove(signalPattern, i)

				if len(d.Five) == TwoThreeFive {
					break
				}
			}
		}
	}
	HandleErr(d.inferFromFive()) // E
	HandleErr(d.inferFromTwo())  // B
	HandleErr(d.inferFromFour()) // D

	// get nine and zero
	if strings.Contains(signalPattern[0], string(d.E)) {
		d.Zero = signalPattern[0]
		d.Nine = signalPattern[1]
	} else {
		d.Zero = signalPattern[1]
		d.Nine = signalPattern[0]
	}

	// get G
	for _, c := range d.Eight {
		if !d.inSegs(c) {
			d.G = c
			d.Segments['g'] = d.G
		}
	}
}

func (d *Digits) inSegs(r rune) bool {
	for _, c := range d.Segments {
		if c == r {
			return true
		}
	}
	return false
}

func remove(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
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

	// split sections
	input := strings.Split(dataStr, " | ")
	hints, _ := strings.Split(input[0], " "), input[1]
	digits := Digits{}
	digits.inferSegemtents(hints)
}
