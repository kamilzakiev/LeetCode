// https://leetcode.com/problems/minimum-operations-to-reduce-x-to-zero/
func minOperations(nums []int, x int) int {
    // 2 pointer approach, but instead of contigious sum, we are calculating sum of elements from both ends
    // Greedy approach
    n := len(nums)
    result := math.MaxInt32
    start := 0

    currSum := 0
    for ;start < n && currSum < x; start++ {
        currSum += nums[start]
        if currSum == x {
            result = min(result, start+1)
        }
    }
    if currSum < x {
        return -1
    }
    start--
    fmt.Println(start)
    end := n-1
    for ;start >= 0; start-- {
        currSum = currSum - nums[start]        
        for ;end > start && currSum < x; end-- {
            currSum += nums[end]            
        }
        if currSum == x {
            //fmt.Println(start, end)
            result = min(result, start + n - end - 1)
       }
    }
    if result == math.MaxInt32 {
        return -1
    }
    return result
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
