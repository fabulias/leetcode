// https://leetcode.com/problems/two-sum
package solution

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for ix, n := range nums {
		if _, ok := m[n]; ok {
			return []int{ix, m[n]}
		}
		m[target-n] = ix
	}
	return []int{0, 0}
}
