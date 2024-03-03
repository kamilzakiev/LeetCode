func findMinHeightTrees(n int, edges [][]int) []int {
    // idea: 
    // 1. Take all leaves and start BFS from them
    // 2. Calculate max distance from leaves for each node
    // 2. Once popped from Queue, node is 'removed' from the tree, if that was last edge
    // 3. Nodes (1 or 2) with max distance (Last level of nodes) are the result
    if n == 1 {
        return []int{0}
    }
    edgeCount := make([]int, n)

    links := make([][]int, n)

    for _, edge := range edges {
        edgeCount[edge[0]]++
        edgeCount[edge[1]]++
        links[edge[0]] = append(links[edge[0]], edge[1])
        links[edge[1]] = append(links[edge[1]], edge[0])
    }

    leaves := []int{}
    for k, v := range edgeCount {
        if v == 1 {
            leaves = append(leaves, k)
        }
    }
    queue := leaves
    visited := make([]bool, n)
    dist := make([]int, n)

    for _, leaf := range leaves {
        visited[leaf] = true
        edgeCount[leaf]--
    }

    maxDist := 0
    maxDistNodes := append([]int{}, leaves...)

    for len(queue) > 0 {
        popped := queue[0]
        queue = queue[1:]

        for _, next := range links[popped] {
            //if !visited[next] {
                //fmt.Println(popped, next)
                visited[next] = true
                dist[next] = dist[popped] + 1
                edgeCount[next]--
                //fmt.Println(next, edgeCount[next])
                if edgeCount[next] == 1 {
                    if dist[next] > maxDist {
                        maxDist = dist[next]
                        maxDistNodes = []int{next}
                    } else if dist[next] == maxDist {
                        maxDistNodes = append(maxDistNodes, next)
                    }
                    queue = append(queue, next)
                }
            //}
        }
    }

    return maxDistNodes
}
