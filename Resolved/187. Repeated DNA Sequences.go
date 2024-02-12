// https://leetcode.com/problems/repeated-dna-sequences/description/
func findRepeatedDnaSequences(s string) []string {    
    mapR := make(map[rune]int)
    mapR['A'] = 0
    mapR['C'] = 1
    mapR['G'] = 2
    mapR['T'] = 3
    curr := 0
    
    result := []string{}
    mapCount := make(map[int]int)

    for i, r := range s {
        curr = curr << 2
        curr = curr | mapR[r]
        if i >= 9 {
            curr = curr & ((1 << 20)-1) // remove info on most left DNA
            //fmt.Println(s[i-9:i+1], strconv.FormatInt(int64(curr), 2), curr)
            mapCount[curr]++
            if mapCount[curr] == 2 {
                result = append(result, s[i-9:i+1])
            }
        }
    }
    return result
}
