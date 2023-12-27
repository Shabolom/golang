package main

import (
	"fmt"
	"timur/hamming_numbers"
)

func main() {
	fmt.Println(hamming_numbers.Hammer(12))
	fmt.Println(9 % 3)
}

func Answord(dano int) uint64 {

	minim := 1
	count := 0

	for count < dano {
		x := minim

		if x%2 > 0 && x%3 > 0 && x%5 > 0 {
			x = 0
		} else {
			for x != 0 && x != 1 {

				if x%2 == 0 {
					x = x / 2

				} else if x%3 == 0 {
					x = x / 3

				} else if x%5 == 0 {
					x = x / 5

				} else {
					x = 0
				}
			}
		}

		if x == 1 {
			count++
		}

		minim++
	}

	return uint64(minim - 1)
}

func Sqe(digit, sqr uint64) (answer uint64) {
	sum := uint64(1)
	if sqr == 0 {
		return uint64(1)
	}
	for count := uint64(1); count <= sqr; count++ {
		sum *= digit
	}
	return sum
}
