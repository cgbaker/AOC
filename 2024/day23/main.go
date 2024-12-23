package main

import (
	"container/list"
	"fmt"
	"os"
	"log"
	"bufio"
	"slices"
	"strings"
)

type Problem struct {
	graph map[string]map[string]bool
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
	icc := map[string]bool{}
	for a, na := range problem.graph {
		if a[0] != 't' {
			continue
		}
		for b := range na {
			for c := range problem.graph[b] {
				if na[c] {
					cc3 := normalize([]string{a,b,c})
					icc[cc3] = true
				}
			}
		}
	}
	fmt.Printf("part1: %d\n",len(icc))
}

func part2(problem *Problem) {
	var longest []string
	for root := range problem.graph {
		visited := map[string]bool{}
		current := map[string]bool{root:true}
		queue := list.New()
		queue.PushBack(root)
		for queue.Len() > 0 {
			next := queue.Front().Value.(string)
			queue.Remove(queue.Front())
			if visited[next] {
				continue
			}
			visited[next] = true
			if isConnected(problem, next, current) || current[next] {
				current[next] = true
				for b := range problem.graph[next] {
					queue.PushBack(b)
				}
			}
		}
		if len(current) > len(longest) {
			longest = []string{}
			for k := range(current) {
				longest = append(longest,k)
			}
		}
	}
	fmt.Printf("part2: %s\n",normalize(longest))
}

func isConnected(problem *Problem, computer string, network map[string]bool) bool {
	for k := range network {
		if !problem.graph[computer][k] {
			return false
		}
	}
	return true
}

func normalize(parts []string) string {
	slices.Sort(parts)	
	return strings.Join(parts,",")
}

func readInput(file *os.File) *Problem {
	graph := map[string]map[string]bool{}
	lineScanner := bufio.NewScanner(file)
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		parts := strings.Split(lineScanner.Text(),"-")
		a, b := parts[0], parts[1]
		if _, ok := graph[a]; !ok {
			graph[a] = map[string]bool{}
		}
		graph[a][b] = true
		if _, ok := graph[b]; !ok {
			graph[b] = map[string]bool{}
		}
		graph[b][a] = true
	}
	return &Problem{
		graph: graph,
	}
}

func intersect(a, b map[string]bool) map[string]bool {
	c := map[string]bool{}
	for k := range b {
		if a[k] {
			c[k] = true
		}
	}
	return c
}
