// https://leetcode.com/problems/maximum-score-of-a-good-subarray/
func maximumScore(nums []int, k int) int {
    // 1. start from index k
    // 2. When choosing left or right, use value that has greater min value from k to it
    n := len(nums)

    mins := make([]int, n)

    minV := nums[k]
    for i := k; i >= 0; i-- {
        minV = min(nums[i], minV)
        mins[i] = minV
    }
    minV = nums[k]
    for i := k+1; i < n; i++ {
        minV = min(nums[i], minV)
        mins[i] = minV
    }
    l, r := k, k
    result := nums[k]
    //fmt.Println(mins)
    for l > 0 || r < n-1 {
        if l == 0 {
            r++
        } else if r == n-1 {
            l--
        } else if mins[l-1] > mins[r+1] {
            l--
        } else {
            r++
        }
        //fmt.Println(l, r)
        minV = min(mins[l], mins[r])
        dist := r - l + 1
        score := minV * dist
        if result < score {
            result = score
        }
    }
    return result
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
