package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	// "github.com/cgbaker/AOC/2024/utils"
	"github.com/schollz/progressbar/v3"
	"strings"
)

type Problem struct {
}

func (p *Problem) size() int {
	return 0
}

func main() {
	file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	problem := readInput(file)
	bar := progressbar.Default(int64(problem.size()))
	part1(problem)
	part2(problem)
	bar.Add(problem.size())
}

func part1(_ *Problem) {
	fmt.Printf("part1: %d\n",0)
}

func part2(_ *Problem) {
	fmt.Printf("part2: %d\n",0)
}

func readInput(file *os.File) *Problem {
	lineScanner := bufio.NewScanner(file)
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		wordScanner := bufio.NewScanner(strings.NewReader(lineScanner.Text()))
		wordScanner.Split(bufio.ScanWords)
		for wordScanner.Scan() {
			// reading, _ := strconv.Atoi(wordScanner.Text())
			// report = append(report, reading)
		}
		// problem.reports = append(problem.reports, report)
	}
	return &Problem{
	}
}


