package main

import (
	"fmt"
	"os"
	"log"
	"github.com/cgbaker/AOC/2024/utils"
)

type Input = utils.CharGrid

type Problem struct {
	antennas map[byte][]utils.Point
	r, c int
}

var (
	EMPTY byte = '.'
	ANTINODE byte = '#'
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	input := utils.ReadCharGrid(file)
	problem := processGrid(input)
	prob1(problem)
	prob2(problem)
}

func prob1(problem Problem) {
	fmt.Printf("prob1: %d\n",solve(problem,false))
}

func prob2(problem Problem) {
	fmt.Printf("prob2: %d\n",solve(problem,true))
}

func solve(problem Problem, harmonics bool) int {
	// use a grid for these in case we need a grid later
	antinodes := utils.NewCharGrid(problem.r, problem.c)
	for i := range antinodes.Chars {
		antinodes.Chars[i] = EMPTY
	}
	for _, antLocs := range problem.antennas {
		for i, a := range antLocs {
			for _, b := range antLocs[i+1:] {
				for _, xn := range makeAntiNodes(harmonics, problem.r, problem.c, a,b) {
					antinodes.SetChar(&xn, ANTINODE)
				}
			}
		}
	}
	// fmt.Println("antinodes")
	// antinodes.Print()
	numAntiNodes := 0
	for _, v := range antinodes.Chars {
		if v == ANTINODE {
			numAntiNodes++
		}
	}
	return numAntiNodes
}

func makeAntiNodes(harmonics bool, nR int, nC int, a, b utils.Point) []utils.Point {
	outOfBounds := func(p utils.Point) bool {
		return p.Row() < 0 || p.Col() < 0 || p.Row() >= nR || p.Col() >= nC
	}
	antiNodes := []utils.Point{}
	// n1 == (b-a) + b == 2b-a             n1-b == 2b-a-b == b-a   n1-a == 2b-a-a == 2b-2a
	// n2 == a-(b-a) == a - (b-a) == 2a-b  n2-b == 2a-b-b == 2a-2b n2-a == 2a-b-a == a-b
	d := b.Minus(a)
	// make "forward" antinodes
	xn := b
	if harmonics {
		antiNodes = append(antiNodes, xn)
	}
	for {
		xn = xn.Plus(d)
		if outOfBounds(xn) {
			break
		}
		antiNodes = append(antiNodes, xn)
		if !harmonics {
			break
		}
	}
	// make "backward" antinodes
	xn = a
	if harmonics {
		antiNodes = append(antiNodes, xn)
	}
	for {
		xn = xn.Minus(d)
		if outOfBounds(xn) {
			break
		}
		antiNodes = append(antiNodes, xn)
		if !harmonics {
			break
		}
	}
	return antiNodes
}

func processGrid(input *Input) Problem {
	antennas := map[byte][]utils.Point{}	
	for i, v := range input.Chars {
		if v != EMPTY {
			antennas[v] = append(antennas[v], utils.NewPoint(input.RowCol(i)))
		}
	}
	return Problem{
		antennas: antennas,
		r: input.NumRows,
		c: input.NumCols,
	}
}
