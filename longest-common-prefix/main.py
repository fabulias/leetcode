class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        common_prefix = ''

        min_word = ''
        min_len_word = float('inf')
        for word in strs:
            if len(word) < min_len_word:
                min_word = word
                min_len_word = len(word)

        i = 0
        for letter in min_word:
            for word in strs:
                if word[i] != letter:
                    return min_word[:i]
            i+=1

        # Time: O(N*M)
        # Space: O(1)
        
        return min_word[:i]
        