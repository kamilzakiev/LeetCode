// https://leetcode.com/problems/gray-code/description/
func grayCode(n int) []int {
    // cannot do much without knowing this - https://www.eetimes.com/gray-code-fundamentals-part-2
    var iter func(currN int) []int 
    iter = func(currN int) []int {
        if currN == 0 {
            return []int{0}
        } else {
            prevResult := iter(currN - 1)
            result := make([]int, len(prevResult))
            copy(result, prevResult)
            for i := len(prevResult)-1; i >= 0; i-- {
                 newVal := 1 << (currN-1) + prevResult[i]                 
                 result = append(result, newVal)
            }
            return result
        }
    }
    return iter(n)
}
