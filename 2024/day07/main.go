package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"github.com/cgbaker/AOC/2024/utils"
	"strings"
)

type Input struct {
	problems []Problem
}

type Problem struct {
	testVal int
	operands []int
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
	for _, p := range input.problems {
		numOps := len(p.operands)-1
		numPossibilities := 2 << (numOps-1)
		for idx := range numPossibilities {
			ops := genOps1(numOps, idx)
			possibility := evaluate(ops, p.operands)
			if possibility == p.testVal {
				// fmt.Printf("winner: %v against %s == %v\n",p.operands,string(ops),p.testVal)
				sum += p.testVal
				break
			}
		}
	}
	fmt.Printf("prob1: %d\n",sum)
}

func prob2(input *Input) {
	sum := 0
	for _, p := range input.problems {
		numOps := len(p.operands)-1
		numPossibilities := utils.Pow(3, numOps)
		for idx := range numPossibilities {
			ops := genOps2(numOps, idx)
			possibility := evaluate(ops, p.operands)
			if possibility == p.testVal {
				// fmt.Printf("winner: %v against %s == %v\n",p.operands,string(ops),p.testVal)
				sum += p.testVal
				break
			}
		}
	}
	fmt.Printf("prob1: %d\n",sum)
}

func readInput(file *os.File) *Input {
	input := &Input{
		problems: []Problem{},
	}
	lineScanner := bufio.NewScanner(file)
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		parts := strings.Split(lineScanner.Text(), ":")
		if len(parts) != 2 {
			panic("that was not what I was expecting")
		}
		p := Problem{
			testVal: utils.Atoi(parts[0]),
			operands: []int{},
		}
		wordScanner := bufio.NewScanner(strings.NewReader(parts[1]))
		wordScanner.Split(bufio.ScanWords)
		for wordScanner.Scan() {
			p.operands = append(p.operands, utils.Atoi(wordScanner.Text()))
		}
		input.problems = append(input.problems, p)
	}
	return input
}

func genOps1(numOps, index int) []rune {
	ops := []rune{}
	for len(ops) < numOps {
		if index & 1 == 0 {
			ops = append(ops,'+')
		} else {
			ops = append(ops,'*')
		}
		index = index >> 1
	}
	return ops
}

func genOps2(numOps, index int) []rune {
	ops := []rune{}
	for len(ops) < numOps {
		if index % 3 == 0 {
			ops = append(ops,'+')
		} else if index % 3 == 1 {
			ops = append(ops,'*')
		}  else {
			ops = append(ops,'|')
		}
		index = index / 3
	}
	return ops
}


func evaluate(ops []rune, operands []int) int {
	// fmt.Printf("checking %v against %s\n", operands, string(ops))
	a, b, operands := operands[0], operands[1], operands[2:]
	op, ops := ops[0], ops[1:]
	a = eval(op, a, b)
	for len(operands) > 0 {
		b, operands = operands[0], operands[1:]
		op, ops = ops[0], ops[1:]
		a = eval(op, a, b)
	}
	return a
}

func eval(op rune, a,b int) (ans int) {
	switch op {
	case '+':
		ans = a+b
	case '*':
		ans = a*b
	case '|':
		ans = utils.Atoi(fmt.Sprintf("%d%d",a,b))
	}
	return
}

