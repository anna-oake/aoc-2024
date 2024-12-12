package main

import (
	"strconv"
)

func (*methods) D9P1(input string) string {
	var drive []int
	for _, c := range input {
		f, err := strconv.Atoi(string(c))
		if err != nil {
			break
		}
		drive = append(drive, f)
	}

	var sum int64
	var li, ri int // lef and right indices
	var lf, rf int // left and right files
	var id int     // file ID
	var pos int    // current block position
	var tomove int // how much of rf is left to move to lf
	var free int   // how much free space is left
	ri = len(drive) - 1
	for {
		if li == ri {
			sum += chksum(id, pos, tomove)
			break
		}
		lf = drive[li]
		if li%2 == 0 {
			// holy moly that's a FILE
			id = li / 2
			sum += chksum(id, pos, lf)
			pos += lf
			li++
			continue
		}
		// Space space wanna go to space yes please space. Space space. Go to space.
		rf = drive[ri]
		id = ri / 2
		if tomove == 0 {
			tomove = rf
		}
		if free == 0 {
			free = lf
		}
		if tomove <= free {
			free -= tomove
			sum += chksum(id, pos, tomove)
			pos += tomove
			tomove = 0
			ri -= 2
		} else {
			tomove -= free
			sum += chksum(id, pos, free)
			pos += free
			free = 0
		}
		if free == 0 {
			li++
		}
	}

	return strconv.FormatInt(sum, 10)
}

type d9entity struct {
	id   int
	pos  int
	size int
}

func (*methods) D9P2(input string) string {
	var files []*d9entity
	var space []*d9entity
	var pos int
	for i, c := range input {
		f, err := strconv.Atoi(string(c))
		if err != nil {
			break
		}
		e := d9entity{
			pos:  pos,
			size: f,
		}
		if i%2 == 0 {
			e.id = i / 2
			files = append(files, &e)
		} else {
			space = append(space, &e)
		}
		pos += f
	}

	for i := len(files) - 1; i >= 0; i-- {
		f := files[i]
		for j := 0; j < len(space); j++ {
			s := space[j]
			if s.pos > f.pos {
				break
			}
			if s.size >= f.size {
				s.size -= f.size
				f.pos = s.pos
				s.pos += f.size
				break
			}
		}
	}

	var sum int64
	for _, f := range files {
		sum += chksum(f.id, f.pos, f.size)
	}

	return strconv.FormatInt(sum, 10)
}

func chksum(id, spos, size int) (sum int64) {
	for pos := spos; pos < spos+size; pos++ {
		sum += int64(pos) * int64(id)
	}
	return
}
