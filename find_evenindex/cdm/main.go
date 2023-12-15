package main

import (
	"fmt"
	evenindex "timur/find_evenindex"
)

func main() {
	fmt.Println(evenindex.Find([]int{-15, 5, 11, 17, 19, -17, 20, -6, 17, -17, 19, 16, -15, -6, 20, 17}))
}
