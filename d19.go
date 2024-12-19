package main

import (
	"strconv"
	"strings"
)

var d19memo = make(map[string]int64)

func d19try(towels []string, design string) int64 {
	options, been := d19memo[design]
	if been {
		return options
	}
	for _, t := range towels {
		t = strings.TrimSpace(t)
		cut, succ := strings.CutPrefix(design, t)
		if !succ {
			continue
		}
		if cut == "" {
			options++
		}
		options += d19try(towels, cut)
	}
	d19memo[design] = options
	return options
}

func (*methods) D19P1(input string) string {
	parts := strings.Split(input, "\n\n")
	towels := strings.Split(parts[0], ",")
	designs := strings.Split(parts[1], "\n")

	var possible int64
	for _, d := range designs {
		if d19try(towels, d) > 0 {
			possible++
		}
	}
	return strconv.FormatInt(possible, 10)
}

func (*methods) D19P2(input string) string {
	parts := strings.Split(input, "\n\n")
	towels := strings.Split(parts[0], ",")
	designs := strings.Split(parts[1], "\n")

	var possible int64
	for _, d := range designs {
		possible += d19try(towels, d)
	}
	return strconv.FormatInt(possible, 10)
}
