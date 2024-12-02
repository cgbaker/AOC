package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	// "github.com/cgbaker/AOC/2024/utils"
	"strings"
)

type Input struct {
}

func main() {
	file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	input := readInput(file)
	prob1(input)
	prob2(input)
}

func prob1(_ *Input) {
	fmt.Printf("prob1: %d\n",0)
}

func prob2(_ *Input) {
	fmt.Printf("prob2: %d\n",0)
}

func readInput(file *os.File) *Input {
	input := &Input{
	}
	lineScanner := bufio.NewScanner(file)
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		wordScanner := bufio.NewScanner(strings.NewReader(lineScanner.Text()))
		wordScanner.Split(bufio.ScanWords)
		for wordScanner.Scan() {
			// reading, _ := strconv.Atoi(wordScanner.Text())
			// report = append(report, reading)
		}
		// input.reports = append(input.reports, report)
	}
	return input
}


