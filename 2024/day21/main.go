package main

import (
	"container/list"
	"fmt"
	"os"
	"log"
	"bufio"
	"github.com/cgbaker/AOC/2024/utils"
	"strings"
)

type Problem struct {
	codes []string
}

func main() {
	NPAD_GRID = utils.NewCharGrid(4,3)
	NPAD_GRID.Chars = []byte{'7','8','9','4','5','6','1','2','3',0,'0','A'}
	DPAD_GRID = utils.NewCharGrid(2,3)
	DPAD_GRID.Chars = []byte{ 0 ,'^','A','<','v','>'}
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
	NUMROBOTS=3
	sum := 0
	for _, code := range problem.codes {
		fmt.Println("problem:",code)
		solution := solve(code)
		fmt.Println("solution:",solution)
		sum += utils.Atoi(code[:len(code)-1])*len(solution)
	}
	fmt.Printf("part1: %d\n",sum)
}

func part2(problem *Problem) {
	sum := 0
	// NUMROBOTS=26
	NUMROBOTS=3
	dirPadPaths := buildDirPadPaths()
	numPadPaths := buildNumPadPaths()
	for _, code := range problem.codes {
		solution := ""
		// last robot: numpad
		src := KEY_A
		for _, dst := range []byte(code) {
			solution += numPadPaths[T{src,dst}]
			solution += string(KEY_A)
			src = dst
		}
		solution += numPadPaths[T{src,KEY_A}]
		solution += string(KEY_A)
		for range NUMROBOTS-1 {
			src = KEY_A
			newSolution := ""
			for _, dst := range []byte(solution) {
				newSolution += dirPadPaths[T{src,dst}]
				newSolution += string(KEY_A)
				src = dst
			}
			solution = newSolution
		}
		fmt.Println("code:",code,"solution:",solution)
		// intermediate robots: dirpads
		sum += utils.Atoi(code[:len(code)-1])*len(solution)
	}
	fmt.Printf("part2: %d\n",sum)
}

func pathNS(step int) string {
	if step > 0 {
		return strings.Repeat("v",step)
	}
	return strings.Repeat("^",-step)
}

func pathEW(step int) string {
	if step > 0 {
		return strings.Repeat(">",step)
	}
	return strings.Repeat("<",-step)
}

/* 
    +---+---+
    | ^ | A |
+---+---+---+
| < | v | > |
+---+---+---+
*/
func buildDirPadPaths() map[T]string {
	paths := map[T]string{}
	for _, src := range []byte{KEY_N,KEY_W,KEY_S,KEY_E,KEY_A} {
		srcXY, _ := DPAD_GRID.Find(src)
		for _, dst := range []byte{KEY_N,KEY_W,KEY_S,KEY_E,KEY_A} {
			if dst == src {
				continue
			}
			dstXY, _ := DPAD_GRID.Find(dst)
			if src == KEY_W {
				paths[T{src,dst}] = pathEW(dstXY.Col()-srcXY.Col()) + pathNS(dstXY.Row()-srcXY.Row())
			} else if dst == KEY_W {
				paths[T{src,dst}] = pathNS(dstXY.Row()-srcXY.Row()) + pathEW(dstXY.Col()-srcXY.Col())
			} else {
				stepNS := dstXY.Row() - srcXY.Row()
				stepEW := dstXY.Col() - srcXY.Col()
				paths[T{src,dst}] = pathNS(stepNS) + pathEW(stepEW)
			}
		}
	}
	return paths
}

