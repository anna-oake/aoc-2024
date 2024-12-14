package main

import (
	"strconv"
	"strings"
)

var d10memo = make(map[int64][]coords)

func d10idx(area [][]int, x, y, dir int64) int64 {
	width := int64(len(area[0]))
	return y*width*4 + x*4 + dir
}

func d10hike(area [][]int, x, y, dir int64) []coords {
	idx := d10idx(area, x, y, dir)
	memo, ok := d10memo[idx]
	if ok {
		return memo
	}
	n1 := area[y][x]
	switch dir {
	case 0:
		y--
	case 1:
		x++
	case 2:
		y++
	case 3:
		x--
	}
	if x < 0 || y < 0 || x >= int64(len(area[0])) || y >= int64(len(area)) {
		d10memo[idx] = nil
		return nil
	}
	n2 := area[y][x]
	if n2-n1 != 1 {
		d10memo[idx] = nil
		return nil
	}
	if n2 == 9 {
		res := []coords{{x: x, y: y}}
		d10memo[idx] = res
		return res
	}
	var res []coords
	for i := 0; i < 4; i++ {
		res = append(res, d10hike(area, x, y, int64(i))...)
	}
	d10memo[idx] = res
	return res
}

func (*methods) D10P1(input string) string {
	var area [][]int
	var starts []coords

	lines := strings.Split(input, "\n")

	for y, line := range lines {
		var row []int
		for x, spot := range line {
			num, _ := strconv.Atoi(string(spot))
			row = append(row, num)
			if num == 0 {
				starts = append(starts, coords{x: int64(x), y: int64(y)})
			}
		}
		area = append(area, row)
	}

	var c int
	for _, start := range starts {
		m := make(map[int64]bool)
		for dir := 0; dir < 4; dir++ {
			res := d10hike(area, start.x, start.y, int64(dir))
			for _, end := range res {
				m[d10idx(area, end.x, end.y, 0)] = true
			}
		}
		c += len(m)
	}
	return strconv.Itoa(c)
}

func (*methods) D10P2(input string) string {
	var area [][]int
	var starts []coords

	lines := strings.Split(input, "\n")

	for y, line := range lines {
		var row []int
		for x, spot := range line {
			num, _ := strconv.Atoi(string(spot))
			row = append(row, num)
			if num == 0 {
				starts = append(starts, coords{x: int64(x), y: int64(y)})
			}
		}
		area = append(area, row)
	}

	var c int
	for _, start := range starts {
		for dir := 0; dir < 4; dir++ {
			c += len(d10hike(area, start.x, start.y, int64(dir)))
		}
	}
	return strconv.Itoa(c)
}
