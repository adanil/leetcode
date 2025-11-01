package main

func merge(nums, result []int, left, mid, right int) {
	it1 := 0
	it2 := 0

	for left+it1 < mid && mid+it2 < right {
		if nums[left+it1] < nums[mid+it2] {
			result[it1+it2] = nums[left+it1]
			it1++
		} else {
			result[it1+it2] = nums[mid+it2]
			it2++
		}
	}

	for left+it1 < mid {
		result[it1+it2] = nums[left+it1]
		it1++
	}

	for mid+it2 < right {
		result[it1+it2] = nums[mid+it2]
		it2++
	}

	for i := 0; i < it1+it2; i++ {
		nums[left+i] = result[i]
	}
}

func mergeSort(nums, result []int, left, right int) {
	if left+1 >= right {
		return
	}

	mid := (left + right) / 2

	mergeSort(nums, result, left, mid)
	mergeSort(nums, result, mid, right)
	merge(nums, result, left, mid, right)
}

func sortArray(nums []int) []int {
	result := make([]int, len(nums))
	mergeSort(nums, result, 0, len(nums))
	return nums
}
