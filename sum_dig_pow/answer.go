package sum

import (
	"strconv"
)

func SqeUp(digit, degree uint64) uint64 {

	back := uint64(1)

	for i := 0; uint64(i) < degree; i++ {

		if i >= 0 {
			back *= digit
		}
	}
	return back
}

// SqerId преобразуем число на его порядок по айди + 1
func SqerId(start, end uint64) []uint64 {

	mass := []uint64{}
	var sums uint64

	for start < end {

		for i, el := range strconv.FormatUint(start, 10) {

			if len(strconv.FormatUint(start, 10)) == 1 {
				mass = append(mass, uint64(el-'0'))
			}

			if len(strconv.FormatUint(start, 10)) > 1 {
				sums += SqeUp(uint64(el-'0'), uint64(i+1))

				if sums == start {
					mass = append(mass, sums)
				}
			}
		}
		sums = 0
		start++
	}
	return mass
}
