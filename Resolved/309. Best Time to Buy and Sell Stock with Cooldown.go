// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
func maxProfit(prices []int) int {
    n := len(prices)
    dp1 := make([]int, n) // we have 1 stock
    dp0 := make([]int, n) // we have 0 stock

    dp0[0] = 0
    dp1[0] = - prices[0]

    for i := 1; i < n; i++ {
        dp0[i] = max(dp0[i-1], dp1[i-1] + prices[i])
        if i >= 2 {
            dp1[i] = max(dp1[i-1], dp0[i-2] - prices[i]) // unique condition for 'with cooldown'case

        } else {
            dp1[i] = max(dp1[i-1], dp0[i-1] - prices[i])
        }
    }

    //fmt.Println(dp0, dp1)
    return max(dp0[n-1], dp1[n-1])
}
