package main

import (
	"fmt"
	"strings"
)

func main() {

	alphabet := []string{}
	for i := 'a'; i <= 'z'; i++ {
		alphabet = append(alphabet, string(i))
	}
	for i := 'A'; i <= 'Z'; i++ {
		alphabet = append(alphabet, string(i))
	}

	text := strings.Join(alphabet, "")
	fmt.Println(text)

}
