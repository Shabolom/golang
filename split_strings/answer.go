package split_strings

// KratTwo принимает в себя строку и возвращает массив строки по 2 буквы
func KratTwo(text string) []string {

	step := ""
	mass := []string{}

	if len(text)%2 != 0 {
		text += "_"
	}

	for i, el := range text {
		step += string(el)
		if (i+1)%2 == 0 {
			mass = append(mass, step)
			step = ""
		}
	}
	return mass
}
