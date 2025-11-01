package main

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := make([][]int, 0, len(intervals))
	currInt := intervals[0]

	for i := 1; i < len(intervals); i++ {
		if currInt[1] >= intervals[i][0] {
			currInt[1] = max(currInt[1], intervals[i][1])
			continue
		}

		result = append(result, currInt)
		currInt = intervals[i]
	}

	result = append(result, currInt)

	return result
}
