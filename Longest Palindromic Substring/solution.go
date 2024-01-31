// https://leetcode.com/problems/longest-palindromic-substring
package solution

func longestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}
	dp := make([][]bool, n)
	for ix := 0; ix < n; ix++ {
		dp[ix] = make([]bool, n)
	}

	// all substrings of large 1 are palindrome
	maxLength := 1
	for ix := 0; ix < len(s); ix++ {
		dp[ix][ix] = true
	}

	// check substrings with large 2
	start := 0
	for ix := 0; ix < n-1; ix++ {
		if s[ix] == s[ix+1] {
			dp[ix][ix+1] = true
			start = ix
			maxLength = 2
		}
	}

	// check for substrings with large bigger than 2
	for ix := 3; ix <= n; ix++ {
		for iy := 0; iy < n-ix+1; iy++ {
			end := iy + ix - 1
			if dp[iy+1][end-1] && s[iy] == s[end] {
				dp[iy][end] = true

				if ix > maxLength {
					start = iy
					maxLength = ix
				}
			}
		}
	}

	return s[start : maxLength+start]
}
