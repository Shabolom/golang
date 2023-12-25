package main

import (
	"fmt"
	"timur/the_baker"
)

func main() {
	//resept := map[string]int{
	//	"flour": 500,
	//	"sugar": 200,
	//	"eggs":  1,
	//}
	//
	//dano := map[string]int{
	//	"flour": 1200,
	//	"sugar": 1200,
	//	"eggs":  5,
	//	"milk":  200,
	//}
	//answer := 0
	//prev := 0
	//
	//if len(resept) > len(dano) {
	//	fmt.Println(answer)
	//}
	//
	//for keyDan, valueDan := range dano {
	//	for keyRes, valueRes := range resept {
	//
	//		if keyDan == keyRes {
	//			prev = valueDan / valueRes
	//		}
	//
	//		if answer == 0 {
	//			answer = prev
	//		}
	//
	//		if answer > prev {
	//			answer = prev
	//		}
	//	}
	//}
	//fmt.Println(answer)

	fmt.Println(the_baker.Cakes(
		map[string]int{"eggs": 1, "flour": 500, "sugar": 200},
		map[string]int{"eggs": 5, "flour": 1200, "milk": 200, "sugar": 1200}))
}
