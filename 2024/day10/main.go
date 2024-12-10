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
		count += dfsScore("", problem.grid, th, visited)
	}
	fmt.Printf("prob1: %d\n",count)
}

func dfsScore(indent string, grid *utils.CharGrid, th utils.Point, visited map[utils.Point]bool) int {
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
				sum += dfsScore(indent + "  ",grid, n, visited)
			}
		}
	}
	return sum
}

func prob2(problem Problem) {
	count := 0
	for _, th := range problem.trailheads {
		// fmt.Println("trailhead",th)
		visited := map[string]bool{}
		count += dfsRating(problem.grid, th, "", visited)
	}
	fmt.Printf("prob2: %d\n",count)
}

func dfsRating(grid *utils.CharGrid, th utils.Point, pathPrefix string, visited map[string]bool) int {
	sum := 0
	ns := getNeighbors(grid, th)
	// fmt.Println("neighbors",ns)
	for _, n := range ns {
		path := pathPrefix + fmt.Sprintf(";%d,%d",n.Row(),n.Col())
		if !visited[path] {
			// fmt.Println("visiting",path)
			visited[path] = true
			if grid.GetChar(&n)-'0' == 9 {
				// fmt.Println("found",n)
				sum++
			} else {
				sum += dfsRating(grid, n, path, visited)
			}
		}
	}
	return sum
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

