package main

import (
	"sort"
)

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {

	nums := append(nums1, nums2...)
	sort.Ints(nums)

	mid := (len(nums) / 2)

	if len(nums)%2 == 1 {
		return float64(nums[mid])
	}

	return (float64(nums[mid-1]) + float64(nums[mid])) / 2
}
