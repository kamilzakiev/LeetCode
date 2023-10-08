// https://leetcode.com/problems/max-dot-product-of-two-subsequences
func maxDotProduct(nums1 []int, nums2 []int) int {
    m, n := len(nums1), len(nums2)

    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
    }

    maxV := math.MinInt32
    for i := m-1; i >=0; i-- {
         maxV = max(maxV, nums1[i] * nums2[n-1])
         dp[i][n-1] = maxV
    }

    maxV = math.MinInt32
    for j := n-1; j >=0; j-- {
         maxV = max(maxV, nums1[m-1] * nums2[j])
         dp[m-1][j] = maxV
    }

    for i := m-2; i >= 0; i-- {
        for j := n-2; j >= 0; j-- {
            dp[i][j] = max(dp[i+1][j], dp[i][j+1])
            dp[i][j] = max(dp[i][j], nums1[i]*nums2[j])
            dp[i][j] = max(dp[i][j], max(0, nums1[i]*nums2[j]) + dp[i+1][j+1])
        }
    }

    fmt.Println(dp)

    return dp[0][0]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
