package main

func minSubArrayLen(target int, nums []int) int {
	minLen := 99999999999
	currentSum := 0
	l, r := 0, 0
	for r < len(nums) || currentSum >= target {
		if l < r && currentSum >= target {
			currentSum -= nums[l]
			l++
		} else if r < len(nums) && currentSum < target {
			currentSum += nums[r]
			r++
		}

		if currentSum >= target && r-l < minLen {
			minLen = r - l
		}
		if minLen == 1 {
			return 1
		}
	}
	if minLen == 99999999999 {
		return 0
	}
	return minLen
}
