// https://leetcode.com/problems/number-of-submatrices-that-sum-to-target/
func numSubmatrixSumTarget(matrix [][]int, target int) int {
    m, n := len(matrix), len(matrix[0])
    prefix := make([][]int, m)
    for i := range matrix {
        prefix[i] = make([]int, n) // 2-d prefix sum array
        colSum := 0
        for j := range prefix[0] {
            colSum += matrix[i][j]
            prefix[i][j] = colSum
            if i > 0 {
                prefix[i][j] += prefix[i-1][j]
            } 
        }
    }
    var calcSum func(iFrom, jFrom, iTo, jTo int) int // func to return sum of submatrix in O(1)
    calcSum = func(iFrom, jFrom, iTo, jTo int) int {
        result := prefix[iTo][jTo]
        if iFrom > 0 {
            result -= prefix[iFrom-1][jTo]
        }
        if jFrom > 0 {
            result -= prefix[iTo][jFrom-1]
        }
        if iFrom > 0 && jFrom > 0 {
            result += prefix[iFrom-1][jFrom-1] // because that part already was subtracted 2 times
        }
        return result
    }
    
    result := 0
    for i1 := range prefix {
        for i2 := i1; i2 < m; i2++ {
            mapSum := make(map[int]int)
            currSum := 0
            for j := range prefix[0] {                
                currSum = calcSum(i1, 0, i2, j)
                
                if currSum == target {
                    result++
                }
                if mapSum[currSum-target] > 0 {
                    result += mapSum[currSum-target]
                }
                mapSum[currSum]++
            }
        }
    }
    return result
}
