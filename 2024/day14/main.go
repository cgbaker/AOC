package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"github.com/cgbaker/AOC/2024/utils"
)

var (
	// FILE string = "sample.txt"
	// ROWS int = 7
	// COLS int = 11

	FILE string = "input.txt"
	ROWS int = 103
	COLS int = 101
)

type Robot struct {
	x, y int
	dx, dy int
}

func (r *Robot) Row() int {
	return r.y
}

func (r *Robot) Col() int {
	return r.x
}

type Problem struct {
	robots []*Robot
}

func main() {
	file, err := os.Open(FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	problem := readInput(file)
	fmt.Println("Num robots:",len(problem.robots))
	// part1(problem)
	part2(problem)
}

func part1(problem *Problem) {
	fmt.Printf("prob1: %d\n",solve(problem, 100))
}

func part2(problem *Problem) {
	fmt.Println("*******\nproblem 2")
	numSteps := 0
	hStep := 160
	vStep := 187
	for {
		var delta int
		if hStep < vStep {
			delta = hStep - numSteps
			hStep += ROWS
		} else {
			delta = vStep - numSteps
			vStep += COLS
		}
		numSteps += delta
		fmt.Println(numSteps)
		solve(problem, delta)
		var s string
		fmt.Scanln(&s)
	}
}

func solve(problem *Problem, seconds int) int {
	var NW, NE, SW, SE int
	g := utils.NewCharGrid(ROWS,COLS)
	for _, r := range problem.robots {
		// move
		r.x = (r.x + seconds*r.dx) % COLS
		if r.x < 0 {
			r.x += COLS
		}
		r.y = (r.y + seconds*r.dy) % ROWS
		if r.y < 0 {
			r.y += ROWS
		}
		g.SetChar(r, g.GetChar(r)+1)
		// determine quadrant
		if r.x < COLS/2 && r.y < ROWS/2 {
			NW++
		} else if r.x > COLS/2 && r.y < ROWS/2 {
			NE++
		} else if r.x < COLS/2 && r.y > ROWS/2 {
			SW++
		} else if r.x > COLS/2 && r.y > ROWS/2 {
			SE++
		}
	}
	for i := range g.Chars {
		if g.Chars[i] == 0 {
			g.Chars[i] = ' '
		} else {
			g.Chars[i] += '0'
		}
	}
	g.Print()
	return NW*NE*SW*SE
}

func readInput(file *os.File) *Problem {
	robots := []*Robot{}
	lineScanner := bufio.NewScanner(file)
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		robot := &Robot{}
		fmt.Sscanf(lineScanner.Text(), "p=%d,%d v=%d,%d", &robot.x, &robot.y, &robot.dx, &robot.dy)
		robots = append(robots,robot)
	}
	return &Problem{
		robots: robots,
	}
}


