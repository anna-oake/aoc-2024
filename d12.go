package main

import (
	"sort"
	"strconv"
	"strings"
)

type d12garden struct {
	garden  []string
	width   int64
	height  int64
	memo    map[int64]*d12region
	regions []*d12region
}

type d12region struct {
	garden *d12garden
	plant  byte
	plots  []coords
}

func (g *d12garden) idx(x, y int64) int64 {
	return y*g.width + x
}

func (g *d12garden) isScanned(x, y int64) bool {
	_, done := g.memo[g.idx(x, y)]
	return done
}

func (g *d12garden) getPlot(x, y int64) byte {
	return g.garden[y][x]
}

func (g *d12garden) isWithinBounds(x, y int64) bool {
	return x >= 0 && y >= 0 && x < g.width && y < g.height
}

func (g *d12garden) scan(x, y int64, plant byte, region *d12region) {
	if !g.isWithinBounds(x, y) || g.isScanned(x, y) {
		return
	}
	if g.getPlot(x, y) != plant {
		return
	}
	g.memo[g.idx(x, y)] = region
	region.plots = append(region.plots, coords{x: x, y: y})
	g.scan(x+1, y, plant, region)
	g.scan(x-1, y, plant, region)
	g.scan(x, y-1, plant, region)
	g.scan(x, y+1, plant, region)
}

func (r *d12region) isWithinRegion(x, y int64) bool {
	return r.garden.memo[r.garden.idx(x, y)] == r
}

func (r *d12region) getPerimeter() (perimeter int) {
	for _, p := range r.plots {
		sides := 4
		if r.isWithinRegion(p.x+1, p.y) {
			sides--
		}
		if r.isWithinRegion(p.x-1, p.y) {
			sides--
		}
		if r.isWithinRegion(p.x, p.y+1) {
			sides--
		}
		if r.isWithinRegion(p.x, p.y-1) {
			sides--
		}
		perimeter += sides
	}
	return
}

func (r *d12region) getSides() (sides int) {
	left := make(map[int64][]int64)
	right := make(map[int64][]int64)
	top := make(map[int64][]int64)
	uwu := make(map[int64][]int64)
	for _, p := range r.plots {
		if !r.isWithinRegion(p.x+1, p.y) {
			right[p.x] = append(right[p.x], p.y)
		}
		if !r.isWithinRegion(p.x-1, p.y) {
			left[p.x] = append(left[p.x], p.y)
		}
		if !r.isWithinRegion(p.x, p.y+1) {
			top[p.y] = append(top[p.y], p.x)
		}
		if !r.isWithinRegion(p.x, p.y-1) {
			uwu[p.y] = append(uwu[p.y], p.x)
		}
	}
	maps := []map[int64][]int64{left, right, top, uwu}
	for _, m := range maps {
		for _, s := range m {
			sort.Slice(s, func(i, j int) bool {
				return s[i] < s[j]
			})
			last := int64(-2)
			for _, c := range s {
				if c != last+1 && c != last {
					sides++
				}
				last = c
			}
		}
	}
	return
}

func (r *d12region) calculatePrice(discount bool) int {
	if discount {
		return len(r.plots) * r.getSides()
	}
	return len(r.plots) * r.getPerimeter()
}

func (*methods) D12P1(input string) string {
	lines := strings.Split(input, "\n")

	g := &d12garden{
		garden: lines,
		width:  int64(len(lines[0])),
		height: int64(len(lines)),
		memo:   make(map[int64]*d12region),
	}

	for y, row := range g.garden {
		for x, plot := range row {
			if g.isScanned(int64(x), int64(y)) {
				continue
			}
			region := &d12region{garden: g, plant: byte(plot)}
			g.scan(int64(x), int64(y), byte(plot), region)
			g.regions = append(g.regions, region)
		}
	}

	var price int
	for _, r := range g.regions {
		price += r.calculatePrice(false)
	}
	return strconv.Itoa(price)
}

func (*methods) D12P2(input string) string {
	lines := strings.Split(input, "\n")

	g := &d12garden{
		garden: lines,
		width:  int64(len(lines[0])),
		height: int64(len(lines)),
		memo:   make(map[int64]*d12region),
	}

	for y, row := range g.garden {
		for x, plot := range row {
			if g.isScanned(int64(x), int64(y)) {
				continue
			}
			region := &d12region{garden: g, plant: byte(plot)}
			g.scan(int64(x), int64(y), byte(plot), region)
			g.regions = append(g.regions, region)
		}
	}

	var price int
	for _, r := range g.regions {
		price += r.calculatePrice(true)
	}
	return strconv.Itoa(price)
}
