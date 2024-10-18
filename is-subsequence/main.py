class Solution:
    def isSubsequence(self, s: str, t: str) -> bool:
        s_i = 0
        t_i = 0
        len_s = len(s)
        len_t = len(t)
        if s == "":
            return True
        if len_s > len_t:
            return False

        while t_i < len(t):
            if s_i == len(s):
                break
            if s[s_i] == t[t_i]:
                s_i += 1
            t_i += 1
        return s_i == len(s)