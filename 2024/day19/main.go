package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"github.com/schollz/progressbar/v3"
	"strings"
)

type Problem struct {
	towels []string
	patterns []string
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	problem := readInput(file)
	sumSolutions := 0
	sumSolvable := 0
	bar := progressbar.Default(int64(len(problem.patterns)))
	for _, pattern := range problem.patterns {
		num := numSolutions(pattern, problem.towels)
		if num > 0 {
			sumSolvable++
		}
		sumSolutions += num
		bar.Add(1)
	}
	fmt.Printf("part1: %d\n",sumSolvable)
	fmt.Printf("part2: %d\n",sumSolutions)
}

var memo = map[string]int{}

func numSolutions(pattern string, towels []string) int {
	if num, visited := memo[pattern]; visited {
		return num
	}
	sum := 0
	for _, t := range towels {
		if strings.Compare(pattern,t) == 0 {
			sum++
		} else if sub, hasPrefix := strings.CutPrefix(pattern, t); hasPrefix {
			sum += numSolutions(sub, towels)
		}
	}
	memo[pattern] = sum
	return sum
}

func pop(stack []string) ([]string,string) {
	return stack[:len(stack)-1], stack[len(stack)-1]
}



func readInput(file *os.File) *Problem {
	towels := []string{}
	patterns := []string{}
	lineScanner := bufio.NewScanner(file)
	lineScanner.Split(bufio.ScanLines)
	lineScanner.Scan()
	for _, t := range strings.Split(lineScanner.Text(),",") {
		towels = append(towels, strings.TrimPrefix(t, " "))
	}
	// eat the newline
	lineScanner.Scan()
	for lineScanner.Scan() {
		patterns = append(patterns, lineScanner.Text())
	}
	return &Problem{
		towels: towels,
		patterns: patterns,
	}
}


