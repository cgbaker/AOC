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

	puzzle := readPuzzle(scanner)

	fmt.Println("Part1:", puzzle.fold(true,scanner))
	fmt.Println("Part2:", puzzle.fold(false,scanner))
}

func readPuzzle(scanner *bufio.Scanner) *Puzzle {
	var puzzle = &Puzzle{
	}
	scanner.Split(bufio.ScanLines)
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
}

func (p *Puzzle) append(x, y int) {
}

func (p *Puzzle) fold(firstFoldOnly bool, scanner *bufio.Scanner) int {
	for {
		axis := ""
		line := -1
		fmt.Sscanf(scanner.Text(), "fold along %s=%d",&axis,&line)

		if firstFoldOnly || !scanner.Scan() {
			break
		}
	}
	return 0
}