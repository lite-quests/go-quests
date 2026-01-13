package solutions

func ProcessScores(scores []int) []int {
	var scoresRes []int
	for i := 0; i < len(scores); i++ {
		if scores[i] < 0 || scores[i] > 100 {
			continue
		} else if scores[i] < 40 {
			scoresRes = append(scoresRes, 40)
		} else {
			scoresRes = append(scoresRes, scores[i])
		}
	}

	if len(scoresRes) > 5 {
		for i := 0; i < len(scoresRes); i++ {
			scoresRes[i] += 5
			if scoresRes[i] > 100 {
				scoresRes[i] = 100
			}
		}
	}
	return scoresRes
}
