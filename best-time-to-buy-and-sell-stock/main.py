class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        min_price = float('inf')
        max_profit = 0
        for price in prices:
            if price < min_price:
                min_price = min(min_price, price)
            else:
                max_profit_curr = price - min_price
                max_profit = max(max_profit, max_profit_curr)

        return max_profit
        # Time: O(N)
        # Space: O(1)