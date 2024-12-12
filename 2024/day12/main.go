package main

import (
	"fmt"
	"os"
	"log"
	"maps"
	"math"
	"github.com/cgbaker/AOC/2024/utils"
)

type Problem = utils.CharGrid

type Perimeter struct {
	// floating point arithmetic is exact for X.0 and X.5
	r, c float64
}

type Plot struct {
	area, perimeter, numSides int
	perimeterMap map[Perimeter]bool
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	problem := utils.ReadCharGrid(file)
	visited := map[int]bool{}
	plots := []Plot{}
	for i, p := range problem.Chars {
		if !visited[i] {
			// fmt.Println("examining plot",i,":",string([]byte{p}))
			plot := solve(problem, visited, i, p)
			// fmt.Println("plot:",plot)
			plot.numSides = computeNumSides(plot.perimeterMap)
			// fmt.Println("numSides:",plot.numSides)
			plots = append(plots, plot)
		}
	}
	fmt.Println("part1:",part1(plots))
	fmt.Println("part2:",part2(plots))
}

func part1(plots []Plot) int {
	cost := 0
	for _, p := range plots {
		cost += p.area*p.perimeter
	}
	return cost
}

func part2(plots []Plot) int {
	cost := 0
	for _, p := range plots {
		cost += p.area*p.numSides
	}
	return cost
}

func solve(problem *Problem, visited map[int]bool, index int, plant byte) Plot {
	// fmt.Println("visiting",index," ",string([]byte{plant}))
	visited[index] = true
	plot := Plot{
		area: 1,
		perimeter: 0,
		perimeterMap: map[Perimeter]bool{},
	}
	p := problem.Point(index)
	neighbors := problem.NSEWPoints(p)
	// fmt.Println("neighbors:",neighbors)
	for _, n := range neighbors {
		// fmt.Println("considering",nIndex,"visited",visited)
		nPlant := problem.GetChar(&n)
		nIndex := problem.Index(n.Row(),n.Col())
		if nPlant != plant {
			plot.perimeter++
			plot.perimeterMap[getPerimeter(p,n)] = true
		} else if !visited[nIndex] {
			neighborPlot := solve(problem, visited, nIndex, nPlant)
			plot.area += neighborPlot.area
			plot.perimeter += neighborPlot.perimeter
			maps.Copy(plot.perimeterMap, neighborPlot.perimeterMap)
		}
	}
	return plot
}

func getPerimeter(in, out utils.Point) Perimeter {
	// four types of perimeters, depending on whether "out" is N, W, E, or W of "in"
	// perimeter is codified as having a non-integer value, 0.25 from the inside
	p := Perimeter{
		r: float64(in.Row()),
		c: float64(in.Col()),
	}
	if in.Row() > out.Row() {
		p.r -= 0.25
	} else if in.Row() < out.Row() {
		p.r += 0.25
	} else if in.Col() < out.Col() {
		p.c += 0.25
	} else {
		p.c -= 0.25
	}
	return p
}

func computeNumSides(pMap map[Perimeter]bool) int {
	numSides := 0
	for len(pMap) != 0 {
		numSides++
		// get an entry
		pEntry := getEntry(pMap)
		// fmt.Println(pMap)
		// fmt.Println("siding",pEntry)
		delete(pMap,pEntry)
		if math.Round(pEntry.r) != pEntry.r {
			// horizontal side
			l := Perimeter{pEntry.r, pEntry.c-1.0}
			for pMap[l] {
				delete(pMap,l)
				l.c = l.c-1.0
			}
			r := Perimeter{pEntry.r, pEntry.c+1.0}
			for pMap[r] {
				delete(pMap,r)
				r.c = r.c+1.0
			}
		} else {
			// vertical side
			u := Perimeter{pEntry.r-1.0, pEntry.c}
			for pMap[u] {
				delete(pMap,u)
				u.r = u.r-1.0
			}
			d := Perimeter{pEntry.r+1.0, pEntry.c}
			for pMap[d] {
				delete(pMap,d)
				d.r = d.r+1.0
			}
		}
	}
	return numSides
}

func getEntry(pMap map[Perimeter]bool) Perimeter {
	for p := range pMap {
		return p
	}
	return Perimeter{-1,-1}
}
