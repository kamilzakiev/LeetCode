func gameOfLife(board [][]int)  {
    // 1 <= m, n
    m, n := len(board), len(board[0])
    // relative coords of all 8 neighbours
    neighbours := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

    // calculate number of alive neighbours for given coordinate
    var numOfAliveNeighbours func(i, j int) int
    numOfAliveNeighbours = func(i, j int) int {
        result := 0
        for _, nbr := range neighbours {
            nbrI, nbrJ := i + nbr[0], j + nbr[1]

            if nbrI >= 0 && nbrI < m && 
                nbrJ >= 0 && nbrJ < n {
                    result += board[nbrI][nbrJ] % 10
            }            
        }
        return result
    }

    for i := range board {
        for j := range board[i] {
            v := board[i][j]
            nbrCount := numOfAliveNeighbours(i, j)

            // apply rules
            newV := v
            if v == 1 {
                if nbrCount < 2 || nbrCount > 3 {
                    newV = 0
                }
            } else {
                if nbrCount == 3 {
                    newV = 1
                }
            }
            board[i][j] = 10 * newV + board[i][j] // new alive flag is stored as 10s
        }
    }
    // once calculated, update all cells simultaneously
    for i := range board {
        for j := range board[i] {
            board[i][j] = board[i][j] / 10
        }
    }
}
