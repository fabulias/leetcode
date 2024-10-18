class Solution:
    def romanToInt(self, s: str) -> int:
        converter={
            "I":1,
            "V":5,
            "X":10,
            "L":50,
            "C":100,
            "D":500,
            "M":1000,
        }
        len_s=len(s)
        sum=0
        i=0
        while i < len_s:
            if i < len_s-1 and converter[s[i]] < converter[s[i+1]]:
                sum+=(converter[s[i+1]]-converter[s[i]])
                i+=1
            else:
                sum+=converter[s[i]]
            i+=1
        return sum