package main

import (
	"strconv"
	"strings"
)

func (a1 coords) findAntinodes(a2 coords) []coords {
	dx := a1.x - a2.x
	dy := a1.y - a2.y
	an1 := coords{x: a1.x + dx, y: a1.y + dy}
	an2 := coords{x: a2.x - dx, y: a2.y - dy}
	return []coords{an1, an2}
}

func (a1 coords) findAllAntinodes(a2 coords, width, height int64) (antinodes []coords) {
	dx := a1.x - a2.x
	dy := a1.y - a2.y
	for {
		if dx%2 != 0 || dy%2 != 0 {
			break
		}
		dx /= 2
		dy /= 2
	}
	antinodes = append(antinodes, a1)
	x := a1.x
	y := a1.y
	for {
		x += dx
		y += dy
		if x < 0 || x >= width || y < 0 || y >= height {
			break
		}
		antinodes = append(antinodes, coords{x: x, y: y})
	}
	x = a1.x
	y = a1.y
	for {
		x -= dx
		y -= dy
		if x < 0 || x >= width || y < 0 || y >= height {
			break
		}
		antinodes = append(antinodes, coords{x: x, y: y})
	}
	return
}

func (*methods) D8P1(input string) string {
	area := make(map[byte][]coords)
	antinodes := make(map[int64]coords)

	lines := strings.Split(input, "\n")
	for y, line := range lines {
		for x, spot := range line {
			if isAlphaNumeric(spot) {
				area[byte(spot)] = append(area[byte(spot)], coords{x: int64(x), y: int64(y)})
			}
		}
	}

	width := int64(len(lines[0]))
	height := int64(len(lines))

	for _, antennas := range area {
		for i, a1 := range antennas {
			for j := i + 1; j < len(antennas); j++ {
				a2 := antennas[j]
				res := a1.findAntinodes(a2)
				for _, c := range res {
					if c.y < 0 || c.y >= height || c.x < 0 || c.x >= width {
						continue
					}
					antinodes[c.y*width+c.x] = c
				}
			}
		}
	}

	return strconv.Itoa(len(antinodes))
}

func (*methods) D8P2(input string) string {
	area := make(map[byte][]coords)
	antinodes := make(map[int64]coords)

	lines := strings.Split(input, "\n")
	for y, line := range lines {
		for x, spot := range line {
			if isAlphaNumeric(spot) {
				area[byte(spot)] = append(area[byte(spot)], coords{x: int64(x), y: int64(y)})
			}
		}
	}

	width := int64(len(lines[0]))
	height := int64(len(lines))

	for _, antennas := range area {
		for i, a1 := range antennas {
			for j := i + 1; j < len(antennas); j++ {
				a2 := antennas[j]
				res := a1.findAllAntinodes(a2, int64(width), int64(height))
				for _, c := range res {
					antinodes[c.y*width+c.x] = c
				}
			}
		}
	}

	return strconv.Itoa(len(antinodes))
}

func isAlphaNumeric(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}
