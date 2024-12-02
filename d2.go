package main

import (
	"math"
	"strconv"
	"strings"
)

func d2isReportSafe(levels []string) bool {
	var lastDiff int64
	for i := 1; i < len(levels); i++ {
		last, _ := strconv.ParseInt(levels[i-1], 10, 64)
		curr, _ := strconv.ParseInt(levels[i], 10, 64)
		diff := curr - last
		if math.Abs(float64(diff)) > 3 {
			return false
		}
		if diff == 0 {
			return false
		}
		if (diff < 0 && lastDiff > 0) || (diff > 0 && lastDiff < 0) {
			return false
		}
		lastDiff = diff
	}
	return true
}

func d2isReportSafeDampened(levels []string, idx int) bool {
	// yeah i know this is horrible but i literally woke up 10 minutes ago
	// and my adhd meds havent kicked in yet and i dont wanna think
	// i just want THE STAR
	if idx == len(levels) {
		return false
	}
	var dampened []string
	// fuck yeah i am doing it this way. don't cry
	for i, level := range levels {
		if i == idx {
			continue
		}
		dampened = append(dampened, level)
	}
	if d2isReportSafe(dampened) {
		return true
	}
	return d2isReportSafeDampened(levels, idx+1)
}

func (*methods) D2P1(input string) string {
	reports := strings.Split(input, "\n")

	var safe int
	for _, r := range reports {
		levels := strings.Split(r, " ")
		if d2isReportSafe(levels) {
			safe++
		}
	}

	return strconv.Itoa(safe)
}

func (*methods) D2P2(input string) string {
	reports := strings.Split(input, "\n")

	var safe int
	for _, r := range reports {
		levels := strings.Split(r, " ")
		if d2isReportSafeDampened(levels, -1) {
			safe++
		}
	}

	return strconv.Itoa(safe)
}
