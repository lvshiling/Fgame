package mathutils

import (
	"fmt"
	"math/rand"
)

//随机
func RandomHit(max int, num int) bool {
	// rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(max)
	return randomNum < num
}

//随机
func RandomHits(max int, num int, times int32) (randomFlags []bool) {
	// rand.Seed(time.Now().UnixNano())
	for i := 0; i < int(times); i++ {
		randomNum := rand.Intn(max)
		randomFlags = append(randomFlags, randomNum < num)
	}
	return randomFlags
}

func RandomOneHit(num float64) bool {
	if num >= 1 {
		return true
	}
	// rand.Seed(time.Now().UnixNano())
	randomNum := rand.Float64()
	return randomNum < num
}

//区间随机
func RandomRange(min int, max int) int {
	if min > max {
		panic(fmt.Errorf("min[%d]应该不超过max[%d]", min, max))
	}
	diff := max - min
	if diff == 0 {
		return min
	}
	return rand.Intn(diff) + int(min)
}

//随机从权重
func RandomWeights(weights []int64) (index int) {
	totalWeight := int64(0)
	for _, weight := range weights {
		if weight < 0 {
			panic("weight should be no less than 0")
		}
		totalWeight += weight
	}
	if totalWeight == 0 {
		return -1
	}
	// now := time.Now().UnixNano()
	// rand.Seed(now)
	randomNum := rand.Int63n(totalWeight) + 1
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

//随机从权重
func RandomWeightsFromTotalWeights(weights []int64, totalWeight int64) (index int) {
	if totalWeight == 0 {
		return -1
	}
	index = -1
	// now := time.Now().UnixNano()
	// rand.Seed(now)
	randomNum := rand.Int63n(totalWeight) + 1
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

//随机从权重
func RandomListFromWeights(weights []int64, num int32) (indexList []int) {
	if num <= 0 {
		panic("num should be no less than 0")
	}
	for i := 0; i < int(num); i++ {
		index := RandomWeights(weights)
		indexList = append(indexList, index)
		weights[index] = 0
	}
	return
}

func RandomList(nums []int32, num int32) (hits []int32) {
	if num <= 0 {
		panic(fmt.Errorf("min[%d]应该不应该小于等于0", num))
	}
	len := int32(len(nums))
	if len <= num {
		return nums
	}
	var weights []int64
	for i := int32(0); i < len; i++ {
		weights = append(weights, 1)
	}

	indexList := RandomListFromWeights(weights, num)

	for _, index := range indexList {
		hits = append(hits, nums[index])
	}
	return
}
