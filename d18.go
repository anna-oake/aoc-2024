package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type d18grid struct {
	start coords
	end   coords
	area  [][]bool
	score int64
	cache map[int64]int64
}

func (g *d18grid) getNext(now coords, dir int) coords {
	next := now
	switch dir {
	case 0: // N
		next.y--
	case 1: // E
		next.x++
	case 2: // S
		next.y++
	case 3: // W
		next.x--
	}
	return next
}

func (g *d18grid) getIdx(c coords, dir int) int64 {
	return c.y*int64(len(g.area[0])*4) + c.x*int64(4) + int64(dir)
}

func (g *d18grid) walk(now coords, dir int, score int64) {
	g.cache[g.getIdx(now, 0)] = score
	if now.x == g.end.x && now.y == g.end.y {
		if score <= g.score {
			g.score = score
		}
		return
	}
	for i := 0; i < 4; i++ {
		if i == 2 {
			continue
		}
		nd := (dir + i) % 4
		n := g.getNext(now, nd)
		if n.x > g.end.x || n.y > g.end.y || n.x < 0 || n.y < 0 {
			continue
		}
		if g.area[n.y][n.x] {
			continue
		}
		ns := score + 1
		paid, been := g.cache[g.getIdx(n, 0)]
		if been && paid <= ns {
			continue
		}
		g.walk(n, nd, ns)
	}
}

func (*methods) D18P1(input string) string {
	width := 71
	height := 71

	var area [][]bool
	for y := 0; y < height; y++ {
		var row []bool
		for x := 0; x < width; x++ {
			row = append(row, false)
		}
		area = append(area, row)
	}

	lines := strings.Split(input, "\n")
	for i, l := range lines {
		if i == 1024 {
			break
		}
		parts := strings.Split(l, ",")
		x, _ := strconv.ParseInt(parts[0], 10, 64)
		y, _ := strconv.ParseInt(parts[1], 10, 64)
		area[y][x] = true
	}

	g := &d18grid{
		start: coords{x: 0, y: 0},
		end:   coords{x: int64(width) - 1, y: int64(height) - 1},
		area:  area,
		score: int64(math.MaxInt64),
		cache: make(map[int64]int64),
	}
	g.walk(g.start, 1, 0)

	return strconv.FormatInt(g.score, 10)
}

func (*methods) D18P2(input string) string {
	width := 71
	height := 71

	var area [][]bool
	for y := 0; y < height; y++ {
		var row []bool
		for x := 0; x < width; x++ {
			row = append(row, false)
		}
		area = append(area, row)
	}

	lines := strings.Split(input, "\n")
	for i, l := range lines {
		parts := strings.Split(l, ",")
		x, _ := strconv.ParseInt(parts[0], 10, 64)
		y, _ := strconv.ParseInt(parts[1], 10, 64)
		area[y][x] = true
		if i < 1024 {
			continue
		}
		g := &d18grid{
			start: coords{x: 0, y: 0},
			end:   coords{x: int64(width) - 1, y: int64(height) - 1},
			area:  area,
			score: int64(math.MaxInt64),
			cache: make(map[int64]int64),
		}
		g.walk(g.start, 1, 0)
		if g.score == int64(math.MaxInt64) {
			return fmt.Sprintf("%d,%d", x, y)
		}
		fmt.Println(i)
	}

	return ""
}
