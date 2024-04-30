// https://leetcode.com/problems/number-of-wonderful-substrings/
func wonderfulSubstrings(word string) int64 {
    // 1. calculate prefixSum of bit representation
    // 2. find previous representation with at most 1 bit differ

    xor := 0
    mapI := make(map[int]int64)

    mapI[0] = 1
    result := int64(0)

    for i := 0; i < len(word); i++ {
        xor = xor ^ (1 << (word[i] - 'a'))
        result += mapI[xor] // no differ
        for i := 0; i < 10; i++ {   
            result += mapI[xor ^ (1<<i)] // 1 bit differ
        }
        //fmt.Println(result)
        mapI[xor]++
    }

    return result   
}
