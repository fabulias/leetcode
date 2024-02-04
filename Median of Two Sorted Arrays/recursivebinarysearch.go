package main

import "fmt"

func findMedianSortedArraysRecursiveBinarySearch(nums1 []int, nums2 []int) float64 {
	k := 0
	aLen, bLen := len(nums1), len(nums2)
	totalLen := aLen + bLen
	if (totalLen)%2 != 0 {
		k = (aLen + bLen) / 2
		return float64(helperFindMedianSortedArrays(nums1, nums2, 0, aLen-1, 0, bLen-1, k))
	} else {
		return float64(helperFindMedianSortedArrays(nums1, nums2, 0, aLen-1, 0, bLen-1, (aLen+bLen)/2)+helperFindMedianSortedArrays(nums1, nums2, 0, aLen-1, 0, bLen-1, (aLen+bLen)/2-1)) / 2
	}
}

func helperFindMedianSortedArrays(A, B []int, aStart, aEnd, bStart, bEnd, k int) int {
	if aStart > aEnd {
		return B[k-aStart]
	} else if bStart > bEnd {
		return A[k-bStart]
	} else {
		aIndex := (aStart + aEnd) / 2
		bIndex := (bStart + bEnd) / 2

		aValue := A[aIndex]
		bValue := B[bIndex]

		if aIndex+bIndex < k {
			if aValue < bValue {
				aStart = aIndex + 1
			} else {
				bStart = bIndex + 1
			}
		} else {
			if bValue > aValue {
				bEnd = bIndex - 1
			} else {
				aEnd = aIndex - 1
			}
		}
	}
	return helperFindMedianSortedArrays(A, B, aStart, aEnd, bStart, bEnd, k)
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArraysRecursiveBinarySearch(nums1, nums2))
	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	fmt.Println(findMedianSortedArraysRecursiveBinarySearch(nums1, nums2))
}
