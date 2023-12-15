package converter

import "strings"

// ConvertString return string in camel case
func ConvertString(text string) string {

	var result string

	for count := 0; count < len(text); {
		if string(text[count]) == "-" || string(text[count]) == "_" {
			result += strings.ToUpper(string(text[count+1]))
			count += 2
		} else {
			result += string(text[count])
			count++
		}
	}

	return result

}
