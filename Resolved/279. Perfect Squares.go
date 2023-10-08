// https://leetcode.com/problems/perfect-squares/
func numSquares(n int) int {
    squares := []int{}
    for i := 1; i*i <= n; i++ {
        squares = append(squares, i*i)
        if i * i == n {
            return 1
        }
    }
    memo := map[int]int{}
    var iter func(x int) int 
    iter = func(x int) int {
        if x == 0 {
            return 0
        }
        if memo[x] > 0 {
            return memo[x]
        }
        result := math.MaxInt32
        for i := 0; i < len(squares) && squares[i] <= x; i++ {
            result = min(result, 1 + iter(x-squares[i]))
        }
        memo[x] = result
        return result

    }
    return iter(n)
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
