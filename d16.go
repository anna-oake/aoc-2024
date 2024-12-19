package main

import (
	"math"
	"strconv"
	"strings"
	"sync"
)

type d16grid struct {
	start coords
	end   coords
	area  [][]bool
	score int64
	cache map[int64]int64
	seats map[int64][]coords
}

func (g *d16grid) getNext(now coords, dir int) coords {
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

func (g *d16grid) getIdx(c coords, dir int) int64 {
	return c.y*int64(len(g.area[0])*4) + c.x*int64(4) + int64(dir)
}

func (g *d16grid) walk(now coords, dir int, score int64, path []coords) {
	path = append(path, now)
	g.cache[g.getIdx(now, dir)] = score
	if now.x == g.end.x && now.y == g.end.y {
		if score <= g.score {
			g.score = score
			seats := g.seats[score]
			seats = append(seats, path...)
			g.seats[score] = seats
		}
		return
	}
	for i := 0; i < 4; i++ {
		if i == 2 {
			continue
		}
		nd := (dir + i) % 4
		n := g.getNext(now, nd)
		if g.area[n.y][n.x] {
			continue
		}
		ns := score + 1
		if nd != dir {
			ns += 1000
		}
		paid, been := g.cache[g.getIdx(n, dir)]
		if been && paid < ns {
			continue
		}
		g.walk(n, nd, ns, path)
	}
}

func (*methods) D16P1(input string) string {
	lines := strings.Split(input, "\n")
	var start, end coords
	var area [][]bool
	for y, l := range lines {
		var row []bool
		for x, s := range l {
			if s == 'S' {
				start.x = int64(x)
				start.y = int64(y)
				row = append(row, false)
				continue
			}
			if s == 'E' {
				end.x = int64(x)
				end.y = int64(y)
				row = append(row, false)
				continue
			}
			row = append(row, s == '#')
		}
		area = append(area, row)
	}

	g := &d16grid{
		start: start,
		end:   end,
		area:  area,
		score: int64(math.MaxInt64),
		m:     &sync.Mutex{},
		cache: make(map[int64]int64),
		seats: make(map[int64][]coords),
	}

	g.walk(start, 1, 0, []coords{})

	return strconv.FormatInt(g.score, 10)
}

func (*methods) D16P2(input string) string {
	lines := strings.Split(input, "\n")
	var start, end coords
	var area [][]bool
	for y, l := range lines {
		var row []bool
		for x, s := range l {
			if s == 'S' {
				start.x = int64(x)
				start.y = int64(y)
				row = append(row, false)
				continue
			}
			if s == 'E' {
				end.x = int64(x)
				end.y = int64(y)
				row = append(row, false)
				continue
			}
			row = append(row, s == '#')
		}
		area = append(area, row)
	}

	g := &d16grid{
		start: start,
		end:   end,
		area:  area,
		score: int64(math.MaxInt64),
		m:     &sync.Mutex{},
		cache: make(map[int64]int64),
		seats: make(map[int64][]coords),
	}

	g.walk(start, 1, 0, []coords{})

	best := make(map[int64]bool)

	for _, s := range g.seats[g.score] {
		best[g.getIdx(s, 0)] = true
	}

	return strconv.FormatInt(int64(len(best)), 10)
}
