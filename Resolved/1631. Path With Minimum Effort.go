// https://leetcode.com/problems/path-with-minimum-effort/
func minimumEffortPath(heights [][]int) int {
    m, n := len(heights), len(heights[0])

    dp := make([][]int, m)
    for i := m - 1; i >= 0; i-- {
        dp[i] = make([]int, n)
        for j := range dp[i] {
            dp[i][j] = math.MaxInt32
        }
    }
    dp[m-1][n-1] = 0

    // BFS+DP approach, will add new cell to the queue, if we found less effort from that cell to end

    queue := [][]int{[]int{m-1, n-1}}

    var checkCell func(prevI, prevJ, i, j int)
    checkCell = func(prevI, prevJ, i, j int) {
        if i < 0 || i > m-1 || j < 0 || j > n-1 {
            return
        }
        prevValue := dp[i][j]
        dp[i][j] = min(dp[i][j], max(dp[prevI][prevJ], abs(heights[prevI][prevJ] - heights[i][j])))
        if dp[i][j] < prevValue {
            queue = append(queue, []int{i, j})
        }
    }


     for len(queue) > 0 {
         popped := queue[0]
         queue = queue[1:]

         i, j := popped[0], popped[1]

         checkCell(i, j, i+1, j)
         checkCell(i, j, i-1, j)
         checkCell(i, j, i, j+1)
         checkCell(i, j, i, j-1)
     }

     return dp[0][0]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
