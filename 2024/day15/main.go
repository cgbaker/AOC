package main

import (
	"fmt"
	"os"
	"log"
	"github.com/cgbaker/AOC/2024/utils"
)

var (
	WALL byte = '#'
	BOX byte = 'O'
	ROBOT byte = '@'
	EMPTY byte = '.'
	BOXL byte = '['
	BOXR byte = ']'
)

type Problem struct {
	grid *utils.CharGrid
	robot utils.Point
	moves []byte
}

func main() {
	// file, err := os.Open("sample_small.txt")
	// file, err := os.Open("sample_small_2.txt")
	// file, err := os.Open("sample_large.txt")
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	problem := readProblem(file)
	problem2 := convertProblem(problem)
	part1(problem)
	part2(problem2)
}

func part1(problem *Problem) {
	// fmt.Println("Initial state:")
	// problem.grid.Print()
	for _, move := range problem.moves {
		//fmt.Println()
		//fmt.Printf("Move %c:\n",rune(move))
		problem.robot = attemptMoveRobot(problem.grid, problem.robot, move)
		//problem.grid.Print()
	}
	sum := 0
	for i, b := range problem.grid.Chars {
		if b == BOX {
			r, c := problem.grid.RowCol(i)
			sum += r*100 + c
		}
	}
	fmt.Printf("prob1: %d\n",sum)
}

func part2(problem *Problem) {
	// fmt.Println("Initial state:")
	// problem.grid.Print()
	for _, move := range problem.moves {
		// fmt.Println()
		// fmt.Printf("Move %c:\n",rune(move))
		problem.robot = attemptMoveRobot(problem.grid, problem.robot, move)
		// problem.grid.Print()
	}
	sum := 0
	for i, b := range problem.grid.Chars {
		if b == BOXL {
			r, c := problem.grid.RowCol(i)
			sum += r*100 + c
		}
	}
	fmt.Printf("prob2: %d\n",sum)
}

func dirToPoint(dir byte) utils.Point {
	switch dir {
	case '>':
		return utils.NewPoint(0,1)
	case '<':
		return utils.NewPoint(0,-1)
	case 'v':
		return utils.NewPoint(1,0)
	case '^':
		return utils.NewPoint(-1,0)
	}
	panic("oops")
}

func attemptMoveRobot(grid *utils.CharGrid, robot utils.Point, move byte) utils.Point {
	// get position in step move... if it can be moved (recurse), move this one and return true
	// else, return false
	dir := dirToPoint(move)
	moveTo := robot.Plus(dir)
	if canBePushed(grid, &moveTo, dir) {
		push(grid, &moveTo, dir)
		grid.SetChar(&moveTo,ROBOT)
		grid.SetChar(&robot,EMPTY)
		return moveTo
	}
	return robot
}

func canBePushed(grid *utils.CharGrid, origin utils.Coord, dir utils.Point) bool {
	c := grid.GetChar(origin)
	switch c {
	case EMPTY:
		return true
	case BOX:
		check := utils.NewPointFromCoord(origin).Plus(dir)
		return canBePushed(grid, &check, dir)
	case BOXL:
		if isEastWest(dir) {
			check := utils.NewPointFromCoord(origin).Plus(dir)
			return canBePushed(grid, &check, dir)
		} else {
			checkL := utils.NewPoint(origin.Row(), origin.Col()  ).Plus(dir)
			checkR := utils.NewPoint(origin.Row(), origin.Col()+1).Plus(dir)
			return canBePushed(grid, &checkL, dir) && canBePushed(grid, &checkR, dir)
		}
	case BOXR:
		if isEastWest(dir) {
			check := utils.NewPointFromCoord(origin).Plus(dir)
			return canBePushed(grid, &check, dir)
		} else {
			checkL := utils.NewPoint(origin.Row(), origin.Col()-1).Plus(dir)
			checkR := utils.NewPoint(origin.Row(), origin.Col()  ).Plus(dir)
			return canBePushed(grid, &checkL, dir) && canBePushed(grid, &checkR, dir)
		}
	}
	return false
}

func isEastWest(dir utils.Point) bool {
	if dir.Row() == 0 {
		return true
	}
	return false
}

// it's safe to move boxes; now move them.
// push the contents of dest one square in the direction dir
func push(grid *utils.CharGrid, dest utils.Coord, dir utils.Point) {
	c := grid.GetChar(dest)
	switch {
	case c == EMPTY || c == WALL:
		return
	case c == BOXL && !isEastWest(dir):
		pair := utils.NewPoint(dest.Row(), dest.Col()+1)
		nextL := utils.NewPoint(dest.Row(), dest.Col()).Plus(dir)
		nextR := utils.NewPoint(pair.Row(), pair.Col()).Plus(dir)
		push(grid, &nextL, dir)
		push(grid, &nextR, dir)
		grid.SetChar(&nextL,BOXL)
		grid.SetChar(&nextR,BOXR)
		grid.SetChar(dest,EMPTY)
		grid.SetChar(&pair,EMPTY)
	case c == BOXR && !isEastWest(dir):
		pair := utils.NewPoint(dest.Row(), dest.Col()-1)
		nextL := utils.NewPoint(pair.Row(), pair.Col()).Plus(dir)
		nextR := utils.NewPoint(dest.Row(), dest.Col()).Plus(dir)
		push(grid, &nextL, dir)
		push(grid, &nextR, dir)
		grid.SetChar(&nextL,BOXL)
		grid.SetChar(&nextR,BOXR)
		grid.SetChar(dest,EMPTY)
		grid.SetChar(&pair,EMPTY)
	default:
		next := utils.NewPoint(dest.Row(), dest.Col()).Plus(dir)
		push(grid, &next, dir)
		grid.SetChar(&next,c)
		grid.SetChar(dest,EMPTY)
	}
}

func readProblem(file *os.File) *Problem {
	grid, lineScanner := utils.ReadCharGrid(file)
	moves := []byte{}
	var robot utils.Point
	for lineScanner.Scan() {
		moves = append(moves, []byte(lineScanner.Text())...)
	}
	for i, c := range grid.Chars {
		if c == ROBOT {
			robot = utils.NewPoint(grid.RowCol(i))
			break
		}
	}
	return &Problem{
		grid: grid,
		moves: moves,
		robot: robot,
	}
}

func convertProblem(problem *Problem) *Problem {
	wideGrid := utils.NewCharGrid(problem.grid.NumRows, 2*problem.grid.NumCols)
	for i, c := range problem.grid.Chars {
		switch c {
		case EMPTY, WALL:
			wideGrid.SetCharFromIndex(2*i,c)
			wideGrid.SetCharFromIndex(2*i+1,c)
		case ROBOT:
			wideGrid.SetCharFromIndex(2*i,c)
			wideGrid.SetCharFromIndex(2*i+1,EMPTY)
		case BOX:
			wideGrid.SetCharFromIndex(2*i,BOXL)
			wideGrid.SetCharFromIndex(2*i+1,BOXR)
		}
	}
	return &Problem{
		grid: wideGrid,
		robot: utils.NewPoint(problem.robot.Row(), problem.robot.Col()*2),
		moves: problem.moves,
	}
}
