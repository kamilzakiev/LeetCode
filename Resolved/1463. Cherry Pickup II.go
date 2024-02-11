// https://leetcode.com/problems/cherry-pickup-ii/
func cherryPickup(grid [][]int) int {
    m, n := len(grid), len(grid[0])

    dp := make([][][]int, m+1)
    for i := range dp {
        dp[i] = make([][]int, n)
        for j := range dp[i] {
            dp[i][j] = make([]int, n)
        }
    }

    for i := m-1; i >= 0; i-- {
        for j1 := range dp[i] {
            for j2 := range dp[i][j1] {
                for a1 := j1-1; a1 <= j1+1; a1++ { // check 3 cells below for robot1
                    for a2 := j2-1; a2 <= j2+1; a2++ { // and robot 2
                        if a1 >= 0 && a1 < n && a2 >= 0 && a2 < n {
                            dp[i][j1][j2] = max(dp[i][j1][j2], dp[i+1][a1][a2]) 
                        }
                    }
                }
                dp[i][j1][j2] += grid[i][j1]
                if j2 != j1 {
                    dp[i][j1][j2] += grid[i][j2]
                }                
            }
        }
    }    
    return dp[0][0][n-1]
}
