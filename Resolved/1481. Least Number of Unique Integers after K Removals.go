// https://leetcode.com/problems/least-number-of-unique-integers-after-k-removals/
func findLeastNumOfUniqueInts(arr []int, k int) int {
    mapI := make(map[int]int)

    for _, v := range arr {
        mapI[v]++
    }

    freq := []int{}
    for _,v := range mapI {
        freq = append(freq, v)
    }

    sort.Ints(freq)

    currCount := k
    for i := 0; i < len(freq); i++ {
        freqV := freq[i]
        if currCount < freqV {
            return len(freq) - i
        } else if currCount == freqV {
            return len(freq) - i - 1
        }
        currCount = currCount - freqV
    }
    return 0
}
