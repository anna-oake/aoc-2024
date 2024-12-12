package main

import (
	"strconv"
	"strings"
)

func (*methods) D6P1(input string) string {
	var area [][]bool
	lines := strings.Split(input, "\n")

	var gx, gy, nx, ny int

	for y, line := range lines {
		var row []bool
		for x, letter := range line {
			obstacle := false
			if letter == '#' {
				obstacle = true
			}
			if letter == '^' {
				gx = x
				gy = y
			}
			row = append(row, obstacle)
		}
		area = append(area, row)
	}

	dir := 0
	visited := make(map[int]bool)
	visited[gy*len(area[0])+gx] = true
	for {
		nx = gx
		ny = gy
		switch dir {
		case 0:
			ny--
		case 1:
			nx++
		case 2:
			ny++
		case 3:
			nx--
		}
		if nx < 0 || ny < 0 || nx >= len(area[0]) || ny >= len(area) {
			return strconv.Itoa(len(visited))
		}
		obstacle := area[ny][nx]
		if obstacle {
			dir++
			if dir > 3 {
				dir = 0
			}
		} else {
			gx = nx
			gy = ny
			visited[gy*len(area[0])+gx] = true
		}
	}

	return ""
}

func d6isLooping(area [][]bool, gx, gy, ox, oy int) bool {
	moves := make(map[int]int)

	var nx, ny, dir int
	for {
		if moves[gy*len(area[0])+gx] > 10 {
			return true
		}

		moves[gy*len(area[0])+gx]++

		nx = gx
		ny = gy
		switch dir {
		case 0:
			ny--
		case 1:
			nx++
		case 2:
			ny++
		case 3:
			nx--
		}
		if nx < 0 || ny < 0 || nx >= len(area[0]) || ny >= len(area) {
			return false
		}

		obstacle := area[ny][nx]
		if obstacle || (ox == nx && oy == ny) {
			dir++
			if dir > 3 {
				dir = 0
			}
		} else {
			gx = nx
			gy = ny
		}
	}
}

func (*methods) D6P2(input string) string {
	var area [][]bool
	lines := strings.Split(input, "\n")

	var ogx, ogy, gx, gy, nx, ny int

	for y, line := range lines {
		var row []bool
		for x, letter := range line {
			obstacle := false
			if letter == '#' {
				obstacle = true
			}
			if letter == '^' {
				gx = x
				gy = y
				ogx = x
				ogy = y
			}
			row = append(row, obstacle)
		}
		area = append(area, row)
	}

	dir := 0
	visited := make(map[int]bool)
	visited[gy*len(area[0])+gx] = true
	for {
		nx = gx
		ny = gy
		switch dir {
		case 0:
			ny--
		case 1:
			nx++
		case 2:
			ny++
		case 3:
			nx--
		}
		if nx < 0 || ny < 0 || nx >= len(area[0]) || ny >= len(area) {
			break
		}
		obstacle := area[ny][nx]
		if obstacle {
			dir++
			if dir > 3 {
				dir = 0
			}
		} else {
			gx = nx
			gy = ny
			visited[gy*len(area[0])+gx] = true
		}
	}

	var total int
	for idx := range visited {
		ox := idx % len(area[0])
		oy := idx / len(area[0])
		if ox == ogx && oy == ogy {
			continue
		}
		if d6isLooping(area, ogx, ogy, ox, oy) {
			total++
		}
	}

	return strconv.Itoa(total)
}
