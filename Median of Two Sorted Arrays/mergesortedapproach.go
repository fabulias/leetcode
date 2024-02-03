package main

import "fmt"

/*

https://leetcode.com/problems/median-of-two-sorted-arrays/


Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).



Example 1:

Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.
Example 2:

Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.


Constraints:

nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-106 <= nums1[i], nums2[i] <= 106
*/

func getMin(nums1, nums2 []int, pointers []int) int {
	number := 0
	if pointers[0] < len(nums1) && pointers[1] < len(nums2) {
		if nums1[pointers[0]] < nums2[pointers[1]] {
			number = nums1[pointers[0]]
			pointers[0]++
		} else {
			number = nums2[pointers[1]]
			pointers[1]++
		}
	} else if pointers[0] >= len(nums1) {
		number = nums2[pointers[1]]
		pointers[1]++
	} else if pointers[1] >= len(nums2) {
		number = nums1[pointers[0]]
		pointers[0]++
	}
	return number
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	pointers := []int{0, 0}
	totalSize := len(nums1) + len(nums2)
	if totalSize%2 != 0 {
		for ix := 0; ix < totalSize/2; ix++ {
			getMin(nums1, nums2, pointers)
		}
		return float64(getMin(nums1, nums2, pointers))
	} else {
		for ix := 0; ix < totalSize/2-1; ix++ {
			getMin(nums1, nums2, pointers)
		}
		n1 := getMin(nums1, nums2, pointers)
		n2 := getMin(nums1, nums2, pointers)
		return float64(n1+n2) / 2
	}
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
