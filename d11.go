package main

import (
	"strconv"
	"strings"
)

var d11memo = make(map[string]int)

func d11blink(n, max, i int64) (c int) {
	s := strconv.FormatInt(n, 10)
	s2 := strconv.FormatInt(i, 10)
	// yeah well it's a key, it works
	memo, ok := d11memo[s+"-"+s2]
	if ok {
		return memo
	}
	if i == max {
		return 1
	}
	if n == 0 {
		res := d11blink(1, max, i+1)
		d11memo[s+"-"+s2] = res
		return res
	}
	if len(s)%2 == 0 {
		n1, _ := strconv.ParseInt(s[:len(s)/2], 10, 64)
		n2, _ := strconv.ParseInt(s[len(s)/2:], 10, 64)
		res1 := d11blink(n1, max, i+1)
		res2 := d11blink(n2, max, i+1)
		d11memo[s+"-"+s2] = res1 + res2
		return res1 + res2
	}
	res := d11blink(n*2024, max, i+1)
	d11memo[s+"-"+s2] = res
	return res
}

func (*methods) D11P1(input string) string {
	strs := strings.Split(input, " ")
	var c int
	for _, s := range strs {
		n, _ := strconv.ParseInt(s, 10, 64)
		c += d11blink(n, 25, 0)
	}

	return strconv.Itoa(c)
}

func (*methods) D11P2(input string) string {
	strs := strings.Split(input, " ")
	var c int
	for _, s := range strs {
		n, _ := strconv.ParseInt(s, 10, 64)
		c += d11blink(n, 75, 0)
	}

	return strconv.Itoa(c)
}
