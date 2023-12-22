package character

import "strings"

func FirstNonRepeating(text string) string {
	//your code here

	text2 := strings.ToLower(text)

	for id, el := range text2 {
		count := 0
		charText := string(text[id])
		for i, char := range text2 {

			if string(el) == string(char) {
				count++
			}
			if count == 1 && i == len(text2)-1 {
				return (charText)
			}
		}
	}
	return ""
}

func FirstNonRepeating2(text string) string {
	for _, el := range text {
		if strings.Count(strings.ToLower(text), strings.ToLower(string(el))) == 1 {
			return string(el)
		}
	}
	return ""
}
