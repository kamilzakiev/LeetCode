// https://leetcode.com/problems/find-all-groups-of-farmland/
func findFarmland(land [][]int) [][]int {
    // simple dfs approach
    m, n := len(land), len(land[0])

    visited := make([][]bool, m)

    for i := range visited {
        visited[i] = make([]bool, n)
    }

    result := [][]int{}

    var dfs func(i, j int, minMax []int) 
    dfs = func(i, j int, minMax []int) {
        if i < 0 || i > m-1 || j < 0 || j > n-1 {
            return
        }
        if visited[i][j] || land[i][j] != 1 {
            return
        }
        visited[i][j] = true
        minMax[0] = min(minMax[0], i)
        minMax[1] = min(minMax[1], j)
        minMax[2] = max(minMax[2], i)
        minMax[3] = max(minMax[3], j)
        dfs(i-1, j, minMax)
        dfs(i+1, j, minMax)
        dfs(i, j-1, minMax)
        dfs(i, j+1, minMax)
    }

    for i := range land {
        for j := range land[i] {
            if land[i][j] == 1 && !visited[i][j] {
                minMax := []int{i, j, i, j}
                dfs(i, j, minMax)
                result = append(result, minMax)                
            }
        }
    }

    return result
}
