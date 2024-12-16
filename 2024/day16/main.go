package main

import (
	"fmt"
	"os"
	"log"
	"math"
	"github.com/cgbaker/AOC/2024/utils"
	"strings"
	"strconv"
)

var (
	EMPTY byte = '.'
)

type Reindeer struct {
	pos utils.Point
	dir utils.Point
}

type Maze = utils.CharGrid

type Problem struct {
	maze *Maze
	start utils.Point
	end   utils.Point
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	maze, _ := utils.ReadCharGrid(file)
	problem := &Problem{
		maze: maze,
		start: utils.NewPoint(maze.NumRows-2,1),
		end: utils.NewPoint(1,maze.NumCols-2),
	}
	bestCost, numTiles := solve(problem)
	fmt.Printf("prob1: %d\n",bestCost)
	fmt.Printf("prob2: %d\n",numTiles)
}

type Node struct {
	r Reindeer
	path string
}

type StackEntry struct {
	node Node
	cost int
}

func str(i int) string {
	return strconv.FormatInt(int64(i),10)
}

func solve(problem *Problem) (int,int) {
	best := math.MaxInt
	bestTiles := map[string]bool{}
	visited := map[Reindeer]int{}
	start := Reindeer{pos: problem.start, dir: utils.DIR_E}
	stack := []StackEntry{ StackEntry{Node{r: start, path: str(problem.maze.Index(start.pos.Row(),start.pos.Col()))}, 0} }
	for len(stack) != 0 {
		curEntry := stack[len(stack)-1]
		curNode, curCost := curEntry.node, curEntry.cost
		stack = stack[0:len(stack)-1]
		// fmt.Printf("visiting %+v (cost: %d)\n",curNode.r,curCost)
		visited[curNode.r] = curCost
		neighbors := getNeighbors(problem.maze, curNode, curCost)
		// fmt.Println("neighbors:",neighbors)
		for _, neighbor := range neighbors {
			node := neighbor.node
			cost := neighbor.cost
			if cost > best {
				continue
			}
			if node.r.pos == problem.end {
				fmt.Printf("found exit with cost %d\n",cost)
				if cost <= best {
					if cost < best {
						// clear the total
						bestTiles = map[string]bool{}
					}
					for _, s := range strings.Split(node.path,",") {
						bestTiles[s] = true
					}
					best = cost
				} 
			} else if prevCost, already := visited[node.r]; !already || cost <= prevCost {
				stack = append(stack, StackEntry{node,cost})
			}
		}
	}
	return best, len(bestTiles)
}

func getNeighbors(maze *Maze, cur Node, cost int) []StackEntry {
	ns := []StackEntry{}
	if cur.r.dir.Col() == 0 {
		ns = append(ns, 
		StackEntry{Node{r: Reindeer{pos: cur.r.pos, dir: utils.DIR_E}, path: cur.path, }, cost+1000},
		StackEntry{Node{r: Reindeer{pos: cur.r.pos, dir: utils.DIR_W}, path: cur.path, }, cost+1000})
	} else {
		ns = append(ns, 
		StackEntry{Node{ r: Reindeer{pos: cur.r.pos, dir: utils.DIR_N}, path: cur.path}, cost+1000},
		StackEntry{Node{ r: Reindeer{pos: cur.r.pos, dir: utils.DIR_S}, path: cur.path}, cost+1000})
	}
	newPos := cur.r.pos.Plus(cur.r.dir)
	if next := maze.GetChar(&newPos); next == EMPTY || next == 'E' {
		index := str(maze.Index(newPos.Row(),newPos.Col()))
		ns = append(ns, StackEntry{Node{r: Reindeer{pos: newPos, dir: cur.r.dir}, path: cur.path+","+index},cost+1})
	}
	return ns
}

func clone(a map[utils.Point]bool, p utils.Point) map[utils.Point]bool {
	b := map[utils.Point]bool{}
	for k, _ := range a {
		b[k] = true
	}
	b[p] = true
	return b
}
