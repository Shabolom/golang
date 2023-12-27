package consecutive

// ProductFib ищем сумму чисел фибоначи наиболее приближенную к дано
func ProductFib(prod uint64) (answer [3]uint64) {

	prev := uint64(1)

	for i := uint64(0); i*prev < prod; {

		i = prev + i
		prev = i - prev

		if i*prev >= prod {
			answer[0], answer[1] = prev, i

			if i*prev == prod {
				answer[2] = 1
			}
		}
	}

	return answer
}
