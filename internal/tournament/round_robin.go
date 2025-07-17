package tournament

type Pair struct {
	Left  int
	Right int
}

func GenerateSchedule(n int, isDoubleRound ...bool) [][]Pair {
	doubleRound := false
	if len(isDoubleRound) > 0 {
		doubleRound = isDoubleRound[0]
	}

	if n%2 == 1 {
		n++
	}

	var rounds [][]Pair
	if doubleRound {
		rounds = make([][]Pair, (n-1)*2)
	} else {
		rounds = make([][]Pair, n-1)
	}

	round := make([]Pair, n/2)
	secondRound := make([]Pair, n/2)

	for i := 0; i < n/2; i++ {
		round[i] = Pair{
			Left:  i + 1,
			Right: n - i,
		}
		if doubleRound {
			secondRound[i] = Pair{
				Left:  round[i].Right,
				Right: round[i].Left,
			}
		}
	}

	rounds[0] = append([]Pair(nil), round...)
	if doubleRound {
		rounds[n-1] = append([]Pair(nil), secondRound...)
	}

	for j := 1; j < n-1; j++ {
		for i := 0; i < n/2; i++ {
			round[i].Left = calc(round[i].Left, n)
			round[i].Right = calc(round[i].Right, n)

			if doubleRound {
				secondRound[i] = Pair{
					Left:  round[i].Right,
					Right: round[i].Left,
				}
			}
		}
		rounds[j] = append([]Pair(nil), round...)

		if doubleRound {
			rounds[j+n-1] = append([]Pair(nil), secondRound...)
		}
	}

	return rounds
}

func calc(val int, n int) int {
	if val == n {
		return val
	}

	val = val + (n / 2)
	if val > (n - 1) {
		return val - (n - 1)
	}

	return val
}
