package main

import (
	"math"
	"strconv"
	"strings"
)

type d20grid struct {
	end   coords
	area  [][]bool
	score int64
	cache map[int64]int64
	seats map[int64][]coords
	path  map[int64]int64
}

func (g *d20grid) getNext(now coords, dir int) coords {
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

func (g *d20grid) getIdx(c coords, dir int) int64 {
	return c.y*int64(len(g.area[0])*4) + c.x*int64(4) + int64(dir)
}

func (g *d20grid) getIdxObs(c coords, dir, obs int) int64 {
	return c.y*int64(len(g.area[0])*4*20) + c.x*int64(4*20) + int64(dir*20) + int64(obs)
}

func (g *d20grid) walk(now coords, dir int, score int64, path []coords) {
	path = append(path, now)
	g.cache[g.getIdx(now, 0)] = score
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
		if i == 2 && score > 0 {
			continue
		}
		nd := (dir + i) % 4
		n := g.getNext(now, nd)
		if g.area[n.y][n.x] {
			continue
		}
		ns := score + 1
		paid, been := g.cache[g.getIdx(n, 0)]
		if been && paid <= ns {
			continue
		}
		g.walk(n, nd, ns, path)
	}
}

func (g *d20grid) cheat(path []coords, skip int) (int, int64) {
	for i, cur := range path {
		if i == len(path)-1 {
			continue
		}
		for d := 0; d < 4; d++ {
			idx := i*4*2 + d*2 + 0
			if idx > skip {
				could := g.getNext(cur, d)
				score, been := g.path[g.getIdx(could, 0)]
				if been {
					save := score - int64(i) - 1
					if save > 0 {
						return i*4*2 + d*2 + 0, save
					}
				}
			}
			idx++
			if idx <= skip {
				continue
			}
			could := g.getNext(g.getNext(cur, d), d)
			score, been := g.path[g.getIdx(could, 0)]
			if been {
				save := score - int64(i) - 2
				if save > 0 {
					return i*4*2 + d*2 + 1, save
				}
			}
		}
	}
	return -1, -1
}

func (*methods) D20P1(input string) string {
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

	g := &d20grid{
		end:   end,
		area:  area,
		score: int64(math.MaxInt64),
		cache: make(map[int64]int64),
		seats: make(map[int64][]coords),
		path:  make(map[int64]int64),
	}

	g.walk(start, 0, 0, []coords{})

	path := g.seats[g.score]
	for i, c := range path {
		g.path[g.getIdx(c, 0)] = int64(i)
	}

	skip := -1
	var cnt int64
	for {
		pos, save := g.cheat(path, skip)
		if pos < 0 {
			break
		}
		skip = pos
		if save >= 100 {
			cnt++
		}
	}

	return strconv.FormatInt(cnt, 10)
}

func (*methods) D20P2(input string) string {
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

	g := &d20grid{
		end:   end,
		area:  area,
		score: int64(math.MaxInt64),
		cache: make(map[int64]int64),
		seats: make(map[int64][]coords),
	}
	g.walk(start, 0, 0, []coords{})

	pt := g.seats[g.score]

	var total int64

	for i := 0; i < len(pt)-20; i++ {
		for j := i + 20; j < len(pt); j++ {
			start := pt[i]
			end := pt[j]
			// manhattan
			dist := int(math.Abs(float64(end.x-start.x)) + math.Abs(float64(end.y-start.y)))
			if dist <= 20 {
				// as j and i are indexes of the points on the original path,
				// they represent the original score at a point in time
				savings := j - i - dist
				if savings >= 100 {
					total++
				}
			}
		}
	}

	return strconv.FormatInt(total, 10)
}
