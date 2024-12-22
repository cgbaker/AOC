package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"github.com/cgbaker/AOC/2024/utils"
)

type Problem = []int

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	problem := readInput(file)
	// part1(problem)
	part2(problem)
}

func part1(problem Problem) {
	sum := 0
	for _, p := range problem {
		for range 2000 {
			p = solve(p)
		}
		sum += p
	}
	fmt.Printf("part1: %d\n",sum)
}

func solve(p0 int) int {
	MASK24 := 0xffffff
	p1 := ((p0 << 6) ^ p0) & MASK24
	p2 := ((p1 >> 5) ^ p1) & MASK24
	p3 := ((p2 << 11) ^ p2) & MASK24
	return p3
}

type Sequence [4]int

func (s Sequence) Shift(v int) Sequence {
	return Sequence{s[1],s[2],s[3],v}
}

type Entry struct {
	price, buyer int
}


func part2(problem Problem) {
	// for a given sequence s, sequences[s] contains two values:
	//    price: the price across all buyers achieved by buying on this sequence
	//    buyer: the max buyer seen so far  
	sequences := map[Sequence]Entry{}
	for buyerId, p := range problem {
		// fmt.Println("buyer:",p)
		seq := Sequence{}
		secret := p
		var price int
		// init the sequence
		for range 3 {
			secret, price, seq = advance(secret, seq)
			// fmt.Println(seq)
		}
		for range 1997 {
			secret, price, seq = advance(secret, seq)
			if e, ok := sequences[seq]; !ok || e.buyer < buyerId {
				sequences[seq] = Entry{
					price: e.price + price,
					buyer: buyerId,
				}
			}
			// fmt.Println(seq)
		}
		// fmt.Println(sequences)
	}
	best := 0
	for _, e := range sequences {
		if e.price > best {
			// fmt.Println("considered",s,"for",e.price)
			best = e.price
		}
	}
	fmt.Printf("part2: %d\n",best)
}

func advance(oldSecret int, oldSequence Sequence) (newSecret int, newPrice int, newSequence Sequence) {
	oldPrice := oldSecret % 10
	newSecret = solve(oldSecret)
	newPrice = newSecret % 10
	diff := newPrice - oldPrice
	newSequence = oldSequence.Shift(diff)
	return
}

func readInput(file *os.File) Problem {
	ints := []int{}
	lineScanner := bufio.NewScanner(file)
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		ints = append(ints, utils.Atoi(lineScanner.Text()))
	}
	return ints
}

