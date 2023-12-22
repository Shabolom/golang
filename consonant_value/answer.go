package consonant_value

import "strings"

// Solve считаем сумму по порядку глассных букв разделяя по гласным
func Solve(str string) int {

	summ := 0
	answer := 0
	del := "aeiou"

	for _, el := range str {

		if strings.Contains(del, string(el)) {
			summ = 0

		} else {
			summ += int(el - 'a' + 1)
			if summ > answer {
				answer = summ
			}
		}
	}
	return answer
}
