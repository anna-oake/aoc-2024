package main

import (
	"sort"
	"strconv"
	"strings"
)

func (*methods) D5P1(input string) string {
	parts := strings.Split(input, "\n\n")

	rlines := strings.Split(parts[0], "\n")
	rules := make(map[string]bool)
	for _, rl := range rlines {
		rules[rl] = true
	}

	updates := strings.Split(parts[1], "\n")
	var sum int
	for _, u := range updates {
		upd := strings.Split(u, ",")
		sort.Slice(upd, func(i, j int) bool {
			a := upd[i]
			b := upd[j]
			_, ok := rules[a+"|"+b]
			if ok {
				return true
			}
			_, ok = rules[b+"|"+a]
			if ok {
				return false
			}
			return false
		})
		u2 := strings.Join(upd, ",")
		if u != u2 {
			continue
		}
		n, _ := strconv.Atoi(upd[len(upd)/2])
		sum += n
	}
	return strconv.Itoa(sum)
}

func (*methods) D5P2(input string) string {
	parts := strings.Split(input, "\n\n")

	rlines := strings.Split(parts[0], "\n")
	rules := make(map[string]bool)
	for _, rl := range rlines {
		rules[rl] = true
	}

	updates := strings.Split(parts[1], "\n")
	var sum int
	for _, u := range updates {
		upd := strings.Split(u, ",")
		sort.Slice(upd, func(i, j int) bool {
			a := upd[i]
			b := upd[j]
			_, ok := rules[a+"|"+b]
			if ok {
				return true
			}
			_, ok = rules[b+"|"+a]
			if ok {
				return false
			}
			return false
		})
		u2 := strings.Join(upd, ",")
		if u == u2 {
			continue
		}
		n, _ := strconv.Atoi(upd[len(upd)/2])
		sum += n
	}
	return strconv.Itoa(sum)
}
