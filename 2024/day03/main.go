package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"github.com/cgbaker/AOC/2024/utils"
	"regexp"
)

type MulOp struct {
	a, b int
}

type Do struct {
}

type DoNot struct {
}

type Input struct {
	ops []interface{}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	input := readInput(file)
	prob1(input)
	prob2(input)
}

func prob1(input *Input) {
	sum := 0
	for _, op := range input.ops {
		switch op := op.(type) {
		case MulOp:
			sum += op.a * op.b
		default:
		}
	}
	fmt.Printf("prob1: %d\n",sum)
}

func prob2(input *Input) {
	sum := 0
	enabled := true
	for _, op := range input.ops {
		switch op := op.(type) {
		case Do:
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

func readInput(file *os.File) *Input {
	pattern := regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d+),(\d+)\)`)
	input := &Input{
		ops: []interface{}{},
	}
	lineScanner := bufio.NewScanner(file)
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		matches := pattern.FindAllStringSubmatch( lineScanner.Text(), -1 )
		for _, m := range matches {
			var op interface{}
			switch m[0] {
			case "do()":
				op = Do{}
			case "don't()":
				op = DoNot{}
			default:
				op = MulOp{utils.Atoi(m[1]), utils.Atoi(m[2])}
			}
			input.ops = append(input.ops, op)
		}
	}
	return input
}


