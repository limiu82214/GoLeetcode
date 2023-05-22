package Q121BestTimetoBuyandSellStockPkg

import "math"

func Solve(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	l := 0
	r := 0
	profit := math.MinInt32

	for r < len(prices) {
		tmp := prices[r] - prices[l]
		if tmp > profit {
			profit = tmp
		}

		if prices[l] > prices[r] {
			l = r
		}

		r++
	}

	if profit < 0 {
		return 0
	}

	return profit
}
