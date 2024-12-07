package main

import (
	"fmt"
	"os"
	"log"
	"github.com/cgbaker/AOC/2024/utils"
)

var (
	DIR_N byte = '^'
	DIR_W byte = '<'
	DIR_E byte = '>'
	DIR_S byte = 'v'
	EMPTY byte = '.'
	OUT_OF_BOUNDS byte = 0
	OBSTACLE byte = '#'
)

type Guard struct {
	dir byte
	r, c int
}

func (g *Guard) Row() int {
	return g.r
}

func (g *Guard) Col() int {
	return g.c
}

func (g *Guard) Clone() Guard {
	return Guard{
		c: g.c,
		r: g.r,
		dir: g.dir,
	}
}

func (guard *Guard) turn() Guard {
	next := guard.Clone()
	switch guard.dir {
	case DIR_S:
		next.dir = DIR_W
	case DIR_W:
		next.dir = DIR_N
	case DIR_N:
		next.dir = DIR_E
	case DIR_E:
		next.dir = DIR_S
	}
	return next
}

func (guard *Guard) stepForward() Guard {
	next := guard.Clone()
	switch guard.dir {
	case DIR_S:
		next.r++
	case DIR_N:
		next.r--
	case DIR_W:
		next.c--
	case DIR_E:
		next.c++
	}
	return next
}

type Problem struct {
	grid *utils.CharGrid
	path map[Guard]bool
	guard Guard
}

func (p *Problem) Clone() *Problem {
	newProblem := &Problem{
		guard: p.guard.Clone(),
		grid: p.grid.Clone(),
		path: map[Guard]bool{},
	}
	for k, v := range p.path {
		newProblem.path[k] = v
	}
	return newProblem
}

func (p *Problem) UpdateGuard(g Guard) {
	p.guard = g
	p.path[g] = true
}

func (p *Problem) HasVisited(g Guard) bool {
	return p.path[g];
}


func (p *Problem) HasVisitedAtAll(coord utils.Coord) bool {
	r, c := coord.Row(), coord.Col()
	return p.path[Guard{r: r, c: c, dir: DIR_N}] ||
	       p.path[Guard{r: r, c: c, dir: DIR_S}] ||
	       p.path[Guard{r: r, c: c, dir: DIR_E}] ||
	       p.path[Guard{r: r, c: c, dir: DIR_W}];
}

type WalkVisitor = func (*Problem, Guard)
func nullVisitor(_ *Problem, _ Guard) {}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	problem := &Problem{
		grid: utils.ReadCharGrid(file),
		path: make(map[Guard]bool),
	}
	problem.UpdateGuard(findGuard(problem.grid))
	// clear the guard
	problem.grid.SetChar(&problem.guard, EMPTY)

	prob1(problem)
	prob2(problem)
}

func prob1(problem *Problem) {
	problem = problem.Clone()
	count := 1
	counter := func(p *Problem, next Guard) {
		v := p.grid.GetChar(&next)
		if v == EMPTY && !p.HasVisitedAtAll(&next) {
			count++
		}
	}
	if walk(problem, counter) {
		fmt.Println("ERROR: DETECTED LOOP IN PART ONE")
	}
	fmt.Printf("prob1: %d\n",count)
}

func prob2(problem *Problem) {
	loopMakers := map[utils.Point]bool{}
	search := func(p *Problem, next Guard) {
		v := p.grid.GetChar(&next)		
		// can only put an ostacle at an empty place where we haven't already been or already put an obstacle
		if v == EMPTY && !loopMakers[utils.NewPoint(next.r,next.c)] && !problem.HasVisitedAtAll(&next) {
			newProblem := problem.Clone()
			newProblem.grid.SetChar(&next, OBSTACLE)
			if walk(newProblem, nullVisitor) {
				loopMakers[utils.NewPoint(next.r,next.c)] = true
			} else {
			}
		}
	}
	if walk(problem, search) {
		fmt.Println("ERROR: DETECTED LOOP IN UNALTERED PART TWO")
	}
	fmt.Printf("prob2: %d\n",len(loopMakers))
}

func findGuard(grid *utils.CharGrid) Guard {
	for i, v := range grid.Chars {
		switch v {
		case DIR_S, DIR_N, DIR_W, DIR_E:
			r := i / grid.NumCols
			c := i - r*grid.NumCols
			return Guard{
				dir: v,
				r: r,
				c: c,
			}
		default:
		}
	}
	panic("couldn't find guard")
}

func walk(p *Problem, visit WalkVisitor) bool {
loop:
	for {
		next := p.guard.stepForward()
		visit(p, next)
		v := p.grid.GetChar(&next)
		// detect loop
		if p.HasVisited(next) {
			return true
		}
		switch v {
		case EMPTY:
			p.UpdateGuard(next)
		case OBSTACLE:
			p.UpdateGuard( p.guard.turn() )
		case OUT_OF_BOUNDS:
			break loop
		}
	}
	return false
}
