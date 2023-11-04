// https://leetcode.com/problems/last-moment-before-all-ants-fall-out-of-a-plank/
func getLastMoment(n int, left []int, right []int) int {
    // just imagine that once 2 ants hit, they are switched
    // it means that it is the same task if ants would cross each other like transparent
    max := 0 // most far away from falling
    for _, v := range right {
        if n - v  > max {
            max = n - v
        }
    }
    for _, v := range left {
        if v > max {
            max = v
        }
    }
    return max
}
