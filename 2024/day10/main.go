package main

import (
	"fmt"
	"os"
	"log"
	"github.com/cgbaker/AOC/2024/utils"
)

type Problem struct {
	grid *utils.CharGrid
	trailheads []utils.Point
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	problem := readInput(file)
	prob1(problem)
	prob2(problem)
}

func prob1(problem Problem) {
	count := 0
	for _, th := range problem.trailheads {
		// fmt.Println("trailhead",th)
		visited := map[utils.Point]bool{}
		count += dfs("", problem.grid, th, visited)
	}
	fmt.Printf("prob1: %d\n",count)
}

func dfs(indent string, grid *utils.CharGrid, th utils.Point, visited map[utils.Point]bool) int {
	sum := 0
	ns := getNeighbors(grid, th)
	// fmt.Println(indent,"neighbors",ns)
	for _, n := range ns {
		if !visited[n] {
			// fmt.Println(indent,"visiting",n)
			visited[n] = true
			if grid.GetChar(&n)-'0' == 9 {
				// fmt.Println(indent,"found",n)
				sum++
			} else {
				sum += dfs(indent + "  ",grid, n, visited)
			}
		}
	}
	return sum
}

func prob2(_ Problem) {
	fmt.Printf("prob2: %d\n",0)
}

func readInput(file *os.File) Problem {
	grid := utils.ReadCharGrid(file) 
	trailheads := []utils.Point{}
	for index, value := range grid.Chars {
		if value == '0' {
			trailheads = append(trailheads, utils.NewPoint(grid.RowCol(index)))
		}
	}
	return Problem{
		grid: grid,
		trailheads: trailheads,
	}
}

func getNeighbors(grid *utils.CharGrid, x utils.Point) []utils.Point {
	xval := grid.GetChar(&x)
	ns := []utils.Point{}
	for _, d := range []utils.Point{ utils.NewPoint(-1,0), utils.NewPoint(1,0), utils.NewPoint(0,-1), utils.NewPoint(0,1) } {
		potN := x.Plus(d)
		if b := grid.GetChar(&potN); b != 0 {
			if b == xval + 1 {
				ns = append(ns, potN)
			}
		}
	}
	return ns
}

