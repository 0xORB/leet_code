package leetcode

func maxProfit(prices []int) int {
	l, r, p := 0, 1, 0
	for r < len(prices) {
		if prices[l] > prices[r] {
			l = r
		} else {
			p = max(p, prices[r]-prices[l])
		}
		r += 1
	}
	return p
}
