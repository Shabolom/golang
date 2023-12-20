package playing_with_digits

import "strconv"

func sqr(digit, count uint64) uint64 {
	sum := digit

	for i := uint64(1); i < count; i++ {
		sum *= digit
	}

	return sum
}

// DigPow ищем нечто
func DigPow(n, p int) int {

	sum := 0
	answer := -1

	for i, el := range strconv.Itoa(n) {
		sqrt := sqr(uint64(el-'0'), uint64(p+i))
		sum += int(sqrt)
	}

	if sum < n {
		return answer
	} else {
		for i := 1; i < sum; i++ {
			if n*i > sum {
				return answer
			}
			if n*i == sum {
				answer = i
				return answer
			}
		}
	}
	return answer
}
