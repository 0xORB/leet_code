package leetcode

func maxProfitTwo(prices []int) int {
	i, p := 1, 0
	for i < len(prices) {
		if prices[i-1] < prices[i] {
			p += prices[i] - prices[i-1]
		}
		i += 1
	}
	return p
}
