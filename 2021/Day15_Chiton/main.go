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

	puzzle := readPuzzle(file)
	//puzzle.print()

	//fmt.Println("Part1:", puzzle.shortestPath())
	fmt.Println("Part2:", puzzle.shortestPath2())
}

func readPuzzle(file *os.File) *Puzzle {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	var puzzle = &Puzzle{
		width: len(scanner.Text()),
		risk: []int8{},
	}
	puzzle.append(scanner.Text())
	for scanner.Scan() {
		puzzle.append(scanner.Text())
	}
	return puzzle
}

type Puzzle struct {
	width, height int
	risk []int8
}

func (p *Puzzle) print() {
	for i, r := range p.risk {
		fmt.Print(r)
		if i % p.width == p.width-1 {
			fmt.Println("")
		}
	}
}

func (p *Puzzle) append(text string) {
	newRow := make([]int8,p.width)
	for i, s := range text {
		newRow[i] = int8(s - '0')
	}
	p.risk = append(p.risk, newRow...)
	p.height++
}

type Coord struct {
	x, y int
}

func (p *Puzzle) shortestPath() int {
	START := Coord{0,0}
	END := Coord{p.width-1,p.height-1}

	minRisk := map[Coord]int{}
	minRisk[START] = 0

	next := []Coord{START}
	for len(next) != 0 {
		c := next[0]
		curMin := minRisk[c]
		next = next[1:]
		nei := p.getNeighbors(c)
		for _, n := range nei {
			proposed := curMin + p.getRisk(n)
			if cur, seen := minRisk[n]; !seen || proposed < cur {
				minRisk[n] = proposed
				next = append(next, n)
			}
		}
	}

	return minRisk[END]
}

func (p *Puzzle) getNeighbors(c Coord) []Coord {
	nei := make([]Coord, 0, 4)
	if c.x > 0 {
		nei = append(nei, Coord{c.x-1,c.y})
	}
	if c.x < p.width-1 {
		nei = append(nei, Coord{c.x+1, c.y})
	}
	if c.y > 0 {
		nei = append(nei, Coord{c.x, c.y-1})
	}
	if c.y < p.height-1 {
		nei = append(nei, Coord{c.x, c.y+1})
	}
	return nei
}

func (p *Puzzle) getNeighbors2(c Coord) []Coord {
	nei := make([]Coord, 0, 4)
	if c.x > 0 {
		nei = append(nei, Coord{c.x-1,c.y})
	}
	if c.x < 5*p.width-1 {
		nei = append(nei, Coord{c.x+1, c.y})
	}
	if c.y > 0 {
		nei = append(nei, Coord{c.x, c.y-1})
	}
	if c.y < 5*p.height-1 {
		nei = append(nei, Coord{c.x, c.y+1})
	}
	return nei
}

func (p *Puzzle) getRisk(n Coord) int {
	return int(p.risk[n.y*p.width + n.x])
}

func (p *Puzzle) getRisk2(n Coord) int {
	tx, ty := n.x / p.width, n.y / p.height
	sx, sy := n.x % p.width, n.y % p.height
	r := int(p.risk[sy*p.width + sx])
	r += tx + ty
	if r > 9 {
		r -= 9
	}
	return r
}

func (p *Puzzle) shortestPath2() int {
	START := Coord{0,0}
	END := Coord{5*p.width-1,5*p.height-1}

	minRisk := map[Coord]int{}
	minRisk[START] = 0

	next := []Coord{START}
	for len(next) != 0 {
		c := next[0]
		curMin := minRisk[c]
		next = next[1:]
		nei := p.getNeighbors2(c)
		for _, n := range nei {
			proposed := curMin + p.getRisk2(n)
			if cur, seen := minRisk[n]; !seen || proposed < cur {
				minRisk[n] = proposed
				next = append(next, n)
			}
		}
	}

	return minRisk[END]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}