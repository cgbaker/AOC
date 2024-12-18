package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type XY struct {
	x,y byte
}

func (xy *XY) Row() int {
	return int(xy.y)
}

func (xy *XY) Col() int {
	return int(xy.x)
}

type Problem struct {
	bytes []XY
	h, w byte
}

func (p *Problem) inBounds(xy XY) bool {
	return xy.x >= 0 && xy.x < p.w && xy.y >= 0 && xy.y < p.h
}

func (p *Problem) getNeighbors(cur XY, numBytes int) []XY {
	neighbors := []XY{}
	for _, pot := range []XY{
		XY{cur.x,cur.y+1},
		XY{cur.x,cur.y-1},
		XY{cur.x-1,cur.y},
		XY{cur.x+1,cur.y}} {
			if p.inBounds(pot) && !p.corrupted(pot,numBytes) {
				neighbors = append(neighbors, pot)
			}
	}
	return neighbors
}

func (p *Problem) corrupted(xy XY, numBytes int) bool {
	for _, b := range p.bytes[:numBytes] {
		if b.x == xy.x && b.y == xy.y {
			return true
		}
	}
	return false
}

func (p *Problem) isEnd(xy XY) bool {
	return xy.x == p.w-1 && xy.y == p.h-1
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	problem := readInput(file)
	problem.h = 71
	problem.w = 71
	part1(problem, 1024)
	part2(problem)
}

type Node struct {
	pos XY
	cost int
}

func pop(stack []Node) ([]Node,Node) {
	return stack[:len(stack)-1], stack[len(stack)-1]
}

func part1(problem *Problem, numBytes int) {
	fmt.Println("part1:",solve(problem, numBytes))
}

func solve(problem *Problem, numBytes int) int {
	best := math.MaxInt
	visited := map[XY]int{}
	stack := []Node{Node{XY{0,0},0}}
	for len(stack) != 0 {
		var cur Node
		stack, cur = pop(stack)
		visited[cur.pos] = cur.cost
		for _, neighbor := range problem.getNeighbors(cur.pos, numBytes) {
			nCost := cur.cost + 1
			if nCost >= best {
				continue
			}
			if problem.isEnd(neighbor) {
				// fmt.Println("found exit with cost ",nCost)
				best = nCost
			} else if prevCost, already := visited[neighbor]; !already || nCost < prevCost {
				stack = append(stack, Node{neighbor,nCost})
			}
		}
	}
	return best
}

func hasPath(problem *Problem, numBytes int) bool {
	solution := solve(problem, numBytes)
	return math.MaxInt != solution
}

func part2(problem *Problem) {
	l, r := 1, len(problem.bytes)
	for {
		if l == r-1 {
			break
		}
		m := (l+r)/2
		if hasPath(problem,m) {
			l = m
		} else {
			r = m
		}
	}
	fmt.Println("part2:",problem.bytes[r-1])
}

func readInput(file *os.File) *Problem {
	bytes := []XY{}
	lineScanner := bufio.NewScanner(file)
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		var x, y int
		fmt.Sscanf(lineScanner.Text(), "%d,%d", &x, &y)
		bytes = append(bytes, XY{byte(x),byte(y)})
	}
	return &Problem{
		bytes: bytes,
	}
}


