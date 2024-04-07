// https://leetcode.com/problems/valid-parenthesis-string/
func checkValidString(s string) bool {
    // DP approach
    // dp[i][sum] = true means that for i-th character we can get sum balance (count of '(' - count of ')')
    dp := make([][]bool, 100)

    for i := range dp {
        dp[i] = make([]bool, 101)
    }

    r0 := s[0]
    if r0 == '(' {
        dp[0][1] = true
    } else if r0 == ')' {
        return false
    } else {
        dp[0][0] = true
        dp[0][1] = true
    }

    for i:= 1; i < len(s); i++ {
        r := s[i]
        if r == '(' {
            for j := 1; j < len(dp[i]); j++ {
                dp[i][j] = dp[i][j] || dp[i-1][j-1]
            }
        } else if r == ')' {
            for j := 0; j < len(dp[i])-1; j++ {
                dp[i][j] = dp[i][j] || dp[i-1][j+1]
            }
        } else { // consider all 3 possibilities
            for j := 0; j < len(dp[i]); j++ {
                dp[i][j] = dp[i][j] || dp[i-1][j]
                if j > 0 {
                    dp[i][j] = dp[i][j] || dp[i-1][j-1]
                }
                if j < len(dp[i])-1 {
                    dp[i][j] = dp[i][j] || dp[i-1][j+1]
                }
            }
        }
    }
    return dp[len(s)-1][0]
}
