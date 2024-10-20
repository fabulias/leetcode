class Solution:
    def summaryRanges(self, nums: List[int]) -> List[str]:
        i = 0
        
        answers = []
        while i < len(nums):
            k = i
            while k < len(nums)-1 and nums[k]+1 == nums[k+1]:
                k+=1
            answers.append(str(nums[i])) if nums[i] == nums[k] else answers.append(str(nums[i])+'->'+str(nums[k]))
            i = k+1
        
        return answers
        # Time: O(n)
        # Space: O(n)