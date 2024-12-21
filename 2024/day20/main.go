package main

import (
	// "container/list"
	"container/heap"
	"fmt"
	"os"
	"log"
	"math"
	"github.com/cgbaker/AOC/2024/utils"
)

var (
	WALL byte = '#'
	EMPTY byte = '.'
	START byte = 'S'
	END byte = 'E'
	NUM_CHEATS int = 20
)

type Problem = utils.CharGrid

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	problem, _ := utils.ReadCharGrid(file)
	start, ok := problem.Find(START)
	if !ok {
		panic("couldn't find start")
	}
	end, ok := problem.Find(END)
	if !ok {
		panic("couldn't find end")
	}
	fmt.Println("start:",start,"\tend:",end)
	costs := shortestPaths(problem,start,end)
	fmt.Println("distance to end:",costs[end])
	fmt.Println("part2:", part2(problem,costs,100))
}

func part2(problem *Problem, costs map[utils.Point]int, threshhold int) int {
	// for every point on the path, check to see whether it's faster to cheat to that point than drive directly to it
	// if it's sufficiently faster, record it

	// cheats: savings -> count
	// don't technically need this, nice to print it for debugging the sample and sanity checking
	cheats := map[int]int{}

	for index := range len(problem.Chars) {
		if problem.GetCharByIndex(index) == WALL {
			continue
		}
		curPos := problem.Point(index)

		// look for a cheat ENDING at curPos, by checking all valid STARTING points within 20 steps
		for dx := range 41 {
			for dy := range 41 {
				cheatPoint := curPos.Plus(utils.NewPoint(dx-20,dy-20))
				cheatCost := utils.TaxiCab(curPos,cheatPoint)
				cheatTile := problem.GetChar(&cheatPoint)
				if cheatTile == 0 || cheatTile == WALL {
					continue
				}
				if cheatCost <= 20 {
					// instead of going 
					//   start -> curPos with costs[curPos]
					// consider going
					//   start -> cheatPoint with costs[cheatPoint]
					//   then
					//   cheatPoint -> curPos with cheatCost
					savings := costs[curPos] - (cheatCost + costs[cheatPoint])
					if savings >= threshhold {
						cheats[savings]++
					}
				}
			}
		}
	}


	numCheatworthy := 0
	fmt.Println("cheats:",cheats)
	for _, v := range cheats {
		numCheatworthy += v
	}
	return numCheatworthy

}


//  Part 1 stuff that probably doesn't work anymore
//
// type Node struct {
// 	xy utils.Point
// 	cheating int
// 	cheatStart utils.Point
// }
// 
// type Entry struct {
// 	node Node
// 	cost int
// }
// 
// func part1(problem *Problem, start Node, limit int) map[int]int {
// 	solByCheat := map[utils.Point]int{}
// 	visited := map[Node]bool{}
// 	queue := list.New()
// 	queue.PushBack(Entry{start,0})
// 	for queue.Len() > 0 {
// 		cur := queue.Front().Value.(Entry)
// 		queue.Remove(queue.Front())
// 		if visited[cur.node] {
// 			continue
// 		}
// 		visited[cur.node] = true
// 		if cur.cost == limit {
// 			continue
// 		}
// 		for _, neighbor := range getNeighbors(problem, cur.node) {
// 			if problem.GetChar(&neighbor.xy) == END {
// 				// fmt.Println("found solution",neighbor.xy)
// 				if prev, ok := solByCheat[neighbor.cheatStart]; !ok || cur.cost+1 < prev {
// 					solByCheat[neighbor.cheatStart] = cur.cost+1
// 				}
// 			} else {
// 				queue.PushBack(Entry{neighbor,cur.cost+1})
// 			}
// 		}
// 	}
// 	numSolutions := map[int]int{}
// 	for _, v := range solByCheat {
// 		numSolutions[v]++
// 	}
// 	return numSolutions
// 
// 	startNode := Node{xy: start, cheating: NUM_CHEATS, cheatStart: utils.NewPoint(-1,-1)}
// 	cheats := solve(problem, startNode, worst-50) 
// 	sum := 0
// 	for k, v := range cheats {
// 		fmt.Printf("%d: %d\n",k,v)
// 		sum += v
// 	}
// }
// 
// func getNeighbors(problem *Problem, node Node) []Node {
// 	neighbors := []Node{}
// 	for _, n := range problem.NSEWPoints(node.xy) {
// 		if problem.IsBorder(&n) {
// 			continue
// 		}
// 		if nextTile := problem.GetChar(&n); nextTile == EMPTY || nextTile == END {
// 			if node.cheating == NUM_CHEATS {
// 				neighbors = append(neighbors, Node{xy: n, cheating: node.cheating, cheatStart: node.cheatStart})
// 			} else {
// 				neighbors = append(neighbors, Node{xy: n, cheating: node.cheating-1, cheatStart: node.cheatStart})
// 			}
// 		} else if nextTile == WALL && node.cheating > 2 {
// 			neighbors = append(neighbors, Node{xy: n, cheating: node.cheating-1, cheatStart: n})
// 		}
// 	}
// 	return neighbors
// }

func neighbors(problem *Problem, point utils.Point) []utils.Point {
	neighbors := []utils.Point{}
	for _, n := range problem.NSEWPoints(point) {
		if problem.IsBorder(&n) {
			continue
		}
		if nextTile := problem.GetChar(&n); nextTile != WALL {
			neighbors = append(neighbors, n)
		}
	}
	return neighbors
}

func shortestPaths(problem *Problem, start, end utils.Point) map[utils.Point]int {
	costs := map[utils.Point]int{}
	for index, value := range problem.Chars {
		if value != WALL {
			costs[problem.Point(index)] = math.MaxInt
		}
	}
	costs[start] = 0

	pq := make(PriorityQueue,0,len(problem.Chars))
	pq.Push(&Item{point: start, cost: 0})
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		if item.cost > costs[item.point] {
			continue
		}
		costs[item.point] = item.cost
		if item.point == end {
			continue
		}
		for _, n := range neighbors(problem, item.point) {
			heap.Push(&pq, &Item{point: n, cost: item.cost+1})
		}
	}
	return costs
}

// An Item is something we manage in a priority queue.
type Item struct {
	point    utils.Point
	cost int
	index int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].cost > pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, point utils.Point, cost int) {
	item.point = point
	item.cost = cost
	heap.Fix(pq, item.index)
}

