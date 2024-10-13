class Solution:
    def mergeAlternately(self, word1: str, word2: str) -> str:
        len1=len(word1)
        len2=len(word2)
        iter1=0
        iter2=0

        a = []
        while iter1 < len1 or iter2 < len2:
            if iter1 < len1:
                a.append(word1[iter1])
                iter1+=1
            if iter2 < len2:
                a.append(word2[iter2])
                iter2+=1

        return ''.join(a)
