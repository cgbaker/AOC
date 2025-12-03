package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/cgbaker/AOC/2024/utils"
)

type Problem struct {
	codes []string
}

type CacheKey struct {
	src, dst, robot byte
}

type Move struct {
	src, dst byte
}

func NewCacheKey(robot byte, move Move) CacheKey {
	return CacheKey{
		src:   move.src,
		dst:   move.dst,
		robot: robot,
	}
}

var (
	NUMPAD     *utils.CharGrid
	DIRPAD     *utils.CharGrid
	CACHE      map[CacheKey]int
	NUM_ROBOTS byte
)

func main() {
	// init globals
	NUMPAD = utils.NewCharGrid(4, 3)
	NUMPAD.Chars = []byte{'7', '8', '9', '4', '5', '6', '1', '2', '3', 0, '0', 'A'}
	DIRPAD = utils.NewCharGrid(2, 3)
	DIRPAD.Chars = []byte{0, '^', 'A', '<', 'v', '>'}
	CACHE = map[CacheKey]int{}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	problem := readInput(file)
	NUM_ROBOTS = 3
	sum := 0
	for _, code := range problem.codes {
		solution := solve(code)
		complexity := utils.Atoi(code[:len(code)-1]) * solution
		fmt.Println("code:", code, "len(solution):", solution, "complexity:", complexity)
		sum += complexity
	}
	fmt.Printf("part1: %d\n", sum)
}

// for the last robot,the numpad, entering the code can be done in a number of steps
// A -> first number
// A to press
// first number -> second number
// A to press
// ....
// last number -> A
// A to press
//
// each of the movements realizes a specific (easily-established) optimal number of inputs
// assume we have that list of inputs for 029A:
//    0 2   9   A
//   <A^A>^^AvvvA
// this is 12 button presses
//
// now add one dirpad indirection, starting on A
// for each dirpad above, we have to move from A to the appropriate spot (using the shortest path) and press A

var (
	KEY_N byte = '^'
	KEY_S byte = 'v'
	KEY_W byte = '<'
	KEY_E byte = '>'

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

	KEY_BLANK byte = 0

	HUMAN byte = 0
)

func solve(code string) int {
	return costToOutput(NUM_ROBOTS, code)
}

func costToOutput(robot byte, output string) int {
	if robot == HUMAN {
		return len(output)
	}
	current := KEY_A
	totalCost := 0
	for _, c := range []byte(output) {
		move := NewMove(current, c)
		key := NewCacheKey(robot, move)
		if entry, ok := CACHE[key]; ok {
			totalCost += entry
		} else {
			input := ComputeInput(move)
			cost := costToOutput(robot-1, input)
			CACHE[key] = cost
			totalCost += cost
		}
		current = c
	}
	return totalCost
}

func ComputeInput(move Move) string {
	return "A"
}

func NewMove(src, dst byte) Move {
	return Move{src: src, dst: dst}
}

/*
+---+---+---+
+   | ^ | A |
+---+---+---+
| < | v | > |
+---+---+---+

+---+---+---+
| 7 | 8 | 9 |
+---+---+---+
| 4 | 5 | 6 |
+---+---+---+
| 1 | 2 | 3 |
+---+---+---+
|   | 0 | A |
+---+---+---+
*/

// ExploreOptimalMoves return a map which indicates the optimal move
//
//	path[{src,dst}] = string of inputs needed for the d-pad robot
func ExploreOptimalMoves(numRobots int, output string) {
	// need to get from src to dst
	// can go:
	//    horizontal then vertical (hor-first)
	// or
	//    vertical then horizontal (ver-first)
	// try both
	// then parent can go hor-first or ver-first
	// and so on
	fmt.Println("Exploring options for ", output, "with", numRobots, "robots")
	numOptions := 1 << numRobots
	for option := range numOptions {
		solution := computeSolution(numRobots, option, output)
		fmt.Printf("index: %2d\tlen(solution): %3d\t%s\n", option, len(solution), solution)
	}
}

func computeSolution(numRobots int, option int, code string) string {
	return ""
}

func DpadMovement(src, dst byte) []string {
	options := []string{}
	srcPt, _ := DIRPAD.Find(src)
	dstPt, _ := DIRPAD.Find(dst)
	// can we go vertical first without stepping on the blank tile?
	if DIRPAD.GetChar(utils.NewCoord(dstPt.Row(), srcPt.Col())) != KEY_BLANK {

	}
	// can we go horizontal first without stepping on the blank tile?
	if DIRPAD.GetChar(utils.NewCoord(srcPt.Row(), dstPt.Col())) != KEY_BLANK {

	}
	return options
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
