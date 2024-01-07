// https://leetcode.com/problems/arithmetic-slices-ii-subsequence/
func numberOfArithmeticSlices(nums []int) int {
    n := len(nums)
    
    // dp[index][diff] = count of each sequence lengths
    dp := make(map[int]map[int]map[int]int)
    result := 0
    for i := n-2; i >= 0; i-- {
        if dp[i] == nil {
            dp[i] = make(map[int]map[int]int)
        }
        
        for j := i+1; j < n; j++ {
            diff := nums[i] - nums[j]
               
            if dp[i][diff] == nil {
                dp[i][diff] = make(map[int]int)
            }
            dp[i][diff][2]++
            if dp[j][diff] != nil {
                for k, v := range dp[j][diff] {
                    result += v
                    dp[i][diff][k+1] += v
                }
            }
        }
    }
    //fmt.Println(dp)
    return result
}
