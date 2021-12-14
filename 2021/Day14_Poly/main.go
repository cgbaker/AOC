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

	puzzle := readPuzzle(file)
	//puzzle.print()

	fmt.Println("Part1:", puzzle.expand(10))
	fmt.Println("Part2:", puzzle.expandMemo(40))
}

func readPuzzle(file *os.File) *Puzzle {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	var puzzle = &Puzzle{
		init:  scanner.Text(),
		rules: map[Pair]uint8{},
	}
	scanner.Scan() // blank line
	for scanner.Scan() {
		var a, b, c uint8
		if _, err := fmt.Sscanf(scanner.Text(), "%c%c -> %c", &a, &b, &c); err != nil {
			break
		}
		puzzle.rules[Pair{a,b}] = c
	}
	return puzzle
}

type Pair struct {
	left, right uint8
}

type Puzzle struct {
	init  string
	rules map[Pair]uint8
}

func (p *Puzzle) expand(steps int) int {
	counts := map[uint8]int{}

	a := p.init[0]
	counts[a]++
	//fmt.Print(string(a))
	for i := 1; i < len(p.init); i++ {
		b := p.init[i]
		counts[b]++
		p.recurse(a, b, steps, counts)
		a = b
		//fmt.Print(string(a))
	}
	fmt.Println("")
	for k, v := range counts {
		fmt.Println(string(k),v)
	}

	vals := make([]int,0,len(counts))
	for _, v := range counts {
		vals = append(vals, v)
	}
	sort.Ints(vals)
	return vals[len(vals)-1] - vals[0]
}

func (p *Puzzle) expandMemo(steps int) int {
	counts := map[uint8]int{}
	memo = map[MemoKey]map[uint8]int{}

	a := p.init[0]
	counts[a]++
	//fmt.Print(string(a))
	for i := 1; i < len(p.init); i++ {
		b := p.init[i]
		counts[b]++
		appendCounts(counts, p.recurseMemo(a, b, steps))
		a = b
		//fmt.Print(string(a))
	}
	fmt.Println("")
	for k, v := range counts {
		fmt.Println(string(k),v)
	}

	vals := make([]int,0,len(counts))
	for _, v := range counts {
		vals = append(vals, v)
	}
	sort.Ints(vals)
	return vals[len(vals)-1] - vals[0]
}


func (p *Puzzle) print() {
	for k, v := range p.rules {
		fmt.Println(string(k.left), string(k.right), " -> ", string(v))
	}
}

func (p *Puzzle) recurse(a uint8, b uint8, steps int, counts map[uint8]int) {
	if steps == 0 {
		return
	}
	mid, ok := p.rules[Pair{a,b}]
	if !ok {
		return
	}
	// a, mid, b
	counts[mid]++
	p.recurse(a,mid,steps-1,counts)
	//fmt.Print(string(mid))
	p.recurse(mid,b,steps-1,counts)
}

type MemoKey struct {
	pair Pair
	steps int
}

var (
	memo map[MemoKey]map[uint8]int
)

func (p *Puzzle) recurseMemo(a uint8, b uint8, steps int) map[uint8]int {
	counts := map[uint8]int{}
	if steps == 0 {
		return counts
	} else if m, ok := memo[MemoKey{Pair{a,b},steps}]; ok {
		return m
	}
	mid, ok := p.rules[Pair{a,b}]
	if !ok {
		return counts
	}
	// a, mid, b
	counts[mid]++
	appendCounts(counts, p.recurseMemo(a,mid,steps-1))
	appendCounts(counts, p.recurseMemo(mid,b,steps-1))
	memo[MemoKey{Pair{a,b},steps}] = counts
	return counts
}

func appendCounts(m1, m2 map[uint8]int) {
	for k, v := range m2 {
		m1[k] += v
	}
}