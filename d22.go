package main

import (
	"fmt"
	"strconv"
	"strings"
)

func d22evolve(in int64) int64 {
	m := in * 64
	in ^= m
	in %= 16777216
	m = in / 32
	in ^= m
	in %= 16777216
	m = in * 2048
	in ^= m
	in %= 16777216
	return in
}

func d22evolveDeep(in int64, depth int) int64 {
	if depth == 0 {
		return in
	}
	in = d22evolve(in)
	return d22evolveDeep(in, depth-1)
}

func d22evolveDeep2(in int64, depth, price int, diffs []int) (int64, []int) {
	if depth == 0 {
		return in, diffs
	}
	in = d22evolve(in)
	np := int(in % 10)
	diffs = append(diffs, np-price)
	return d22evolveDeep2(in, depth-1, np, diffs)
}

func (*methods) D22P1(input string) string {
	var sum int64
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		b, _ := strconv.ParseInt(l, 10, 64)
		sum += d22evolveDeep(b, 2000)
	}
	return strconv.FormatInt(sum, 10)
}

func (*methods) D22P2(input string) string {
	var buyers []int64
	lines := strings.Split(input, "\n")
	for _, l := range lines {
		b, _ := strconv.ParseInt(l, 10, 64)
		buyers = append(buyers, b)
	}

	allDiffs := make(map[string]bool)
	var diffsToPriceByBuyer []map[string]int

	for _, b := range buyers {
		var diffs []int
		diffsToPrice := make(map[string]int)
		op := int(b % 10)
		_, diffs = d22evolveDeep2(b, 2000, op, diffs)
		for i, d := range diffs {
			op += d
			if i > 2 {
				ds := fmt.Sprintf("%d,%d,%d,%d", diffs[i-3], diffs[i-2], diffs[i-1], d)
				allDiffs[ds] = true
				_, ok := diffsToPrice[ds]
				if !ok {
					diffsToPrice[ds] = op
				}
			}
		}
		diffsToPriceByBuyer = append(diffsToPriceByBuyer, diffsToPrice)
	}

	var maxPrice int
	for diff := range allDiffs {
		var price int
		for _, buyer := range diffsToPriceByBuyer {
			p := buyer[diff]
			price += p
		}
		if price > maxPrice {
			maxPrice = price
		}
	}

	return strconv.Itoa(maxPrice)
}
