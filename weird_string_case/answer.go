package string_case

import (
	"strings"
)

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
