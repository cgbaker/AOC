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
	scanner.Split(bufio.ScanLines)

	puzzle := readPuzzle(scanner)
	//puzzle.print()

	//fmt.Println("Part1:", puzzle.fold(true,scanner))
	fmt.Println("Part2:", puzzle.fold(false,scanner))
	puzzle.print()
}

func readPuzzle(scanner *bufio.Scanner) *Puzzle {
	var puzzle = &Puzzle{
		entries: map[int]map[int]bool{},
	}
	for scanner.Scan() {
		x, y := 0, 0
		if _, err := fmt.Sscanf(scanner.Text(), "%d,%d", &x, &y); err != nil {
			break
		}
		puzzle.append(x,y)
	}
	return puzzle
}

type Puzzle struct {
	width, height int
	entries map[int]map[int]bool
}

func (p *Puzzle) append(x, y int) {
	p.width = max(p.width, x)
	p.height = max(p.height, y)
	if v, ok := p.entries[y]; ok {
		v[x] = true
	} else {
		p.entries[y] = map[int]bool{
			x: true,
		}
	}
}

func (p *Puzzle) foldLeft(line int) {
	fmt.Println("folding left at", line)
	newEntries := map[int]map[int]bool{}
	for y, oldRow := range p.entries {
		newRow := map[int]bool{}
		for x, _ := range oldRow {
			if x < line {
				newRow[x] = true
			} else if x > line {
				// line - (x - line) == 2x - x
				newRow[2*line-x] = true
			}
		}
		newEntries[y] = newRow
	}
	p.width = line - 1
	p.entries = newEntries
}

func (p *Puzzle) foldUp(line int) {
	fmt.Println("folding up at",line)
	for y := line+1; y <= p.height; y++ {
		if from, ok := p.entries[y]; ok {
			to := p.entries[2*line-y]
			if to == nil {
				to = map[int]bool{}
			}
			for x, _ := range from {
				to[x] = true
			}
			p.entries[2*line-y] = to
			delete(p.entries, y)
		}
	}
	p.height = line-1
}

func (p *Puzzle) count() int {
	count := 0
	for _, v := range p.entries {
		count += len(v)
	}
	return count
}

func (p *Puzzle) fold(firstFoldOnly bool, scanner *bufio.Scanner) int {
	for scanner.Scan() {
		axis := ""
		line := -1
		fmt.Println("fold instruction: ",scanner.Text())
		fmt.Sscanf(scanner.Text(), "fold along %1s=%d",&axis,&line)
		switch axis {
		case "x":
			p.foldLeft(line)
		case "y":
			p.foldUp(line)
		default:
			panic("bad fold " + axis)
		}
		//p.print()
		if firstFoldOnly {
			break
		}
	}
	return p.count()
}

func (p *Puzzle) print() {
	for y := 0; y <= p.height; y++ {
		row, ok := p.entries[y]
		if !ok {
			row = map[int]bool{}
		}
		for x := 0; x <= p.width; x++ {
			if row[x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}