package perimeter_of_squares

// Perimeter считаем сумму периметров
func Perimeter(n int) int {
	prev := 0
	answer := 4
	count := 0

	for i := 1; count < n; {
		i = prev + i
		prev = i - prev
		answer += i * 4
		count++
	}
	return answer
}
