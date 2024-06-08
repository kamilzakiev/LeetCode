// https://leetcode.com/problems/continuous-subarray-sum/
func checkSubarraySum(nums []int, k int) bool {
    // use prefix sum, but instead of sum calculate reminder
    // if the same reminder is met twice, then sum between them is divisible by k
    rems := map[int]int{}

    remSum := 0

    rems[0] = 0

    for i, v := range nums {
        remSum = int((int64(remSum) + int64(v)) % int64(k))
        if prevCount, exists := rems[remSum]; exists {
            if i + 1 - prevCount >= 2 { // its length is at least two
                return true
            }
        } else {
            rems[remSum] = i+1
        }
    }
    return false
}
