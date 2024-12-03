package main

import (
	"strconv"
	"strings"
)

// yes. fuck regex, fuck token based parsing
// split split split
// i woke up after some silly dreams and the very next minute i grabbed my laptop
func (*methods) D3P1(input string) string {
	parts := strings.Split(input, "mul(")
	var total int64
	for i, p := range parts {
		if i == 0 {
			continue
		}
		parts := strings.Split(p, ")")
		parts = strings.Split(parts[0], ",")
		if len(parts) != 2 {
			continue
		}
		n1, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			continue
		}
		n2, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			continue
		}
		total += n1 * n2
	}
	return strconv.FormatInt(total, 10)
}

// i hope you are enjoying the beauty of golang error handling
func (*methods) D3P2(input string) string {
	parts := strings.Split(input, "don't()")
	var filtered string

	for i, p := range parts {
		if i == 0 {
			filtered += p
			continue
		}
		parts := strings.Split(p, "do()")
		if len(parts) > 1 {
			filtered += strings.Join(parts[1:], "")
		}
	}

	parts = strings.Split(filtered, "mul(")
	var total int64
	for i, p := range parts {
		if i == 0 {
			continue
		}
		parts := strings.Split(p, ")")
		parts = strings.Split(parts[0], ",")
		if len(parts) != 2 {
			continue
		}
		n1, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			continue
		}
		n2, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			continue
		}
		total += n1 * n2
	}
	return strconv.FormatInt(total, 10)
}
