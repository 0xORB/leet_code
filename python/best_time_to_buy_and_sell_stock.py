"""
You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing
a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction.
If you cannot achieve any profit, return 0.


Example 1:

Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
Example 2:

Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.


Constraints:

1 <= prices.length <= 105
0 <= prices[i] <= 104
"""

from typing import List

# def maxProfit(prices: List[int]) -> int:
#     """
#     :type prices: List[int]
#     :rtype: int
#     """
#     max_price: int = 0
#     min_price: int = 0
#     buy = None
#     for i in range(len(prices)):
#         cheat = max(prices)
#         if prices[0] or prices[-1] == cheat:
#             prices[0] = 0

#         if i == 0:
#             if min_price == 0 or prices[i] < min_price:
#                 min_price = prices[i]
#             continue

#         if prices[i] < min_price:
#             min_price = prices[i]

#         if prices[i] > max_price:
#             max_price = prices[i]

#         if buy != min_price and min_price == 1 and i != len(prices) - 1 and max_price == max(prices):
#             buy = min_price
#             print(f"Buy  @ ${buy} on Day {i+1}")

#         # print(f"Iteration {i+1}")
#         # print(f"{min_price=}")
#         # print(f"{max_price=}")
#         # print(f"{buy=}\n")

#         if buy is not None and (max_price > min_price) and buy != prices[i]:
#             print(f"Sell @ ${prices[i]} on Day {i+1}")
#             return i + 1
#     return 0

def maxProfit_v2(prices: List[int]) -> int:
    min_price = prices[0]
    max_profit = 0
    for price in prices[1:]:
        max_profit = max(max_profit, price - min_price)
        min_price  = min(min_price, price)

def maxProfit_v3(prices: List[int]) -> int:
    if len(prices) <= 1:
        return 0
    min_price = prices[0]
    max_profit = 0

    for i in range(1, len(prices)):
        if prices[i] < min_price:
            min_price = prices[i]
        elif prices[i] - min_price > max_profit:
            max_profit = prices[i] - min_price
    return max_profit

if __name__ == '__main__':
    prices = [7,1,5,3,6,4] # 5
    # prices = [7,6,4,3,1] # 0
    result =  maxProfit_v3(prices)
    print(f"Day {result}")
