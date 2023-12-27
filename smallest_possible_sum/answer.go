package smallest_possible_sum

import (
	"fmt"
	"slices"
)

func Solution(ar []int) int {
	id := 0

	if slices.Max(ar) == slices.Min(ar) {
		return ar[0] * len(ar)
	}

	for slices.Max(ar) > slices.Min(ar) {

		if id > len(ar)-1 {
			id = 0
		}

		for i, el := range ar {

			if el < 1 {
				return 0
			}

			if el > ar[id] {
				ar[i] = el - ar[id]

			}
		}
		id++
	}
	fmt.Println(ar)

	if slices.Max(ar) == slices.Min(ar) {
		return ar[0] * len(ar)
	}

	return 0
}

func Min(dano []int) (min int) {
	min = -1

	for _, el := range dano {
		if min == -1 {
			min = el
		}
		if min > el {
			min = el
		}
	}
	return min
}

func Max(dano []int) (max int) {
	max = -1

	for _, el := range dano {
		if max == -1 {
			max = el
		}
		if max < el {
			max = el
		}
	}
	return max
}

func Solution2(ar []int) int {
	result := ar[len(ar)-1]

	for i := len(ar) - 2; i >= 0; i-- {
		result = gcd(ar[i], result)
	}

	return result * len(ar)
}

func gcd(x int, y int) int {
	for y != 0 {
		x, y = y, x%y
	}

	return x
}
