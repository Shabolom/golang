package printer

import (
	"fmt"
	"strconv"
)

// PrinterErrors смотрим сколько букв посе m и считаем их

func PrintErrors(text string) string {

	var count int
	for _, el := range text {
		if el > 'm' {
			count += 1
		}
	}
	return fmt.Sprintf(strconv.Itoa(count) + "/" + strconv.Itoa(len(text)))
}
