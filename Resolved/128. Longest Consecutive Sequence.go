// https://leetcode.com/problems/longest-consecutive-sequence/
func longestConsecutive(nums []int) int {
    // add all nums to map
    // run through map and check if num is the beginning of the new sequence
    // if so, calculate length of sequence
    mapI := make(map[int]bool)

    for _, v := range nums {
        mapI[v] = true
    }

    result := 0

    for v := range mapI {
        if v == math.MinInt32 || !mapI[v-1] {
            // start of new sequence
            newV := v

            for mapI[newV] {
                if newV == math.MaxInt32 {
                    break
                }
                newV++
            }
            count := newV - v
            if result < count {
                result = count
            }
        }
    }

    return result
}
