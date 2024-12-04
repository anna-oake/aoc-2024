package main

import (
	"strconv"
	"strings"
)

var d4needle = []rune{'X', 'M', 'A', 'S'}

func d4search(letters [][]rune, x, y, dir, idx int) bool {
	switch dir {
	case 0:
		y--
	case 1:
		y--
		x++
	case 2:
		x++
	case 3:
		x++
		y++
	case 4:
		y++
	case 5:
		y++
		x--
	case 6:
		x--
	case 7:
		x--
		y--
	}
	if x < 0 || y < 0 || x >= len(letters[0]) || y >= len(letters) {
		return false
	}
	if d4needle[idx] != letters[y][x] {
		return false
	}
	if idx == len(d4needle)-1 {
		return true
	}
	return d4search(letters, x, y, dir, idx+1)
}

func d4check(letters [][]rune, x, y int) (c int) {
	for i := 0; i < 8; i++ {
		if letters[y][x] == d4needle[0] && d4search(letters, x, y, i, 1) {
			c++
		}
	}
	return
}

func d4shift(letters [][]rune, x, y, dir int) rune {
	switch dir {
	case 0:
		y--
		x++
	case 1:
		x++
		y++
	case 2:
		y++
		x--
	case 3:
		x--
		y--
	}
	if x < 0 || y < 0 || x >= len(letters[0]) || y >= len(letters) {
		return 0
	}
	return letters[y][x]
}

func d4checkMas(letters [][]rune, x, y int) bool {
	if d4shift(letters, x, y, 0) == 'M' {
		if d4shift(letters, x, y, 2) != 'S' {
			return false
		}
	} else if d4shift(letters, x, y, 0) == 'S' {
		if d4shift(letters, x, y, 2) != 'M' {
			return false
		}
	} else {
		return false
	}
	if d4shift(letters, x, y, 1) == 'M' {
		if d4shift(letters, x, y, 3) != 'S' {
			return false
		}
	} else if d4shift(letters, x, y, 1) == 'S' {
		if d4shift(letters, x, y, 3) != 'M' {
			return false
		}
	} else {
		return false
	}
	return true
}

func (*methods) D4P1(input string) string {
	var letters [][]rune
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		var row []rune
		for _, letter := range line {
			row = append(row, letter)
		}
		letters = append(letters, row)
	}

	var c int
	for y, r := range letters {
		for x := range r {
			c += d4check(letters, x, y)
		}
	}

	return strconv.Itoa(c)
}

func (*methods) D4P2(input string) string {
	var letters [][]rune
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		var row []rune
		for _, letter := range line {
			row = append(row, letter)
		}
		letters = append(letters, row)
	}

	var c int
	for y, r := range letters {
		if y == 0 {
			continue
		}
		for x, l := range r {
			if l == 'A' && d4checkMas(letters, x, y) {
				c++
			}
		}
	}

	return strconv.Itoa(c)
}
