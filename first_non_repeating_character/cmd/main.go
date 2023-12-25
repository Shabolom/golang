package main

import (
	"fmt"
	"timur/find_evenindex"
	character "timur/first_non_repeating_character"
)

func main() {
	fmt.Println(character.FirstNonRepeating("sTreSS"))
	fmt.Println(character.FirstNonRepeating2("sTreSS"))
	fmt.Println(find_evenindex.Find([]int{-15, 5, 11, 17, 19, -17, 20, -6, 17, -17, 19, 16, -15, -6, 20, 17}))
}
