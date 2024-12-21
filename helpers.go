package main

type coords struct {
	x, y int64
}

func (c coords) getIdx(w int64) int64 {
	return c.y*w + c.x
}

func (c coords) move(dir int, steps int64) coords {
	x := c.x
	y := c.y
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
	return coords{
		x: x,
		y: y,
	}
}

func (c coords) inBounds(w, h int64) bool {
	return c.x >= 0 && c.y >= 0 && c.x < w && c.y < h
}

func coordsFromIdx(idx, w int64) coords {
	return coords{
		x: idx % w,
		y: idx / w,
	}
}

type coords32 struct {
	x, y int
}

func (c coords32) getIdx(w int) int {
	return c.y*w + c.x
}

func (c coords32) move(dir int, steps int) coords32 {
	x := c.x
	y := c.y
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
	return coords32{
		x: x,
		y: y,
	}
}

func (c coords32) inBounds(w, h int) bool {
	return c.x >= 0 && c.y >= 0 && c.x < w && c.y < h
}

func coords32FromIdx(idx, w int) coords32 {
	return coords32{
		x: idx % w,
		y: idx / w,
	}
}