func buildNumPadPaths() map[T]string {
	paths := map[T]string{}
	for _, src := range []byte{KEY_0,KEY_1,KEY_2,KEY_3,KEY_4,KEY_5,KEY_6,KEY_7,KEY_8,KEY_9,KEY_A} {
		srcXY, _ := NPAD_GRID.Find(src)
		for _, dst := range []byte{KEY_0,KEY_1,KEY_2,KEY_3,KEY_4,KEY_5,KEY_6,KEY_7,KEY_8,KEY_9,KEY_A} {
			if dst == src {
				continue
			}
			dstXY, _ := NPAD_GRID.Find(dst)
			if dstXY.Row() == 3 && srcXY.Col() == 0 {
				paths[T{src,dst}] = pathEW(dstXY.Col()-srcXY.Col()) + pathNS(dstXY.Row()-srcXY.Row())
			} else if dstXY.Col() == 0 && srcXY.Row() == 3 {
				paths[T{src,dst}] = pathNS(dstXY.Row()-srcXY.Row()) + pathEW(dstXY.Col()-srcXY.Col())
			} else {
				stepNS := dstXY.Row() - srcXY.Row()
				stepEW := dstXY.Col() - srcXY.Col()
				paths[T{src,dst}] = pathNS(stepNS) + pathEW(stepEW)
			}
		}
	}
	return paths
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
)

type GlobalState struct {
	output string
	robots string
}

func (g GlobalState) Clone() GlobalState {
	return GlobalState{
		output: g.output,
		robots: g.robots,
	}
}

func InitState() GlobalState {
	init := make([]byte,NUMROBOTS)
	for i := range init {
		init[i] = KEY_A
	}
	return GlobalState {
		output: "",
		robots: string(init),
	}
}

func (g GlobalState) advance(choice byte, id int) GlobalState {
	newState := g.Clone()
	switch id {
	case HUMAN:
		return g.advance(choice, id+1)
	case NUMROBOTS-1: // d-pad robot
		if choice == KEY_A {
			newState.output = newState.output + string(g.robots[id])
		} else {
			newState.robots = newState.robots[:id] + string(NumpadMove(newState.robots[id],choice)) + newState.robots[id+1:]
		}
	default:
		if choice == KEY_A {
			return g.advance(g.robots[id], id+1)
		}
		newState.robots = newState.robots[:id] + string(DpadMove(newState.robots[id],choice)) + newState.robots[id+1:]
	}
	return newState
}

var (
	HUMAN int = -1
	NUMROBOTS int = -1
)

func solve(code string) string {
	// need to find the shortest set of commands
	// Human: direct control over dirpad, controlling robot 1
	// Robot1: movement control over dirpad, controlling robot 2
	// Robot2: movement control over dirpad, controlling robot 3
	// Robot3: movement control over keypad at door
	// A solution must result in Robot3 entering the code
	// Robot3 

	root := InitState()
	parent := map[GlobalState]struct{GlobalState;byte}{}
	queue := list.New()
	queue.PushBack(root)
	numChars := 0
	for queue.Len() > 0 {
		cur := queue.Front().Value.(GlobalState)
		queue.Remove(queue.Front())
		// fmt.Printf("visiting %+v\n",cur)
		if len(cur.output) > numChars {
			fmt.Println("got",cur.output)
			numChars = len(cur.output)
		}
		if cur.output == code {
			solution := ""
			for cur != root {
				p := parent[cur]
				solution = string(p.byte) + solution
				cur = p.GlobalState
			}
			return solution
		}
		for _, choice := range getChoices(code, cur, HUMAN) {
			// fmt.Printf("\tchoice: %c\n",rune(choice))
			newState := cur.advance(choice,HUMAN)
			if newState.output != code[0:len(newState.output)] {
				panic(fmt.Sprintf("bad output: '%s'\n",newState.output))
			}
			// fmt.Printf("\tchild %+v\n",newState)
			if _, visited := parent[newState]; !visited {
				parent[newState] = struct{GlobalState;byte}{cur,choice}
				queue.PushBack(newState)
			}
		}
	}
	panic("did not find solution")
}

type T struct {
	src, dst byte
}

