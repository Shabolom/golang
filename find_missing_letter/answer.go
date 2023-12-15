package finder

import (
	"sort"
	"strings"
)

func AsciiLetters() string {
	alphabet := []string{}
	for i := 'A'; i <= 'Z'; i++ {
		alphabet = append(alphabet, string(i))
	}
	for i := 'a'; i <= 'z'; i++ {
		alphabet = append(alphabet, string(i))
	}

	text := strings.Join(alphabet, "")

	return text
}

// FindMissingLetter возвращает недостающую букву по порядку алфовита

func MissingLetter(text string) string {

	alf := AsciiLetters()
	var massText []string
	answer := ""

	for _, el := range text {
		massText = append(massText, string(el))
		sort.Strings(massText)
		text = strings.Join(massText, "")
	}

	slice := alf[strings.Index(alf, string(text[0])):strings.Index(alf, string(text[len(text)-1]+1))]

	for _, el := range slice {
		char := strings.Contains(text, string(el))
		if char == false {
			answer += string(el)
		}
	}
	return answer
}
