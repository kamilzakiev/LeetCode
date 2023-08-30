// https://leetcode.com/problems/surrounded-regions/
func solve(board [][]byte)  {   
    m := len(board)
    n := len(board[0])
    neighbours := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    visited := make([][]bool, m)

    // returns whether current cell with value '0' is in surrounded region
    var isIsland func(i, j int) bool

    isIsland = func(i, j int) bool {        
        if i == 0 || i == m - 1 || j == 0 || j == n - 1 {
            return false
        }
        visited[i][j] = true
        result := true

        for _, neighbour := range neighbours {
            iDelta, jDelta := neighbour[0], neighbour[1]
            if board[i+iDelta][j+jDelta] == 'O' && !visited[i+iDelta][j+jDelta] {
                if !isIsland(i+iDelta, j+jDelta) {
                    // why we do not return here? because we should mark all connected 0s as visited
                    result = false
                }
            }
        }        

        return result
    }

    var capture func(i, j int) 
    capture = func(i, j int) {   
        board[i][j] = 'X'
        for _, neighbour := range neighbours {
            iDelta, jDelta := neighbour[0], neighbour[1]
            if board[i+iDelta][j+jDelta] == 'O' {
                capture(i+iDelta, j+jDelta)
            }
        }
    }

    for i := range board {
        visited[i] = make([]bool, n)
    }
    for i := range board {
        for j := range board[i] {
            if board[i][j] == 'O' && !visited[i][j] {
                if isIsland(i, j) {   
                    capture(i, j)
                }
            }
        }
    }
}
