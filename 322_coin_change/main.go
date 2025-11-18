package main

import "math"

const INF = math.MaxInt32

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	dp := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		dp[i] = INF
	}
	for _, coin := range coins {
		if coin > amount {
			continue
		}
		dp[coin] = 1
	}

	for i := 0; i < amount+1; i++ {
		for j := 0; j < len(coins); j++ {
			if i-coins[j] < 0 || dp[i-coins[j]] == INF {
				continue
			}
			dp[i] = min(dp[i], dp[i-coins[j]]+1)
		}
	}

	ans := dp[amount]
	if ans == INF {
		return -1
	}
	return ans
}

// coins = [1 2 5] amount = 11
// dp = [INF 1 1 INF INF 1 INF INF INF INF INF INF]
// dp[i] = min(dp[i], dp[i - coins[j]] + 1), dp[i - coins[j] > 0
// O (amount * len(coins))
