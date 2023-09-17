// https://leetcode.com/problems/shortest-path-visiting-all-nodes/
func shortestPathLength(graph [][]int) int {
    n := len(graph)
    // 1. Calculate min distance between every 2 nodes using Floyd Warshall Algorithm:
    
    W := make([][]int, n)

    for i := range W {
        W[i] = make([]int, n)
    }
    for i := range W {
        for j := 0; j < n; j++ {
            W[i][j] = 1001 // default values
        }
    }
    for i := range W {
        for _, j := range graph[i] {
            W[i][j] = 1 // distance between 2 adjacent nodes
        }
        W[i][i] = 0
    }

    // Floyd Warshall Algorithm:
    for i := range W {
        for j := range W {
            for k := range W {
                W[i][j] = min(W[i][j] , W[i][k] + W[k][j])
            }
        }
    }

    // 2. Use DP approach to find all pathes visiting all nodes
    dp := make([][]int, 1<<n) // dp[mask][last] - mask = visited nodes, last = last visited node

    for mask := 0; mask < 1<<n; mask++ {
        dp[mask] = make([]int, n)
    }

    for i := 0; i < n ; i++ {
        for mask := 0; mask < 1<<n; mask++ {
            dp[mask][i] = 1000
        }
        dp[1<<i][i] = 0
    }

    for mask := 0; mask < 1<<n; mask++ {
        visited := []int{} // visited nodes in mask
        for j := 0; j < n; j++ { // loop through visited nodes in mask
            if mask & (1<<j) > 0 {
                visited = append(visited, j)
            }
        }        

        for _, j := range visited {
            for _, k := range visited {
                if j != k {
                    neib := dp[mask ^ (1 << j)][k] // mask excluding j, ending with k
                    dp[mask][j] = min(dp[mask][j], neib + W[k][j])
                }
            }
        }
    }
    result := math.MaxInt32

    for i := 0; i < n; i++ {
        result = min(result, dp[(1<<n) - 1][i])
        fmt.Println(i, dp[(1<<n) - 1][i])
    }
    return result
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
