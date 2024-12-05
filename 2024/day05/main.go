package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"github.com/cgbaker/AOC/2024/utils"
	"slices"
	"sort"
	"strings"
)

type Input struct {
	rules map[int][]int
	updates [][]int
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
	for _, u := range input.updates {
		if isCorrect(input.rules, u) {
			sum += utils.GetMiddle(u)
		}
	}
	fmt.Printf("prob1: %d\n",sum)
}

func prob2(input *Input) {
	sum := 0
	for _, u := range input.updates {
		if !isCorrect(input.rules, u) {
			newUpdate := reorder(input.rules, u)
			sum += utils.GetMiddle(newUpdate)
		}
	}
	fmt.Printf("prob2: %d\n",sum)
}

func readInput(file *os.File) *Input {
	input := &Input{
		rules: make(map[int][]int),
		updates: [][]int{},
	}
	lineScanner := bufio.NewScanner(file)
	lineScanner.Split(bufio.ScanLines)
	// scan rules
	for lineScanner.Scan() {
		rule := lineScanner.Text()
		if len(rule) == 0 {
			break
		}
		var before, after int
		fmt.Sscanf(rule, "%d|%d", &before, &after)
		if r, ok := input.rules[before]; ok {
			input.rules[before] = append(r, after)
		} else {
			input.rules[before] = []int{after}
		}
	}
	// scan updates
	for lineScanner.Scan() {
		update := lineScanner.Text()
		strPages := strings.Split(update, ",")
		intPages := make([]int, 0, len(strPages))
		for _, s := range strPages {
			intPages = append(intPages, utils.Atoi(s))
		}
		input.updates = append(input.updates, intPages)
	}
	// sort rules for speed
	for _, v := range input.rules {
		sort.Ints(v)
	}
	return input
}

func isCorrect(rules map[int][]int, update []int) bool {
	for i, b := range update {
		for _, a := range update[i:] {
			// b is before a
			// if there is a rule saying that a must be before b, then this is not a value update
			if goesBefore(rules, a, b) {
				return false
			}
		}
	}
	return true
}

func goesBefore(rules map[int][]int, a, b int) bool {
	if r, ok := rules[a]; ok {
		if _, found := slices.BinarySearch(r, b); found {
			return true
		}
	}
	return false
}

// return negative when a is "less than" b
// return zero when equal
// return positive when b is "less than" a
func RuleOrder(rules map[int][]int, a, b int) int {
	if goesBefore(rules, a,b) {
		return -1
	} else if goesBefore(rules, b,a) {
		return 1
	}
	return 0
}

func reorder(rules map[int][]int, update []int) []int {
	newUpdate := slices.Clone(update)
	slices.SortFunc(newUpdate, func (a, b int) int {return RuleOrder(rules,a,b)})
	return newUpdate
}
