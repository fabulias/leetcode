// https://leetcode.com/problems/longest-substring-without-repeating-characters
package solution

func lengthOfLongestSubstring(s string) int {
	chars := map[byte]int{}
	ans := 0
	left, right := 0, 0

	for right < len(s) {
		char := s[right]

		if v, ok := chars[char]; ok && v >= left && v < right {
			left = v + 1
		}

		chars[char] = right
		ans = max(ans, right+1-left)
		right++
	}
	return ans
}
