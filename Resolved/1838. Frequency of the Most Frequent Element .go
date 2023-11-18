// https://leetcode.com/problems/frequency-of-the-most-frequent-element/
func maxFrequency(nums []int, k int) int {
    // sort, and then use 2-pointer approach from end of the array
    n := len(nums)
    sort.Ints(nums)
    //fmt.Println(n, nums)
    if n == 1 {
        return 1
    }
    result := 0
    freeK := k
    j := n-2
    currFreq := 1
    for i := n-1; i >= 0; i-- {
        for j >= 0 && freeK >= (nums[i] - nums[j]) {
            freeK = freeK - (nums[i] - nums[j])
            currFreq++
            j--
        }
        //fmt.Println(i, j, freeK, currFreq)
        if i > 0 {
            freeK = freeK + (nums[i] - nums[i-1]) * (i-j-1)
        }
        //fmt.Println(i, freeK)
        if result < currFreq {
            result = currFreq
        }
        currFreq--
    }
    return result
}
