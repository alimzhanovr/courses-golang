package main

import (
	"math/rand"
	"time"
)

func shuffleInt(nums []int) {
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
}
