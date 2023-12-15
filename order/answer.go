package order

import (
	"strings"
	"unicode"
)

// Answer возвращает текст упорядочевая его по цифрам
func Answer(text string) string {

	textSplit := strings.Split(text, " ")
	textSplit2 := make([]string, len(textSplit))

	for i, _ := range textSplit {
		for _, el := range textSplit[i] {
			if unicode.IsDigit(rune(el)) {
				chis := el - '0'
				textSplit2[chis-1] = textSplit[i]
			}
		}
	}
	return strings.Join(textSplit2, " ")
}
