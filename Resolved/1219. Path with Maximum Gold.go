// https://leetcode.com/problems/path-with-maximum-gold/
func getMaximumGold(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    var dfs func(i, j int, visited *map[int]bool) int
    dfs = func(i, j int, visited *map[int]bool) int {
        if i < 0 || i > m-1 || j < 0 || j > n-1 {
            return 0
        }
        if grid[i][j] == 0 {
            return 0
        }
        hash := i*20 + j
        if (*visited)[hash] {
            return 0
        }
        (*visited)[hash] = true
        maxPath := 0
        maxPath = max(maxPath, dfs(i+1, j, visited))
        maxPath = max(maxPath, dfs(i-1, j, visited))
        maxPath = max(maxPath, dfs(i, j+1, visited))
        maxPath = max(maxPath, dfs(i, j-1, visited))
        (*visited)[hash] = false
        return maxPath + grid[i][j]
    }
    result := 0
    for i := range grid {
        for j := range grid[0] {
            visited := make(map[int]bool) 
            curr := dfs(i, j, &visited)
            if curr > result {
                result = curr
            }

        }
    }

    return result
}
