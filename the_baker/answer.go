package the_baker

// Cakes печем торты
func Cakes(recipe, available map[string]int) int {

	answer := -1
	prev := 0
	count := len(recipe)

	if len(recipe) > len(available) {
	}

	for keyDan, valueDan := range available {
		for keyRes, valueRes := range recipe {

			if keyDan == keyRes {

				prev = valueDan / valueRes
				count--

				if answer == -1 {
					answer = prev

				} else if answer > prev {
					answer = prev
				}
			}
		}
	}

	if count != 0 {
		answer = 0
		return answer
	}

	return answer
}

func Cakes2(recipe, available map[string]int) int {
	var least int = 1e9
	for material, need := range recipe {
		if available[material]/need < least {
			least = available[material] / need
		}
	}
	return least
}
