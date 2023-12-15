package words

import "strings"

// Revers разворачиваем слова с сохранением их порядка
func Revers(text string) string {

	textSplit := strings.Split(text, " ")
	preAnsw := []string{}

	for count := 0; count < len(textSplit); count++ {
		word := textSplit[count]
		ansWord := []string{}

		for i := 0; i < len(word); i++ {
			last := string(word[len(word)-1-i])
			ansWord = append(ansWord, last)
		}

		preAnsw = append(preAnsw, strings.Join(ansWord, ""))
		ansWord = []string{}

	}
	answer := strings.Join(preAnsw, " ")

	return answer
}

// Revers2 разворачиваем слова с сохранением их порядка
func Revers2(text string) string {

	textSplit := strings.Split(text, " ")
	answer := []string{}

	for i := 0; i < len(textSplit); i++ {
		newWord := ""
		for j := len(textSplit[i]) - 1; j > -1; j-- {
			newWord += string(textSplit[i][j])
		}
		answer = append(answer, newWord)
	}
	return strings.Join(answer, " ")
}
