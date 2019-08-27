package randomutils

import (
	"math/rand"
	"time"
)

func RandomWeights(weights []int) (index int) {
	totalWeight := 0
	for _, weight := range weights {
		if weight < 0 {
			panic("weight should be no less than 0")
		}
		totalWeight += weight
	}
	if totalWeight == 0 {
		return -1
	}
	now := time.Now().UnixNano()
	rand.Seed(now)
	randomNum := rand.Intn(totalWeight) + 1
	for i, weight := range weights {
		if randomNum > weight {
			randomNum -= weight
			continue
		}
		index = i
		break
	}
	return
}
