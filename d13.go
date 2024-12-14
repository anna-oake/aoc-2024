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

func (m d13machine) calculate() int64 {
	maxA := math.Max(float64(m.prize.x/m.adiff.x), float64(m.prize.y/m.adiff.y))
	maxB := math.Max(float64(m.prize.x/m.bdiff.x), float64(m.prize.y/m.bdiff.y))
	max := int64(maxA)
	diff := m.adiff
	diff2 := m.bdiff
	cost := int64(3)
	cost2 := int64(1)
	if maxB > maxA {
		max = int64(maxB)
		diff = m.bdiff
		diff2 = m.adiff
		cost = 1
		cost2 = 3
	}

	price := int64(math.MaxInt64)
	var i int64
	for i = 0; i <= max; i++ {
		prize := coords{x: m.prize.x - diff.x*i, y: m.prize.y - diff.y*i}
		if (prize.x%diff2.x == 0) && (prize.y%diff2.y == 0) {
			j := prize.x / diff2.x
			if j != prize.y/diff2.y {
				continue
			}
			p := i*cost + j*cost2
			if p < price {
				price = p
			}
		}
	}
	if price == math.MaxInt {
		return 0
	}
	return price
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
		tokens += m.calculate()
	}
	return strconv.FormatInt(tokens, 10)
}

func (*methods) D13P2(input string) string {
	// TODO: maybe solve this
	return ""
}
