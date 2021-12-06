package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//fmt.Println(part1(file))
	fmt.Println(part2(file))
}

type Puzzle []int8

func NewPuzzle() Puzzle {
	return make([]int8, 0, 25)
}

var (
	masks = []int32{
		0b1111100000000000000000000,
		0b0000011111000000000000000,
		0b0000000000111110000000000,
		0b0000000000000001111100000,
		0b0000000000000000000011111,
		0b1000010000100001000010000,
		0b0100001000010000100001000,
		0b0010000100001000010000100,
		0b0001000010000100001000010,
		0b0000100001000010000100001,
	}
)

func part1(file *os.File) int {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	numbers := readNumbers(scanner)
	fmt.Println("Numbers: ", len(numbers))
	winningScore := -1
	minToWin := len(numbers) + 1
	for {
		puzzle := readPuzzle(scanner)
		if puzzle == nil {
			break
		}
		numToWin, score := checkPuzzle(numbers, puzzle)
		if numToWin < minToWin {
			winningScore = score
			minToWin = numToWin
		}
	}
	fmt.Println("Winning number: ", numbers[minToWin-1])
	fmt.Println("Winning score: ", winningScore)
	return int(numbers[minToWin-1]) * winningScore
}

func part2(file *os.File) int {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	numbers := readNumbers(scanner)
	fmt.Println("Numbers: ", len(numbers))
	lastToWin := -1
	losersScore := 0
	for {
		puzzle := readPuzzle(scanner)
		if puzzle == nil {
			break
		}
		numToWin, score := checkPuzzle(numbers, puzzle)
		if numToWin > lastToWin {
			losersScore = score
			lastToWin = numToWin
		}
	}
	fmt.Println("Losers ball: ", numbers[lastToWin-1])
	fmt.Println("Losers score: ", losersScore)
	return int(numbers[lastToWin-1]) * losersScore
}

func readNumbers(scanner *bufio.Scanner) []int8 {
	if ok := scanner.Scan(); !ok  {
		panic("scanning first word")
	}
	fmt.Println(scanner.Text())
	tokens := strings.Split(scanner.Text(), ",")
	numbers := make([]int8,0,len(tokens))
	for _, t := range tokens {
		var v int8
		if n, err := fmt.Sscan(t, &v); n == 1 && err == nil {
			numbers = append(numbers, v)
		}
	}
	return numbers
}

func readPuzzle(scanner *bufio.Scanner) Puzzle {
	puzzle := NewPuzzle()
	for len(puzzle) < 25 && scanner.Scan() {
		var v int8
		if n, err := fmt.Sscan(scanner.Text(), &v); n == 1 && err == nil {
			puzzle = append(puzzle, v)
		} else {
			panic(err.Error() + fmt.Sprint(n))
		}
	}
	switch len(puzzle) {
	case 25:
		return puzzle
	case 0:
		return nil
	default:
		panic("incomplete puzzle")
	}
}

func checkPuzzle(numbers []int8, puzzle Puzzle) (int, int) {
	var card int32
	for call, ball := range numbers {
		for i, p := range puzzle {
			if p == ball {
				card |= 1 << (24 - i)
				if winner(card) {
					return call+1, score(card, puzzle)
				}
			}
		}
	}
	return len(numbers)+1,-1
}

func winner(card int32) bool {
	for _, m := range masks {
		if m & card == m {
			return true
		}
	}
	return false
}

func score(card int32, puzzle Puzzle) int {
	score := 0
	var bit int32 = 1 << 24
	for _, p := range puzzle {
		if card & bit == 0 {
			score += int(p)
		}
		bit >>= 1
	}
	return score
}