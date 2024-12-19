package main

import (
	"math"
	"strconv"
	"strings"
)

type d17puter struct {
	a, b, c  int64
	program  []int
	ip       int64
	out      []int64
	quine    bool
	quineCnt int
}

func (p *d17puter) load(input []string) {
	for i, l := range input {
		if i == 3 {
			continue
		}
		_, l, _ = strings.Cut(l, ": ")
		if i == 4 {
			nums := strings.Split(l, ",")
			for _, num := range nums {
				n, _ := strconv.Atoi(num)
				p.program = append(p.program, n)
			}
			continue
		}
		n, _ := strconv.ParseInt(l, 10, 64)
		switch i {
		case 0:
			p.a = n
		case 1:
			p.b = n
		case 2:
			p.c = n
		}
	}
	p.quine = true
	p.quineCnt = 0
}

func (p *d17puter) combo(val int64) int64 {
	if val < 4 {
		return val
	}
	switch val {
	case 4:
		return p.a
	case 5:
		return p.b
	case 6:
		return p.c
	}
	panic("invalid program")
}

func (p *d17puter) cycle() bool {
	if p.ip >= int64(len(p.program)) || p.ip+1 >= int64(len(p.program)) {
		if p.quineCnt < len(p.program) {
			p.quine = false
		}
		return false
	}
	op := int64(p.program[p.ip])
	val := int64(p.program[p.ip+1])

	switch op {
	case 0: // adv
		val = p.combo(val)
		p.a = int64(float64(p.a) / math.Pow(2, float64(val)))
	case 1: // bxl
		p.b ^= val
	case 2: // bst
		val = p.combo(val)
		p.b = val % 8
	case 3: // jnz
		if p.a == 0 {
			break
		}
		p.ip = val
		return true
	case 4: // bxc
		p.b ^= p.c
	case 5: // out
		val = p.combo(val)
		p.out = append(p.out, val%8)
		if p.quine && len(p.out) <= len(p.program) && p.out[len(p.out)-1] == int64(p.program[len(p.out)-1]) {
			p.quineCnt++
		} else {
			p.quine = false
		}
	case 6: // bdv
		val = p.combo(val)
		p.b = int64(float64(p.a) / math.Pow(2, float64(val)))
	case 7: // cdv
		val = p.combo(val)
		p.c = int64(float64(p.a) / math.Pow(2, float64(val)))
	}

	p.ip += 2

	return true
}

func (p *d17puter) run(quine bool) bool {
	for {
		running := p.cycle()
		if quine && !p.quine {
			return false
		}
		if !running {
			return true
		}
	}
}

func (p *d17puter) reset() {
	p.a = 0
	p.b = 0
	p.c = 0
	p.ip = 0
	p.quine = true
	p.out = nil
	p.quineCnt = 0
}

func (p *d17puter) printOutput() string {
	var outs []string
	for _, o := range p.out {
		outs = append(outs, strconv.FormatInt(o, 10))
	}
	return strings.Join(outs, ",")
}

func (*methods) D17P1(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) != 5 {
		return "invalid input (unexpected amount of lines)"
	}

	p := &d17puter{}
	p.load(lines)
	p.run(false)

	return p.printOutput()
}

//TODO: D17P2
