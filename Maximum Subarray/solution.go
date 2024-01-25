// https://leetcode.com/problems/maximum-subarray
package solution

import "math"

func maxSubArray(nums []int) int {
	maxSubArray, maxCurrSumArray := nums[0], nums[0]
	for ix := 1; ix < len(nums); ix++ {
		maxCurrSumArray = max(nums[ix], maxCurrSumArray+nums[ix])
		maxSubArray = max(maxSubArray, maxCurrSumArray)
	}
	return maxSubArray
}

// Divide and conquer approach

func maxSubArray2(nums []int) int {
	return helperMaxSubArray(nums, 0, len(nums)-1)
}

func helperMaxSubArray(nums []int, left, right int) int {
	if left > right {
		return math.MinInt
	}

	middle := (right + left) / 2

	sumLeft := 0
	curr := 0
	for ix := middle - 1; ix >= left; ix-- {
		curr += nums[ix]
		sumLeft = max(sumLeft, curr)
	}

	sumRight := 0
	curr = 0
	for ix := middle + 1; ix <= right; ix++ {
		curr += nums[ix]
		sumRight = max(sumRight, curr)
	}

	bestCombined := nums[middle] + sumRight + sumLeft

	leftHalf := helperMaxSubArray(nums, left, middle-1)
	rightHalf := helperMaxSubArray(nums, middle+1, right)

	return max(bestCombined, max(leftHalf, rightHalf))
}
