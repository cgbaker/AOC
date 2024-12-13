package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"math"
)

type XY struct {
	x, y float64
}

type Machine struct {
	buttonA XY
	buttonB XY
	prize XY
}

type Problem struct {
	machines []Machine
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
	sum := 0
	for _, machine := range problem.machines {
		sum += solve(machine)
	}
	fmt.Printf("prob1: %v\n",sum)
}

func part2(problem *Problem) {
	OFFSET := 10000000000000.0
	sum := 0
	for _, machine := range problem.machines {
		sum += solve(Machine{
			buttonA: machine.buttonA,
			buttonB: machine.buttonB,
			prize: XY{
				x: machine.prize.x+OFFSET, 
				y: machine.prize.y+OFFSET,
			},
		})
	}
	fmt.Printf("prob2: %v\n",sum)
}

func solve(machine Machine) int {
	// A = [a11 a12]   [buttonA.x buttonB.x] 
	//     [a21 a22] = [buttonA.y buttonB.y]
	//
	// A [numA;numB] = [prize.x;prize.y]
	alpha := machine.buttonA.y / machine.buttonA.x
	aprime22 := machine.buttonB.y - alpha*machine.buttonB.x
	numB := (machine.prize.y - alpha*machine.prize.x) / aprime22
	numA := (machine.prize.x - machine.buttonB.x * numB) / machine.buttonA.x 
	// fmt.Printf("solution: %v,%v\n",numA,numB)
	solA, solB := int(math.Round(numA)), int(math.Round(numB))
	if math.Abs(float64(solA)-numA) > .01 || math.Abs(float64(solB)-numB) > 0.01 {
		return 0
	}
	return solA*3 + solB
}

func readInput(file *os.File) *Problem {
	machines := []Machine{}
	lineScanner := bufio.NewScanner(file)
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		machine := Machine{}
		var x, y int
		fmt.Sscanf(lineScanner.Text(),"Button A: X+%d, Y+%d",&x,&y)
		machine.buttonA = XY{float64(x),float64(y)}
		lineScanner.Scan()
		fmt.Sscanf(lineScanner.Text(),"Button B: X+%d, Y+%d",&x,&y)
		machine.buttonB = XY{float64(x),float64(y)}
		lineScanner.Scan()
		fmt.Sscanf(lineScanner.Text(),"Prize: X=%d, Y=%d",&x,&y)
		machine.prize = XY{float64(x),float64(y)}
		lineScanner.Scan()
		machines = append(machines,machine)
	}
	return &Problem{
		machines: machines,
	}
}


