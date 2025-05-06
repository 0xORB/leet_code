package leetcode

func maximumWealth(accounts [][]int) int {
	var wealth int
	for i := range accounts {
		temp := 0
		for j := range accounts[i] {
			temp += accounts[i][j]
		}
		wealth = max(wealth, temp)
	}
	return wealth
}
