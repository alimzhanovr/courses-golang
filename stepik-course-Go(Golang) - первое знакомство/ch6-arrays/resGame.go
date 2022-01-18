package main

func calculateScore(results ...string) int {
	score := 0
	scores := map[string]int{
		"w": 3,
		"l": 0,
		"d": 1,
	}
	for _, res := range results {
		score += scores[res]
	}
	return score
}
