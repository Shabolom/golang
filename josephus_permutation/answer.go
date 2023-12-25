package josephus_permutation

// Josephus заждачка про суицыдных римлян
func Josephus(items []interface{}, k int) []interface{} {

	answer := []interface{}{}
	count := 1
	kill := 1

	for lens := 0; lens < len(items); {
		if count > len(items) {
			count = 1
		}
		if kill%k == 0 {
			answer = append(answer, items[count-1])
			items = append(items[:count-1], items[count:]...)
			count--
		}
		count++
		kill++
	}
	return answer
}
