// https://leetcode.com/problems/reduction-operations-to-make-the-array-elements-equal/
func reductionOperations(nums []int) int {
    mapI := make(map[int]int)
    for _, v := range nums {
        mapI[v]++
    }
    sortedKeys := []int{}
    for k := range mapI {
        sortedKeys = append(sortedKeys, k)
    } 
    sort.Ints(sortedKeys)
    result := 0
    for i := len(sortedKeys)-1; i >= 1; i-- {
        result += mapI[sortedKeys[i]] * i
    }
    return result
}
