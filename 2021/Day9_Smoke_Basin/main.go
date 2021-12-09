package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	puzzle := readPuzzle(scanner)
	fmt.Println("Part1:", part1(puzzle))
	fmt.Println("Part2:", part2(puzzle))
}

func part1(puzzle *Puzzle) int {
	localMins := puzzle.getLocalMins()
	totalRisk := 0
	for _, l := range localMins {
		totalRisk += int(puzzle.depths[l]+1)
	}
	return totalRisk
}

func part2(puzzle *Puzzle) int {
	// find the low points
	locals := puzzle.getLocalMins()
	basinSizes := []int{}
	// BFS of each basin
	for _, d := range locals {
		basinSizes = append(basinSizes, puzzle.exploreBasin(d))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	return basinSizes[0]*basinSizes[1]*basinSizes[2]
}

func (puzzle *Puzzle) exploreBasin(min int) int {
	visited := map[int]bool{min: true}
	basin := []int{min}
	b := 0
	for len(basin) != 0 {
		b, basin = basin[0], basin[1:]
		neighbors := puzzle.getNeighbors(b)
		for _, n := range neighbors {
			if puzzle.depths[b] <= puzzle.depths[n] && puzzle.depths[n] < 9 && !visited[n] {
				visited[n] = true
				basin = append(basin, n)
			}
		}
	}
	return len(visited)
}

func readPuzzle(scanner *bufio.Scanner) *Puzzle {
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	first := scanner.Text()
	puzzle := &Puzzle{
		width: len(first),
		depths: []int8{},
	}
	puzzle.append(parseLine(first))
	for scanner.Scan() {
		puzzle.append(parseLine(scanner.Text()))
	}
	return puzzle
}

type Puzzle struct {
	width, height int
	depths []int8
}

func (p *Puzzle) append(vals []int8) {
	p.depths = append(p.depths, vals...)
	p.height++
}

func parseLine(line string) []int8 {
	ints := make([]int8, len(line))
	for i, s := range line {
		ints[i] = int8(s - '0')
	}
	return ints
}

func (p *Puzzle) getLocalMins() []int {
	locals := []int{}
	for i, d := range p.depths {
		neighbors := p.getNeighbors(i)
		isLocal := true
		for _, n := range neighbors {
			if d >= p.depths[n] {
				isLocal = false
				break
			}
		}
		if isLocal {
			locals = append(locals, i)
		}
	}
	return locals
}

func (p *Puzzle) getNeighbors(i int) []int {
	n := []int{}
	if i >= p.width {
		n = append(n, i-p.width)
	}
	if i < p.width*p.height - p.width {
		n = append(n, i+p.width)
	}
	if i % p.width != 0 {
		n = append(n, i-1)
	}
	if i % p.width != p.width-1 {
		n = append(n, i+1)
	}
	return n
}