/* 
    +---+---+
    | ^ | A |
+---+---+---+
| < | v | > |
+---+---+---+
*/
func DpadDownhill(src, dst byte) []byte {
	// press A
	if src == dst {
		return []byte{KEY_A}
	}
	// single move
	t := T{src,dst}
	switch t {
	case T{KEY_S,KEY_W}, T{KEY_A,KEY_N}, T{KEY_E,KEY_S}:
		return []byte{KEY_W}
	case T{KEY_S,KEY_N}, T{KEY_E,KEY_A}:
		return []byte{KEY_N}
	case T{KEY_N,KEY_S}, T{KEY_A,KEY_E}:
		return []byte{KEY_S}
	case T{KEY_S,KEY_E}, T{KEY_N,KEY_A}:
		return []byte{KEY_E}
	case T{}:
	}
	// double moves
	if src == KEY_E {
		if dst == KEY_N {
			return []byte{KEY_N,KEY_W}
		} else {
			return []byte{KEY_W}
		}
	}
	if src == KEY_S {
		return []byte{KEY_N,KEY_E}
	}
	if src == KEY_N {
		if dst == KEY_E {
			return []byte{KEY_E,KEY_S}
		} else {
			return []byte{KEY_S}
		}
	}
	if src == KEY_A {
		return []byte{KEY_S,KEY_W}
	}
	// KEY_W
	return []byte{KEY_E}
}

/* 
    +---+---+
    | ^ | A |
+---+---+---+
| < | v | > |
+---+---+---+
*/
func DpadMove(src, movement byte) byte {
	switch src {
	case KEY_A:
		switch movement {
		case KEY_S:
			return KEY_E;
		case KEY_W:
			return KEY_N;
		}
	case KEY_N:
		switch movement {
		case KEY_E:
			return KEY_A;
		case KEY_S:
			return KEY_S;
		}
	case KEY_S:
		switch movement {
		case KEY_N:
			return KEY_N;
		case KEY_E:
			return KEY_E;
		case KEY_W:
			return KEY_W;
		}
	case KEY_E:
		switch movement {
		case KEY_N:
			return KEY_A;
		case KEY_W:
			return KEY_S;
		}
	case KEY_W:
		switch movement {
		case KEY_E:
			return KEY_S;
		}
	}
	panic(fmt.Sprintf("bad d-pad movement: %c by %c",rune(src),rune(movement)))
}

