package main

import (
	"math/bits"
	"sort"
	"strconv"
	"strings"
)

type d24 struct {
	ops   []*d24op
	sus1  []*d24op
	sus2  []*d24op
	input map[string]bool

	buf  map[string]bool
	done map[string]bool

	result []bool
}

type d24op struct {
	a      string
	b      string
	op     string
	target string
	idx    int
	sus    int
}

func (d *d24) process() int64 {
	d.buf = make(map[string]bool)
	d.done = make(map[string]bool)
	d.result = make([]bool, len(d.result))
	for k, v := range d.input {
		d.buf[k] = v
	}
	for {
		if len(d.done) == len(d.ops) {
			return d.getResult()
		}
		acted := false
		for _, op := range d.ops {
			key := op.a + op.b + op.target
			if d.done[key] {
				continue
			}
			a, aok := d.buf[op.a]
			b, bok := d.buf[op.b]
			if !aok || !bok {
				continue
			}
			var val bool
			switch op.op {
			case "AND":
				val = a && b
			case "OR":
				val = a || b
			case "XOR":
				val = (a && !b) || (!a && b)
			}
			d.done[key] = true
			d.buf[op.target] = val
			if op.idx != -1 {
				d.result[op.idx] = val
			}
			acted = true
		}
		if !acted {
			return d.getResult()
		}
	}
}

func (d *d24) getResult() int64 {
	var num string
	for i := len(d.result) - 1; i >= 0; i-- {
		if d.result[i] {
			num += "1"
		} else {
			num += "0"
		}
	}

	res, _ := strconv.ParseInt(num, 2, 64)
	return res
}

func (o1 *d24op) swapWith(o2 *d24op) {
	o1.target, o2.target = o2.target, o1.target
	o1.idx, o2.idx = o2.idx, o1.idx
}

func initd24(input string) *d24 {
	d := &d24{
		input: make(map[string]bool),
	}
	parts := strings.Split(input, "\n\n")

	// parse input
	inp := strings.Split(parts[0], "\n")
	for _, i := range inp {
		kv := strings.Split(i, ": ")
		d.input[kv[0]] = kv[1] == "1"
	}

	// parse operations
	ops := strings.Split(parts[1], "\n")
	for _, op := range ops {
		sides := strings.Split(op, " -> ")
		parts = strings.Split(sides[0], " ")
		tg, _, _ := strings.Cut(sides[1], "#")
		op := parts[1]
		o := d24op{
			a:      parts[0],
			b:      parts[2],
			target: tg,
			op:     op,
			idx:    -1,
		}
		if strings.HasPrefix(tg, "z") {
			d.result = append(d.result, false) // make space in the result array
			resultIdx, _ := strconv.Atoi(strings.TrimPrefix(tg, "z"))
			o.idx = resultIdx
		}
		d.ops = append(d.ops, &o)
	}

	for _, op := range d.ops {
		if op.idx != -1 && op.idx != len(d.result)-1 && op.op != "XOR" {
			op.sus = 1
			d.sus1 = append(d.sus1, op)
		}
		ab := string([]byte{op.a[0], op.b[0]})
		if op.idx == -1 && ab != "xy" && ab != "yx" && op.op == "XOR" {
			op.sus = 2
			d.sus2 = append(d.sus2, op)
		}
	}
	return d
}

func (*methods) D24P1(input string) string {
	d := initd24(input)
	res := d.process()

	return strconv.FormatInt(res, 10)
}

func (*methods) D24P2(input string) string {
	var x, y []byte
	parts := strings.Split(input, "\n\n")
	inp := strings.Split(parts[0], "\n")
	for _, i := range inp {
		parts := strings.Split(i, ": ")
		if strings.HasPrefix(parts[0], "x") {
			x = append(x, parts[1][0])
		} else {
			y = append(y, parts[1][0])
		}
	}

	for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
		x[i], x[j] = x[j], x[i]
		y[i], y[j] = y[j], y[i]
	}

	xn, _ := strconv.ParseInt(string(x), 2, 64)
	yn, _ := strconv.ParseInt(string(y), 2, 64)
	ezn := xn + yn // expected Z number

	d := initd24(input)

	confs := getSwapConfigurations(len(d.sus1))
	var lowestDiff int64
	for _, swaps := range confs {
		for _, swap := range swaps {
			d.sus1[swap[0]].swapWith(d.sus2[swap[1]])
		}
		diff := d.process() ^ ezn
		if lowestDiff == 0 || diff < lowestDiff {
			lowestDiff = diff
		}
		for _, swap := range swaps {
			d.sus1[swap[0]].swapWith(d.sus2[swap[1]])
		}
	}
	faultyIdx := bits.TrailingZeros64(uint64(lowestDiff))
	fs := strconv.Itoa(int(faultyIdx))

	var culprits []*d24op
	for _, op := range d.ops {
		a := strings.TrimLeft(op.a, "xy")
		b := strings.TrimLeft(op.b, "xy")
		if a == fs && b == fs {
			culprits = append(culprits, op)
		}
	}
	culprits = append(culprits, d.sus1...)
	culprits = append(culprits, d.sus2...)

	var outs []string
	for _, c := range culprits {
		outs = append(outs, c.target)
	}
	sort.StringSlice(outs).Sort()

	return strings.Join(outs, ",")
}
