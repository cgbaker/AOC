package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	X = 1000
	Y = 1000
)

type Board [X][Y]int

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var board Board
	numOverlap := 0
	for scanner.Scan() {
		l := Line{}
		fmt.Sscanf(scanner.Text(), "%d,%d -> %d,%d", &l.x1, &l.y1, &l.x2, &l.y2)
		//if !l.isDiagonal() {
			numOverlap += l.render(&board)
		//}
	}
	//board.print()
	fmt.Println(numOverlap)
}

type Line struct {
	x1, y1 int
	x2, y2 int
}

func step(a, b int) int {
	if a == b {
		return 0
	} else if a < b {
		return 1
	} else {
		return -1
	}
}

func (l *Line) render(board *Board) int {
	newOverlaps := 0
	xStep := step(l.x1, l.x2)
	yStep := step(l.y1, l.y2)
	x := l.x1
	y := l.y1
	numSteps := max2(
		xStep * (l.x2 - l.x1),
		yStep * (l.y2 - l.y1)) + 1
	for i := 0; i < numSteps; i++ {
		switch board[x][y] {
		case 0:
			board[x][y] = 1
		case 1:
			board[x][y] = 2
			newOverlaps++
		}
		x += xStep
		y += yStep
	}
	return newOverlaps
}

func (l *Line) isDiagonal() bool {
	return l.x1 != l.x2 && l.y1 != l.y2
}

func max2(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func (b *Board) print() {
	for y := 0; y < Y; y++ {
		for x := 0; x < X; x++ {
			switch v := b[x][y]; v {
			case 0:
				fmt.Print(".")
			default:
				fmt.Print(v)
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}