package words

import (
	"strings"
)

// Wave принимает текст и возвращает массив текстов с заглавными буквами по их айди + 1
func Wave(words string) []string {

	word := ""
	massWords := []string{}

	for i, elem := range words {
		if string(elem) != " " {
			for id, el := range words {
				if i == id {
					word += strings.ToUpper(string(el))
				} else {
					word += string(el)
				}
			}
			massWords = append(massWords, word)
			word = ""
		}
	}

	return massWords
}
