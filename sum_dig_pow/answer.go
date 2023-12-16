package sum

import (
	"math"
	"strconv"
)

// SqerId преобразуем число на его порядок по айди + 1
func SqerId(start, end int) []uint64 {

	sum := float64(0)
	mass := []uint64{}
	var sums float64

	for start < end {

		for i, el := range strconv.Itoa(start) {

			if len(strconv.Itoa(start)) == 1 {
				mass = append(mass, uint64(el-'0'))
			}

			if len(strconv.Itoa(start)) > 1 {
				sums += math.Pow(float64(el-'0'), float64(i+1))

				if int(sums) == start {
					sum = math.Pow(float64(el-'0'), float64(i+1))
					mass = append(mass, uint64(sum))
				}
			}
		}
		sums = 0
		sum = 0
		start++
	}
	return mass
}
