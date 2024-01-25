// https://leetcode.com/problems/maximum-subarray
package solution

func maxSubArray(nums []int) int {
	maxSubArray, maxCurrSumArray := nums[0], nums[0]
	for ix := 1; ix < len(nums); ix++ {
		maxCurrSumArray = max(nums[ix], maxCurrSumArray+nums[ix])
		maxSubArray = max(maxSubArray, maxCurrSumArray)
	}
	return maxSubArray
}
