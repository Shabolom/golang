package bouncing_balls

// BallUp считаем сколько раз пролетит мяч мимо окна
func BallUp(h, bounce, window float64) int {

	answer := -1
	hightJump := h * bounce

	if h <= 0 || (bounce <= 0 && bounce >= 1) || window >= h {
		return answer

	} else if hightJump > window {
		answer = 3

		for hightJump > window {
			hightJump *= bounce

			if hightJump > window {
				answer += 2
			}
		}

	} else {
		answer = 1
	}
	return answer
}
