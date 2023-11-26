func largestSubmatrix(matrix [][]int) int {
    // 1. For each row, count how many consecutive 1's we have for each column ending at this row
    // 2. For each row, sort columns by that number and count max matrix
    result := 0
    consecs := make([][]int, len(matrix))
    for i := range matrix {
        consecs[i] = make([]int, len(matrix[i]))
        for j := range matrix[i] {
            if matrix[i][j] == 1 {
                if i == 0 {
                    consecs[i][j] = 1
                } else {
                    consecs[i][j] = consecs[i-1][j] + 1
                }
            }
        }
    }
    for i := range matrix {
        sort.Sort(sort.Reverse(sort.IntSlice(consecs[i])))
        for j :=0; j < len(matrix[i]) && consecs[i][j] > 0; j++ {
            currRowResult := consecs[i][j] * (j+1)
            if currRowResult > result {
                result = currRowResult
            }
        }
    }
    

    return result
}
