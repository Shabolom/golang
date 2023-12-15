package find_missing_letter

import (
	"strings"
)

func AsciiLetters() string {
	alphabet := []string{}
	for i := 'a'; i <= 'z'; i++ {
		alphabet = append(alphabet, string(i))
	}
	for i := 'A'; i <= 'Z'; i++ {
		alphabet = append(alphabet, string(i))
	}

	text := strings.Join(alphabet, "")

	return text
}

// findMissingLetter возвращает недостающую букву по порядку алфовита
func findMissingLetter(text []string) string {
	alf := AsciiLetters()
	slis := alf[strings.Index(alf, text)]

	for _, el := range text {

	}

}
