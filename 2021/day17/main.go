package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	target := readTarget(file)
	fmt.Println(target.countTrajectories())
}

func readTarget(file *os.File) *Target {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	tgt := &Target{}
	n, err := fmt.Sscanf(scanner.Text(), "target area: x=%d..%d, y=%d..%d",
		&tgt.minX, &tgt.maxX, &tgt.minY, &tgt.maxY)
	if err != nil || n != 4 {
		panic("bad read")
	}
	return tgt
}

type Target struct {
	minX, maxX int
	minY, maxY int
}

func (t *Target) findHighestTrajectory() int {
	maxApex := 0
	for dx := 1; dx <= 1000; dx++ {
		for dy := 0; dy <= 1000; dy++ {
			if apex, hit := t.simulate(dx,dy); hit {
				maxApex = max(maxApex, apex)
			}
		}
	}
	return maxApex
}

func (t *Target) countTrajectories() int {
	count := 0
	for dx := 1; dx <= 1000; dx++ {
		for dy := -1000; dy <= 1000; dy++ {
			if _, hit := t.simulate(dx,dy); hit {
				count++
			}
		}
	}
	return count
}


func (t *Target) simulate(dx int, dy int) (int, bool) {
	x, y, apex := 0, 0, 0
	for !t.overshot(x,y) {
		x, y, dx, dy = step(x, y, dx, dy)
		apex = max(apex, y)
		if t.hit(x, y) {
			return apex, true
		}
	}
	return 0, false
}

func step(x int, y int, dx int, dy int) (int, int, int, int) {
	x += dx
	y += dy
	if dx > 0 {
		dx--
	}
	dy--
	return x, y, dx, dy
}

func (t *Target) overshot(x, y int) bool {
	if x > t.maxX || y < t.minY {
		return true
	}
	return false
}

func (t *Target) hit(x int, y int) bool {
	return t.minX <= x && x <= t.maxX &&
		t.minY <= y && y <= t.maxY
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}