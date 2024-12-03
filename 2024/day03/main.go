package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"github.com/cgbaker/AOC/2024/utils"
	"regexp"
)

var (
	re *regexp.Regexp
)

type Op interface {}

type MulOp struct {
	a, b int
}

type DoOp struct {
}

type DoNot struct {
}

type Input struct {
	ops []Op
}

func main() {
	re = regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d+),(\d+)\)`)
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	input := readInput(file)
	// prob1(input)
	prob2(input)
}

func prob1(input *bufio.Scanner) {
	sum := 0
	for input.Scan() {
		switch op := parseOp(input.Text()).(type) {
		case MulOp:
			sum += op.a * op.b
		default:
		}
	}
	fmt.Printf("prob1: %d\n",sum)
}

func prob2(input *bufio.Scanner) {
	sum := 0
	enabled := true
	for input.Scan() {
		switch op := parseOp(input.Text()).(type) {
		case DoOp:
			enabled = true
		case DoNot:
			enabled = false
		case MulOp:
			if enabled {
				sum += op.a * op.b
			}
		}
	}
	fmt.Printf("prob2: %d\n",sum)
}

func readInput(file *os.File) *bufio.Scanner {
	regexScanner := bufio.NewScanner(file)
	regexScanner.Split(utils.SplitRegex(re))
	return regexScanner
}

func parseOp(s string) (op Op) {
	m := re.FindStringSubmatch(s)
	if m == nil {
		panic("This should have been a valid match")
	}
	switch m[0] {
	case "do()":
		op = DoOp{}
	case "don't()":
		op = DoNot{}
	default:
		op = MulOp{utils.Atoi(m[1]), utils.Atoi(m[2])}
	}
	return
}

