package main

import (
	"fmt"
	"strconv"
	"strings"
)

type d14robot struct {
	g   *d14grid
	pos *coords
	vel *coords
}

type d14grid struct {
	width  int64
	height int64
	cycle  int
	robots []*d14robot
}

func d14parseBounds(bounds string) (width int64, height int64) {
	parts := strings.Split(bounds, ",")
	width, _ = strconv.ParseInt(parts[0], 10, 64)
	height, _ = strconv.ParseInt(parts[1], 10, 64)
	return
}

func d14parseRobot(line string) *d14robot {
	line = strings.TrimPrefix(line, "p=")
	line = strings.ReplaceAll(line, " v=", ",")
	parts := strings.Split(line, ",")
	var nums []int64
	for _, p := range parts {
		num, _ := strconv.ParseInt(p, 10, 64)
		nums = append(nums, num)
	}
	return &d14robot{
		pos: &coords{x: nums[0], y: nums[1]},
		vel: &coords{x: nums[2], y: nums[3]},
	}
}

func d14parseGrid(input string) *d14grid {
	lines := strings.Split(input, "\n")
	w, h := d14parseBounds(lines[0])
	g := &d14grid{
		width:  w,
		height: h,
	}
	for _, line := range lines[1:] {
		r := d14parseRobot(line)
		r.g = g
		g.robots = append(g.robots, r)
	}
	return g
}

func (r *d14robot) walk(cycles int) {
	r.pos.x += r.vel.x * int64(cycles) % r.g.width
	r.pos.y += r.vel.y * int64(cycles) % r.g.height
	if r.pos.x < 0 {
		r.pos.x += r.g.width
	}
	if r.pos.y < 0 {
		r.pos.y += r.g.height
	}
	if r.pos.x >= r.g.width {
		r.pos.x -= r.g.width
	}
	if r.pos.y >= r.g.height {
		r.pos.y -= r.g.height
	}
}

func (g *d14grid) run(cycles int) {
	for _, r := range g.robots {
		r.walk(cycles)
	}
	g.cycle += cycles
}

func (g *d14grid) calculateSafetyFactor() int64 {
	hw := g.width / 2
	hh := g.height / 2
	var q1, q2, q3, q4 int64
	for _, r := range g.robots {
		if r.pos.x < hw && r.pos.y < hh {
			q1++
		}
		if r.pos.x > hw && r.pos.y < hh {
			q2++
		}
		if r.pos.x < hw && r.pos.y > hh {
			q3++
		}
		if r.pos.x > hw && r.pos.y > hh {
			q4++
		}
	}
	return q1 * q2 * q3 * q4
}

func (g *d14grid) print() (out string) {
	var grid [][]bool
	for y := int64(0); y < g.height; y++ {
		var row []bool
		for x := int64(0); x < g.width; x++ {
			row = append(row, false)
		}
		grid = append(grid, row)
	}

	for _, r := range g.robots {
		grid[r.pos.y][r.pos.x] = true
	}

	for _, row := range grid {
		for _, spot := range row {
			if spot {
				out += "*"
			} else {
				out += "."
			}
		}
		out += "\n"
	}

	return
}

func (g *d14grid) getTreeProbability() int {
	var prob int
	m := make(map[int64]bool)

	for _, r := range g.robots {
		if r.pos.y == 42 {
			_, ok := m[r.pos.x]
			if !ok {
				prob++
				m[r.pos.x] = true
			}
		}
	}

	return prob
}

func (*methods) D14P1(input string) string {
	grid := d14parseGrid(input)
	grid.run(100)
	return strconv.FormatInt(grid.calculateSafetyFactor(), 10)
}

// half manual solution...
func (*methods) D14P2(input string) string {
	grid := d14parseGrid(input)
	var maxprob int
	for {
		grid.run(1)
		prob := grid.getTreeProbability()
		if prob >= maxprob {
			maxprob = prob
			fmt.Println(grid.print())
			fmt.Printf("Cycle %d, prob %d\n", grid.cycle, prob)
			fmt.Scanln()
		}
	}
	return ""
}
