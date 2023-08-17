// https://leetcode.com/problems/01-matrix/
func updateMatrix(mat [][]int) [][]int {
    m := len(mat)
    n := len(mat[0])

    result := make([][]int, m)

    for i := range result {
        result[i] = make([]int, n)
        for j := range result[i] {
            if mat[i][j] == 1 {
                result[i][j] = math.MaxInt32
            }
        }
    }
        
    queue := [][]int{}

    for i := range mat {
        for j := range mat[i] {
            if mat[i][j] == 0 { 
                queue = append(queue, []int{i, j})
            }
        }
    }

    visited := make([][]bool, m)
    for i := range visited {
        visited[i] = make([]bool, n)
    }
    // regular BFS approach
    for len(queue) > 0 {
        popped := queue[0]
        queue = queue[1:]

        i, j := popped[0], popped[1]
        //fmt.Println(i, j)
        // can be wrapped into separate function with 4 calls of it
        if i > 0 && mat[i-1][j] == 1 && !visited[i-1][j] {
            queue = append(queue, []int{i-1, j})
            result[i-1][j] = result[i][j] + 1
            visited[i-1][j] = true
        }
        if i < m - 1  && mat[i+1][j] == 1 && !visited[i+1][j] {
            queue = append(queue, []int{i+1, j})
            result[i+1][j] = result[i][j] + 1
            visited[i+1][j] = true
        }
        if j > 0  && mat[i][j - 1] == 1 && !visited[i][j-1] {
            queue = append(queue, []int{i, j-1})
            result[i][j-1] = result[i][j] + 1
            visited[i][j-1] = true
        }
        if j < n - 1  && mat[i][j + 1] == 1 && !visited[i][j+1] {
            queue = append(queue, []int{i, j + 1})
            result[i][j+1] = result[i][j] + 1
            visited[i][j+1] = true
        }        
    }    

    return result
}
