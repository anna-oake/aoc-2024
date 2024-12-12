package main

import (
	"sort"
	"strconv"
	"strings"
)

type d12garden struct {
	garden  []string
	width   int
	height  int
	memo    map[int]*d12region
	regions []*d12region
}

type d12region struct {
	garden *d12garden
	plant  byte
	plots  []coords
}

func (g *d12garden) idx(x, y int) int {
	return y*g.width + x
}

func (g *d12garden) isScanned(x, y int) bool {
	_, done := g.memo[g.idx(x, y)]
	return done
}

func (g *d12garden) getPlot(x, y int) byte {
	return g.garden[y][x]
}

func (g *d12garden) isWithinBounds(x, y int) bool {
	return x >= 0 && y >= 0 && x < g.width && y < g.height
}

func (g *d12garden) scan(x, y int, plant byte, region *d12region) {
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

func (r *d12region) isWithinRegion(x, y int) bool {
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
	left := make(map[int][]int)
	right := make(map[int][]int)
	top := make(map[int][]int)
	uwu := make(map[int][]int)
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
	maps := []map[int][]int{left, right, top, uwu}
	for _, m := range maps {
		for _, s := range m {
			sort.IntSlice(s).Sort()
			last := -2
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
		width:  len(lines[0]),
		height: len(lines),
		memo:   make(map[int]*d12region),
	}

	for y, row := range g.garden {
		for x, plot := range row {
			if g.isScanned(x, y) {
				continue
			}
			region := &d12region{garden: g, plant: byte(plot)}
			g.scan(x, y, byte(plot), region)
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
		width:  len(lines[0]),
		height: len(lines),
		memo:   make(map[int]*d12region),
	}

	for y, row := range g.garden {
		for x, plot := range row {
			if g.isScanned(x, y) {
				continue
			}
			region := &d12region{garden: g, plant: byte(plot)}
			g.scan(x, y, byte(plot), region)
			g.regions = append(g.regions, region)
		}
	}

	var price int
	for _, r := range g.regions {
		price += r.calculatePrice(true)
	}
	return strconv.Itoa(price)
}
