// https://leetcode.com/problems/binary-subarrays-with-sum/
func numSubarraysWithSum(nums []int, goal int) int {
    sum := 0

    prefixHash := map[int]int{}
    result := 0

    for _, v := range nums {
        sum += v
        prevSearchSum := sum - goal
        result += prefixHash[prevSearchSum]
        if sum == goal {
            result++
        }
        prefixHash[sum]++
    }
    
    return result
}
