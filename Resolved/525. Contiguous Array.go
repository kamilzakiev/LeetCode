// https://leetcode.com/problems/contiguous-array/
func findMaxLength(nums []int) int {
    // we need to find 2 points max fara away from each other, that has the same diff between number of 1 and 0
    //prefixSum := make([]int, n)
    result := 0
    diffOneZero := map[int]int{} // diif between count of 1 and 0

    oneCount := 0
    diffOneZero[0] = -1
    for i, v := range nums {
        oneCount += v
        //prefixSum[i] = sum
        diff := 2 * oneCount - i - 1
        if prevIndex, exists := diffOneZero[diff]; exists {
            result = max(result, i - prevIndex)
        } else {
            diffOneZero[diff] = i
        }
    }

    return result
}
