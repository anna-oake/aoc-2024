package main

import (
	"math"
	"strconv"
	"strings"
)

func (*methods) D7P1(input string) string {
	lines := strings.Split(input, "\n")
	var eqs [][]int64
	for _, l := range lines {
		l := strings.ReplaceAll(l, ":", "")
		nums := strings.Split(l, " ")
		var members []int64
		for _, num := range nums {
			member, _ := strconv.ParseInt(num, 10, 64)
			members = append(members, member)
		}
		eqs = append(eqs, members)
	}

	var possible []int64
	for _, eq := range eqs {
		c := len(eq) - 2
		max := int64(math.Pow(2, float64(c)))

		var i int64
		for i = 0; i < max; i++ {
			res := eq[1]
			for o := 0; o < c; o++ {
				op := i & (1 << o)
				if op > 0 {
					res *= eq[o+2]
				} else {
					res += eq[o+2]
				}
			}
			if res == eq[0] {
				possible = append(possible, eq[0])
				break
			}
		}
	}

	var sum int64
	for _, p := range possible {
		sum += p
	}

	return strconv.FormatInt(sum, 10)
}

func (*methods) D7P2(input string) string {
	lines := strings.Split(input, "\n")
	var eqs [][]int64
	for _, l := range lines {
		l := strings.ReplaceAll(l, ":", "")
		nums := strings.Split(l, " ")
		var members []int64
		for _, num := range nums {
			member, _ := strconv.ParseInt(num, 10, 64)
			members = append(members, member)
		}
		eqs = append(eqs, members)
	}

	var possible []int64

	for _, eq := range eqs {
		c := len(eq) - 2
		max := int64(math.Pow(3, float64(c)))

		var i int64
		for i = 0; i < max; i++ {
			s := strconv.FormatInt(i, 3)
			res := eq[1]
			for o := 0; o < c; o++ {
				d := c - len(s)
				var op byte
				if o < d {
					op = '0'
				} else {
					op = s[o-d]
				}
				switch op {
				case '0':
					res += eq[o+2]
				case '1':
					res *= eq[o+2]
				case '2':
					res, _ = strconv.ParseInt(strconv.FormatInt(res, 10)+strconv.FormatInt(eq[o+2], 10), 10, 64)
				}
			}
			if res == eq[0] {
				possible = append(possible, eq[0])
				break
			}
		}
	}

	var sum int64
	for _, p := range possible {
		sum += p
	}

	return strconv.FormatInt(sum, 10)
}
