package main

import (
	"fmt"
)

var maze = [][]int{
	{0, 1, 0, 0, 0},
	{0, 0, 0, 1, 0},
	{0, 1, 0, 1, 0},
	{1, 1, 1, 0, 0},
	{0, 1, 0, 0, 1},
	{0, 1, 0, 0, 0},
}

type point struct {
	i int
	j int
}

var ds = [4]point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func (p point) move(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

func walk(start, end point) [][]int {
	steps := make([][]int, len(maze))
	for l, _ := range steps {
		steps[l] = make([]int, len(maze[l]))
	}

	Q := []point{start}
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]
		if cur.i == end.i && cur.j == end.j {
			break
		}
		for _, d := range ds {
			next := cur.move(d)
			v, ok := next.at(maze)
			if v == 1 || !ok {
				continue
			}
			if next == start {
				continue
			}

			v, ok = next.at(steps)
			if !ok || v != 0 {
				continue
			}
			v, _ = cur.at(steps)
			steps[next.i][next.j] = v + 1
			Q = append(Q, next)
			fmt.Println(next.i, next.j, steps[next.i][next.j])
		}
	}
	return steps
}

func main() {
	fmt.Println("--------------map-------------------------------")
	for _, v := range maze {
		for _, n := range v {
			fmt.Printf("%3d", n)
		}
		fmt.Println()
	}
	steps := walk(point{0, 0}, point{5, 4})
	fmt.Println("--------------after map-------------------------")
	for _, v := range steps {
		for _, n := range v {
			fmt.Printf("%3d", n)
		}
		fmt.Println()
	}
	return
}
