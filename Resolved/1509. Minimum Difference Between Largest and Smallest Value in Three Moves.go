// https://leetcode.com/problems/minimum-difference-between-largest-and-smallest-value-in-three-moves/
func minDifference(nums []int) int {
    n := len(nums)
    sort.Ints(nums)

    if len(nums) <= 4 {
        return 0
    }
    // need to check only start and end part of the sorted array
    // because that provides the max diff
    list := []int{}

    list = append(list, nums[n-1] - nums[3])
    list = append(list, nums[n-2] - nums[2])
    list = append(list, nums[n-3] - nums[1])
    list = append(list, nums[n-4] - nums[0])
    
    return slices.Min(list)
}
