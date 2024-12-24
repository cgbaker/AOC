package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	// "github.com/cgbaker/AOC/2024/utils"
)

type Op struct {
	op string
	a, b string
}

type Problem struct {
	values map[string]int
	graph map[string]Op
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	problem := readInput(file)
	part1(problem)
	part2(problem)
}

func part1(problem *Problem) {
	answer := 0
	zind := 0
	for {
		zstr := fmt.Sprintf("z%2.2d",zind)
		zval := solve(problem,zstr)
		fmt.Printf("%s is %v\n",zstr,zval)
		if zval == -1 {
			break
		}
		answer += zval * (1 << zind)
		zind++
	}
	fmt.Printf("part1: %d\n",answer)
}

func solve(problem *Problem, wire string) int {
	if v, ok := problem.values[wire]; !ok {
		return -1
	} else if v != -1 {
		return v
	}
	op, exists := problem.graph[wire]
	if !exists {
		panic(fmt.Sprintf("no graph entry for wire %s",wire))
	}
	a, b := solve(problem, op.a), solve(problem, op.b)
	var val int
	switch op.op {
	case "AND":
		val = a & b
	case "XOR":
		val = a ^ b
	case "OR":
		val = a | b
	}
	problem.values[wire] = val
	return val
}

func part2(_ *Problem) {
	fmt.Printf("part2: %d\n",0)
}

func readInput(file *os.File) *Problem {
	lineScanner := bufio.NewScanner(file)
	lineScanner.Split(bufio.ScanLines)
	values := map[string]int{}
	ops := map[string]Op{}
	for lineScanner.Scan() {
		var a, b, c, op string
		var val int
		if n, _ := fmt.Sscanf(lineScanner.Text(), "%3s: %d",&a,&val); n == 2 {
			values[a] = val
		} else if n, _ := fmt.Sscanf(lineScanner.Text(), "%s %s %s -> %s",&a,&op,&b,&c); n == 4 {
			ops[c] = Op{
				a: a, 
				b: b,
				op: op,
			}
			values[c] = -1
		}
	}
	return &Problem{
		graph: ops,
		values: values,
	}
}


