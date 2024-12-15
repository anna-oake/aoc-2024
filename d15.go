package main

import (
	"fmt"
	"strconv"
	"strings"
)

type d15warehouse struct {
	width  int64
	height int64
	area   [][]byte
	robot  *coords
	moves  []byte
	step   int
	wide   bool
}

func (wh *d15warehouse) getVel() *coords {
	if wh.step == len(wh.moves) {
		return nil
	}
	c := coords{}
	switch wh.moves[wh.step] {
	case '^':
		c.y = -1
	case '>':
		c.x = 1
	case 'v':
		c.y = 1
	case '<':
		c.x = -1
	}
	return &c
}

func (wh *d15warehouse) push(x, y int64, dry bool) bool {
	s := wh.area[y][x]
	if s == '#' {
		return false
	}
	if s == '.' {
		return true
	}
	v := wh.getVel()
	if v == nil {
		return false
	}
	nx := x + v.x
	ny := y + v.y
	if !wh.wide {
		if !wh.push(nx, ny, dry) {
			return false
		}
		if !dry {
			wh.area[y][x] = '.'
			wh.area[ny][nx] = 'O'
		}
		return true
	}
	if v.x != 0 {
		nx2 := nx + v.x
		ny2 := ny + v.y
		if !wh.push(nx2, ny2, dry) {
			return false
		}
		if !dry {
			wh.area[ny2][nx2] = wh.area[ny][nx]
			wh.area[ny][nx] = wh.area[y][x]
			wh.area[y][x] = '.'
		}
		return true
	}
	x2 := x + 1
	y2 := y
	if s == ']' {
		x2 -= 2
	}
	nx2 := x2 + v.x
	ny2 := y2 + v.y
	if !wh.push(nx, ny, true) || !wh.push(nx2, ny2, true) {
		return false
	}
	wh.push(nx, ny, dry)
	wh.push(nx2, ny2, dry)
	if !dry {
		wh.area[ny2][nx2] = wh.area[y2][x2]
		wh.area[ny][nx] = wh.area[y][x]
		wh.area[y2][x2] = '.'
		wh.area[y][x] = '.'
	}
	return true
}

func (wh *d15warehouse) move() bool {
	v := wh.getVel()
	if v == nil {
		return false
	}
	x := wh.robot.x + v.x
	y := wh.robot.y + v.y
	if wh.push(x, y, false) {
		wh.robot.x = x
		wh.robot.y = y
	}
	wh.step++
	return true
}

func (wh *d15warehouse) getSum() (sum int) {
	for y, row := range wh.area {
		for x, s := range row {
			if s == 'O' || s == '[' {
				sum += y*100 + x
			}
		}
	}
	return
}

func (wh *d15warehouse) print() (out string) {
	out += fmt.Sprintf("Next move %s, step %d:\n", string(wh.moves[wh.step]), wh.step)
	for y, row := range wh.area {
		for x, spot := range row {
			if wh.robot.x == int64(x) && wh.robot.y == int64(y) {
				out += string(wh.moves[wh.step])
				continue
			}
			out += string(spot)
		}
		out += "\n"
	}

	return
}

func (*methods) D15P1(input string) string {
	parts := strings.Split(input, "\n\n")
	lines := strings.Split(parts[0], "\n")
	var robot coords
	var area [][]byte
	for y, l := range lines {
		var row []byte
		for x, s := range l {
			if s == '@' {
				robot.x = int64(x)
				robot.y = int64(y)
				row = append(row, '.')
				continue
			}
			row = append(row, byte(s))
		}
		area = append(area, row)
	}
	var moves []byte
	for _, m := range parts[1] {
		moves = append(moves, byte(m))
	}
	wh := &d15warehouse{
		width:  int64(len(area[0])),
		height: int64(len(area)),
		area:   area,
		robot:  &robot,
		moves:  moves,
	}

	for {
		if !wh.move() {
			break
		}
	}
	return strconv.Itoa(wh.getSum())
}

func (*methods) D15P2(input string) string {
	parts := strings.Split(input, "\n\n")
	lines := strings.Split(parts[0], "\n")
	var robot coords
	var area [][]byte
	for _, l := range lines {
		var row []byte
		for _, s := range l {
			if s == '@' {
				robot.x = int64(len(row))
				robot.y = int64(len(area))
				row = append(row, '.')
				row = append(row, '.')
				continue
			}
			if s == 'O' {
				row = append(row, '[')
				row = append(row, ']')
				continue
			}
			row = append(row, byte(s))
			row = append(row, byte(s))
		}
		area = append(area, row)
	}
	var moves []byte
	for _, m := range parts[1] {
		if m != '\n' {
			moves = append(moves, byte(m))
		}
	}
	wh := &d15warehouse{
		width:  int64(len(area[0])),
		height: int64(len(area)),
		area:   area,
		robot:  &robot,
		moves:  moves,
		wide:   true,
	}

	for {
		if !wh.move() {
			break
		}
	}
	return strconv.Itoa(wh.getSum())
}
