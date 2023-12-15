package evenindex

// Find находим индекс где по левую и по правую сторону равные значения
func Find(mass []int) int {

	for i, _ := range mass {

		leftSum := 0
		rightSum := 0
		count := 0

		for count < i {
			leftSum += mass[count]
			count++
		}

		count = i + 1

		for count < len(mass) {
			rightSum += mass[count]
			count++
		}
		if leftSum == rightSum {
			return i
		}
	}

	return -1
}
