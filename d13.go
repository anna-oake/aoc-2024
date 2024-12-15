package main

import (
	"math"
	"strconv"
	"strings"
)

type d13machine struct {
	adiff coords
	bdiff coords
	prize coords
}

func (m d13machine) calculateFull() int64 {
	a := m.calculate(false)
	b := m.calculate(true)
	if a < b {
		return a
	}
	return b
}

func (m d13machine) calculate(b bool) int64 {
	maxA := math.Max(float64(m.prize.x/m.adiff.x), float64(m.prize.y/m.adiff.y))
	maxB := math.Max(float64(m.prize.x/m.bdiff.x), float64(m.prize.y/m.bdiff.y))
	max := int64(maxA)
	diff := m.adiff
	diff2 := m.bdiff
	cost := int64(3)
	cost2 := int64(1)
	if b {
		max = int64(maxB)
		diff = m.bdiff
		diff2 = m.adiff
		cost = 1
		cost2 = 3
	}

	i := max

	prize := coords{x: m.prize.x - diff.x*i, y: m.prize.y - diff.y*i}
	val := diff2.y*prize.x - diff2.x*prize.y
	factor := int64(1)
	if val < 0 {
		factor = -1
	}

	var min int64
	for {
		prize := coords{x: m.prize.x - diff.x*i, y: m.prize.y - diff.y*i}
		val := diff2.y*prize.x - diff2.x*prize.y
		if val == 0 {
			j := prize.x / diff2.x
			p := i*cost + j*cost2
			return p
		}
		if min > max || (max == min && (i == max-1 || i == 1)) {
			return 0
		}
		if val*factor > 0 {
			max = i - 1
			if min == max {
				i = min
				continue
			}
			i -= (i - min) / 2
		}
		if val*factor < 0 {
			min = i + 1
			if min == max {
				i = min
				continue
			}
			i += (max - i) / 2
		}
	}
}

func (*methods) D13P1(input string) string {
	parts := strings.Split(input, "\n\n")
	var machines []d13machine
	for _, part := range parts {
		lines := strings.Split(part, "\n")
		var numbers []int64
		for _, line := range lines {
			_, line, _ = strings.Cut(line, ": ")
			line = strings.ReplaceAll(line, "=", "+")
			line = strings.ReplaceAll(line, "X", "")
			line = strings.ReplaceAll(line, "Y", "")
			line = strings.ReplaceAll(line, "+", "")
			nums := strings.Split(line, ", ")
			for _, s := range nums {
				num, _ := strconv.ParseInt(s, 10, 64)
				numbers = append(numbers, num)
			}
		}
		machines = append(machines, d13machine{
			adiff: coords{x: numbers[0], y: numbers[1]},
			bdiff: coords{x: numbers[2], y: numbers[3]},
			prize: coords{x: numbers[4], y: numbers[5]},
		})
	}

	var tokens int64
	for _, m := range machines {
		tokens += m.calculateFull()
	}
	return strconv.FormatInt(tokens, 10)
}

func (*methods) D13P2(input string) string {
	parts := strings.Split(input, "\n\n")
	var machines []d13machine
	for _, part := range parts {
		lines := strings.Split(part, "\n")
		var numbers []int64
		for _, line := range lines {
			_, line, _ = strings.Cut(line, ": ")
			line = strings.ReplaceAll(line, "=", "+")
			line = strings.ReplaceAll(line, "X", "")
			line = strings.ReplaceAll(line, "Y", "")
			line = strings.ReplaceAll(line, "+", "")
			nums := strings.Split(line, ", ")
			for _, s := range nums {
				num, _ := strconv.ParseInt(s, 10, 64)
				numbers = append(numbers, num)
			}
		}
		machines = append(machines, d13machine{
			adiff: coords{x: numbers[0], y: numbers[1]},
			bdiff: coords{x: numbers[2], y: numbers[3]},
			prize: coords{x: 10000000000000 + numbers[4], y: 10000000000000 + numbers[5]},
		})
	}

	var tokens int64
	for _, m := range machines {
		tokens += m.calculateFull()
	}
	return strconv.FormatInt(tokens, 10)
}
