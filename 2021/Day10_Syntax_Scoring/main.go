package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	part1, part2 := checkBlocks(scanner)
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}

func checkBlocks(scanner *bufio.Scanner) (int, int) {
	corruptionScore := 0
	completionScores := []int{}
	for scanner.Scan() {
		corrupt, complete := checkLineCompletion(scanner.Text())
		if complete != 0 {
			completionScores = append(completionScores, complete)
		} else {
			corruptionScore += corrupt
		}
	}
	sort.Ints(completionScores)
	return corruptionScore, completionScores[len(completionScores)/2]
}

func checkLineCompletion(line string) (int, int) {
	stack := []int32{}
	for _, b := range line {
		switch (b) {
		case '(', '[', '{', '<':
			stack = append(stack,b)
		case ')', ']', '}', '>':
			if len(stack) == 0 || !match(stack[len(stack)-1],b) {
				// corrupted
				return scoreCorrupted(b), 0
			}
			stack = stack[0:len(stack)-1]
		default:
			panic(fmt.Sprint("unknown character",b))
		}
	}
	if len(stack) == 0 {
		return 0, 0
	}
	// complete
	score := 0
	for i := len(stack)-1; i >= 0; i-- {
		score *= 5
		switch (stack[i]) {
		case '(':
			score += 1
		case '[':
			score += 2
		case '{':
			score += 3
		case '<':
			score += 4
		}
	}
	return 0, score
}

func scoreCorrupted(b int32) int {
	switch (b) {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	}
	panic("scoring panic")
}

func match(a, b int32) bool {
	return (a == '(' && b == ')') ||
		(a == '{' && b == '}') ||
		(a == '[' && b == ']') ||
		(a == '<' && b == '>')
}