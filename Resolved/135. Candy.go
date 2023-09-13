// https://leetcode.com/problems/candy
func candy(ratings []int) int {
    n := len(ratings)

    leftCandies := make([]int, n)
    rightCandies := make([]int, n)

    for i := 0; i < n; i ++ {
        if i > 0 && ratings[i] > ratings[i-1] {
            leftCandies[i] = leftCandies[i-1] + 1
        } else {
            leftCandies[i] = 1
        }

        j := n - 1 - i
        if j < n - 1 && ratings[j] > ratings[j+1] {
            rightCandies[j] = rightCandies[j+1] + 1
        } else {
            rightCandies[j] = 1
        }
    }
    //fmt.Println(leftCandies)
    //fmt.Println(rightCandies)
    
    result := 0
    for i := 0; i < n; i ++ {
        result += max(leftCandies[i], rightCandies[i])
    }
    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
