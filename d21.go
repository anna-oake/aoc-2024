package main

import (
	"strconv"
	"strings"
)

type d21arm struct {
	w      int
	h      int
	keypad []byte
	start  byte

	shortest    int
	predictions []string
	pcache      map[int]int
	ways        map[string][]string

	ffcache map[string]int64
}

func (a *d21arm) predictPath(pos coords32, dir int, target byte, path string) {
	a.pcache[pos.getIdx(a.w)] = len(path)
	if a.keypad[pos.getIdx(a.w)] == target {
		if a.shortest == 0 || len(path) <= a.shortest {
			a.shortest = len(path)
			a.predictions = append(a.predictions, path)
		}
	}
	for i := 0; i < 4; i++ {
		if i == 2 && path != "" {
			continue
		}
		ndir := (dir + i) % 4
		npos := pos.move(ndir, 1)
		if !npos.inBounds(a.w, a.h) {
			continue
		}
		if a.keypad[npos.getIdx(a.w)] == 0 {
			continue
		}
		c, been := a.pcache[npos.getIdx(a.w)]
		if been && c < len(path)+1 {
			continue
		}
		npath := path
		switch ndir {
		case 0:
			npath += "^"
		case 1:
			npath += ">"
		case 2:
			npath += "v"
		case 3:
			npath += "<"
		}
		a.predictPath(npos, ndir, target, npath)
	}
}

func (a *d21arm) predictBestPaths(start, target byte) []string {
	var spos coords32
	for i, key := range a.keypad {
		if key == start {
			spos = coords32FromIdx(i, a.w)
			break
		}
	}
	a.shortest = 0
	a.predictions = nil
	a.pcache = make(map[int]int)
	a.predictPath(spos, 0, target, "")
	var res []string
	for _, p := range a.predictions {
		if len(p) > a.shortest {
			continue
		}
		res = append(res, p)
	}
	return res
}

func (a *d21arm) warmup() {
	a.ways = make(map[string][]string)
	a.ffcache = make(map[string]int64)
	for i, k1 := range a.keypad {
		for j, k2 := range a.keypad {
			if i == j || k1 == 0 || k2 == 0 {
				continue
			}
			p := a.predictBestPaths(k1, k2)
			a.ways[string([]byte{k1, k2})] = p
		}
	}
}

func (a *d21arm) enterKeysRecurse(keys string, curr byte, press byte, path string) []string {
	if keys == "" {
		return []string{path}
	}
	var paths []string
	ways := []string{""}
	if curr != keys[0] {
		ways = a.ways[string([]byte{curr, keys[0]})]
	}
	for _, w := range ways {
		res := a.enterKeysRecurse(keys[1:], keys[0], press, path+w+string(press))
		paths = append(paths, res...)
	}
	return paths
}

func (a *d21arm) enterKeys(keys string, press byte) []string {
	ways := a.enterKeysRecurse(keys, a.start, press, "")
	var shortest int
	for _, w := range ways {
		if shortest == 0 || len(w) < shortest {
			shortest = len(w)
		}
	}
	var sways []string
	for _, w := range ways {
		if len(w) == shortest {
			sways = append(sways, w)
		}
	}
	return sways
}

func (a *d21arm) enterKeysDeep(start byte, keys string, press byte, depth int) int64 {
	if depth == 0 {
		return int64(len(keys))
	}
	ck := keys + strconv.Itoa(depth)
	res, ok := a.ffcache[ck]
	if ok {
		return res
	}
	var path int64
	for i, k := range keys {
		var k1 byte
		if i == 0 {
			k1 = start
		} else {
			k1 = keys[i-1]
		}
		k2 := byte(k)
		ways := []string{""}
		if k1 != k2 {
			ways = a.ways[string([]byte{k1, k2})]
		}
		var sh int64
		for _, w := range ways {
			res = a.enterKeysDeep(start, w+string(press), press, depth-1)
			if sh == 0 || res < sh {
				sh = res
			}
		}
		path += sh
	}
	a.ffcache[ck] = path
	return path
}

func (*methods) D21P1(input string) string {
	na := &d21arm{
		w:      3,
		h:      4,
		keypad: []byte{'7', '8', '9', '4', '5', '6', '1', '2', '3', 0, '0', 'A'},
		start:  'A',
	}
	na.warmup()

	da := &d21arm{
		w:      3,
		h:      2,
		keypad: []byte{0, '^', 'A', '<', 'v', '>'},
		start:  'A',
	}
	da.warmup()

	lines := strings.Split(input, "\n")
	var total int64
	for _, l := range lines {
		code, _ := strconv.Atoi(strings.TrimSuffix(l, "A"))
		p1s := na.enterKeys(l, 'A')
		var shortest int64
		for _, p := range p1s {
			res := da.enterKeysDeep(da.start, p, 'A', 2)
			if shortest == 0 || res < shortest {
				shortest = res
			}
		}
		total += int64(shortest) * int64(code)
	}

	return strconv.FormatInt(total, 10)
}

func (*methods) D21P2(input string) string {
	na := &d21arm{
		w:      3,
		h:      4,
		keypad: []byte{'7', '8', '9', '4', '5', '6', '1', '2', '3', 0, '0', 'A'},
		start:  'A',
	}
	na.warmup()

	da := &d21arm{
		w:      3,
		h:      2,
		keypad: []byte{0, '^', 'A', '<', 'v', '>'},
		start:  'A',
	}
	da.warmup()

	lines := strings.Split(input, "\n")
	var total int64
	for _, l := range lines {
		code, _ := strconv.Atoi(strings.TrimSuffix(l, "A"))
		p1s := na.enterKeys(l, 'A')
		var shortest int64
		for _, p := range p1s {
			res := da.enterKeysDeep(da.start, p, 'A', 25)
			if shortest == 0 || res < shortest {
				shortest = res
			}
		}
		total += int64(shortest) * int64(code)
	}

	return strconv.FormatInt(total, 10)
}
