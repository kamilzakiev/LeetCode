// https://leetcode.com/problems/out-of-boundary-paths/
func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
    const mod = 1e9+7
    dp := make([][][]int, m)

    for i := range dp {
        dp[i] = make([][]int, n)
        for j := range dp[i] {
            dp[i][j] = make([]int, maxMove + 1)

            for x := range dp[i][j] {
                dp[i][j][x] = -1
            }
        }
    }

    var iter func(i, j, moveLeft int) int

    iter = func(i, j, moveLeft int) int {
        if i < 0 || i == m || j < 0 || j == n || moveLeft == 0 {
            return 0
        }
        
        if dp[i][j][moveLeft] > -1 {
            return dp[i][j][moveLeft]
        }
        result := 0
        if i == 0 {
            result++
        }
        if i == m-1 {
            result++
        } 
        if j == 0 {
            result++
        }
        if j == n-1 {
            result++
        }
        result += int((int64(iter(i-1, j, moveLeft-1)) + int64(iter(i+1, j, moveLeft-1)) + int64(iter(i, j-1, moveLeft-1)) + int64(iter(i, j+1, moveLeft-1))) % mod)
        dp[i][j][moveLeft] = result
        return result
    } 

    return iter(startRow, startColumn, maxMove)   
}
