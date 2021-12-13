package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	START = "start"
	END = "end"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	puzzle := readPuzzle(scanner)
	//fmt.Println(puzzle)

	//fmt.Println("Part1:", part1(puzzle))
	fmt.Println("Part2:", part2(puzzle))
}

func part1(puzzle *Puzzle) int {
	seen := map[Path]bool{}
	numPaths := 0
	stack := []Path{START}
	for len(stack) != 0 {
		p := stack[0]
		stack = stack[1:]
		seen[p] = true
		for n := range puzzle.graph[p.end()] {
			nextPath := p.append(n)
			if !seen[nextPath] && !p.containsSmall(n) {
				stack = append(stack, nextPath)
				if n == END {
					numPaths++
				}
			}
		}
	}
	return numPaths
}

func part2(puzzle *Puzzle) int {
	seen := map[Path]bool{}
	numPaths := 0
	stack := []Path{START}
	seen[START] = true
	for len(stack) != 0 {
		p := stack[0]
		stack = stack[1:]
		for n := range puzzle.graph[p.end()] {
			if n == START { // don't go back to START
				continue
			}
			nextPath := p.append(n)
			if !seen[nextPath] && !nextPath.containsMoreThanOneSmallVisit() {
				seen[nextPath] = true
				if n == END {
					numPaths++
				} else {
					stack = append(stack, nextPath)
				}
			}
		}
	}
	return numPaths
}

type Path string

func (p Path) end() Cave {
	caves := strings.Split(string(p), ",")
	return Cave(caves[len(caves)-1])
}

type Cave string

func (p Path) append(c Cave) Path {
	return Path(string(p) + "," + string(c))
}

func (c Cave) isLarge() bool {
	if string(c) == strings.ToUpper(string(c)) {
		return true
	}
	return false
}

func readPuzzle(scanner *bufio.Scanner) *Puzzle {
	var puzzle = &Puzzle{
		graph: map[Cave]map[Cave]struct{}{},
	}
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		ret := strings.Split(scanner.Text(), "-")
		puzzle.append(Cave(ret[0]), Cave(ret[1]))
	}
	return puzzle
}

type Puzzle struct {
	graph map[Cave]map[Cave]struct{}
}

func (p *Puzzle) append(a, b Cave) {
	if _, ok := p.graph[a]; ok {
		p.graph[a][b] = struct{}{}
	} else {
		p.graph[a] = map[Cave]struct{}{
			b: {},
		}
	}
	if _, ok := p.graph[b]; ok {
		p.graph[b][a] = struct{}{}
	} else {
		p.graph[b] = map[Cave]struct{}{
			a: {},
		}
	}
}

func (p Path) containsSmall(s Cave) bool {
	caves := strings.Split(string(p), ",")
	for _, c := range caves {
		if c == string(s) && !Cave(c).isLarge() {
			return true
		}
	}
	return false
}

func (p Path) containsMoreThanOneSmallVisit() bool {
	caves := strings.Split(string(p), ",")
	visits := map[string]int{}
	for _, c := range caves {
		if !Cave(c).isLarge() {
			visits[c]++
		}
	}
	moreThanOne := 0
	for _, v := range visits {
		if v == 2 {
			moreThanOne++
		} else if v > 2 {
			return true
		}
		if moreThanOne > 1 {
			return true
		}
	}
	return false
}