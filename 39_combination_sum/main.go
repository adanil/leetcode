package main

import (
	"fmt"
	"sort"
)

var result [][]int

func combinationSum(candidates []int, target int) [][]int {
	result = make([][]int, 0)
	numbers := make(map[int]any, len(candidates))
	for _, cand := range candidates {
		numbers[cand] = struct{}{}
	}
	step(numbers, []int{}, target)
	res := deduplicate(result)
	return res
}

func step(numbers map[int]any, path []int, current int) {
	for k, _ := range numbers {
		if k <= current {
			currPath := make([]int, len(path))
			path = append(path, k)
			copy(currPath, path)
			step(numbers, path, current-k)
			path = currPath
		}
	}
	if current == 0 {
		result = append(result, path)
	}
}

func deduplicate(arr [][]int) [][]int {
	r := make(map[string][]int)
	_ = r
	for i := 0; i < len(arr); i++ {
		sort.Slice(arr[i], func(k, j int) bool { return arr[i][k] < arr[i][j] })
		r[fmt.Sprint(arr[i])] = arr[i]
	}

	res := make([][]int, 0)
	for _, v := range r {
		res = append(res, v)
	}
	return res

}