/* 
+---+---+---+
| 7 | 8 | 9 |
+---+---+---+
| 4 | 5 | 6 |
+---+---+---+
| 1 | 2 | 3 |
+---+---+---+
    | 0 | A |
    +---+---+
*/    
func NumpadDownhill(src, dst byte) []byte {
	// press A
	if src == dst {
		return []byte{KEY_A}
	}
	// direct singles
	t := T{src,dst}
	switch t {
	case T{KEY_7,KEY_8}, T{KEY_8,KEY_9}, T{KEY_4,KEY_5}, T{KEY_5,KEY_6}, T{KEY_1,KEY_2}, T{KEY_2,KEY_3}, T{KEY_0,KEY_A}:
		return []byte{KEY_E}
	case T{KEY_7,KEY_4}, T{KEY_8,KEY_5}, T{KEY_9,KEY_6}, T{KEY_4,KEY_1}, T{KEY_5,KEY_2}, T{KEY_6,KEY_3}, T{KEY_2,KEY_0}, T{KEY_3,KEY_A}:
		return []byte{KEY_S}
	case T{KEY_8,KEY_7}, T{KEY_9,KEY_8}, T{KEY_5,KEY_4}, T{KEY_6,KEY_5}, T{KEY_2,KEY_1}, T{KEY_3,KEY_2}, T{KEY_0,KEY_A}:
		return []byte{KEY_W}
	case T{KEY_4,KEY_7}, T{KEY_5,KEY_8}, T{KEY_6,KEY_9}, T{KEY_1,KEY_4}, T{KEY_2,KEY_5}, T{KEY_3,KEY_6}, T{KEY_0,KEY_2}, T{KEY_3,KEY_A}:
		return []byte{KEY_N}
	}
	// indirect single
	switch t {
	case T{KEY_7,KEY_1}, T{KEY_8,KEY_2}, T{KEY_8,KEY_0}, T{KEY_9,KEY_3}, T{KEY_9,KEY_A}, T{KEY_5,KEY_0}, T{KEY_6,KEY_A}:
		return []byte{KEY_S}
	case T{KEY_7,KEY_9}, T{KEY_4,KEY_6}, T{KEY_1,KEY_3}:
		return []byte{KEY_E}
	case T{KEY_1,KEY_7}, T{KEY_2,KEY_8}, T{KEY_0,KEY_8}, T{KEY_3,KEY_9}, T{KEY_A,KEY_9}, T{KEY_0,KEY_5}, T{KEY_A,KEY_6}:
		return []byte{KEY_N}
	case T{KEY_9,KEY_7}, T{KEY_6,KEY_4}, T{KEY_3,KEY_1}:
		return []byte{KEY_W}
	}
	// N or W
	if src == KEY_A || (src == KEY_3 && dst != KEY_0) || (src == KEY_6 && (dst == KEY_7 || dst == KEY_8)) ||
		(src == KEY_2 && (dst == KEY_4 || dst == KEY_7)) || (src == KEY_5 && dst == KEY_7) {
		return []byte{KEY_W,KEY_N}
	}
	// S or W
	if src == KEY_3 || src == KEY_9 || src == KEY_6 || (src == KEY_5 && dst == KEY_1) || (src == KEY_8 && (dst == KEY_4 || dst == KEY_1)) {
		return []byte{KEY_S,KEY_W}
	}
	// N or E
	if (src == KEY_2 && (dst == KEY_6 || dst == KEY_9)) || (src == KEY_5 && dst == KEY_9) ||
		(src == KEY_1 && (dst == KEY_5 || dst == KEY_8 || dst == KEY_6 || dst == KEY_9)) || 
		(src == KEY_4 && (dst == KEY_8 || dst == KEY_9)) {
		return []byte{KEY_N,KEY_E}
	}
	// S or E
	if src == KEY_8 || src == KEY_5 || src == KEY_2 || src == KEY_7 || src == KEY_4 {
		return []byte{KEY_S,KEY_E}
	}
	if src == KEY_0 {
		switch dst {
		case KEY_3, KEY_6, KEY_9:
			return []byte{KEY_N,KEY_E}
		case KEY_1, KEY_4, KEY_7:
			return []byte{KEY_N}
		}
	}
	if src == KEY_1 {
		return []byte{KEY_E}
	}
	panic(fmt.Sprintf("missing numpad option: %c -> %c",rune(src),rune(dst)))
}

/* 
+---+---+---+
| 7 | 8 | 9 |
+---+---+---+
| 4 | 5 | 6 |
+---+---+---+
| 1 | 2 | 3 |
+---+---+---+
    | 0 | A |
    +---+---+
*/    
var (
	NPAD_GRID *utils.CharGrid
	DPAD_GRID *utils.CharGrid
)
func NumpadMove(src, movement byte) byte {
	c, ok := NPAD_GRID.Find(src)
	if !ok {
		panic("oops")
	}
	var dir utils.Point
	switch movement {
	case KEY_N:
		dir = utils.DIR_N
	case KEY_W:
		dir = utils.DIR_W
	case KEY_E:
		dir = utils.DIR_E
	case KEY_S:
		dir = utils.DIR_S
	}
	nextLoc := c.Plus(dir)
	next := NPAD_GRID.GetChar(&nextLoc)
	if next == 0 {
		panic(fmt.Sprintf("bad numpad movement: %c by %c (%+v by %+v)",rune(src),rune(movement),c,dir))
	}
	return next
}


func getChoices(code string, state GlobalState, id int) []byte {
	// fmt.Printf("getChoices(%s, %+v, %d)\n",code,state,id)
	nextOutput := code[len(state.output)]
	choices := map[byte]bool{}
	switch id {
	case NUMROBOTS-1: // d-pad robot
		return NumpadDownhill(state.robots[id], nextOutput)
	case HUMAN:
		desirable := getChoices(code, state, id+1)
		for _, out := range desirable {
			choices[out] = true
		}
	default:
		desirable := getChoices(code, state, id+1)
		for _, out := range desirable {
			for _, in := range DpadDownhill(state.robots[id],out) {
				choices[in] = true
			}
		}
	}
	output := []byte{}
	for k := range choices {
		output = append(output,k)
	}
	return output
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
