// https://leetcode.com/problems/make-sum-divisible-by-p/
func minSubarray(nums []int, p int) int {
    n := len(nums)
    prefixRem := make([]int, n + 1)

    for i, v := range nums {
        prefixRem[i + 1] = (prefixRem[i] + v) % p
    }
    excessiveRem := prefixRem[n] // need to find subarray that has a sum with such remainder
    if excessiveRem == 0 {
        return 0
    }

    result := math.MaxInt32
    mostRight := make(map[int]int) // indicies of most right remainders of prefix sum
    for i, v := range prefixRem {    
        searchedPrev := (v - excessiveRem + p) % p
        if prevIndex, exists := mostRight[searchedPrev]; exists {
            diff := i - prevIndex
            if diff < result {
                result = diff
            }
        }    
        mostRight[prefixRem[i]] = i
    }

    if result == n || result == math.MaxInt32 {
        return -1
    }
    return result
}
