// https://leetcode.com/problems/word-search
func exist(board [][]byte, word string) bool {
    // simple recursive approach
    m := len(board)
    n := len(board[0])

    var findNext func(i, j int, wordPart string, visited [][]bool) bool 

    findNext = func(i, j int, wordPart string, visited [][]bool) bool {
        if len(wordPart) == 0 {
            return true
        }
        if visited[i][j] {
            return false
        }
        visited[i][j] = true
        defer func(){ 
            visited[i][j] = false
        }()

        currSym := wordPart[0]   
        nextPart := wordPart[1:]

        var validNeighbour func(x, y int) bool
        validNeighbour = func(x, y int) bool {
            return x >= 0 && x < m && y >= 0 && y < n && !visited[x][y] && board[x][y] == currSym && findNext(x, y, nextPart, visited) 
        }

        if validNeighbour(i-1, j) {
            return true
        }
        if validNeighbour(i+1, j) {
            return true
        }
        if validNeighbour(i, j-1) {
            return true
        }
        if validNeighbour(i, j+1) {
            return true
        }
        
        return false
    } 
    for i := range board {
        for j := range board[i] {
            if board[i][j] == word[0] {
                visited := make([][]bool, m)
                for i := range board {
                    visited[i] = make([]bool, n)
                }
                if findNext(i, j, word[1:], visited) {
                    return true
                }
            }
        }
    }
    return false
}
