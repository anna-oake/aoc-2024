package main

type coords struct {
	x, y int64
}

func (c coords) getIdx(w int64) int64 {
	return c.y*w + c.x
}

func (c coords) move(dir int, steps int64) coords {
	x := c.x
	y := c.y
	switch dir {
	case 0:
		y--
	case 1:
		x++
	case 2:
		y++
	case 3:
		x--
	}
	return coords{
		x: x,
		y: y,
	}
}

func (c coords) inBounds(w, h int64) bool {
	return c.x >= 0 && c.y >= 0 && c.x < w && c.y < h
}

func coordsFromIdx(idx, w int64) coords {
	return coords{
		x: idx % w,
		y: idx / w,
	}
}

type coords32 struct {
	x, y int
}

func (c coords32) getIdx(w int) int {
	return c.y*w + c.x
}

func (c coords32) move(dir int, steps int) coords32 {
	x := c.x
	y := c.y
	switch dir {
	case 0:
		y--
	case 1:
		x++
	case 2:
		y++
	case 3:
		x--
	}
	return coords32{
		x: x,
		y: y,
	}
}

func (c coords32) inBounds(w, h int) bool {
	return c.x >= 0 && c.y >= 0 && c.x < w && c.y < h
}

func coords32FromIdx(idx, w int) coords32 {
	return coords32{
		x: idx % w,
		y: idx / w,
	}
}

// i asked chatgpt to generate the following functions, because i was too cooked and lazy
// let's agree no one wants to write this manually
// maybe there is a better way to implement these. if so, address your comments and suggestions to chatgpt

// combinations generates all combinations of k elements from the input slice.
// It returns a slice of combinations, each combination being a slice of integers.
func combinations(arr []int, k int) [][]int {
	var result [][]int
	var comb func(start int, current []int)

	comb = func(start int, current []int) {
		if len(current) == k {
			// Make a copy of current slice to avoid mutation
			combination := make([]int, k)
			copy(combination, current)
			result = append(result, combination)
			return
		}
		for i := start; i <= len(arr)-k+len(current); i++ {
			comb(i+1, append(current, arr[i]))
		}
	}

	comb(0, []int{})
	return result
}

// permutations generates all permutations of the input slice.
// It returns a slice of permutations, each permutation being a slice of integers.
func permutations(arr []int) [][]int {
	var result [][]int
	var permute func(start int)

	permute = func(start int) {
		if start == len(arr)-1 {
			// Make a copy of arr to avoid mutation
			permutation := make([]int, len(arr))
			copy(permutation, arr)
			result = append(result, permutation)
			return
		}
		for i := start; i < len(arr); i++ {
			arr[start], arr[i] = arr[i], arr[start]
			permute(start + 1)
			arr[start], arr[i] = arr[i], arr[start] // backtrack
		}
	}

	permute(0)
	return result
}

// getSwapConfigurations generates all unique swap configurations between sus1 and sus2.
// Each configuration is a slice of swaps, where each swap is a slice of two integers:
// [sus1 index, sus2 index]. The function ensures that each element is involved
// in at most one swap per configuration.
func getSwapConfigurations(n int) [][][]int {
	var configurations [][][]int

	// All possible indices for sus1 and sus2
	sus1Indices := make([]int, n)
	sus2Indices := make([]int, n)
	for i := 0; i < n; i++ {
		sus1Indices[i] = i
		sus2Indices[i] = i
	}

	// Iterate over all possible number of swaps (0 to n)
	for k := 0; k <= n; k++ {
		// Generate all combinations of k indices from sus1 and sus2
		sus1Comb := combinations(sus1Indices, k)
		sus2Comb := combinations(sus2Indices, k)

		for _, comb1 := range sus1Comb {
			for _, comb2 := range sus2Comb {
				// Generate all permutations of comb2 to pair with comb1
				perms := permutations(comb2)
				for _, perm := range perms {
					// Create a configuration by pairing comb1[i] with perm[i]
					var config [][]int
					for i := 0; i < k; i++ {
						swap := []int{comb1[i], perm[i]}
						config = append(config, swap)
					}
					configurations = append(configurations, config)
				}
			}
		}
	}

	return configurations
}
