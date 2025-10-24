package main

import "sort"

type Triplet struct {
	a, b, c int
}

func makeTriplet(a, b, c int) Triplet {
	return Triplet{
		a: a,
		b: b,
		c: c,
	}
}

// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k,
// and nums[i] + nums[j] + nums[k] == 0.
// Notice that the solution set must not contain duplicate triplets.
func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	result := make([][]int, 0)

	uniqTriplets := make(map[Triplet]bool)

	m := make(map[int][]int, len(nums))
	for pos, e := range nums {
		m[e] = append(m[e], pos)
	}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // Skip duplicate values for i
		}
		for j := i + 1; j < len(nums); j++ {
			rem := -nums[i] - nums[j]
			poses, ok := m[rem]

			if ok {
				t := makeTriplet(nums[i], nums[j], rem)
				if _, ok = uniqTriplets[t]; ok {
					continue
				}

				for _, pos := range poses {
					if pos == i || pos == j || pos < j {
						continue
					}

					result = append(result, []int{nums[i], nums[j], rem})
					uniqTriplets[t] = true
					break
				}
			}

		}
	}
	return result
}

func main() {

}
