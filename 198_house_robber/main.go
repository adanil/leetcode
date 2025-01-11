package main

func rob(nums []int) int {
	dp := make([]int, len(nums))
	mx := 0
	for i, _ := range nums {
		if i == 0 {
			dp[i] = nums[0]
			mx = dp[i]
			continue
		}

		if i-2 < 0 {
			dp[i] = max(nums[i], nums[i-1])
		} else {
			dp[i] = max(nums[i]+dp[i-2], dp[i-1])
		}
		if dp[i] > mx {
			mx = dp[i]
		}
	}
	return mx
}
