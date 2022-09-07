package array

func maxProfit(prices []int) int {
	maxvalue := 0
	firstIndex := 0
	currentMax := 0
	for secondIndex := 1; secondIndex < len(prices); secondIndex++ {
		if prices[secondIndex] < prices[firstIndex] {
			maxvalue = maxInt(maxvalue, currentMax)
		}
		currentMax += prices[secondIndex] - prices[firstIndex]
		if currentMax < 0 {
			currentMax = 0
		}
		firstIndex++
	}
	return maxInt(maxvalue, currentMax)
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
