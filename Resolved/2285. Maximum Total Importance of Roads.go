// https://leetcode.com/problems/maximum-total-importance-of-roads/
func maximumImportance(n int, roads [][]int) int64 {
    // 1. Calculate number of roads connected to each city
    roadCount := make([]int, n)
    for _, road := range roads {
        roadCount[road[0]]++
        roadCount[road[1]]++
    }

    // 2. Sort cities by number of connected to it roads
    citySorted := [][]int{}
    for i, v := range roadCount {
        citySorted = append(citySorted, []int{i, v})
    }
    sort.Slice(citySorted, func(i, j int) bool {
        return citySorted[i][1] < citySorted[j][1]
    })

    result := int64(0)
    // 3. Calculate result - assign 1..n number to each city according to previous sorting
    for i := 1; i <= n; i++ {
        result += int64(citySorted[i-1][1]) * int64(i)
    }
    return result
}
