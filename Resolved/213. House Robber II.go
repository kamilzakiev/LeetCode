// https://leetcode.com/problems/house-robber-ii/
func rob(nums []int) int {
    n := len(nums)
    if n == 1 {
        return nums[0]
    }
    firstDp := make([]int, n+3) // starting from first house, then dismiss last
    secondDp := make([]int, n+3) // dismiss first, can include last one

    for i := n-1; i >= 0; i-- {
        if i < n-1 {
            firstDp[i] = max(nums[i] + firstDp[i+2], firstDp[i+1])
        }
        if i > 0 {
            secondDp[i] = max(nums[i] + secondDp[i+2], secondDp[i+1])
        }
    }
    return max(firstDp[0], secondDp[1])
}
