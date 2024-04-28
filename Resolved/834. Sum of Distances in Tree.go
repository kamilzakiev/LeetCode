// https://leetcode.com/problems/sum-of-distances-in-tree/
func sumOfDistancesInTree(n int, edges [][]int) []int {
    links := make([][]int, n)

    for _, edge := range edges {
        a, b := edge[0], edge[1]
        links[a] = append(links[a], b)
        links[b] = append(links[b], a)
    }
    // 1.
    // count(i, j) is number of nodes in subtree of node i if link between i and j is removed
    var count func(i, j int) int
    memoCount := make(map[int]int, n)

    count = func(i, j int) int {
        hash := i*n+j
        if memoCount[hash] > 0 {
            return memoCount[hash]
        }
        result := 1
        for _, next := range links[i] {
            if next != j {
                result += count(next, i)
            }
        }
        memoCount[hash] = result
        memoCount[j*n+i] = n - result
        return result
    }

    // 2 calculate sum[0] - sum of the distances between the 0 node and all other nodes
    sum := make([]int, n)
    queue := []int{0}
    visited := make([]bool, n)
    dist := make([]int, n)
    visited[0] = true
    for len(queue) > 0 {
        popped := queue[0]
        queue = queue[1:]

        for _, next := range links[popped] {
            if !visited[next] {
                visited[next] = true
                dist[next] = dist[popped] + 1
                sum[0] += dist[next]
                queue = append(queue, next)
            }
        }
    }

    // 3. sum[j] = sum[i] - count[j,i] + count[i,j]
    queue = []int{0}
    visited = make([]bool, n)
    visited[0] = true
    for len(queue) > 0 {
        popped := queue[0]
        queue = queue[1:]

        for _, next := range links[popped] {
            if !visited[next] {
                visited[next] = true
                sum[next] = sum[popped] - count(next, popped) + count(popped, next)
                
                queue = append(queue, next)
            }
        }
    }

    

    return sum
}
