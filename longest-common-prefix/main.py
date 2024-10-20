class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        common_prefix = ''

        min_len_word = min(len(s) for s in strs)

        i = 0
        for i in range(0, min_len_word):
            for word in strs:
                if word[i] != strs[0][i]:
                    return strs[0][:i]
            i+=1

        # Time: O(N*M)
        # Space: O(1)
        
        return strs[0][:i]
        