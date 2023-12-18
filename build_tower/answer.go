package build

func Tower(nFloors int) []string {

	urov := ""
	elkaMass := []string{}

	for i := 0; i < nFloors+1; {

		for j := 0; j < nFloors*2-1; j++ {

			if j < (nFloors*2-1)-(nFloors*2-1-i) {
				urov += " "

			} else if j > (nFloors*2-1)-(i+1) {
				urov += " "

			} else {
				urov += "*"
			}
		}

		elkaMass = append(elkaMass, urov+"\n")
		urov = ""
		i++
	}

	revElka := []string{}

	for i := 0; i < len(elkaMass); {

		revElka = append(revElka, elkaMass[len(elkaMass)-1-i])
		i++
	}
	return revElka
}
