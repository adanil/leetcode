package main

import "math"

const INF = math.MaxInt32

func binSearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	var mid int

	for {
		if low+1 > high {
			return low
		}
		mid = (low + high) / 2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			low = mid + 1
			continue
		}
		if arr[mid] > target {
			high = mid
		}
	}
}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums)+1)
	for i := 0; i < len(nums)+1; i++ {
		dp[i] = INF
	}
	dp[0] = -INF

	ans := 1

	for i := 0; i < len(nums); i++ {
		val := nums[i]
		pos := binSearch(dp, val)
		dp[pos] = min(val, dp[pos])
		ans = max(ans, pos)
	}

	return ans
}

//[10,9,2,5,3,7,101,18]
//
//dp = [-INF, INF ,INF , INF ...]
//dp = [-INF, 10, INF, INF, INF ...]
//dp = [-INF, 9, INF, INF, INF ...]
//dp = [-INF, 2, INF, INF, INF ...]
//dp = [-INF, 2, 5, INF, INF ...]
//dp = [-INF, 2, 3, INF, INF ...]
//dp = [-INF, 2, 3, 7, INF ...]
//dp = [-INF, 2, 3, 7, 101 ...]
