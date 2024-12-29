package main

import (
	"errors"
	"fmt"
	"os"
	"log"
	"bufio"
	// "github.com/cgbaker/AOC/2024/utils"
	"slices"
	"strings"
)

type Gate struct {
	op string
	a, b string
}

var (
	XOR string = "XOR"
	AND string = "AND"
	OR string = "OR"
)

type Problem struct {
	values map[string]int
	graph map[string]Gate
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	problem := readInput(file)
	part1(problem)
	fmt.Println("")
	part2(problem)
}

func part1(problem *Problem) {
	answer := 0
	zind := 0
	for {
		zstr := fmt.Sprintf("z%2.2d",zind)
		zval := solve(problem,zstr)
		// fmt.Printf("%s is %v\n",zstr,zval)
		if zval == -1 {
			break
		}
		answer += zval * (1 << zind)
		zind++
	}
	fmt.Printf("part1: %d\n",answer)
}

func fmtnode(lbl byte, ind int) string {
	return fmt.Sprintf("%c%2.2d",rune(lbl),ind)
}

func get(problem *Problem, xy byte) (int,int) {
	answer := 0
	ind := 0
	for {
		lbl := fmtnode(xy,ind)
		if val, ok := problem.values[lbl]; !ok {
			break
		} else {
			// fmt.Printf("%s is %v\n",lbl,val)
			answer += val * (1 << ind)
			ind++
		}
	}
	return answer,ind
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
	case AND:
		val = a & b
	case XOR:
		val = a ^ b
	case OR:
		val = a | b
	}
	problem.values[wire] = val
	return val
}

var (
	lastZ string
)

func part2(problem *Problem) {
	x, _ := get(problem, 'x')
	y, _ := get(problem, 'y')
	fmt.Printf("x: %d\n",x)
	fmt.Printf("y: %d\n",y)
	z, LENZ := get(problem,'z')
	fmt.Printf("z(exp): %b\n",x+y)
	fmt.Printf("z(act): %b\n",z)
	fmt.Printf("%d-bit adder\n", LENZ)
	lastZ = fmt.Sprintf("z%2.2d",LENZ-1)

	bad := map[string]bool{}
	inputsum := map[string]bool{}
	numAnds := 0

	// xyand := map[string]bool{}
	fmt.Println("\nChecking graph:\n")
	for out, gate := range problem.graph {
		switch gate.op {
		case AND:
			numAnds++
			if out[0] == 'z' {
				fmt.Printf("bad output gate: %+v -> %s\n",gate,out)
				bad[out] = true
			}
			if inputPair(gate.a, gate.b) {
			}
		case OR:
			if out[0] == 'z' && out != lastZ {
				fmt.Printf("bad output gate: %+v -> %s\n",gate,out)
				bad[out] = true
			} 
			if problem.graph[gate.a].op != AND {
				fmt.Printf("bad OR carry input: %s\n", gate.a)
				bad[gate.a] = true
			} 
			if problem.graph[gate.b].op != AND {
				fmt.Printf("bad OR carry input: %s\n", gate.b)
				bad[gate.b] = true
			}
		case XOR:
			if inputPair(gate.a, gate.b) {
				inputsum[out] = true
			} else {
				if out[0] != 'z' {
					fmt.Printf("unexpected XOR: %+v -> %s\n",gate,out)
					bad[out] = true
				}
			}
		}
	}

	fmt.Println("num input sums:",len(inputsum))
	fmt.Println("num ANDs:",numAnds)

	for out, gate := range problem.graph {
		if out[0] != 'z' {
			continue
		}
		if err := matchOutputGraph(out, gate, problem); err != nil {
			fmt.Printf("problem with sub-circuit %s: %v\n",out, err)
		}
	}

	output := []string{}
	for k := range bad {
		output = append(output,k)
	}
	slices.Sort(output)
	fmt.Println("\npart2:", strings.Join(output,","))
}

func matchOutputGraph(output string, gate Gate, problem *Problem) error {
	switch gate.op {
	case XOR:
		// if first output bit, only two parents: the XOR of the first input bits
		if output == "z00" {
			if inputPair(gate.a, gate.b) && gate.a[1:3] == "00" {
				return nil
			}
			return errors.New("bad first bit")
		}
		// otherwise, two parents: a carry (itself from an OR) and XOR(x,y)
		clbl, inlbl := gate.a, gate.b
		gin, gcarry := problem.graph[inlbl], problem.graph[clbl]
		if gin.op != XOR {
			gin, gcarry = gcarry, gin
			clbl, inlbl = inlbl, clbl
		}
		if gin.op != XOR {
			return errors.New(fmt.Sprintf("parent gates not XOR and OR: %s:%s",inlbl,gin.op))
		} else if gcarry.op != OR {
			return errors.New(fmt.Sprintf("parent gates not XOR and OR: %s:%s",clbl,gcarry.op))
		}
		if !inputPair(gin.a, gin.b) || gin.a[1:3] != output[1:3] {
			return errors.New("input XOR parent gate is not input-tied")
		}
		carryp1, carryp2 := problem.graph[gcarry.a], problem.graph[gcarry.b]
		if carryp1.op != AND || carryp2.op != AND {
			return errors.New(fmt.Sprintf("carry OR parent gate does not have AND parents: %s and %s",gcarry.a,gcarry.b))
		}
		// TODO: check that they're the correct inputs (previous bit)
		if !inputPair(carryp1.a,carryp1.b) && !inputPair(carryp2.a,carryp2.b) {
			return errors.New("carry AND grandparent is not input-tied")
		}
	case OR:
		// zlast is simply the last carry bit
		if output != lastZ {
			return errors.New("OR gate on non-terminal bit")
		}
	default:
		return errors.New("bad gate (swap)")
	}
	return nil
}

func inputPair(a,b string) bool {
	if (a[0] == 'x' && b[0] == 'y') || (a[0] == 'y' && b[0] == 'x') {
		if a[1:2] == b[1:2] {
			return true
		}
	}
	return false
}

func readInput(file *os.File) *Problem {
	lineScanner := bufio.NewScanner(file)
	lineScanner.Split(bufio.ScanLines)
	values := map[string]int{}
	ops := map[string]Gate{}
	for lineScanner.Scan() {
		var a, b, c, op string
		var val int
		if n, _ := fmt.Sscanf(lineScanner.Text(), "%3s: %d",&a,&val); n == 2 {
			values[a] = val
		} else if n, _ := fmt.Sscanf(lineScanner.Text(), "%s %s %s -> %s",&a,&op,&b,&c); n == 4 {
			ops[c] = Gate{
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


