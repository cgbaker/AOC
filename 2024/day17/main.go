package main

import (
	"errors"
	"fmt"
	"os"
	"log"
	"bufio"
	"github.com/cgbaker/AOC/2024/utils"
	"slices"
	"strings"
	"strconv"
)

var (
	REG_A int = 0
	REG_B int = 1
	REG_C int = 2

	OP_ADV int = 0
	OP_BXL int = 1
	OP_BST int = 2
	OP_JNZ int = 3
	OP_BXC int = 4
	OP_OUT int = 5
	OP_BDV int = 6
	OP_CDV int = 7
)

type Computer struct {
	registers []int
	iPtr int
}

func (c Computer) Fetch(operand int) int {
	if operand < 0 || operand > 6 {
		panic("bad operand")
	}
	if operand < 4 {
		return operand
	}
	return c.registers[operand-4]
}

type Instruction struct {
	icode int
	operand int
}

type Problem struct {
	computer Computer
	program []int
}
func (p *Problem) GetInstruction(iPtr int) (Instruction,error) {
	if iPtr < 0 || iPtr >= len(p.program) {
		return Instruction{}, errors.New("out of bounds")
	}
	inst := Instruction{
		icode: p.program[iPtr],
	}
	if iPtr+1 < len(p.program) {
		inst.operand = p.program[iPtr+1]
	}
	return inst,nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	problem := readInput(file)
	// fmt.Printf("%+v\n",problem.computer)
	// fmt.Printf("%+v\n",problem.program)
	part1(problem)
	part2(problem)
}

func part1(problem *Problem) {
	outputs := []int{}
	computer := problem.computer
	for {
		inst, err := problem.GetInstruction(computer.iPtr)
		if err != nil {
			break
		}
		var out *int
		// fmt.Println("step:",inst)
		computer, out = step(computer, inst)
		// fmt.Println("state:",computer)
		if out != nil {
			outputs = append(outputs,*out)
		}
	}

	outputStr := ""
	for i, o := range outputs {
		if i > 0 {
			outputStr = outputStr + ","
		}
		outputStr += strconv.FormatInt(int64(o),10)
	}
	fmt.Println("prob1:",outputStr)
}

func part2(problem *Problem) {
	candidates := []int{0}
	program := slices.Clone(problem.program)
	for len(program) > 0 {
		target := program[len(program)-1]
		program = program[:len(program)-1]
		children := []int{}
		for _, c := range candidates {
			for i := range 8 {
				a := (c << 3) + i
				b := a & 7
				b ^= 1
				c := a / (1<<b)
				b ^= c
				b ^= 4
				if b&7 == target {
					children = append(children,a)
				}
			}
		}
		candidates = children
	}
	fmt.Println("prob2:",candidates)
}


func step(curr Computer, inst Instruction) (Computer,*int) {
	next := Computer{
		iPtr: curr.iPtr+2,
		registers: slices.Clone(curr.registers),
	}
	switch inst.icode {
	case OP_ADV:
		d := 1 << next.Fetch(inst.operand)
		next.registers[REG_A] = next.registers[REG_A] / d
	case OP_BXL:
		next.registers[REG_B] ^= inst.operand
	case OP_BST:
		next.registers[REG_B] = next.Fetch(inst.operand) & 7
	case OP_JNZ:
		if next.registers[REG_A] != 0 {
			next.iPtr = inst.operand
		}
	case OP_BXC:
		next.registers[REG_B] ^= next.registers[REG_C]
	case OP_OUT:
		out := next.Fetch(inst.operand) & 7
		return next, &out
	case OP_BDV:
		d := 1 << next.Fetch(inst.operand)
		next.registers[REG_B] = next.registers[REG_A] / d
	case OP_CDV:
		d := 1 << next.Fetch(inst.operand)
		next.registers[REG_C] = next.registers[REG_A] / d
	}
	return next, nil
}

func readInput(file *os.File) *Problem {
	lineScanner := bufio.NewScanner(file)
	lineScanner.Split(bufio.ScanLines)
	lines := []string{}
	for lineScanner.Scan() {
		lines = append(lines, lineScanner.Text())
	}
	var regA, regB, regC int
	fmt.Sscanf(lines[0],"Register A: %d", &regA)
	fmt.Sscanf(lines[1],"Register B: %d", &regB)
	fmt.Sscanf(lines[2],"Register C: %d", &regC)
	computer := Computer{
		registers: []int{regA,regB,regC},
		iPtr: 0,
	}
	var programTxt string
	fmt.Sscanf(lines[4],"Program: %s",&programTxt)
	insts := []int{}
	for _, s := range strings.Split(programTxt,",") {
		insts = append(insts, utils.Atoi(s))
	}
	return &Problem{
		computer: computer,
		program: insts,
	}
}


