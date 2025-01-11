package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	idx1, idx2, resultIdx := 0, 0, 0
	result := make([]int, n+m)
	for idx1 < m || idx2 < n {
		if idx1 == m {
			result[resultIdx] = nums2[idx2]
			idx2++
		} else if idx2 == n {
			result[resultIdx] = nums1[idx1]
			idx1++
		} else if nums1[idx1] <= nums2[idx2] {
			result[resultIdx] = nums1[idx1]
			idx1 += 1
		} else {
			result[resultIdx] = nums2[idx2]
			idx2 += 1
		}
		resultIdx += 1
	}
	copy(nums1, result)
}

func main() {

}
