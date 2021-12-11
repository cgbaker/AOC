package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	puzzle := readPuzzle(scanner)
	//fmt.Println("Part1:", part1(100,puzzle))
	fmt.Println("Part2:", part2(puzzle))
}

func part1(numSteps int, puzzle *Puzzle) int {
	numFlashes := 0
	for s := 0; s < numSteps; s++ {
		flashers := puzzle.incrementAll()
		for i := 0; i < len(flashers); i++ {
			flashers = append(flashers, puzzle.incNeighbors(flashers[i])...)
		}
		numFlashes += puzzle.resetFlashers()
	}
	return numFlashes
}

func part2(puzzle *Puzzle) int {
	step := 0
	for {
		step++
		flashers := puzzle.incrementAll()
		for i := 0; i < len(flashers); i++ {
			flashers = append(flashers, puzzle.incNeighbors(flashers[i])...)
		}
		numFlashes := puzzle.resetFlashers()
		if numFlashes == 100 {
			return step
		}
	}
}


func readPuzzle(scanner *bufio.Scanner) *Puzzle {
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	first := scanner.Text()
	puzzle := &Puzzle{
		energy: []int8{},
	}
	puzzle.append(parseLine(first))
	for scanner.Scan() {
		puzzle.append(parseLine(scanner.Text()))
	}
	return puzzle
}

type Puzzle struct {
	energy []int8
}

func (p *Puzzle) append(vals []int8) {
	p.energy = append(p.energy, vals...)
}

func parseLine(line string) []int8 {
	ints := make([]int8, len(line))
	for i, s := range line {
		ints[i] = int8(s - '0')
	}
	return ints
}

func (p *Puzzle) incNeighbors(i int8) []int8 {
	nei := make([]int8,0,9)
	// north, nw, ne
	if i >= 10 {
		nei = append(nei, i-10)
		if i % 10 != 0 {
			nei = append(nei, i-11)
		}
		if i % 10 != 9 {
			nei = append(nei, i-9)
		}
	}
	// south, sw, se
	if i < 100 - 10 {
		nei = append(nei, i+10)
		if i % 10 != 0 {
			nei = append(nei, i+9)
		}
		if i % 10 != 9 {
			nei = append(nei, i+11)
		}
	}
	// east
	if i % 10 != 0 {
		nei = append(nei, i-1)
	}
	// west
	if i % 10 != 9 {
		nei = append(nei, i+1)
	}
	flashed := make([]int8, 0, 9)
	for _, n := range nei {
		p.energy[n]++
		if p.energy[n] == 10 {
			flashed = append(flashed, n)
		}
	}
	return flashed
}

func (puzzle *Puzzle) incrementAll() []int8 {
	flashes := make([]int8,0,100)
	for i:=0; i<100; i++ {
		puzzle.energy[i]++
		if puzzle.energy[i] > 9 {
			flashes = append(flashes, int8(i))
		}
	}
	return flashes
}

func (puzzle *Puzzle) resetFlashers() int {
	count := 0
	for i:=0; i<100; i++ {
		if puzzle.energy[i] > 9 {
			puzzle.energy[i] = 0
			count++
		}
	}
	return count
}

func (p *Puzzle) print() {
	for i, e := range p.energy {
		fmt.Print(e)
		if i % 10 == 9 {
			fmt.Println("")
		}
	}
}