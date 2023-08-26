https://leetcode.com/problems/bitwise-and-of-numbers-range/submissions
func rangeBitwiseAnd(left int, right int) int {
    // find mostleft bit=1 for left and right
    // if that's not the same bit, return 0
    // if left and right nums start with same bit=1, compose result from same bits from left to right
    if left == right {
        return left
    }
    i := 1 << 31
    for i >= 1 {
        if left & i > 0 && right & i > 0 {
            result := 0
            for i >= 1 && left & i == right & i {
                if left & i > 0 {
                    result = result | i
                }
                i = i >> 1
            }
            return result
        } else if left & i > 0 || right & i > 0 {
            return 0
        }
        i = i >> 1
    }
    return 0
}
