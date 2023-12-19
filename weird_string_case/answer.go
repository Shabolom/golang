package string_case

import (
	"strings"
	"unicode"
)

// ToWeirdCase преобразуем каждую не четную букву в заглавную а чет в обычную
func ToWeirdCase(text string) string {

	textSplit := strings.Split(text, " ")
	textAnswer := []string{}
	word := ""

	for i := 0; i < len(textSplit); i++ {
		for id, el := range textSplit[i] {

			if (id+1)%2 == 1 {
				word += strings.ToUpper(string(el))
			} else {
				word += strings.ToLower(string(el))
			}
		}
		textAnswer = append(textAnswer, word)
		word = ""
	}
	return strings.Join(textAnswer, " ")
}

func ToWeirdCase2(text string) string {
	answer := ""
	count := 1

	for _, el := range text {

		if unicode.IsLetter(el) == false {
			count = 1
			answer += string(el)

		} else if count%2 == 1 {
			answer += strings.ToUpper(string(el))
			count++
		} else {
			answer += strings.ToLower(string(el))
			count++
		}
	}
	return answer
}
