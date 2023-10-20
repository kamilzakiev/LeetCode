// https://leetcode.com/problems/partition-labels
func partitionLabels(s string) []int {
    mapRune := make(map[rune]int)

    for i, r := range s {
        mapRune[r] = max(mapRune[r], i)
    }

    result := []int{}

    l, r := 0, 0
    for i, rune := range s{
        r = max(r, mapRune[rune])
        if i == r {
            result = append(result, r - l + 1)
            l = r + 1
        }
    }

    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
