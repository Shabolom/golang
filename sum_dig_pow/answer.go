package sum

import (
	"fmt"
	"math"
	"strconv"
)

// SqerId преобразуем число на его порядок по айди + 1
func SqerId(start, end uint64) []uint64 {

	mass := []uint64{}
	var sums float64

	for start < end {

		for i, el := range strconv.FormatUint(start, 10) {

			if len(strconv.FormatUint(start, 10)) == 1 {
				mass = append(mass, uint64(el-'0'))
			}

			if len(strconv.FormatUint(start, 10)) > 1 {
				sums += math.Pow(float64(el-'0'), float64(i+1))
				if uint64(sums) == start {
					fmt.Println(start, sums)
					mass = append(mass, uint64(sums))
				}
			}
		}
		sums = 0
		start++
	}
	return mass
}
