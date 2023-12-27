package hamming_numbers

import "fmt"

// Hammer код хэмминга
func Hammer(n int) uint {
	hammingNumbers := []int{1}
	i, j, k := 0, 0, 0

	for len(hammingNumbers) < n {

		fmt.Println(hammingNumbers)
		fmt.Println(Mina(2*hammingNumbers[i], 3*hammingNumbers[j], 5*hammingNumbers[k]), i, j, k)

		nextHamming := Mina(2*hammingNumbers[i], 3*hammingNumbers[j], 5*hammingNumbers[k])
		hammingNumbers = append(hammingNumbers, nextHamming)

		if nextHamming == 2*hammingNumbers[i] {
			i++
		}
		if nextHamming == 3*hammingNumbers[j] {
			j++
		}
		if nextHamming == 5*hammingNumbers[k] {
			k++
		}
	}

	return uint(hammingNumbers[len(hammingNumbers)-1])
}

func Mina(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	} else {
		return c
	}
}
