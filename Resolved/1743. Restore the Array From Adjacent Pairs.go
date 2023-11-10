// https://leetcode.com/problems/restore-the-array-from-adjacent-pairs/
func restoreArray(adjacentPairs [][]int) []int {
    adj := make(map[int][]int)
    for _, pair := range adjacentPairs {        
        adj[pair[0]] = append(adj[pair[0]], pair[1])
        adj[pair[1]] = append(adj[pair[1]], pair[0])
    }
    start := math.MinInt32
    end := math.MinInt32
    for k, v := range adj {
        if len(v) == 1 {
            if start == math.MinInt32 {
                start = k
            } else {
                end = k
            }
        }
    }
    //fmt.Println(start, end)
    //fmt.Println(adj)
    result := []int{}
    used := make(map[int]bool)
    curr := start
    for curr != end {
        result = append(result, curr)
        prev := curr
        list := adj[curr]
        if !used[list[0]] {
            curr = list[0]
        } else {
            curr = list[1]
        }
        used[prev] = true
    }
    result = append(result, end)

    return result
}
