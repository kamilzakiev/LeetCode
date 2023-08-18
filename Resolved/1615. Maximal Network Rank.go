func maximalNetworkRank(n int, roads [][]int) int {
    // simple greedy solution - count roads from each city and then pick up pair with max rank
    // Scale of n (<=100) allows n^2 solution
    roadCount := make([]int, n)

    links := make([]map[int]bool, n)

    for _, road := range roads {
        roadCount[road[0]]++
        roadCount[road[1]]++

        if links[road[0]] == nil {
            links[road[0]] = make(map[int]bool)
        }
        links[road[0]][road[1]] = true
        if links[road[1]] == nil {
            links[road[1]] = make(map[int]bool)
        }
        links[road[1]][road[0]] = true
    }

    result := 0
    for i := range roadCount {
        for j := 0; j < i; j++ {
            rank := roadCount[i] + roadCount[j]
            if links[i][j] {
                rank--
            }
            if result < rank {
                result = rank
            }
        }
    }
    return result
}
