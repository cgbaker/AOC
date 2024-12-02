package main

import (
	"errors"
	"fmt"
	"os"
	"log"
	"bufio"
	"strconv"
	"strings"
	"github.com/cgbaker/AOC/2024/utils"
)

type Input struct {
	reports [][]int
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
	numSafeReports := 0
	for _, report := range input.reports {
		if isSafeReport(report) {
			numSafeReports++
		}
	}
	fmt.Printf("prob1: %d\n",numSafeReports)
}

func prob2(input *Input) {
	numSafeReports := 0
	for _, report := range input.reports {
		for skip := 0; skip < len(report); skip++ {
			if isSafeDampedReport(report, skip) {
				numSafeReports++
				break
			}
		}
	}
	fmt.Printf("prob2: %d\n",numSafeReports)
}

func isSafeReport(report []int) bool {
	if len(report) <= 1 {
		return true
	}
	globalTrend, err := getTrend(report[0], report[1])
	if err != nil {
		return false
	}
	for i := 1; i < len(report)-1; i++ {
		newTrend, err := getTrend(report[i], report[i+1])
		if err != nil || newTrend != globalTrend {
			return false
		}
	}
	return true
}

func isSafeDampedReport(report []int, skipIndex int) bool {
	if len(report) <= 1 || (len(report) <= 2 && skipIndex != -1) {
		return true
	}
	i := 0
	if i == skipIndex {
		i++
	}
	j := i+1
	if j == skipIndex {
		j++
	}
	globalTrend, err := getTrend(report[i], report[j])
	if err != nil {
		return false
	}
	for {
		i = j
		j++
		if j == skipIndex {
			j++
		}
		if j >= len(report) { 
			break
		}
		newTrend, err := getTrend(report[i], report[j])
		if err != nil || newTrend != globalTrend {
			return false
		}
	}
	return true
}

func getTrend(x, y int) (int, error) {
	d := utils.AbsInt(y-x)
	if d < 1 || d > 3 {
		return 0, errors.New("not safe")
	}
	return utils.SignInt(y-x), nil
}

func readInput(file *os.File) *Input {
	input := &Input{
		reports: [][]int{},
	}
	lineScanner := bufio.NewScanner(file)
	for lineScanner.Scan() {
		report := []int{}
		wordScanner := bufio.NewScanner(strings.NewReader(lineScanner.Text()))
		wordScanner.Split(bufio.ScanWords)
		for wordScanner.Scan() {
			reading, _ := strconv.Atoi(wordScanner.Text())
			report = append(report, reading)
		}
		input.reports = append(input.reports, report)
	}
	return input
}


