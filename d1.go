package main

import (
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func (*methods) D1P1(input string) string {
	lines := strings.Split(input, "\n")
	var left, right []int64
	for _, line := range lines {
		sides := strings.Split(line, "   ")
		if len(sides) != 2 {
			continue
		}
		l, _ := strconv.ParseInt(sides[0], 10, 64)
		r, _ := strconv.ParseInt(sides[1], 10, 64)
		left = append(left, l)
		right = append(right, r)
	}
	sort.Slice(left, func(i, j int) bool { return left[i] < left[j] })
	sort.Slice(right, func(i, j int) bool { return right[i] < right[j] })

	var distance int64
	for i, l := range left {
		r := right[i]
		distance += int64(math.Abs(float64(l - r)))
	}
	return strconv.FormatInt(distance, 10)
}

func (*methods) D1P2(input string) string {
	lines := strings.Split(input, "\n")
	var left, right []int64
	for _, line := range lines {
		sides := strings.Split(line, "   ")
		if len(sides) != 2 {
			continue
		}
		l, _ := strconv.ParseInt(sides[0], 10, 64)
		r, _ := strconv.ParseInt(sides[1], 10, 64)
		left = append(left, l)
		right = append(right, r)
	}

	ld := make([]int64, len(left))
	copy(ld, left)
	sort.Slice(ld, func(i, j int) bool { return ld[i] < ld[j] })
	ld = slices.Compact(ld)
	m := make(map[int64]int)
	for _, l := range ld {
		c := m[l]
		for _, r := range right {
			if l == r {
				c++
			}
		}
		m[l] = c
	}

	var score int64
	for _, l := range left {
		score += l * int64(m[l])
	}
	return strconv.FormatInt(score, 10)
}
