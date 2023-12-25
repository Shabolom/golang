package main

import (
	"fmt"
	"timur/josephus_permutation"
)

func main() {

	//dano := []int{1, 2, 3, 4, 5, 6, 7}
	//kill := 3
	//
	//for i, el := range dano {
	//	if el == kill {
	//		dano[i], dano[0] = dano[0], dano[i]
	//	}
	//}
	//
	//fmt.Println(dano)
	//for i := 1; i < len(dano)-1; i++ {
	//	count := i
	//	for sicle := 0; sicle < kill+1; sicle++ {
	//		if count+kill > len(dano)-1 {
	//			for count+kill > len(dano)-1 {
	//				count = (i + kill - len(dano)) + i - kill
	//			}
	//		}
	//		if sicle == kill {
	//			dano[i], dano[count+kill] = dano[count+kill], dano[i]
	//			fmt.Println(dano)
	//		}
	//	}
	//	print("done\n")
	//}
	//fmt.Println(dano)
	items := []interface{}{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(josephus_permutation.Josephus(items, 2))
}
