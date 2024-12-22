package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"github.com/cgbaker/AOC/2024/utils"
)

type Problem struct {
	codes []string
}

func main() {
	file, err := os.Open("sample.txt")
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
	for _, code := range problem.codes {
		solution := solve(code)
		sum += utils.Atoi(code[:len(code)-2])*len(solution)
	}
	fmt.Printf("part1: %d\n",sum)
}

func part2(_ *Problem) {
	fmt.Printf("part2: %d\n",0)
}

var (
	KEY_UP byte = '^'
	KEY_DOWN byte = 'v'
	KEY_LEFT byte = '<'
	KEY_RIGHT byte = '>'

	KEY_A byte = 'A'

	KEY_0 byte = '0'
	KEY_1 byte = '1'
	KEY_2 byte = '2'
	KEY_3 byte = '3'
	KEY_4 byte = '4'
	KEY_5 byte = '5'
	KEY_6 byte = '6'
	KEY_7 byte = '7'
	KEY_8 byte = '8'
	KEY_9 byte = '9'
)

type RobotState struct {
	key byte
}

type GlobalState struct {
	robot1, robot2, robot3 RobotState
}

func InitState() GlobalState {
	return GlobalState {
		robot1: RobotState{KEY_A},
		robot2: RobotState{KEY_A},
		robot3: RobotState{KEY_A},
	}
}

func solve(code string) string {
	// need to find the shortest set of commands
	// Human: direct control over dpad, controlling robot 1
	// Robot1: movement control over dpad, controlling robot 2
	// Robot2: movement control over dpad, controlling robot 3
	// Robot3: movement control over keypad at door
	// A solution must result in Robot3 entering the code
	// Robot3 

	state := InitState()

	return code
}

func readInput(file *os.File) *Problem {
	codes := []string{}
	lineScanner := bufio.NewScanner(file)
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		codes = append(codes, lineScanner.Text())
	}
	return &Problem{
		codes: codes,
	}
}


